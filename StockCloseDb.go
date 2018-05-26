package main

import (
	"sort"
	"time"
)

type StockCloseDb struct {
	File StockCloseFile
	Data []StockCloseEnt
}


func CreateStockCloseDb() StockCloseDb {

	db := StockCloseDb{}
	file, err := CreateStockCloseFile()

	if err !=  nil {
		panic(err)

	} else {
		db.File = file
	}

	return db
}


func (db *StockCloseDb) Add(ent StockCloseEnt) error {

	db.File.Append(ent.ToString())

	db.Data = append(db.Data, ent)

	return nil
}


func (db *StockCloseDb) parseRowsFile() {

	db.Data = []StockCloseEnt{}
	rows := db.File.Rows

	if len(rows) >0 {

		for _, row := range rows {

			ent := StockCloseEnt{}
			e := ent.ParseRow(row)

			if e!= nil {
				panic(e)
			}

			db.Data = append(db.Data, ent)
		}
	}

}



func (db *StockCloseDb) List() []StockCloseEnt {

	db.File.ReadAllStockFile()

	db.parseRowsFile()

	return db.Data
}


func (db *StockCloseDb) SortByDate() {
	sort.Sort(StockCloseByDate(db.Data))
}


func (db *StockCloseDb) ResetAction() {
	for i, _ := range db.Data {
		db.Data[i].Action = ""
	}
}

func (db *StockCloseDb) ResetMax() {
	for i, _ := range db.Data {
		db.Data[i].Max = false
	}
}


func (db *StockCloseDb) RowsToString() []string {

	var rows []string
	for _, row :=range db.Data {
		rows = append(rows, row.ToString())
	}
	return rows
}



func (db *StockCloseDb) UpdateAction(id string, action string) {

	for i, v := range db.Data {
		if v.ID == id {
			db.Data[i].Action = action
			break
		}
	}

	db.File.Update(db.RowsToString())
}


func (db *StockCloseDb) UpdateMax(id string) {

	for i, v := range db.Data {
		if v.ID == id {
			db.Data[i].Max = true
			break
		}
	}

	db.File.Update(db.RowsToString())

}


func (db StockCloseDb) FindByDate(year int, month time.Month, day int) (StockCloseEnt, bool) {

	var row StockCloseEnt
	var found bool = false

	for _, stock := range db.Data {

		stockDate := stock.StockDate

		if stockDate.Year() == year &&
			stockDate.Month() == month &&
			stockDate.Day() == day {

			row = stock
			found = true
			break
		}
	}

	return row, found

}


func (db *StockCloseDb) Update(ent StockCloseEnt) bool {

	found := false

	for i, v := range db.Data {
		if v.ID == ent.ID {

			db.Data[i].Open = ent.Open
			db.Data[i].Close = ent.Close
			db.Data[i].High = ent.High
			db.Data[i].Low = ent.Low
			db.Data[i].VolumeTrade = ent.VolumeTrade

			found = true
			break

		}
	}

	if found == true {
		db.File.Update(db.RowsToString())
	}

	return found
}



func (db *StockCloseDb) Delete(id string) bool {

	var stacks []StockCloseEnt
	found := false

	for _, v := range db.Data {
		if v.ID == id {
			found = true
		} else {
			stacks = append(stacks, v)
		}
	}

	if found {
		db.Data = stacks
		rows := db.RowsToString()
		db.File.Update(rows)
	}

	return found

}

func (db *StockCloseDb) TruncateData() {
	db.File.Truncate()
	db.Data = []StockCloseEnt{}
}


func (db StockCloseDb) Find(id string) (bool, StockCloseEnt) {

	var row StockCloseEnt
	var found = false
	for _, v := range db.Data {
		if v.ID == id {
			row = v
			found = true
			break
		}
	}
	return found, row

}



type StockCloseByDate [] StockCloseEnt
func (a StockCloseByDate) Len() int { return len(a) }
func (a StockCloseByDate) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a StockCloseByDate) Less(i, j int) bool {
	diff := a[i].StockDate.Sub(a[j].StockDate)
	if diff < 0 {
		return true
	} else {
		return false
	}
}

