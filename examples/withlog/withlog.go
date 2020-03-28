package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/kolo/xmlrpc"
)

type request struct {
	First  string
	Second string
}

type response struct {
	Joined string
}

func main() {
	client, _ := xmlrpc.NewClient("http://127.0.0.1:8000/RPC2", new())
	req := request{"foo", "bar"}
	res := response{}
	if err := client.Call("join", &req, &res); err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.Joined)
}

func new() http.RoundTripper {
	return &logTripper{transport: http.DefaultTransport}
}

type logTripper struct {
	transport http.RoundTripper
}

func (lg *logTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	reqBytes, err := httputil.DumpRequest(req, true)
	if err != nil {
		return nil, err
	}
	log.Println(string(reqBytes))
	res, err := lg.transport.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	resBytes, err := httputil.DumpResponse(res, true)
	if err != nil {
		return nil, err
	}
	log.Println(string(resBytes))
	return res, nil
}
