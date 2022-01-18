package main

import (
	"flag"

	"github.com/lucas-clemente/quic-go/http3"
)

var (
	port   int
	server string
)

func main() {

	flag.IntVar(&port, "p", 443, "server port")
	flag.StringVar(&server, "s", "", "server address")

}