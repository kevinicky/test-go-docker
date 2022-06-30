package main

import "net/http"

func main() {
	mux := http.NewSer1veM1ux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world!"))
	})

	http.ListenAndServe(":8080", mux)
}
