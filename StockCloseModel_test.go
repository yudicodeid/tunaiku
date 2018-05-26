package main

import (
	"testing"
	"time"
	"strconv"
	"fmt"
)


func TestValidate(t *testing.T) {

	model := StockCloseModel{}
	model.TruncateData()

	model.StockDate = time.Now()
	model.Open = 110
	model.High = 111
	model.Low = 99
	model.Close = 120
	model.VolumeTrade = 50

	_, err := model.Add()

	if err!=nil {
		t.Log("TestValidate :" + err.Error() + ": Success")
	} else {
		t.Error("Fail to validate Close must be between High and Low")
	}

}


func TestInsert(t *testing.T) {

	model := StockCloseModel{}
	model.TruncateData()

	model.StockDate = time.Now()
	model.Open = 110
	model.High = 111
	model.Low = 99
	model.Close = 100
	model.VolumeTrade = 50

	id, err := model.Add()

	if id != "" {

		model.ID = id
		_, err := model.Find()
		if err != nil {
			t.Error(err.Error())
		} else {
			t.Log("TestInsert success")
		}

	} else {
		t.Error(err.Error())
	}

}


func TestMaxProfit(t *testing.T) {

	model := StockCloseModel{}
	model.TruncateData()

	model.StockDate = time.Now()
	fmt.Println(model.StockDate.Format("02/01/2006"))

	model.Open = 110
	model.High = 111
	model.Low = 99
	model.Close = 100
	model.VolumeTrade = 50
	id, err := model.Add()
	if id == "" {
		t.Error(err.Error())
	}


	model1 := StockCloseModel{}
	model1.StockDate = time.Now().AddDate(0,0,1)
	fmt.Println(model1.StockDate.Format("02/01/2006"))

	model1.Open = 150
	model1.High = 180
	model1.Low = 145
	model1.Close = 180
	model1.VolumeTrade = 60
	id, err = model1.Add()

	if id == "" {
		t.Error(err.Error())
	}


	model2 := StockCloseModel{}
	model2.StockDate  = time.Now().AddDate(0,0,2)
	fmt.Println(model2.StockDate.Format("02/01/2006"))

	model2.Open = 230
	model2.High = 265
	model2.Low = 223
	model2.Close = 260
	model2.VolumeTrade = 30
	id, err = model2.Add()

	if id == "" {
		t.Error(err.Error())
	}

	stocks := model2.List()
	if len(stocks.Models) != 3 {
		t.Error("Invalid Stocks.Length=0")
	}


	p := stocks.Models[2]
	if p.Action == "S" && p.Max == true {
		t.Log("Calc Profit Success." + fmt.Sprint(p))

	} else {
		t.Log("Calc Profit Error. Stocks[2].Action=" + p.Action + ",Stocks[2].Max=" + strconv.FormatBool(p.Max))
	}


}
