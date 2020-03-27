package main

import (
	"fmt"
	"log"

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
	client, _ := xmlrpc.NewClient("http://127.0.0.1:8000/RPC2", nil)
	req := request{"foo", "bar"}
	res := response{}
	if err := client.Call("join", &req, &res); err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.Joined)
}
