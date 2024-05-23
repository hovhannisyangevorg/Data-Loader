package DPClient

import (
	"bytes"
	"fmt"
	"github.com/hovhannisyangevorg/Data-Procesor/internal/utils"
	"net/http"
)

type HTTPClient struct {
	httpClient  *http.Client
	httpRequest *http.Request
}

func NewClient() *HTTPClient {
	return &HTTPClient{
		httpClient:  &http.Client{},
		httpRequest: nil,
	}
}

func (c *HTTPClient) SetRequest(method, url, apiKey string, body []byte) error {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return utils.WrapError("SetRequest", err)
	}
	if apiKey != "" {
		req.Header.Set("X-Api-Key", apiKey)
	}
	c.httpRequest = req
	return nil
}

func (c *HTTPClient) Do() (*http.Response, error) {
	if c.httpRequest == nil {
		return nil, fmt.Errorf("do: http.Request is not set")
	}

	resp, err := c.httpClient.Do(c.httpRequest)
	if err != nil {
		return nil, utils.WrapError("do: error sending request", err)
	}
	return resp, nil
}
