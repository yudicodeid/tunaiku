package main

type StockCloseIDb interface {
	Add(ent StockCloseEnt) error
	List() []StockCloseEnt
}

