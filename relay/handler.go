package relay

import (
	"net/http"
)

func RelayHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" && r.Method != "POST" {
		w.WriteHeader(400)
		return
	}

	var _, hasKey = r.Header["Gorelay-Targeturi"]
	if !hasKey {
		w.WriteHeader(400)
		return
	}

	var targetURI = r.Header.Get("Gorelay-Targeturi")
	var info = ParseRequest(r)
	go SendRequest(targetURI, info.Method, info.Header, info.Body)

	w.WriteHeader(200)
}