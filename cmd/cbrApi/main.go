package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/tealeg/xlsx/v3"
	"golang.org/x/text/encoding/charmap"
)

type CurrencyData struct {
	Date     time.Time
	CharCode string
	Rate     float64
}

// ValCurs представляет корневой элемент XML
type ValCurs struct {
	Date    string   `xml:"Date,attr"`
	Valutes []Valute `xml:"Valute"`
}

// Valute представляет информацию о валюте
type Valute struct {
	CharCode string `xml:"CharCode"`
	Rate     string `xml:"VunitRate"`
}

type Answer struct {
	charCode string
	maxRate  float64
	maxDate  time.Time
	minRate  float64
	minDate  time.Time
	avg      float64
}

type Rate struct {
	Date time.Time
	Rate float64
}

func main() {

	to := time.Now()
	from := to.AddDate(0, 0, -90)

	currencyData, err := getCurrency(from, to)
	if err != nil {
		fmt.Println("Ошибка получения данных:", err)
	}

	answerData := fillAnswer(currencyData)

	for i, val := range answerData {
		fmt.Println(i, val.charCode, val.maxRate, val.minRate, val.avg, val.maxDate, val.minDate)
	}
	fmt.Println()

	filename := fmt.Sprintf("Курс_валют_с_%s_по_%s.xlsx", from.Format("02-01-2006"), to.Format("02-01-2006"))
	err = makeExcel(answerData, filename)
	if err != nil {
		fmt.Println("Не удалось создать exel файл:", err)
	}

	fmt.Scanln()
}

func makeExcel(answerData []Answer, filename string) error {
	file := xlsx.NewFile()
	sheet, _ := file.AddSheet("Currency")

	exePath, _ := os.Executable()
	exeDir := filepath.Dir(exePath)
	fullPath := filepath.Join(exeDir, filename)

	headers := []string{
		"Номер", "Валюта", "Max", "Min", "Avg", "Дата max курса", "Дата min курса",
	}

	headerRow := sheet.AddRow()
	for _, header := range headers {
		cell := headerRow.AddCell()
		cell.Value = header
		cell.GetStyle().Font.Bold = true
	}

	// Заполняем данные
	for i, val := range answerData {
		row := sheet.AddRow()

		//Номер
		row.AddCell().SetInt(i)
		//Валюта
		row.AddCell().Value = val.charCode
		//Max
		row.AddCell().SetFloat(val.maxRate)
		//Min
		row.AddCell().SetFloat(val.minRate)
		//Avg
		row.AddCell().SetFloat(val.avg)
		//Дата max курса
		row.AddCell().Value = val.maxDate.Format("02.01.2006")
		//Дата min курса
		row.AddCell().Value = val.minDate.Format("02.01.2006")
	}

	err := file.Save(fullPath)
	if err != nil {
		return fmt.Errorf("error saving Excel file: %v", err)
	}

	fmt.Println("Файл сохранен по адресу:", fullPath)
	return nil
}

func fillAnswer(CurrencyData map[string]map[time.Time]float64) (resData []Answer) {

	for charCode, val := range CurrencyData {
		var maxRate float64
		var maxDate time.Time
		var minRate float64 = 1000 // Чтобы начинать сравнивать не с 0
		var minDate time.Time
		var totalSum float64
		totalCount := 0

		for date, rate := range val {
			totalCount++
			totalSum += rate

			if maxRate < rate {
				maxRate = rate
				maxDate = date
			}

			if minRate > rate {
				minRate = rate
				minDate = date
			}

		}
		answer := Answer{
			charCode: charCode,
			maxRate:  maxRate,
			maxDate:  maxDate,
			minRate:  minRate,
			minDate:  minDate,
			avg:      totalSum / float64(totalCount),
		}

		resData = append(resData, answer)

	}
	return resData
}

func getCurrency(from, to time.Time) (map[string]map[time.Time]float64, error) {

	rateMap := make(map[string]map[time.Time]float64)

	for from.Before(to.AddDate(0, 0, 1)) { // пофиксить на числовой цикл
		dateStr := from.Format("02/01/2006")

		url := fmt.Sprintf("https://cbr.ru/scripts/XML_daily.asp?date_req=%s", dateStr)

		resp, err := http.Get(url)
		if err != nil {
			return nil, fmt.Errorf("HTTP error:%v", err)
		}

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("the server returned the status: %d", resp.StatusCode)
		}

		body, _ := io.ReadAll(resp.Body)
		resp.Body.Close()

		valCurs := parseBody(body)
		for _, val := range valCurs.Valutes {

			newDate, err := time.Parse("02.01.2006", valCurs.Date)
			if err != nil {
				return nil, fmt.Errorf("error getting Date: %v", err)
			}
			newRate, err := strconv.ParseFloat(strings.ReplaceAll(val.Rate, ",", "."), 64)
			if err != nil {
				return nil, fmt.Errorf("error getting Rate: %v", err)
			}

			currencyRate := CurrencyData{
				Date:     newDate,
				CharCode: val.CharCode,
				Rate:     newRate,
			}

			if _, exists := rateMap[currencyRate.CharCode]; !exists {
				rateMap[currencyRate.CharCode] = make(map[time.Time]float64)
			}

			rateMap[currencyRate.CharCode][currencyRate.Date] = currencyRate.Rate
		}

		from = from.AddDate(0, 0, 1)
	}

	return rateMap, nil
}

func parseBody(body []byte) (valCurs ValCurs) {

	decoder := xml.NewDecoder(bytes.NewReader(body))
	decoder.CharsetReader = charsetReader

	err := decoder.Decode(&valCurs)
	if err != nil {
		fmt.Println("XML parsing error:", err)
	}

	return valCurs
}

// Преобразование windows-1251 в UTF-8
func charsetReader(charset string, input io.Reader) (io.Reader, error) {
	if charset == "windows-1251" {
		return charmap.Windows1251.NewDecoder().Reader(input), nil
	}
	return nil, fmt.Errorf("неизвестная кодировка: %s", charset)
}
