package main

import (
	"time"
	"fmt"
	"strconv"
	"sort"
)

type ProfitModel struct {

	BuyID      string
	BuyDate    time.Time
	SellID     string
	SellDate   time.Time
	BuyPrice   int
	BuyVolume  int
	SellPrice  int
	SellVolume int
	Value      int

}

func (p ProfitModel) Log() {

	fmt.Print("BuyDate:" + p.BuyDate.Format("2006-01-02") + ",")
	fmt.Print("BuyPrice:" + strconv.Itoa(p.BuyPrice) + ",")
	fmt.Print("BuyVolume:" + strconv.Itoa(p.BuyVolume) + ",")
	fmt.Println("BuyAmount:" + strconv.Itoa(p.BuyVolume*p.BuyPrice) + "")


	fmt.Print("SellDate:" + p.SellDate.Format("2006-01-02") + ",")
	fmt.Print("SellPrice:" + strconv.Itoa(p.SellPrice) + ",")
	fmt.Print("SellVolume:" + strconv.Itoa(p.SellVolume) + ",")
	fmt.Println("SellAmount:" + strconv.Itoa(p.SellVolume*p.SellPrice) + "")

	fmt.Println("Profit: " + strconv.Itoa(p.Value))
	fmt.Println("")

}


type Profits []ProfitModel
func (a Profits) Len() int  { return len(a) }
func (a Profits) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Profits) Less(i, j int) bool {
	return a[i].Value > a[j].Value
}

func (p ProfitModel) findLowest(rows []StockCloseEnt) int {

	min := 0
	i:=0

	for a:=0; a< len(rows); a++ {

		buyPrice := rows[a].Close
		buyVol := rows[a].VolumeTrade
		buyAmount := buyPrice*buyVol

		if a == 0 {
			min = buyAmount
			i=a
		} else {

			if buyAmount < min {
				min = buyAmount
				i= a
			}

		}

	}

	return i
}



func (p ProfitModel) calcCombination(rows []StockCloseEnt) Profits {

	var profits []ProfitModel

	a := 0

	for a < (len(rows) - 1) {

		b:= a+1

		for b < len(rows) {

			buyPrice := rows[a].Close
			buyVol := rows[a].VolumeTrade

			sellPrice := rows[b].Close
			sellVol := rows[b].VolumeTrade

			buyAmount := buyPrice*buyVol
			sellAmount := sellPrice*sellVol

			profit := sellAmount - buyAmount

			if sellVol <= buyVol  && profit > 0 {

				var p = ProfitModel {

					BuyID:     rows[a].ID,
					BuyDate: rows[a].StockDate,
					BuyPrice:  rows[a].Close,
					BuyVolume: rows[a].VolumeTrade,

					SellID:    rows[b].ID,
					SellDate : rows[b].StockDate,
					SellPrice: rows[b].Close,
					SellVolume: rows[b].VolumeTrade,

					Value:     profit,

				}

				profits = append(profits, p)

			}
			b++
		}
		a++
	}

	if len(profits) == 0 {

		lowestIndex := p.findLowest(rows)

		var p = ProfitModel {
			BuyID:     rows[lowestIndex].ID,
			BuyPrice:  rows[lowestIndex].Close,
			BuyVolume: rows[lowestIndex].VolumeTrade,
			BuyDate:  rows[lowestIndex].StockDate,
		}

		profits = append(profits, p)

	}

	return profits

}


func (p ProfitModel) getPossibleProfits() ([]ProfitModel, bool) {

	var profits []ProfitModel
	lenStock := len(stockCloseDb.Data)

	if lenStock > 0 {

		stockCloseDb.SortByDate()

		profits = p.calcCombination(stockCloseDb.Data)

		if len(profits)> 0 {

			sort.Sort(Profits(profits))

			for _,p := range profits {
				p.Log()
			}

			return profits, true

		} else {
			return profits , false
		}

	} else {
		return profits , false
	}

}



func (p ProfitModel) sliceInvalidProfits(profits []ProfitModel) []ProfitModel {

	var cleanProfits []ProfitModel

	if len(profits) >0 {

		maxProfit := profits[0]
		cleanProfits = append(cleanProfits, maxProfit)

		for a:=1; a<len(profits); a++ {

			profit := profits[a]

			if profit.SellDate.Sub(maxProfit.BuyDate) > 0 {

				cleanProfits = append(cleanProfits, profit)

			}

		}
	}

	return cleanProfits

}


func (p ProfitModel) AnalyzeProfits() {

	stockCloseDb.ResetAction()
	stockCloseDb.ResetMax()

	profits, b := p.getPossibleProfits()
	if b == true {

		cleanProfits := p.sliceInvalidProfits(profits)

		maxProfit := cleanProfits[0]
		stockCloseDb.UpdateAction(maxProfit.BuyID, "B")

		if maxProfit.SellID != "" {
			stockCloseDb.UpdateAction(maxProfit.SellID, "S")
			stockCloseDb.UpdateMax(maxProfit.SellID)
		}

		if len(cleanProfits) >0 {

			for a:=1; a<len(cleanProfits); a++ {

				profit := cleanProfits[a]
				stockCloseDb.UpdateAction(profit.SellID, "S")
			}

		}

	}

}

