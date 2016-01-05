package monitor

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andreandradecosta/rpimonitor/models"

	"gopkg.in/unrolled/render.v1"
)

func setup(t *testing.T) (*Monitor, *httptest.ResponseRecorder, *http.Request) {
	req, err := http.NewRequest("GET", "http://teste.com/", nil)
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()
	m := &Monitor{Render: render.New()}
	return m, res, req
}

func TestIndex(t *testing.T) {
	m, res, req := setup(t)
	m.Index(res, req)
	info := make(models.Info)
	err := json.NewDecoder(res.Body).Decode(&info)
	if err != nil {
		t.Fatalf("Can't decode the result: %s", err)
	}
}

func TestSnapshot(t *testing.T) {
	m, res, req := setup(t)
	m.Snapshot(res, req)
	sample := &models.Sample{}
	err := json.NewDecoder(res.Body).Decode(sample)
	if err != nil {
		t.Fatalf("Can't decode the result: %s", err)
	}

}
