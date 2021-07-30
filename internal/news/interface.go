package news

type RSS interface {
	Parse(b []byte) (RSS, error)
	Print(lastTimeInformadoWasRun int64) error
}
