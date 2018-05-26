package main

import (
	"net/http"
	"fmt"
	"log"
)

func hello(w http.ResponseWriter, _ *http.Request){
	fmt.Fprintln(w,"Hello Tunaiku")
}

func main() {

	http.HandleFunc("/hello", hello)

	log.Println("Listening to port 5001. Access : http:/localhost:5001")
	err := http.ListenAndServe(":5001", nil)
	if err!= nil {
		log.Fatal(err)
	}

}
