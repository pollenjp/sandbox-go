package app

import "log"

var (
	Info info
)

type info struct {
	channelIDForNotification string
	channelIDForPomodoroVC string
}

func InitInfo(
	channelIDForNotification string,
	channelIDForPomodoroVC string,
) {
	Info = info {
		channelIDForNotification: channelIDForNotification,
		channelIDForPomodoroVC: channelIDForPomodoroVC,
	}
}


func (i *info) SetChannelIDForNotification(channelID string) {
	if len(channelID) == 0 {
		log.Fatal("no channel id for notification exists.")
	}
	i.channelIDForNotification = channelID
}

func (i *info) GetChannelIDForNotification() string {
	if len(i.channelIDForNotification) == 0 {
		log.Fatal("no channel id for notification exists.")
	}
	return i.channelIDForNotification
}
func (i *info) SetChannelIDForPomodoroVC(channelID string) {
	if len(channelID) == 0 {
		log.Fatal("no channel id for notification exists.")
	}
	i.channelIDForPomodoroVC = channelID
}

func (i *info) GetChannelIDForPomodoroVC() string {
	if len(i.channelIDForPomodoroVC) == 0 {
		log.Fatal("no channel id for notification exists.")
	}
	return i.channelIDForPomodoroVC
}
