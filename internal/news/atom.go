package news

import (
	"encoding/xml"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/slack-go/slack"
)

type Link struct {
	Href string `xml:"href,attr"`
}
type Entry struct {
	XMLName xml.Name `xml:"entry"`
	Updated string   `xml:"updated"`
	Title
	Link Link `xml:"link"`
}
type Atom struct {
	Feed  xml.Name `xml:"feed"`
	Entry []Entry  `xml:"entry"`
}

func (a Atom) Parse(b []byte) (RSS, error) {
	if err := xml.Unmarshal(b, &a); err != nil {
		return Atom{}, err
	}
	return a, nil
}

func (a Atom) Print(lastTimeInformadoWasRun int64) error {
	for i := 0; i < len(a.Entry); i++ {
		updated := a.Entry[i].Updated
		updatedInt64, err := dateToEpoch(updated)
		if err != nil {
			return err
		}

		if updatedInt64 > lastTimeInformadoWasRun {
			msg := updated + " " + a.Entry[i].Title.Name + " " + a.Entry[i].Link.Href
			fmt.Println(msg)
			if false {
				if err := sendMessage("x", msg, "y"); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func sendMessage(channelID, msg, token string) error {
	if channelID == "" || token == "" || msg == "" {
		return fmt.Errorf("channelID, slack_token or msg should not be empty")
	}

	log.Info("Sending message to Slack...")
	api := slack.New(token)
	channelID, timestamp, err := api.PostMessage(
		channelID,
		slack.MsgOptionText(msg, false),
		slack.MsgOptionAsUser(false),
	)
	if err != nil {
		return err
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)

	return nil
}
