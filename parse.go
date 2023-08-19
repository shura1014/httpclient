package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/clbanning/mxj/v2"
	"github.com/shura1014/common/utils/reflectutil"
	"github.com/shura1014/common/utils/stringutil"
	"net/http"
	"net/url"
	"reflect"
)

// 解析Json
func (client *Client) prepareJson(method, url string, data any) (request *Request, err error) {
	if data != nil {
		switch value := data.(type) {
		case string, []byte:
			content := stringutil.AnyToByte(value)
			req, err := http.NewRequest(method, url, bytes.NewReader(content))
			if err != nil {
				return nil, err
			}

			return &Request{Request: req, content: content}, nil
		default:
			if b, err := json.Marshal(data); err != nil {
				return nil, err
			} else {
				req, err := http.NewRequest(method, url, bytes.NewReader(b))
				if err != nil {
					return nil, err
				}
				return &Request{Request: req, content: b}, nil
			}
		}
	}

	return nil, nil
}

// 解析xml，标准库的xml解析太简单，后面需要替换
func (client *Client) prepareXml(method, url string, data any) (request *Request, err error) {
	var content []byte
	if data != nil {
		switch value := data.(type) {
		case string, []byte:
			content = stringutil.AnyToByte(value)
			req, err := http.NewRequest(method, url, bytes.NewReader(content))
			if err != nil {
				return nil, err
			}
			return &Request{Request: req, content: content}, nil
		case map[string]any:
			content, err = mxj.Map(value).Xml()
		case Map:
			content, err = mxj.Map(value).Xml()
		default:
			content, err = mxj.AnyXml(data)
		}
		if err != nil {
			return nil, err
		} else {
			req, err := http.NewRequest(method, url, bytes.NewReader(content))
			if err != nil {
				return nil, err
			}
			return &Request{Request: req, content: content}, nil
		}
	}

	return nil, nil
}

func ParseParam(data any) string {
	if data != nil {
		switch value := data.(type) {
		case string, []byte:
			return stringutil.ToString(value)
		case map[string]string:
			values := url.Values{}
			for k, v := range value {
				values.Set(k, v)
			}
			return values.Encode()
		case map[string]any:
			values := url.Values{}
			for k, v := range value {
				values.Set(k, stringutil.ToString(v))
			}
			return values.Encode()
		default:
			v := reflectutil.Indirect(data)
			if v.Kind() == reflect.Map {
				values := url.Values{}
				mapRange := v.MapRange()
				for mapRange.Next() {
					// 最多支持一层map
					value = mapRange.Value().Interface()
					vv := reflectutil.Indirect(mapRange.Value().Interface())
					if vv.Kind() == reflect.Map {
						k := mapRange.Key().String()
						sonMapRange := vv.MapRange()
						for sonMapRange.Next() {
							// http://127.0.0.1:8888/user/add?person=%7B%22xx%22%3A%22hh%22%7D&sex=1
							values.Set(fmt.Sprintf("%s[%s]", k, sonMapRange.Key().String()), stringutil.ToString(sonMapRange.Value().Interface()))
						}
					} else {
						values.Set(mapRange.Key().String(), stringutil.ToString(value))
					}
				}
				return values.Encode()
			}
		}
	}
	return ""
}
