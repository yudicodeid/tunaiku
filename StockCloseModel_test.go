package main

import (
	"testing"
	"time"
	"strconv"
	"fmt"
	"webstock/entity"
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

	err := model.Add()

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

	ent := entity.StockCloseEnt{}
	ent.TruncateData()

	ent.StockDate = time.Now()
	fmt.Println(ent.StockDate.Format("02/01/2006"))

	ent.Open = 110
	ent.High = 111
	ent.Low = 99
	ent.Close = 100
	ent.VolumeTrade = 50
	id, err := ent.Insert()
	if id == "" {
		t.Error(err.Error())
	}


	ent1 := entity.StockCloseEnt{}
	ent1.StockDate = time.Now().AddDate(0,0,1)
	fmt.Println(ent1.StockDate.Format("02/01/2006"))

	ent1.Open = 150
	ent1.High = 180
	ent1.Low = 145
	ent1.Close = 180
	ent1.VolumeTrade = 60
	id, err = ent1.Insert()

	if id == "" {
		t.Error(err.Error())
	}


	ent2 := entity.StockCloseEnt{}
	ent2.StockDate  = time.Now().AddDate(0,0,2)
	fmt.Println(ent2.StockDate.Format("02/01/2006"))

	ent2.Open = 230
	ent2.High = 265
	ent2.Low = 223
	ent2.Close = 260
	ent2.VolumeTrade = 30
	id, err = ent2.Insert()

	if id == "" {
		t.Error(err.Error())
	}

	stocks := ent2.GetList()
	if len(stocks) != 3 {
		t.Error("Invalid Stocks.Length=0")
	}


	p := stocks[2]
	if p.Action == "S" && p.Max == true {
		t.Log("Calc Profit Success." + fmt.Sprint(p))

	} else {
		t.Log("Calc Profit Error. Stocks[2].Action=" + p.Action + ",Stocks[2].Max=" + strconv.FormatBool(p.Max))
	}


}
