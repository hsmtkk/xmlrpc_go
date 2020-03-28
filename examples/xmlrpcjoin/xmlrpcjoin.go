package main

import (
	"fmt"
	"log"

	"github.com/hsmtkk/xmlrpc_go/pkg/logtransport"
	"github.com/hsmtkk/xmlrpc_go/pkg/xmlrpcjoin"
)

type request struct {
	First  string
	Second string
}

type response struct {
	Joined string
}

func main() {
	// without specifying client
	xj := xmlrpcjoin.New("http://127.0.0.1:8000/RPC2")
	req := request{First: "foo", Second: "bar"}
	res := response{}
	if err := xj.Join(&req, &res); err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.Joined)

	// with specifying client
	xj = xmlrpcjoin.NewWithClient(logtransport.New(), "http://127.0.0.1:8000/RPC2")
	if err := xj.Join(&req, &res); err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.Joined)
}
