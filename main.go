package main

import (
	"github.com/pion/webrtc/v3"
	"log"
)

func main() {
	http.HandleFunc("/signal", server.SignalHandler)

	http.Handle("/", http.FileServer(http.Dir("./html")))

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
