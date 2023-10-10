package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/get_data_from_service_b", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("http://localhost:8081/send_data_to_service_a")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		fmt.Fprintf(w, "Data from service B: %s", data)
	})

	http.ListenAndServe(":8080", nil)
}
