package main

import (
	"fmt"
	"net/http"
)

// HelloServer func
func HelloServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":38080", nil)
}
