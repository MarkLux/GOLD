package http

//interface for http client
type GoldHttpClient interface {
	Get(url string, body interface{}, headers map[string]string) (GoldHttpResponse, err error)
	Post(url string, body interface{}, headers map[string]string) (GoldHttpResponse, err error)
	Put(url string, body interface{}, headers map[string]string) (GoldHttpResponse, err error)
	Delete(url string, body interface{}, headers map[string]string) (GoldHttpResponse, err error)
}

//type for response
type GoldHttpResponse struct {
	Headers map[string]string
	Body interface{}
}

//common used errors
type HttpRequestError struct {
	Message string
}

func (e HttpRequestError) Error() string {
	return e.Message
}
