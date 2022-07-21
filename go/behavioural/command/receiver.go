package command

import "log"

type (
	// Receiver - external code
	Receiver interface {
		call(url string, requestBody interface{})
	}
)

var sai Invoker

// HttpClient - receiver
type HttpClient struct{}

func newHttpClient() *HttpClient {
	return &HttpClient{}
}

func (hc *HttpClient) call(url string, requestBody interface{}) {
	log.Printf("url : %s, with requestBody: %s\n", url, requestBody)
}
