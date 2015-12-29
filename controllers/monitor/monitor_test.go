package monitor

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"gopkg.in/unrolled/render.v1"
)

func Test_Index(t *testing.T) {
	req, err := http.NewRequest("GET", "http://teste.com/", nil)
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()
	m := &Monitor{Render: render.New()}
	m.Index(res, req)

	exp := "Index"
	act := res.Body.String()
	if exp != act {
		t.Fatalf("Expected %s got %s", exp, act)
	}
}
