package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/sending_to_serviceA", func(w http.ResponseWriter, r *http.Request) {
		// You can customize the data to be sent to service A here.
		data := "Hello from service B"
		fmt.Fprintf(w, "Info TO service A: %s", data)
	})

	http.ListenAndServe(":8081", nil)
}
