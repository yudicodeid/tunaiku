package main

import (
	"time"
	"github.com/google/uuid"
	"errors"
)

type StockCloseModel struct {

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

type StockCloseModelList struct {
	ResponseModel
	Models []StockCloseModel
}


func (model StockCloseModel) modelToEnt() (StockCloseEnt, error){

	ent := StockCloseEnt{}
	ent.ID = uuid.New().String()
	ent.StockDate = model.StockDate
	ent.Open = model.Open
	ent.High = model.High
	ent.Low = model.Low
	ent.Close = model.Close
	ent.VolumeTrade = model.VolumeTrade

	return ent, nil

}

func (model *StockCloseModel) entToModel(ent StockCloseEnt) error {

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

	return nil

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


func (model StockCloseModel) Add() (error){

	err := model.validateInsert()
	if err != nil {
		return err
	}

	stock,  found := stockCloseDb.FindByDate(model.StockDate.Year(),
		model.StockDate.Month(),
		model.StockDate.Day())

	ent, err := model.modelToEnt()
	if err!= nil {
		return err
	}

	updated := false

	if found == true {
		ent.ID = stock.ID
		updated = stockCloseDb.Update(ent)

		if updated == true {

			profitModel := ProfitModel{}
			profitModel.AnalyzeProfits()

		}

	} else {

		err = stockCloseDb.Add(ent)
		if err!= nil {
			return err
		} else {
			profitModel := ProfitModel{}
			profitModel.AnalyzeProfits()
		}

	}

	return nil
}


func (model StockCloseModel) List() (StockCloseModelList)  {

	modelList := StockCloseModelList{}
	modelList.Success("")

	entities := stockCloseDb.List()

	for _, ent := range entities {

		model := StockCloseModel{}
		model.entToModel(ent)

		modelList.Models = append(modelList.Models, model)
	}

	return modelList
}


func (model StockCloseModel) Delete() (bool, error) {

	if model.ID == "" {
		return false, errors.New("Invalid ID")
	}

	b := stockCloseDb.Delete(model.ID)

	if b == true {

		profitModel := ProfitModel{}
		profitModel.AnalyzeProfits()

		return true, nil
	} else {
		return false, errors.New("failed to delete data")
	}


}
