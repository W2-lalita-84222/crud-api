package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println(mux.Vars(r)["id"])
	w.Write([]byte("Hello world !!!!"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/home/{id}", HomeHandler)

	http.ListenAndServe(":8888", r)

}
