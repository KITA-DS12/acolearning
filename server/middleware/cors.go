package middleware

import "net/http"

func CorsMiddleware(w http.ResponseWriter) {
	protocol := "http://"
	host := "localhost:5173"
	headers := w.Header()
	headers.Set("Access-Control-Allow-Origin", protocol+host)
	headers.Set("Access-Control-Allow-Methods", "OPTIONS")
}

func OptionMiddleware(w http.ResponseWriter) {
	headers := w.Header()
	headers.Set("Access-Control-Allow-Methods", "OPTIONS, POST")
	headers.Set("Access-Control-Allow-Headers", "Content-Type")
	headers.Set("Access-Control-Max-Age", "86400")
	w.WriteHeader(http.StatusOK)
}
