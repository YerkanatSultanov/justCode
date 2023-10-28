package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/get_serviceB", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("http://localhost:8081/sending_to_serviceA")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()
		data, _ := ioutil.ReadAll(resp.Body)
		fmt.Fprintf(w, "Info from B: %s", data)
	})

	http.ListenAndServe(":8080", nil)
}
