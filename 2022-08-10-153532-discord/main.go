package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main () {
	discordToken := loadToken()

	session, err := discordgo.New("Bot " + discordToken)
	if err != nil {
		log.Fatal("Error in create session")
	}

	session.AddHandler(messageHandler)

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
