package http

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestName(t *testing.T) {
	//req := httptest.NewRequest(http.MethodGet, "/someplace", nil)
	w := httptest.NewRecorder()

	JSONError(w, "hello world", 400)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	expected := "{\"Code\":400,\"Error\":\"hello world\"}\n"
	if string(data) != expected {
		t.Errorf("expected [%s] got [%v]", expected, string(data))
	}

}
