package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to Too Tasty of a Trade!")

	r := mux.NewRouter()

	var JsonMiddleware mux.MiddlewareFunc = func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")

			next.ServeHTTP(w, r)
		})
	}

	HealthCheck := func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{}{
			"message": "Welcome to the dark side",
		}

		w.WriteHeader(http.StatusOK)

		jsonData, _ := json.Marshal(data)

		fmt.Fprint(w, string(jsonData))
	}

	r.Use(JsonMiddleware)

	r.HandleFunc("/health-check", HealthCheck).Methods("GET")

	http.ListenAndServe(":3333", r)
}
