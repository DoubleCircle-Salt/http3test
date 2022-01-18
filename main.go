package main

import (
	"crypto/tls"
	"flag"
	"net/http"
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




	roundTripper := &http3.RoundTripper{
		TLSClientConfig: &tls.Config{
			ServerName:         "giasstest.ecn.zenlayer.net",
			ClientSessionCache: tls.NewLRUClientSessionCache(64),
			InsecureSkipVerify: true,
		},
		QuicConfig: &quic.Config{
			KeepAlive:      true,
			Versions:       []quic.VersionNumber{quic.VersionDraft29},
			MaxIdleTimeout: 3 * time.Second,
		},
	}

	request, err := http.NewRequest("GET", "https://128.14.141.109/test.file", nil)
	if err != nil {
		println("create request failed, error:", err.Error())
		return
	}

	response, err := roundTripper.RoundTrip(request)
	if err != nil {
		println("get response failed, error:", err.Error())
		return
	}

	println("status:", response.Status)

	buffer := make([]byte, 4096)
	n := 0
	for {
		nn, err := response.Body.Read(buffer)
		n += nn
		if err != nil {
			println("read failed, err:", err.Error())
			return
		}
	}

	println("read", n, "bytes")
}
