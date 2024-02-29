package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprint(w, "Hello, World!")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
	})

	_ = http.ListenAndServe(":8080", nil)
}
