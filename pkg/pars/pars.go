package pars

import (
	"encoding/csv"
	"encoding/json"
	"io"
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

func ParseResponseToJSON(body io.ReadCloser) (map[string]interface{}, error) {
	var (
		message map[string]interface{}
		err     error
	)
	if content, err := io.ReadAll(body); err == nil {
		err = json.Unmarshal(content, &message)
	}
	return message, err
}
