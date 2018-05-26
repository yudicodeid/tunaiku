package main

import (
	"net/http"
	"fmt"
	"log"
)

func hello(w http.ResponseWriter, _ *http.Request){
	fmt.Fprintln(w,"Hello Tunaiku")
}

var stockCloseService StockCloseService = CreateStockCloseService()

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/", stockCloseService.Index)
	http.HandleFunc("/list", stockCloseService.List)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Listening to port 5001. Access : http:/localhost:5001")
	err := http.ListenAndServe(":5001", nil)
	if err!= nil {
		log.Fatal(err)
	}

}
