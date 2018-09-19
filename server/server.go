package server

import (
	"net/http"
	"fmt"
	"os"
)

func Start() {

	router := http.NewServeMux()
	router.HandleFunc("/", defaultHandler)
	router.HandleFunc("/go", goHandler)

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	http.ListenAndServe("localhost:"+port, router)

}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello, Web from Go!")
	fmt.Fprintln(w, "Hello, Web from Go!")
}

func goHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I <3 Go!")
	fmt.Fprintln(w, "I <3 Go!")
}
