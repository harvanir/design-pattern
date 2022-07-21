package command_before

import "log"

// HttpClient - receiver
type HttpClient struct{}

func newHttpClient() *HttpClient {
	return &HttpClient{}
}

func (hc *HttpClient) call(url string, requestBody interface{}) {
	log.Printf("url : %s, with requestBody: %s\n", url, requestBody)
}
