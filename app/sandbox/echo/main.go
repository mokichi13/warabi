package main

import (
	"fmt"
	"html"
	"net/http"
)

func main() {
	// Echo Service.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Echo, %q\n", html.EscapeString(r.URL.Path))
		fmt.Fprintf(w, "User Agent, %s\n", html.EscapeString(r.UserAgent()))
	})

	http.ListenAndServe(":8080", nil)
}
