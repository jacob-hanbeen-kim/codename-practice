package app

import (
	"net/http"
)

func Start() {
	//define routes
	http.HandleFunc("/greet", greet)

	//starting server
	http.ListenAndServe("localhost:8000", nil)
}
