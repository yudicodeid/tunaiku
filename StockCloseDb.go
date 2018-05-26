package main

import "sort"

type StockCloseDb struct {
	Data []StockCloseEnt
}

func (db *StockCloseDb) Add(ent StockCloseEnt) error {
	db.Data = append(db.Data, ent)
	return nil
}

func (db *StockCloseDb) List() []StockCloseEnt {
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

func (db *StockCloseDb) UpdateAction(id string, action string) {

	for i, v := range db.Data {
		if v.ID == id {
			db.Data[i].Action = action
			break
		}
	}
}


func (db *StockCloseDb) UpdateMax(id string) {

	for i, v := range db.Data {
		if v.ID == id {
			db.Data[i].Max = true
			break
		}
	}

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

