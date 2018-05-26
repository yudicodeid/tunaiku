package main

import "time"

type StockCloseEnt struct {

	ID string
	StockDate time.Time
	Open int
	High int
	Low int
	Close int
	Action string

}
