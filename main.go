package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path !="/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method !="GET" {
		http.Error(w, "Methos is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello")
}

func formHander (w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "Post request Successfull")
	name := r.FormValue("name")
	password := r.FormValue("password")

	fmt.Fprintf(w, "Name =%s\n", name)
	fmt.Fprintf(w, "Password =%s\n", password)
}

func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/hello", fileServer)

	http.HandleFunc("/form", formHander)
	http.HandleFunc("./", helloHandler)


	fmt.Printf(("Starting server at port 3000\n"))

	if err:= http.ListenAndServe(":3000", nil); err !=nil {
		log.Fatal(err)
}

}