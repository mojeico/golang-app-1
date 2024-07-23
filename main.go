package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/request", requestToApp2)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	_, err := fmt.Fprint(w, "Hello, Golang-app-1 ! trugger udpater 1 12")
	if err != nil {
		w.WriteHeader(http.StatusOK)
	}
}

func requestToApp2(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/request" {
		http.NotFound(w, r)
		return
	}

	url := os.Getenv("APP2_HOST")

	if url == "" {
		url = "http://localhost:8081"
	}

	res, err := http.Get(url + "/app2")

	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	if err != nil {
		w.WriteHeader(http.StatusOK)
	}
}
