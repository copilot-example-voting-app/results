// Package server provides the HTTP server to render the "results" web page.
package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"math"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
)

// Server is the Results server.
type Server struct {
	Router *mux.Router
}

// page holds the data needed to render the "results" web page.
type page struct {
	Winner string
	Percentages map[string]int
}

// resultCount is a pair of a vote result and the sum of votes for the result.
type resultCount struct {
	Result string `json:"result"`
	Count  int    `json:"count"`
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.HandleFunc("/_healthcheck", s.handleHealthCheck())
	s.Router.HandleFunc("/results", s.handleView()).Methods(http.MethodGet)

	s.Router.ServeHTTP(w, r)
}

func (s *Server) handleHealthCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}

func (s *Server) handleView() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resultCounts, err := getResults()
		if err != nil {
			http.Error(w, "get results", http.StatusInternalServerError)
			return
		}
		renderTemplate(w, "index", page{
			Winner: getWinner(resultCounts),
			Percentages: getPercentages(resultCounts),
		})
	}
}


func getResults() ([]resultCount, error) {
	endpoint := fmt.Sprintf("http://api.%s:8080/results", os.Getenv("COPILOT_SERVICE_DISCOVERY_ENDPOINT"))
	resp, err := http.Get(endpoint)
	if err != nil {
		log.Printf("WARN: server: coudln't get results: %v\n", err)
		return nil, fmt.Errorf("server: get results: %v\n", err)
	}
	defer resp.Body.Close()
	data := struct {
		Results []resultCount `json:"results"`
	}{}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&data); err != nil {
		log.Printf("ERROR: server: decode results data: %v\n", err)
		return nil, fmt.Errorf("server: decode results: %v",err)
	}
	log.Printf("INFO: server: received results of votes")
	return data.Results, nil
}

func getWinner(rcs []resultCount) string {
	catCount, dogCount  := 0, 0
	for _, rc := range rcs {
		switch rc.Result {
		case "dog":
			dogCount = rc.Count
		case "cat":
			catCount = rc.Count
		}
	}
	if catCount > dogCount {
		return "cat"
	}
	return "dog"
}

func getPercentages(rcs []resultCount) map[string]int {
	sum := 0
	for _, rc := range rcs {
		if rc.Result == "cat" || rc.Result == "dog" {
			sum += rc.Count
		}
	}

	m := make(map[string]int)
	for _, rc := range rcs {
		m[rc.Result] = int(math.Round(float64(rc.Count)/(float64(sum))))
	}
	return m
}

func renderTemplate(w http.ResponseWriter, tmpl string, p page) {
	t, err := template.ParseFiles(filepath.Join("templates", tmpl + ".html"))
	if err != nil {
		log.Fatalf("parse file: %v\n", err)
	}
	t.Execute(w, p)
}