package logtransport

import (
	"log"
	"net/http"
	"net/http/httputil"
)

func New() *http.Client {
	return &http.Client{Transport: &logTransport{}}
}

type logTransport struct{}

func (lg *logTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	reqBytes, err := httputil.DumpRequest(req, true)
	if err != nil {
		return nil, err
	}
	log.Println(string(reqBytes))
	res, err := http.DefaultTransport.RoundTrip(req)
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
