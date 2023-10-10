package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/send_data_to_service_a", func(w http.ResponseWriter, r *http.Request) {
		// You can customize the data to be sent to service A here.
		data := "Hello from service B"
		fmt.Fprintf(w, "Data to service A: %s", data)
	})

	http.ListenAndServe(":8081", nil)
}
