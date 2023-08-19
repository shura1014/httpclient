package httpclient

import (
	"bytes"
	"github.com/shura1014/common/goerr"
	"net/http"
	"strings"
	"time"
)

// Get 建议类型是 string|[]byte|map[string]string|StringMap
func (client *Client) Get(url string, data ...any) (Response, error) {
	return client.Request(http.MethodGet, url, data...)
}

func (client *Client) Post(url string, data ...any) (Response, error) {
	return client.Request(http.MethodPost, url, data...)
}

func (client *Client) Put(url string, data ...any) (Response, error) {
	return client.Request(http.MethodPut, url, data...)
}

func (client *Client) Patch(url string, data ...any) (Response, error) {
	return client.Request(http.MethodPatch, url, data...)
}

func (client *Client) Head(url string, data ...any) (Response, error) {
	return client.Request(http.MethodHead, url, data...)
}

func (client *Client) Request(method, url string, args ...any) (Response, error) {
	var (
		req    *Request
		err    error
		rawReq *http.Request
		data   any
	)
	if len(args) > 0 {
		data = args[0]
	}
	switch client.Header[ContentType] {
	case ApplicationXml:
		req, err = client.prepareXml(method, url, data)
	case ApplicationJson:
		req, err = client.prepareJson(method, url, data)
	default:
		param := ParseParam(data)
		switch method {
		case http.MethodGet:
			if strings.Contains(url, "?") {
				url = url + "&" + param
			} else {
				url = url + "?" + param
			}
			rawReq, err = http.NewRequest(method, url, nil)
			req = &Request{Request: rawReq, content: []byte{}}
		default:
			rawReq, err = http.NewRequest(method, url, bytes.NewReader([]byte(param)))
			req = &Request{Request: rawReq, content: []byte(param)}
		}
	}
	if err != nil || req == nil || req.Request == nil {
		return nil, err
	}

	return client.Do(req)
}

func (client *Client) Do(req *Request) (Response, error) {
	// 设置请求头
	for k, v := range client.Header {
		req.Header.Set(k, v)
	}
	// 请求前做一些处理
	for _, handler := range client.RequestHandler {
		handler(req)
	}
	var (
		resp *http.Response
		err  error
	)
	for {
		// 是否需要dump
		if client.Dump {
			req.Dump()
		}
		if resp, err = client.Client.Do(req.Request); err != nil {
			err = goerr.Wrapf(err, "请求失败")
			if resp != nil {
				_ = resp.Body.Close()
			}
			if client.RetryCount > 0 {
				client.RetryCount--
				time.Sleep(client.RetryInterval)
			} else {
				break
			}
		} else {
			break
		}
	}
	response := client.NewDefaultResponse(resp)
	// 是否需要dump
	if client.Dump {
		response.Dump()
	}

	return response, err
}
