package main

import (
	"github.com/elazarl/goproxy"
	"log"
	"net"
	"net/http"
)

func main() {
	log.Print("Starting proxy")
	proxy := goproxy.NewProxyHttpServer()
	proxy.Tr.Dial = func(network, addr string) (c net.Conn, err error) {
		c, err = net.Dial(network, addr)
		if c, ok := c.(*net.TCPConn); err == nil && ok {
			log.Println("Set keep alive")
			c.SetKeepAlive(true)
			c.SetNoDelay(true)
		}
		return
	}
	proxy.Verbose = true
	log.Fatal(http.ListenAndServe("127.0.0.1:8123", proxy))
}