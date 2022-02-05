package main

import (
	"io"
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {

	svc := Service{}
	// Hello world, the web server

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	convertAppleToSpotifyHandler := httptransport.NewServer(
		makeConvertAppleToSpotifyEndpoint(svc),
		decodeUppercaseRequest,
		encodeResponse,
	)

	convertSpotifyToAppleHandler := httptransport.NewServer(
		makeConvertSpotifyToAppleEndpoint(svc),
		decodeCountRequest,
		encodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)

	http.HandleFunc("/hello", helloHandler)
	log.Println("Listing for requests at http://localhost:8080/hello")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
