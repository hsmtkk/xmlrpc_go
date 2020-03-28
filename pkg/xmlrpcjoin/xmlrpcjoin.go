package xmlrpcjoin

import (
	"net/http"

	"github.com/kolo/xmlrpc"
)

type StringJoiner interface {
	Join(request, response interface{}) error
}

func New(url string) StringJoiner {
	return &stringJoinerImpl{http.DefaultClient, url}
}

func NewWithClient(client *http.Client, url string) StringJoiner {
	return &stringJoinerImpl{client, url}
}

type stringJoinerImpl struct {
	client *http.Client
	url    string
}

func (sj *stringJoinerImpl) Join(req, res interface{}) error {
	xmlClt, err := xmlrpc.NewClient(sj.url, sj.client.Transport)
	if err != nil {
		return err
	}
	if err := xmlClt.Call("join", req, res); err != nil {
		return err
	}
	return nil
}
