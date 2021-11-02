package news

import (
	"informado/internal/pkg/slack"
)

type RSS interface {
	Parse(b []byte) (RSS, error)
	Print(lastTimeInformadoWasRun int64, slackChannel slack.Channel) error
}
