package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HelloWorldResponse struct {
	Message string
}

func HandleRequests(w http.ResponseWriter, r *http.Request) {
	hello := &HelloWorldResponse{Message: "Hello World"}
	response, err := json.Marshal(hello)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, string(response))
}
func main() {
	http.HandleFunc("/", HandleRequests)
	http.ListenAndServe(":80", nil)
}
