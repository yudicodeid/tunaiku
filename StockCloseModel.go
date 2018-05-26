package main

import "time"

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

}

type StockCloseModelList struct {
	ResponseModel
	Models []StockCloseModel
}

var stockCloseDb = StockCloseDb{}

func (model StockCloseModel) modelToEnt() (StockCloseEnt, error){

	ent := StockCloseEnt{}
	ent.ID = model.ID
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
	model.Action = ent.Action

	return nil

}

func (model StockCloseModel) Add() (error){

	ent, err := model.modelToEnt()
	if err!= nil {
		return err
	}

	err = stockCloseDb.Add(ent)
	if err!= nil {
		return err
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
