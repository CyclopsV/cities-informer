package pars

import (
	"encoding/csv"
	"log"
	"os"
)

func ParseCSV(filePath string) [][]string {
	log.Println("Извлечение данных из файла")
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var resultList [][]string
	for {
		row, err := reader.Read()
		if err != nil {
			break
		}
		resultList = append(resultList, row)
	}
	return resultList
}
