package main

import (
	"net/http"
	"strings"
)

type StockCloseService struct {

}

func CreateStockCloseService() StockCloseService {
	svc := StockCloseService{}
	return svc
}

func (svc StockCloseService) List(w http.ResponseWriter, r *http.Request) {

}

func (svc StockCloseService) Post(w http.ResponseWriter, r *http.Request) {

}

func (svc StockCloseService) Index(w http.ResponseWriter, r *http.Request) {

	method := strings.ToUpper(r.Method)
	if method == "GET" {

	} else if method == "POST" {

	}

}
