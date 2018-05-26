package main

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

