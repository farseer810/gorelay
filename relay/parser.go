package relay

import (
	"net/http"
	"io/ioutil"
)

type RequestInfo struct {
	Method string
	Header http.Header
	Body []byte
}

func ParseRequest(r *http.Request) *RequestInfo {
	var info = new(RequestInfo)
	info.Method = r.Method
	info.Header = parseHeader(r)
	info.Body = parseBody(r)
	return info
}

func parseBody(r *http.Request) []byte {
	var body, _ = ioutil.ReadAll(r.Body)
	if len(body) == 0 {
		return nil
	}
	return body
}

func parseHeader(r *http.Request) http.Header {
	var header http.Header
	deepCopy(&header, r.Header)
	delete(header, "Gorelay-Targeturi")
	return header
}