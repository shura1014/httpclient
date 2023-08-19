package main

import (
	"fmt"
	"github.com/shura1014/httpclient"
	"log"
)

func main() {
	//FormUrlencoded()
	//Json()
	//Xml()
	Simple()
}

func Simple() {
	client := httpclient.NewClient()
	response, _ := client.Get("http://127.0.0.1:8989/api/remoteIP")
	fmt.Println(response.GetBodyString())
}
func FormUrlencoded() {
	client := httpclient.NewClient()
	// 开启dump
	client.EnableDump()
	// 设置请求头
	client.SetHeader(httpclient.ContentType, httpclient.ApplicationXWWWFormUrlencoded)
	// 设置agent
	client.SetAgent("httpclient")
	response, err := client.Post("http://127.0.0.1:8989/user/add", httpclient.Map{
		"sex":    "1",
		"person": httpclient.Map{"xx": "hh"},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response.GetBodyString())
}

func Json() {
	client := httpclient.NewClient()
	client.SetHeader(httpclient.ContentType, httpclient.ApplicationJson)
	client.EnableDump()
	response, err := client.Post("http://127.0.0.1:8989/user/bind/json", httpclient.Map{"name": "shura", "age": 20})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response.GetBodyString())
}

func Xml() {
	client := httpclient.NewClient()
	client.EnableDump()
	client.SetHeader(httpclient.ContentType, httpclient.ApplicationXml)
	// 自定义rootTag User
	response, err := client.Post("http://127.0.0.1:8989/user/bind/xml", httpclient.Map{"User": httpclient.Map{"name": "shura", "age": "25"}})
	//response, err := client.Post("http://127.0.0.1:8989/user/bind/xml", httpclient.Map{"name": "shura", "age": "25"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response.GetBodyString())
}
