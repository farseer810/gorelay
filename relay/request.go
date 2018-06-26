package relay

import (
	"fmt"
	"net/http"
	"bytes"
	"io/ioutil"
)

type Body struct {
	buf *bytes.Buffer
}

func (b *Body) Close() error {
	return nil
}

func (b *Body) Read(p []byte) (n int, err error) {
	return 
}

func SendRequest(URI, method string, header http.Header, body []byte) {
	var client = &http.Client{}

	var r, _ = http.NewRequest(method, URI, bytes.NewBuffer(body))
	r.Header = header

	var res, err = client.Do(r)
	if err != nil {
		fmt.Println(res, err)
	} else {
		var result, err = ioutil.ReadAll(res.Body)
		fmt.Println(err)
		fmt.Println(string(result))
	}
}