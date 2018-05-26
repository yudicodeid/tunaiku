package main

import (
	"net/http"
	"strings"
	"encoding/json"
	"html/template"
	"time"
	"strconv"
)


type StockCloseService struct {
	model StockCloseModel
}

func CreateStockCloseService(db *StockCloseDb) StockCloseService {

	svc := StockCloseService{}
	svc.model = CreateStockCloseModel(db)
	return svc
}




func (svc StockCloseService) Post(w http.ResponseWriter, r *http.Request) {

	stockDate := r.PostFormValue("StockDate")
	open := r.PostFormValue("Open")
	high := r.PostFormValue("High")
	low := r.PostFormValue("Low")
	close := r.PostFormValue("Close")
	vol := r.PostFormValue("VolumeTrade")

	svc.model.StockDate, _ = time.Parse("2006-01-02",stockDate)
	svc.model.Open, _ = strconv.Atoi(open)
	svc.model.High, _ = strconv.Atoi(high)
	svc.model.Low, _ = strconv.Atoi(low)
	svc.model.Close, _ = strconv.Atoi(close)
	svc.model.VolumeTrade, _ = strconv.Atoi(vol)

	_ , err := svc.model.Add()
	response :=ResponseModel{}
	if err!= nil {
		response.Error(err)
	} else {
	 	response.Success("")
	}
	json.NewEncoder(w).Encode(response)

}




func (svc StockCloseService) Delete(w http.ResponseWriter, r *http.Request) {

	id := r.PostFormValue("ID")

	svc.model.ID = id
	_, err := svc.model.Delete()
	response := ResponseModel{}

	if err != nil {
		response.Error(err)
	} else {
		response.Success("")
	}
	json.NewEncoder(w).Encode(response)

}



func (svc StockCloseService) Index(w http.ResponseWriter, r *http.Request) {

	method := strings.ToUpper(r.Method)
	if method == "GET" {

		t, _ := template.ParseFiles("static/Index.html")
		t.Execute(w,nil)

	} else if method == "POST" {
		svc.Post(w, r)

	} else if method == "PATCH" {
		svc.Delete(w, r)
	}
}


func (svc StockCloseService) List(w http.ResponseWriter, r *http.Request) {

	method := strings.ToUpper(r.Method)
	if method == "GET" {
		modelList := svc.model.List()
		json.NewEncoder(w).Encode(modelList)
	}

}
