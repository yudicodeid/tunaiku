package main

import (
	"os"
	"errors"
	"bufio"
	"encoding/json"
)

type StockCloseFile struct {
	Rows []string
	DirName string
	Filename string
}

func checkDirectory(dirname string) {


	_, err1 := os.Stat(dirname)
	if err1 != nil {
		if os.IsNotExist(err1) {
			os.Mkdir(dirname, 0666)
		}
	}

}

func (f *StockCloseFile) ReadAllStockFile() error {

	osFile, err := os.OpenFile(f.getFullPath(), os.O_RDONLY, 0666)

	if err == nil {

		scanner := bufio.NewScanner(osFile)
		scanner.Split(bufio.ScanLines)

		f.Rows = []string{}
		for scanner.Scan() {
			text := scanner.Text()
			if len(text)>0 {
				f.Rows = append(f.Rows, scanner.Text())
			}
		}
		osFile.Close()
		return nil

	}else {
		return err
	}
}


func readConfig() (StockCloseFile) {

	stockCloseFile := StockCloseFile{}
	f, err := os.Open("config.json")
	defer f.Close()
	if err != nil {
		panic(err)
	}

	jsonParser := json.NewDecoder(f)
	jsonParser.Decode(&stockCloseFile)

	return stockCloseFile
}


func CreateStockCloseFile() (StockCloseFile, error){

	stockCloseFile := readConfig()

	checkDirectory(stockCloseFile.DirName)

	fullpath := stockCloseFile.getFullPath()

	_, err := os.Stat(fullpath)
	if err != nil {

		if os.IsNotExist(err) {

			osFile, err1 := os.Create(fullpath) //create osFile here

			if err1 != nil {
				return stockCloseFile, errors.New(err1.Error())

			} else {
				defer osFile.Close()
				return stockCloseFile, nil

			}

		} else {
			return stockCloseFile, nil
		}

	} else {
		return stockCloseFile, nil
	}

}

func (f StockCloseFile) getFullPath() string {
	return f.DirName + f.Filename
}

func (f StockCloseFile) Append(row string) (bool, error) {

	file, err := os.OpenFile(f.getFullPath(), os.O_APPEND, 0666)

	if err==nil {
		defer file.Close()

		file.WriteString(row + "\r\n")
		return true, nil

	} else {
		return false, err
	}

}

func (f *StockCloseFile) Update(rows []string) (bool,error) {

	file, err := os.OpenFile( f.getFullPath(), os.O_TRUNC | os.O_WRONLY, 0644)

	if err==nil {

		defer file.Close()
		for _, row :=range rows {
			if len(row)>0 {
				file.WriteString(row + "\r\n")
			}
		}
		return true, nil

	} else {
		return false, err
	}

}

func (f *StockCloseFile) Truncate() {

	file, _ :=os.OpenFile(f.getFullPath(), os.O_TRUNC, 0666)
	defer file.Close()

}

