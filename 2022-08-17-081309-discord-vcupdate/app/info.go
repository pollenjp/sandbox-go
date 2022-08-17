package app

import "log"

var (
	Info info
)

type info struct {
	channelIDForNotification string
	channelIDForPomodoroVC   string
}

func InitInfo(
	channelIDForNotification string,
	channelIDForPomodoroVC string,
) {
	Info = info{
		channelIDForNotification: channelIDForNotification,
		channelIDForPomodoroVC:   channelIDForPomodoroVC,
	}
}

func (i *info) GetChannelIDForNotification() string {
	if len(i.channelIDForNotification) == 0 {
		log.Fatal("no channel id for notification exists.")
	}
	return i.channelIDForNotification
}

func (i *info) GetChannelIDForPomodoroVC() string {
	if len(i.channelIDForPomodoroVC) == 0 {
		log.Fatal("no channel id for notification exists.")
	}
	return i.channelIDForPomodoroVC
}
