// Server stores main dependencies to use with handlers.
package app

import (
	"log"
	"net/http"
)

func HandleIndex(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Length", "12")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	if _, err := w.Write([]byte("Hello World!")); err != nil {
		log.Fatalf("Failed to execute `HandleIndex`: %v", err)
	}
}
