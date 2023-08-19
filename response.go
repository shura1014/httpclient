package httpclient

import (
	"io"
	"net/http"
)

type ResponseHandler func(bytes []byte) []byte

type Response interface {
	RawResponse() *http.Response
	Wrap(*http.Response)
	GetBody() []byte
	GetBodyString() string
	Header() http.Header
	StatusCode() int
	Close() error
	Dump()
}

type BaseResponse struct {
	*http.Response
	client *Client
}

func (base *BaseResponse) RawResponse() *http.Response {
	return base.Response
}

func (base *BaseResponse) Wrap(resp *http.Response) {
	base.Response = resp
}

type DefaultResponse struct {
	*BaseResponse
}

func (client *Client) NewDefaultResponse(resp *http.Response) *DefaultResponse {
	return &DefaultResponse{
		BaseResponse: &BaseResponse{resp, client},
	}
}

// GetBody 返回nil代表异常
func (resp *DefaultResponse) GetBody() []byte {
	if resp.Response == nil {
		return []byte{}
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		Error("%+v", err)
		return nil
	}

	for _, handler := range resp.client.ResponseHandler {
		body = handler(body)
	}
	return body
}

func (resp *DefaultResponse) GetBodyString() string {
	return string(resp.GetBody())
}

func (resp *DefaultResponse) StatusCode() int {
	return resp.Response.StatusCode
}

func (resp *DefaultResponse) Header() http.Header {
	return resp.Response.Header
}

// Close 方便关闭resp.Response
func (resp *DefaultResponse) Close() error {
	if resp == nil || resp.Response == nil {
		return nil
	}
	return resp.Response.Body.Close()
}
