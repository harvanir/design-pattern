package command_before

import "log"

type httpClient struct{}

func newHttpClient() HttpClient {
	return &httpClient{}
}

func (hc *httpClient) Call(url string, requestBody interface{}) {
	log.Printf("url : %s, with requestBody: %s\n", url, requestBody)
}
