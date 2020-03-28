package xmlrpcjoin_test

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/hsmtkk/xmlrpc_go/pkg/xmlrpcjoin"
	"github.com/stretchr/testify/assert"
)

type request struct {
	First  string
	Second string
}

type response struct {
	Joined string
}

func TestJoin(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotRequest, err := ioutil.ReadAll(r.Body)
		assert.Nil(t, err, "error should not occur")
		assert.Equal(t, wantRequest, string(gotRequest), "should match")
		fmt.Fprintln(w, responseBody)
	}))
	defer ts.Close()

	xrj := xmlrpcjoin.NewWithClient(ts.Client(), ts.URL)
	req := request{First: "foo", Second: "bar"}
	res := response{}
	err := xrj.Join(&req, &res)
	assert.Nil(t, err, "error should not occur")
	assert.Equal(t, "foobar", res.Joined, "should match")
}

func equalXML(want, got []byte) bool {
	var xa interface{}
	var xb interface{}
	if err := xml.Unmarshal([]byte(want), xa); err != nil {
		return false
	}
	if err := xml.Unmarshal([]byte(want), xb); err != nil {
		return false
	}
	return reflect.DeepEqual(xa, xb)
}

const wantRequest = `<?xml version="1.0" encoding="UTF-8"?><methodCall><methodName>join</methodName><params><param><value><struct><member><name>First</name><value><string>foo</string></value></member><member><name>Second</name><value><string>bar</string></value></member></struct></value></param></params></methodCall>`

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
