package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
)

type trnAction struct {
	valueDate    time.Time
	baseCurrency string
	debAccount   string
	credAccount  string
	amount       float64
	parseError   string
}

func (x trnAction) String() (s string) {
	s = "dt=" + x.valueDate.Format("2006-01-02") + ", curr=" + x.baseCurrency + ", sum=" +
		strconv.FormatFloat(x.amount, 'f', -1, 64) + ": " +
		"DT account = " + x.debAccount + " / CT account = " + x.credAccount
	if x.parseError != "" {
		s += "\n" + x.parseError
	}
	return
}

type trnArray []trnAction

func (x trnArray) String() (s string) {
	s = ""
	for _, value := range x {
		s += value.String() + "\n"
	}
	return
}

//LoadFromFile -
func (x *trnArray) LoadFromFile(filename string) error {
	buffer, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	//
	csvReader := csv.NewReader(bytes.NewReader(buffer))
	csvReader.Comma = ';'
	csvArray, err := csvReader.ReadAll()
	if err != nil {
		return err
	}
	//
	badDate, err := time.Parse("2006-01-02", "1900-01-01")
	if err != nil {
		return err
	}
	*x = make(trnArray, len(csvArray))
	for idx, csvRec := range csvArray {
		(*x)[idx].parseError = ""
		if (*x)[idx].valueDate, err = time.Parse("2006-01-02", csvRec[0]); err != nil {
			(*x)[idx].valueDate = badDate
			(*x)[idx].parseError += err.Error() + "\n"
		}
		(*x)[idx].baseCurrency = csvRec[1]
		(*x)[idx].debAccount = csvRec[2]
		(*x)[idx].credAccount = csvRec[3]
		if (*x)[idx].amount, err = strconv.ParseFloat(csvRec[4], 64); err != nil {
			(*x)[idx].amount = 0
			(*x)[idx].parseError += err.Error() + "\n"
		}
	}
	return nil
}

//SaveToFile -
func (x trnArray) SaveToFile(filename string) error {
	//
	csvArray := make([][]string, len(x))
	for idx, value := range x {
		csvArray[idx] = make([]string, 5)
		csvArray[idx][0] = value.valueDate.Format("2006-01-02")
		csvArray[idx][1] = value.baseCurrency
		csvArray[idx][2] = value.debAccount
		csvArray[idx][3] = value.credAccount
		csvArray[idx][4] = strconv.FormatFloat(value.amount, 'f', -1, 64)
	}
	//
	var buffer bytes.Buffer
	csvWriter := csv.NewWriter(&buffer)
	csvWriter.Comma = ';'
	csvWriter.WriteAll(csvArray)
	if err := csvWriter.Error(); err != nil {
		return err
	}
	//
	if err := ioutil.WriteFile(filename, buffer.Bytes(), 0644); err != nil {
		return err
	}
	return nil
}

func main() {
	var x trnArray

	x = append(x,
		trnAction{time.Now(), "RUR", "30109", "30110", 23423.56, ""},
		trnAction{time.Now(), "RUR", "30127", "30110", 23400.47, ""},
		trnAction{time.Now(), "RUR", "30145", "30112", 1231223423.25, ""},
		trnAction{time.Now(), "RUR", "40109", "40110", 23423423423.00, ""})
	fmt.Println(x)
	if err := x.SaveToFile("c:/ice/trn1.csv"); err != nil {
		panic(err.Error())
	}

	x = append(x,
		trnAction{time.Now(), "RUR", "30127", "30109", 123423.56, ""},
		trnAction{time.Now(), "RUR", "30127", "30109", 123400.47, ""},
		trnAction{time.Now(), "RUR", "30112", "30112", 3423.25, ""},
		trnAction{time.Now(), "RUR", "40109", "30109", 992423423.00, ""})
	fmt.Println(x)
	if err := x.SaveToFile("c:/ice/trn2.csv"); err != nil {
		panic(err.Error())
	}

	if err := x.LoadFromFile("c:/ice/trn1.csv"); err != nil {
		panic(err.Error())
	}
	fmt.Println(x)
}
