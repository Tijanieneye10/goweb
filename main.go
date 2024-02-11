package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fprintf, err := fmt.Fprintf(w, "Hello world")
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(fprintf)
	})

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		return
	}

}
