package main

import (
	"github.com/sevlyar/go-daemon"
	"github.com/farseer810/gorelay/relay"
	"log"
	"net/http"
	"flag"
	"fmt"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "port to be listened")
	flag.Parse()
	var listenAddr string
	listenAddr = fmt.Sprintf("0.0.0.0:%d", port)

	cntxt := &daemon.Context{
		PidFileName: "gorelay.pid",
		PidFilePerm: 0644,
		// LogFileName: "gorelay.log",
		// LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
		Args:        []string{listenAddr},
	}

	d, err := cntxt.Reborn()
	if err != nil {
		log.Fatal("Unable to run: ", err)
	} 
	if d != nil {
		fmt.Println("child process id:", d.Pid)
		return
	}
	defer cntxt.Release()

	listenAddr = cntxt.Args[0]
	fmt.Println("child listening on", listenAddr)

	http.HandleFunc("/", relay.RelayHandler)
	log.Fatalln(http.ListenAndServe(listenAddr, nil))
}