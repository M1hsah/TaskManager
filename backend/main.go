package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// healthHandler responds to GET /health with a small JSON payload
// so we (or a deploy platform) can check the server is alive.
func healthHandler(w http.ResponseWriter, r *http.Request) {
	// TODO 1: Tell the client it's getting JSON back.
	// Hint: w has a Header() method that returns something you can Set() on.
	// The header name is "Content-Type", the value should be "application/json".

	// TODO 2: Build a response body and write it as JSON to w.
	// Hint: json.NewEncoder(...) wraps a writer and gives you an Encode(v) method.
	// You can encode any Go value that marshals to JSON — a map[string]string
	// like {"status": "ok"} is a fine starting point.
}

func main() {
	// TODO 3: Register healthHandler for the "/health" path.
	// Hint: look at http.HandleFunc(pattern, handler).

	// TODO 4: Start the server listening on a port, e.g. ":8080".
	// Hint: http.ListenAndServe(addr, handler) — what should the handler
	// argument be if you registered routes with http.HandleFunc above?
	// It returns an error — this is a good first look at Go's habit of
	// returning errors instead of throwing exceptions. log.Fatal(err) is
	// a common way to bail out if the server fails to start.
}
