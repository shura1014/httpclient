package httpclient

import (
	"net/http"
	"time"
)

const (
	ContentType                   = "Content-Type"
	ApplicationJson               = "application/json"
	ApplicationXml                = "application/xml"
	ApplicationXWWWFormUrlencoded = "application/x-www-form-urlencoded"
	UserAgent                     = "User-Agent"
)

type Map map[string]any
type StringMap map[string]string

type RequestHandler func(request *Request)

type Client struct {
	http.Client
	*ClientConfig
}

func NewClient(config ...*ClientConfig) *Client {
	if len(config) > 0 {
		return &Client{ClientConfig: config[0]}
	}
	return &Client{ClientConfig: &ClientConfig{Header: make(map[string]string)}}
}

func Default() *Client {

	return &Client{
		ClientConfig: &ClientConfig{
			RetryCount:    0,
			RetryInterval: 0,
			Header:        StringMap{ContentType: ApplicationJson},
		},
	}
}

func (client *Client) SetHeader(k, v string) {
	client.Header[k] = v
}

func (client *Client) SetHeaderMap(header map[string]string) {
	for k, v := range header {
		client.Header[k] = v
	}
}

func (client *Client) SetRetryCount(retryCount int) {
	client.RetryCount = retryCount
}

func (client *Client) SetRetryInterval(interval time.Duration) {
	client.RetryInterval = interval
}

func (client *Client) AddRequestHandler(handler ...RequestHandler) {
	client.RequestHandler = append(client.RequestHandler, handler...)
}

func (client *Client) AddResponseHandler(handler ...ResponseHandler) {
	client.ResponseHandler = append(client.ResponseHandler, handler...)
}

func (client *Client) EnableDump() {
	client.Dump = true
}
