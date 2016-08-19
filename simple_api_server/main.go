package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	Text string
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome, %s!", r.URL.Path[1:])
}

func about(w http.ResponseWriter, r *http.Request) {
	m := Message{"Welcome in my first ever API server"}
	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	w.Write(b)
}
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/about/", about)
	http.ListenAndServe(":8081", nil)

}
