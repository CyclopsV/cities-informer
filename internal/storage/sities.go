package storage

import (
	"fmt"
	"github.com/CyclopsV/cities-informer-skillbox/internal/models"
	"log"
	"math"
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
		if err := city.CreateFromRAW(raw); err != nil {
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

func (cs *Cities) Add(city *models.City) error {
	if checkCity, ok := (*cs)[city.ID]; ok {
		return fmt.Errorf("Город с таким ID уже есть: %v\n", checkCity)
	}
	(*cs)[city.ID] = city
	return nil
}

func (cs *Cities) Drop(id uint16) *models.City {
	if targetCity, ok := (*cs)[id]; ok {
		delete(*cs, id)
		return targetCity
	}
	return nil
}

/*
GetCitiesByRegionOrDistrict

	mode (bool) - режим работы функции
		true - поиск по области (district)
		false - поиск по региону (region)
*/
func (cs *Cities) GetCitiesByRegionOrDistrict(target string, mode bool) []*models.City {
	var (
		cities     []*models.City
		chekTarget bool
	)
	for _, city := range *cs {
		chekTarget = false
		if mode {
			chekTarget = city.GetDistrict() == target
		} else {
			chekTarget = city.GetRegion() == target
		}
		if chekTarget {
			cities = append(cities, city)
		}
	}
	return cities
}

/*
GetCitiesByPopulationOrFoundation

	mode (bool) - режим работы функции
		true - поиск по дате основания (Foundation)
		false - поиск по числености населения (Population)
*/
func (cs *Cities) GetCitiesByPopulationOrFoundation(from, to uint32, mode bool) []*models.City {
	var (
		cities []*models.City
		target uint32
	)

	if to == 0 {
		if mode {
			to = math.MaxInt16
		} else {
			to = math.MaxInt32
		}
	}
	for _, city := range *cs {
		if mode {
			target = uint32(city.GetFoundation())
		} else {
			target = city.GetPopulation()
		}
		if target >= from && target <= to {
			cities = append(cities, city)
		}
	}
	return cities
}
