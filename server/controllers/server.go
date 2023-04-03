package controllers

import "net/http"

func Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/aco", runAco)

	return mux
}
