package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/ingrid/data"
)

func RoutesFinder(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		url := req.URL
		urlString := url.String()
		paramsStr := strings.Split(urlString, "?")[1]
		params := strings.Split(paramsStr, "&")
		var input [][]string
		for _, v := range params {
			value := strings.Split(v, "=")[1]
			dst := strings.Split(value, ",")
			input = append(input, dst)
		}
		data, err := data.Request(input)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(400)
			if err := json.NewEncoder(w).Encode(err); err != nil {
				panic(err)
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)

	} else {
		http.Error(w, "Bad Request method. Need method GET", 400)
	}

}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, it's route distance service")
	})
	http.HandleFunc("/routes", RoutesFinder)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
