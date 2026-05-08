package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	"github.com/SaurabPoudel/smart-parking/types"
)

func main() {
	listenAddr := flag.String("listenAddr", ":3000", "the listen address of the HTTP server")
	flag.Parse()
	var (
		store = NewMemoryStore()
		svc   = NewInvoiceAggregator(store)
	)
	svc = NewLogMiddleware(svc)
	makeHTTPTransport(*listenAddr, svc)
	fmt.Println("working fine")
}

func makeHTTPTransport(listenAddr string, svc Aggregator) {
	fmt.Println("starting HTTP server on ", listenAddr)
	http.HandleFunc("/aggregate", handleAggregate(svc))
	http.ListenAndServe(listenAddr, nil)
}

func handleAggregate(svc Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var p types.Parking
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			writeJSON(w, http.StatusBadGateway, map[string]string{"error": err.Error()})
			return
		}
		if err := svc.AggregateParking(p); err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}

	}

}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}
