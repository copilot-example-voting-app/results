// Package results starts the results server.
package results

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/copilot-example-voting-app/results/server"
	"github.com/gorilla/mux"
)

// Run starts the server.
func Run() error {
	addr := flag.String("addr", ":8080", "port to listen on")
	flag.Parse()

	s := http.Server{
		Addr: *addr,
		Handler: &server.Server{
			Router: mux.NewRouter(),
		},
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}
	log.Printf("INFO: vote: listen on port %s\n", *addr)
	return s.ListenAndServe()
}