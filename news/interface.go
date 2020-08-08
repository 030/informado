package news

type RSS interface {
	Parse(b []byte) (RSS, error)
	Print() error
}
