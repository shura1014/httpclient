package httpclient

import "time"

type ClientConfig struct {
	RetryCount      int               // 重试次数
	RetryInterval   time.Duration     // 重试间隔
	Header          map[string]string // 每个请求都将携带的header 例如 application/json
	RequestHandler  []RequestHandler  // 对请求前做一些处理
	ResponseHandler []ResponseHandler // 对响应做一些处理
	Dump            bool
}

// SetTimeout 超时
func (client *Client) SetTimeout(t time.Duration) *Client {
	client.Client.Timeout = t
	return client
}

// SetAgent User-Agent
func (client *Client) SetAgent(agent string) *Client {
	client.Header[UserAgent] = agent
	return client
}
