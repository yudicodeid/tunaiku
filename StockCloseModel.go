package main

import "time"

type StockCloseModel struct {

	ResponseModel

	ID string
	StockDate time.Time
	Open int
	High int
	Low int
	Close int
	VolumeTrade int
	Action string

}

var stockCloseDb StockCloseDb = StockCloseDb{}

func (model StockCloseModel) modelToEnt() (StockCloseEnt, error){

	ent := StockCloseEnt{}
	ent.ID = model.ID
	ent.StockDate = model.StockDate
	ent.Open = model.Open
	ent.High = model.High
	ent.Low = model.Low
	ent.Close = model.Close
	ent.Action = model.Action

	return ent, _

}

func (model StockCloseModel) entToModel(ent StockCloseEnt) error {

	model.ID = ent.ID
	model.StockDate = ent.StockDate
	model.Open = ent.Open
	model.High = ent.High
	model.Low = ent.Low
	model.Close = ent.Close
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

func (model StockCloseModel) List() ([]StockCloseModel)  {

	var models []StockCloseModel
	entities := stockCloseDb.List()

	for _, ent := range entities {

		model := StockCloseModel{}
		model.entToModel(ent)

		models = append(models, model)
	}

	return models
}
