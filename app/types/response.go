package types

type Response struct {
	Protocol   string
	Status     int
	StatusText string
	Headers    map[string]string
	Body       string
}
