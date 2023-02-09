package pars

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
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

/*
CheckFields

	data - Присланные данные
	fields - Поля которые должны присутсвовать в присланных данных
		формат: map[name]type
*/
func CheckFields(data map[string]interface{}, fields map[string]string) error {
	for fieldName := range fields {
		if _, ok := data[fieldName]; !ok {
			return fmt.Errorf("отсутсвует поле `%v`", fieldName)
		}
	}

	for valName, val := range data {
		if _, ok := fields[valName]; !ok {
			continue
		}

		valFloat, ok := val.(float64)
		statusType := true
		switch fields[valName] {
		case "uint16":
			if !ok || valFloat > math.MaxUint16 || valFloat < 0 {
				statusType = false
			}
		case "uint32":
			if !ok || valFloat > math.MaxInt32 || valFloat < 0 {
				statusType = false
			}
		case "string":
			if _, ok = val.(string); !ok {
				statusType = false
			}
		}
		if !statusType {
			return fmt.Errorf("поле `%v` имеет неверный формат", valName)
		}
	}
	return nil
}
