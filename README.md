# httpclient

一个简单的http客户端调用工具

# 快速使用
client.EnableDump() 开启dump，可以打印出请求的详细信息
client.SetHeader() 设置全局请求头
client.SetAgent() 设置全局agent

## 最简单的使用
```go
func Simple() {
	client := httpclient.NewClient()
	response, _ := client.Get("http://127.0.0.1:8989/api/remoteIP")
	fmt.Println(response.GetBodyString())
}

{"code":200,"msg":"OK","data":"127.0.0.1"}
```

## FormUrlencoded


```go
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
```

结果

```text
+--------------------Request-------------------------+
POST /user/add HTTP/1.1
Host: 127.0.0.1:8989
User-Agent: httpclient
Content-Length: 23
Content-Type: application/x-www-form-urlencoded
Accept-Encoding: gzip


person%5Bxx%5D=hh&sex=1

+--------------------Response-------------------------+
HTTP/1.1 200 OK
Transfer-Encoding: chunked
Date: Sat, 19 Aug 2023 03:17:50 GMT
```


## Json请求

```go
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

```

结果
```text
+--------------------Request-------------------------+
POST /user/bind/json HTTP/1.1
Host: 127.0.0.1:8989
User-Agent: Go-http-client/1.1
Content-Length: 25
Content-Type: application/json
Accept-Encoding: gzip


{"age":20,"name":"shura"}
 
+--------------------Response-------------------------+
HTTP/1.1 200 OK
Transfer-Encoding: chunked
Content-Type: text/plain; charset=utf-8
Date: Sat, 19 Aug 2023 03:22:35 GMT


{"code":200,"msg":"OK","data":"success"}
```

## Xml
```go
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
```

结果

```text
+--------------------Request-------------------------+
POST /user/bind/xml HTTP/1.1
Host: 127.0.0.1:8989
User-Agent: Go-http-client/1.1
Content-Length: 42
Content-Type: application/xml
Accept-Encoding: gzip


<User><age>25</age><name>shura</name></User>
 
+--------------------Response-------------------------+
HTTP/1.1 200 OK
Transfer-Encoding: chunked
Content-Type: text/plain; charset=utf-8
Date: Sat, 19 Aug 2023 03:23:54 GMT


{"code":200,"msg":"OK","data":"success"}
```