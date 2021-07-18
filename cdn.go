package uoa

type cdn interface {
	sign(original string) (url string, err error)
}
