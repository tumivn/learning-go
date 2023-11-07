package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloWorld)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		println(err)
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`Hello, World!`))
}
