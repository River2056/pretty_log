package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		temp, _ := template.ParseFiles("./index.html")

		err := temp.Execute(w, nil)
		if err != nil {
			panic(err)
		}
	})

	fmt.Println("running server at port 9876...")
	err := http.ListenAndServe(":9876", nil)
	if err != nil {
		panic(err)
	}
}
