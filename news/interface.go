package news

type RSS interface {
	Parse(url string) (error, RSS)
	Print(date string) error
}
