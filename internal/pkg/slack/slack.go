package slack

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/slack-go/slack"
)

type Channel struct {
	ID, Token string
	Msg       string `validate:"nonzero"`
}

func (c *Channel) Send() error {
	log.Info("Sending message to Slack...")
	api := slack.New(c.Token)
	channelID, timestamp, err := api.PostMessage(
		c.ID,
		slack.MsgOptionText(c.Msg, false),
		slack.MsgOptionAsUser(false),
	)
	if err != nil {
		return err
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)

	return nil
}
