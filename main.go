package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	// var b struct{
	// 	Id int
	// 	Name string
	// }
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	//t := rand.Intn(5000)
	//time.Sleep(time.Duration(t) * time.Millisecond)
	w.Write(data)

}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", handleRequest).Methods("POST")
	fmt.Println("Webserver started at 4000 port...")
	log.Fatal(http.ListenAndServe(":4000", r))
}
