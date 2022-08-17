package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/pollenjp/sandbox-go/2022-08-17-081309-discord-vcupdate/app"
)

func init() {
	app.InitInfo(
		os.Getenv("CHANNEL_ID_FOR_NOTIFICATION"),
		os.Getenv("CHANNEL_ID_FOR_POMODORO_VC"),
	)
}

func main() {
	discordToken := loadToken()

	fmt.Printf("Info: %+v\n", app.Info)

	session, err := discordgo.New("Bot " + discordToken)
	if err != nil {
		log.Fatal("Error in create session")
	}

	session.AddHandler(messageHandler)
	session.AddHandler(onVoiceStateUpdate)

	if err = session.Open(); err != nil {
		panic(err)
	}
	defer session.Close()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	log.Print("booted!!!")

	<-sc

}

func loadToken() string {
	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		log.Fatal("no discord token exists.")
	}
	return token
}

func messageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.Bot {
		// interrupt conversation with "bot"
		return
	}
	log.Printf("%20s %20s > %s", message.ChannelID, message.Author.Username, message.Content)

	switch {
	case message.Content == "ping":
		sendMessage(session, message.ChannelID, "pong")
	}

}

func sendMessage(s *discordgo.Session, channelID string, msg string) {
	_, err := s.ChannelMessageSend(channelID, msg)

	log.Print(channelID, msg)
	if err != nil {
		log.Print("Error sending message: ", err)
	}
}

func onVoiceStateUpdate(session *discordgo.Session, updated *discordgo.VoiceStateUpdate) {
	var user *discordgo.User
	user, err := session.User(updated.UserID)
	if err != nil {
		log.Print("Error getting user: ", err)
		return
	}

	var beforeChName string = "unknown"
	if updated.BeforeUpdate != nil {
		beforeCh, err := session.Channel(updated.BeforeUpdate.ChannelID)
		if err == nil {
			beforeChName = beforeCh.Name
		}
	}

	var currentChName string = "unknown"
	currentCh, err := session.Channel(updated.ChannelID)
	if err == nil {
		currentChName = currentCh.Name
	}

	sendMessage(
		session,
		app.Info.GetChannelIDForNotification(),
		fmt.Sprintf(
			"%s's voice status was updated (before: %s ) -> (after: %s )",
			user.Username,
			beforeChName,
			currentChName,
		),
	)
}
