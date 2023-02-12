package storage

import (
	"encoding/csv"
	"fmt"
	"os"
)

func (cs *Cities) Save() error {
	var citiesList [][]string
	for _, city := range *cs {
		cityList := city.ToList()
		citiesList = append(citiesList, cityList)
	}

	f, err := os.OpenFile("sources/Cities.csv", os.O_WRONLY|os.O_CREATE, 0222)
	defer f.Close()
	if err != nil {
		return fmt.Errorf("Не удалось открыть файл\n---\n%v", err.Error())
	}
	w := csv.NewWriter(f)
	if err := w.WriteAll(citiesList); err != nil {
		return fmt.Errorf("Не удалось сохранить данные\n---\n%v", w.Error())
	}
	return nil
}
