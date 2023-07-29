package main

import (
	"fmt"
	"log"
	"net/http"
)

type FileController struct {
}

func (c FileController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s-%s", r.Method, r.URL)
	w.Header().Set("Content-Type", "application/vnd.ms-excel")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("wu"))
}

func main() {

	http.HandleFunc("/fuck", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("caonim"))
	})
	http.Handle("/word", FileController{})
	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}
}
