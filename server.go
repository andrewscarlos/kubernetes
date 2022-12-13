package main

import "net/http"

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":3000", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world !"))
}

// sudo docker tag andrews/hello-go 18257047/hello-go:1.0
//                  nome da img    login dockerhub    repositorio/ versão
