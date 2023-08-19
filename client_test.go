package httpclient

import (
	"testing"
)

func TestGet(t *testing.T) {
	client := NewClient()
	response, err := client.Get("http://127.0.0.1:8989/user/add", Map{
		"id":     "1001",
		"person": Map{"xx": "hh"},
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf(response.GetBodyString())
}

func Test(t *testing.T) {
	client := NewClient()
	client.EnableDump()
	response, err := client.Get("http://127.0.0.1:8989/api/remoteIP", nil)
	if err != nil {
		t.Error(err)
	}
	t.Logf(response.GetBodyString())
}

func TestFormUrl(t *testing.T) {
	client := NewClient()
	// 开启dump
	client.EnableDump()
	// 设置请求头
	client.SetHeader(ContentType, ApplicationXWWWFormUrlencoded)
	// 设置agent
	client.SetAgent("httpclient")
	response, err := client.Post("http://127.0.0.1:8989/user/add", Map{
		"sex":    "1",
		"person": Map{"xx": "hh"},
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf(response.GetBodyString())
}

func TestJSON(t *testing.T) {
	client := NewClient()
	client.SetHeader(ContentType, ApplicationJson)
	client.EnableDump()
	response, err := client.Post("http://127.0.0.1:8989/user/bind/json", Map{"name": "shura", "age": 20})
	if err != nil {
		t.Error(err)
	}
	t.Logf(response.GetBodyString())
}

func TestXml(t *testing.T) {
	client := NewClient()
	client.EnableDump()
	client.SetHeader(ContentType, ApplicationXml)
	response, err := client.Post("http://127.0.0.1:8989/user/bind/xml", Map{"name": "shura", "age": "25"})
	if err != nil {
		t.Error(err)
	}
	t.Logf(response.GetBodyString())
}
