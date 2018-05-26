package main

import (
	"time"
	"github.com/google/uuid"
	"errors"
)


type StockCloseModel struct {

	BaseModel

	ID string
	StockDate time.Time
	StockDateString string
	Open int
	High int
	Low int
	Close int
	VolumeTrade int
	Action string
	Max bool

}

func CreateStockCloseModel(db *StockCloseDb) StockCloseModel{
	model := StockCloseModel{}
	model.Db = db
	return model
}

type StockCloseModelList struct {
	ResponseModel
	Models []StockCloseModel
}


func (model StockCloseModel) modelToEnt() (StockCloseEnt, error){

	ent := StockCloseEnt{}
	ent.StockDate = model.StockDate
	ent.Open = model.Open
	ent.High = model.High
	ent.Low = model.Low
	ent.Close = model.Close
	ent.VolumeTrade = model.VolumeTrade

	return ent, nil

}

func (model *StockCloseModel) entToModel(ent StockCloseEnt) {

	model.ID = ent.ID
	model.StockDate = ent.StockDate
	model.StockDateString = ent.StockDate.Format("02/01/2006")
	model.Open = ent.Open
	model.High = ent.High
	model.Low = ent.Low
	model.Close = ent.Close
	model.VolumeTrade = ent.VolumeTrade

	if ent.Action == "B" {
		model.Action = "Buy"
	} else if ent.Action == "S" {
		model.Action = "Sell"
	}

	model.Max = ent.Max


}

func (model StockCloseModel) validateInsert()  error {

	currentTime := time.Now()
	elapsed := currentTime.Sub(model.StockDate)

	h := int(elapsed.Hours()/24)

	if h > 0 {
		return errors.New("stock Date must be greater than or equals to Today's date")
	}

	if model.Open ==0 {
		return errors.New("Invalid BuyPrice value")
	}

	if model.Close == 0 {
		return errors.New("Invalid SellPrice value")
	}

	if model.High ==0 {
		return errors.New("Invalid High value")
	}

	if model.Low ==0 {
		return errors.New("Invalid Low value")
	}

	if model.High <= model.Low {
		return errors.New("High must be greater than Low")
	}

	if model.Close > model.High || model.Close < model.Low {
		return errors.New("Close must be between High and Low")
	}

	return nil

}


func (model *StockCloseModel) Add() (string, error){

	err := model.validateInsert()
	if err != nil {
		return "", err
	}

	stock,  found := model.Db.FindByDate(model.StockDate.Year(),
		model.StockDate.Month(),
		model.StockDate.Day())

	ent, err := model.modelToEnt()
	if err!= nil {
		return "", err
	}

	updated := false

	if found == true {

		ent.ID = stock.ID
		updated = model.Db.Update(ent)

		if updated == true {

			profitModel := CreateProfitModel(model.Db)
			profitModel.AnalyzeProfits()

		}

		return ent.ID, nil

	} else {

		ent.ID = uuid.New().String()
		err = model.Db.Add(ent)
		if err!= nil {
			return "", err
		} else {

			profitModel := CreateProfitModel(model.Db)
			profitModel.AnalyzeProfits()

			return ent.ID, nil
		}

	}

	return "", nil
}


func (model StockCloseModel) List() (StockCloseModelList)  {

	modelList := StockCloseModelList{}
	modelList.Success("")

	entities := model.Db.List()

	for _, ent := range entities {

		model := StockCloseModel{}
		model.entToModel(ent)

		modelList.Models = append(modelList.Models, model)
	}

	return modelList
}


func (model *StockCloseModel) Delete() (bool, error) {

	if model.ID == "" {
		return false, errors.New("Invalid ID")
	}

	b := model.Db.Delete(model.ID)

	if b == true {

		profitModel := CreateProfitModel(model.Db)
		profitModel.AnalyzeProfits()

		return true, nil
	} else {
		return false, errors.New("failed to delete data")
	}


}


func (model *StockCloseModel) TruncateData() {
	model.Db.TruncateData()
}



func (model StockCloseModel) Find() (StockCloseModel, error) {

	found, ent := model.Db.Find(model.ID)
	var e = StockCloseModel{}

	if found == false {
		return  e, errors.New("Cannot find Stock wtih ID:" + model.ID)

	} else {
		model.entToModel(ent)
		return e, nil
	}

}