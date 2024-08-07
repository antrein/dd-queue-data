package rest

import (
	"antrein/dd-queue-data/application/common/repository"
	"antrein/dd-queue-data/model/config"
	"compress/gzip"
	"fmt"
	"net/http"
	"strings"
)

func setupCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE, PATCH")
	w.Header().Set("Agccess-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
}

type gzipResponseWriter struct {
	http.ResponseWriter
	Writer *gzip.Writer
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func compressHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}
		w.Header().Set("Content-Encoding", "gzip")
		gz, _ := gzip.NewWriterLevel(w, gzip.BestSpeed)
		defer gz.Close()

		gzw := gzipResponseWriter{ResponseWriter: w, Writer: gz}
		next.ServeHTTP(gzw, r)
	})
}

func ApplicationDelegate(cfg *config.Config, repo *repository.CommonRepository) (http.Handler, error) {
	mux := http.NewServeMux()

	mux.HandleFunc("/dd/queue/data", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Makan nasi pagi-pagi, ngapain kamu disini?")
	})
	mux.HandleFunc("/dd/queue/data/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "pong!")
	})

	handlerWithMiddleware := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		setupCORS(w)

		if r.Method == "OPTIONS" {
			return
		}

		compressHandler(mux).ServeHTTP(w, r)
	})

	return handlerWithMiddleware, nil
}
