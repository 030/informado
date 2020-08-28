package news

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Title struct {
	Name string `xml:"title"`
}

func dateToEpoch(date string) (int64, error) {
	var t time.Time
	var err error

	dateZ := regexp.MustCompile(`^.*\dT\d.*$`)
	dateZone := regexp.MustCompile(`^.*GMT$`)
	dateNumericZone := regexp.MustCompile(`^.*(\-|\+)\d+$`)

	switch {
	case dateZ.MatchString(date):
		t, err = time.Parse(time.RFC3339, date)
	case dateZone.MatchString(date):
		t, err = time.Parse(time.RFC1123, date)
	case dateNumericZone.MatchString(date):
		t, err = time.Parse(time.RFC1123Z, date)
	default:
		return 0, fmt.Errorf("'%v' cannot be parsed. Check whether the date matches the regex", date)
	}

	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

func lastTimeInformadoWasRun() (int64, error) {
	date, err := ioutil.ReadFile("/tmp/some-repo3/.informado")
	if err != nil {
		return 0, err
	}

	n, err := strconv.ParseInt(strings.TrimSuffix(string(date), "\n"), 10, 64)
	if err != nil {
		return 0, err
	}

	return n, nil
}
