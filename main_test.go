package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMainReturnsJson(t *testing.T) {

	var (
		contentType string = "application/json"
	)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	HandleRequests(w, req)

	res := w.Result()

	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		t.Error(err)
	}

	if w.Result().Header.Get("Content-Type") != contentType {
		t.Errorf("Expected Content-Type %s but got %s", contentType, w.Result().Header.Get("Content-Type"))
	}

	var helloWorldMessage HelloWorldResponse
	dataBytes := []byte(data)
	err = json.Unmarshal(dataBytes, &helloWorldMessage)

	if err != nil {
		t.Error(err)
	}
}
