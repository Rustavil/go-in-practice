package main

import (
	"github.com/braintree/manners"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type handler struct {
}

func (handler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	log.Println("Start")
	for i := 1; i <= 10; i++ {
		time.Sleep(5 * time.Second)
		log.Println("Ping", i)
	}
	log.Println("Done")
	resp.Write([]byte("Response"))
}

func main() {
	log.Println("PID", os.Getpid())
	handler := newHandler()
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill, syscall.SIGTERM)
	go listenShutdown(ch)
	err := manners.ListenAndServe(":8888", handler)
	if err != nil {
		panic(err)
	}
}

func listenShutdown(signals chan os.Signal) {
	<-signals
	log.Println("Received signal")
	manners.Close()
}

func newHandler() handler {
	return handler{}
}
