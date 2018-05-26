package main

import (
	"time"
	"strings"
	"strconv"
)

type StockCloseEnt struct {

	ID string
	StockDate time.Time
	Open int
	High int
	Low int
	Close int
	VolumeTrade int
	Action string
	Max bool

}


const (
	CountFields = 9
)

func (t *StockCloseEnt) ParseRow(row string) (error)  {

	cols := strings.Split(row, ",")

	if len(cols) == CountFields {

		t.ID = cols[0]
		t.StockDate, _ = time.Parse( "20060102", cols[1])
		t.Open, _ = strconv.Atoi(cols[2])
		t.High, _ = strconv.Atoi(cols[3])
		t.Low, _ = strconv.Atoi(cols[4])
		t.Close, _ = strconv.Atoi(cols[5])
		t.VolumeTrade, _ = strconv.Atoi(cols[6])
		t.Action = cols[7]
		t.Max, _ = strconv.ParseBool(cols[8])

		return nil

	} else {

		return nil
		//return errors.New("Invalid Fields Count=" + strconv.Itoa(len(cols)))
	}

}

func(t StockCloseEnt) ToString() string {

	delimiter := ","

	s := []string{t.ID, t.StockDate.Format("20060102"),
		strconv.Itoa(t.Open),
		strconv.Itoa(t.High),
		strconv.Itoa(t.Low),
		strconv.Itoa(t.Close),
		strconv.Itoa(t.VolumeTrade),
		t.Action,
		strconv.FormatBool(t.Max),
	}

	str := strings.Join(s, delimiter)

	return str

}