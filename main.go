package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/lucas-clemente/quic-go"
	"github.com/lucas-clemente/quic-go/http3"
)

var (
	port   int
	server string
	count  int
	typ    string
	size   int
)

type SS struct {
}

func (ss *SS) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	buffer := make([]byte, size)
	w.Header().Add("Content-Length", fmt.Sprintf("%d", size))
	w.WriteHeader(200)
	w.Write(buffer)
}



func serverHandler() {

	err := http3.ListenAndServeQUIC("0.0.0.0:443", "giaclient.crt", "giaclient.key", &SS{})
	println("listen server failed, err:", err.Error())
}

func main() {

	flag.IntVar(&port, "p", 443, "server port")
	flag.StringVar(&server, "s", "", "server address")
	flag.IntVar(&count, "c", 1, "count")
	flag.StringVar(&typ, "t", "client", "programe type client/server")
	flag.IntVar(&size, "size", 1024*1024, "size")

	flag.Parse()

	if server == "" {
		println("with no server")
		return
	}

	if typ == "server" {
		serverHandler()
	}

	for i := 0; i < count; i++ {
		go func() {
			for {


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

				request, err := http.NewRequest("GET", fmt.Sprintf("https://%s/test.file", server), nil)
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
		}()
	}
	time.Sleep(time.Hour)
}
