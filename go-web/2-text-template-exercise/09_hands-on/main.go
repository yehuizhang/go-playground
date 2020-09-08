package main

import (
	"encoding/csv"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type record struct {
	Date time.Time
	High float64
	Low  float64
}

func parseCsv(filePath string) []record {
	src, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer src.Close()

	reader := csv.NewReader(src)
	rows, err := reader.ReadAll()
	if err != nil {
		log.Fatalln(err)
	}

	records := make([]record, 0, len(rows))

	for i, row := range rows {
		if i == 0 {
			continue
		}

		date, _ := time.Parse("2006-01-02", row[0])
		high, _ := strconv.ParseFloat(row[2], 64)
		low, _ := strconv.ParseFloat(row[3], 64)

		records = append(records, record{
			Date: date,
			High: high,
			Low:  low,
		})
	}
	return records
}

func loadTemplate(res http.ResponseWriter, req *http.Request) {
	records := parseCsv("table.csv")

	tpl := template.Must(template.ParseFiles("index.gohtml"))

	err := tpl.Execute(res, records)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	http.HandleFunc("/", loadTemplate)
	http.ListenAndServe(":3000", nil)
}
