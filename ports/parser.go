package ports

type Parser interface {
	Check(URL string) bool
	Parse(URL string) (interface{}, error)
}
