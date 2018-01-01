package rest

import (
	"net/url"
	"testing"
)

func TestNewRequest(t *testing.T) {
	rq := NewRequest("text/xml", "text/xml")

	t.Log(rq)
}

func TestGet(t *testing.T) {
	rq := NewRequest("text/xml", "text/xml")
	v := url.Values{}
	v.Add("a", "t-a")
	v.Add("b", "t-b")
	//rs, err := rq.Get("http://master:9099/test/row1", &v)
	rs, err := rq.Get("http://master:9099/?" + v.Encode())

	t.Log(rs)
	t.Log(err)
}
