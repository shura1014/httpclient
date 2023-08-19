package httpclient

import (
	"net/http/httputil"
)

const dumpTextFormat = `
+--------------------%s-------------------------+
%s
%s
`

// Dump 导出请求
func (req *Request) Dump() {
	bs, err := httputil.DumpRequestOut(req.Request, false)
	if err != nil {
		Error("%+v", err)
		return
	}
	Debug(
		dumpTextFormat,
		"Request",
		string(bs),
		string(req.content),
	)
}

// Dump 导出响应
func (resp *DefaultResponse) Dump() {
	if resp == nil || resp.Response == nil {
		return
	}
	bs, err := httputil.DumpResponse(resp.Response, false)
	if err != nil {
		Error("%+v", err)
		return
	}
	Debug(
		dumpTextFormat,
		"Response",
		string(bs),
		resp.GetBodyString(),
	)
}
