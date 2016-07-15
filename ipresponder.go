package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	println("Server starting...")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("<html><head><title>Get your IP!</title></head><body>Your IP is: %s</body></html>", strings.Split(r.RemoteAddr, ":")[0])))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
