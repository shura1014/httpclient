package httpclient

import "net/http"

type Request struct {
	*http.Request
	content []byte
}
