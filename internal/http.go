package internal

import (
	"io"
	"net/http"
)

const defaultSchema = "http://"

// HTTPWrap can do simple http requests.
type HTTPWrap struct {
	client *http.Client
}

func NewHTTPWrap(cl *http.Client) *HTTPWrap{
	return &HTTPWrap{client: cl}
}


// MakeRequest do HTTP GET request, write result in channel in any case.
func (h *HTTPWrap) MakeRequest(url string) []byte{
	resp, err := h.client.Get(defaultSchema + url)
	if err != nil {
		return []byte{}
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}
	}
	return body
}
