package main

import (
	"net/http"
	"strings"
	"encoding/json"
	"html/template"
)

type StockCloseService struct {
	model StockCloseModel
}

func CreateStockCloseService() StockCloseService {
	svc := StockCloseService{}
	svc.model = StockCloseModel{}
	return svc
}


func (svc StockCloseService) Post(w http.ResponseWriter, r *http.Request) {

	err := svc.model.Add()
	response :=ResponseModel{}
	if err!= nil {
		response.Error(err)
	} else {
	 	response.Success("")
	}
	json.NewEncoder(w).Encode(response)

}

func (svc StockCloseService) Index(w http.ResponseWriter, r *http.Request) {

	method := strings.ToUpper(r.Method)
	if method == "GET" {

		t, _ := template.ParseFiles("/static")


	} else if method == "POST" {
		svc.Post(w, r)
	}
}


func (svc StockCloseService) List(w http.ResponseWriter, r *http.Request) {

	method := strings.ToUpper(r.Method)
	if method == "GET" {
		models := svc.model.List()
		json.NewEncoder(w).Encode(models)
	}

}
