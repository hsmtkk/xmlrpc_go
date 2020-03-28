package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"

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
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, responseBody)
	}))
	defer ts.Close()

	httpClient := ts.Client()
	xmlClient, err := xmlrpc.NewClient(ts.URL, httpClient.Transport)
	if err != nil {
		log.Fatal(err)
	}
	req := request{"foo", "bar"}
	res := response{}
	if err := xmlClient.Call("join", &req, &res); err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.Joined)
}

const responseBody = `<?xml version='1.0'?>
<methodResponse>
<params>
<param>
<value><struct>
<member>
<name>Joined</name>
<value><string>foobar</string></value>
</member>
</struct></value>
</param>
</params>
</methodResponse>`
