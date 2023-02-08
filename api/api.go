package api

import (
	"encoding/json"
	"github.com/CyclopsV/cities-informer-skillbox/internal/storage"
	"github.com/CyclopsV/cities-informer-skillbox/pkg/pars"
	"math"
	"net/http"
)

var (
	cities = storage.Cities{}
)

func init() {
	raw := pars.ParseCSV("sources/cities.csv")
	cities.Create(raw)
}

func getCityByIdHandler(w http.ResponseWriter, r *http.Request) {
	jsonBuf, err := pars.ParseResponseToJSON(r.Body)
	if err != nil {
		statusBadRequest(&w, "Не распознано тело запроса:\n"+err.Error())
		return
	}
	idInterface, ok := jsonBuf["id"]
	if !ok {
		statusBadRequest(&w, "Не найдено поле `id`")
		return
	}
	id, ok := idInterface.(float64)
	if !ok || id > math.MaxInt16 {
		statusBadRequest(&w, "Неверный тип данных")
		return
	}
	city := cities.GetCityById(uint16(id))
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
