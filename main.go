package main

import (
	"crypto/tls"
	"flag"
	"time"

	"github.com/lucas-clemente/quic-go"
	"github.com/lucas-clemente/quic-go/http3"
)

var (
	port   int
	server string
)

func main() {

	flag.IntVar(&port, "p", 443, "server port")
	flag.StringVar(&server, "s", "", "server address")


	rt := &http3.RoundTripper{
		TLSClientConfig: &tls.Config{

		},
		QuicConfig: &quic.Config{
			KeepAlive:      true,
			Versions:       []quic.VersionNumber{quic.VersionDraft29},
			MaxIdleTimeout: 3 * time.Second,
		},
	}
}