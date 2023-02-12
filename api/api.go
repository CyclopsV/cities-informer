package api

import (
	"encoding/json"
	"fmt"
	"github.com/CyclopsV/cities-informer-skillbox/internal/models"
	"github.com/CyclopsV/cities-informer-skillbox/internal/storage"
	"github.com/CyclopsV/cities-informer-skillbox/pkg/pars"
	"github.com/go-chi/chi"
	"net/http"
)

var (
	Cities = storage.Cities{}
)

func init() {
	raw := pars.ParseCSV("sources/Cities.csv")
	Cities.Create(raw)
}

func getCityByIdHandler(w http.ResponseWriter, r *http.Request) {
	jsonBuf, err := pars.ParseResponseToJSON(r.Body)
	if err != nil {
		statusBadRequest(&w, err.Error())
		return
	}
	fields := map[string]string{
		"id": "uint16",
	}
	if err = pars.CheckFields(jsonBuf, fields); err != nil {
		statusBadRequest(&w, err.Error())
		return
	}
	city := Cities.GetCityById(uint16(jsonBuf["id"].(float64)))
	if city == nil {
		statusBadRequest(&w, "Город не найден")
		return
	}
	cityMap := city.ToMap()
	cityBytes, err := json.Marshal(cityMap)
	if err != nil {
		statusInternalServerError(&w, "Не уодалось создать ответ:\n"+err.Error())
		return
	}
	statusOK(&w, cityBytes)
}

func createCityHandler(w http.ResponseWriter, r *http.Request) {
	jsonBuf, err := pars.ParseResponseToJSON(r.Body)
	if err != nil {
		statusBadRequest(&w, err.Error())
		return
	}
	fields := map[string]string{
		"id":         "uint16",
		"name":       "string",
		"region":     "string",
		"district":   "string",
		"foundation": "uint16",
		"population": "uint32",
	}
	if err = pars.CheckFields(jsonBuf, fields); err != nil {
		statusBadRequest(&w, err.Error())
		return
	}
	id := uint16(jsonBuf["id"].(float64))
	name := jsonBuf["name"].(string)
	region := jsonBuf["region"].(string)
	district := jsonBuf["district"].(string)
	population := uint32(jsonBuf["population"].(float64))
	foundation := uint16(jsonBuf["foundation"].(float64))
	city := models.City{}
	city.Create(id, foundation, population, name, region, district)
	if check := Cities.Add(&city); check != nil {
		statusBadRequest(&w, fmt.Sprintf("город с id %v уже существует\n{%v}", city.ID, city))
		return
	}
	statusCreated(&w, []byte{})
}

func deleteCityHandler(w http.ResponseWriter, r *http.Request) {
	jsonBuf, err := pars.ParseResponseToJSON(r.Body)
	if err != nil {
		statusBadRequest(&w, err.Error())
		return
	}
	fields := map[string]string{
		"id": "uint16",
	}
	if err = pars.CheckFields(jsonBuf, fields); err != nil {
		statusBadRequest(&w, err.Error())
		return
	}
	city := Cities.Drop(uint16(jsonBuf["id"].(float64)))
	if city == nil {
		statusBadRequest(&w, "Город не найден")
		return
	}
	statusOK(&w, []byte{})
}

func updateCityHandler(w http.ResponseWriter, r *http.Request) {
	jsonBuf, err := pars.ParseResponseToJSON(r.Body)
	if err != nil {
		statusBadRequest(&w, err.Error())
		return
	}
	fields := map[string]string{
		"id":         "uint16",
		"population": "uint32",
	}
	if err = pars.CheckFields(jsonBuf, fields); err != nil {
		statusBadRequest(&w, err.Error())
		return
	}

	id := uint16(jsonBuf["id"].(float64))
	populationNew := uint32(jsonBuf["population"].(float64))
	city := Cities.GetCityById(id)
	if city == nil {
		statusBadRequest(&w, "Город не найден")
		return
	}
	city.PopulateUpdate(populationNew)

	statusOK(&w, []byte{})
}

func getRegionOrDistrictHandler(w http.ResponseWriter, r *http.Request) {
	jsonBuf, err := pars.ParseResponseToJSON(r.Body)
	if err != nil {
		statusBadRequest(&w, err.Error())
		return
	}
	targetMode := chi.URLParam(r, "regionOrDistrict")
	fields := map[string]string{
		targetMode: "string",
	}
	if err = pars.CheckFields(jsonBuf, fields); err != nil {
		statusBadRequest(&w, err.Error())
		return
	}
	target := jsonBuf[targetMode].(string)
	var citiesList []*models.City
	if targetMode == "district" {
		citiesList = Cities.GetCitiesByRegionOrDistrict(target, true)
	} else {
		citiesList = Cities.GetCitiesByRegionOrDistrict(target, false)
	}
	var citiesMapList []map[string]interface{}
	for _, city := range citiesList {
		citiesMapList = append(citiesMapList, city.ToMap())
	}
	citiesMap := map[string]interface{}{
		"cities": citiesMapList,
	}
	citiesBytes, err := json.Marshal(citiesMap)
	if err != nil {
		statusInternalServerError(&w, "Не уодалось создать ответ:\n"+err.Error())
		return
	}
	statusOK(&w, citiesBytes)
}

func getPopulationOrFoundationHandler(w http.ResponseWriter, r *http.Request) {
	jsonBuf, err := pars.ParseResponseToJSON(r.Body)
	if err != nil {
		statusBadRequest(&w, err.Error())
		return
	}
	targetMode := chi.URLParam(r, "populationOrFoundation")
	fields := map[string]string{
		"from": "uint32",
		"to":   "uint32",
	}
	if err = pars.CheckFields(jsonBuf, fields); err != nil {
		statusBadRequest(&w, err.Error())
		return
	}
	from := uint32(jsonBuf["from"].(float64))
	to := uint32(jsonBuf["to"].(float64))
	var citiesList []*models.City
	if targetMode == "population" {
		citiesList = Cities.GetCitiesByPopulationOrFoundation(from, to, false)
	} else {
		citiesList = Cities.GetCitiesByPopulationOrFoundation(from, to, true)
	}
	var citiesMapList []map[string]interface{}
	for _, city := range citiesList {
		citiesMapList = append(citiesMapList, city.ToMap())
	}
	citiesMap := map[string]interface{}{
		"cities": citiesMapList,
	}
	citiesBytes, err := json.Marshal(citiesMap)
	if err != nil {
		statusInternalServerError(&w, "Не уодалось создать ответ:\n"+err.Error())
		return
	}
	statusOK(&w, citiesBytes)
}
