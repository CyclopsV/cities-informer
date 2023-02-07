package storage

import (
	"github.com/CyclopsV/cities-informer-skillbox/internal/models"
	"log"
)

type Cities map[uint16]*models.City

func (cs *Cities) String() string {
	sitiesInfoStr := "Иформация о городах:\n"
	for _, city := range *cs {
		cityStr := city.String()
		sitiesInfoStr += "\t" + cityStr
	}
	return sitiesInfoStr
}

func (cs *Cities) Create(rawInfo [][]string) {
	for _, raw := range rawInfo {
		city := models.City{}
		if err := city.Create(raw); err != nil {
			log.Printf("Ошибка распознания информации о городе: %#v\n\terr: %v\n", raw, err)
			continue
		}
		(*cs)[city.ID] = &city
	}
}

func (cs *Cities) GetCityById(id uint16) *models.City {
	targetCity, ok := (*cs)[id]
	if !ok {
		return nil
	}
	return targetCity
}
