package models

import (
	"fmt"
	"strconv"
)

type City struct {
	ID         uint16
	name       string
	region     string
	district   string
	population uint32
	foundation uint16
}

func (c *City) String() string {
	return fmt.Sprintf("%v\n", *c)
}

func (c *City) Create(id, foundation uint16, population uint32, name, region, district string) {
	c.ID = id
	c.name = name
	c.region = region
	c.district = district
	c.foundation = foundation
	c.population = population
}

func (c *City) CreateFromRAW(rawInfo []string) error {
	id, err := strconv.ParseUint(rawInfo[0], 10, 16)
	if err != nil {
		return err
	}
	population, err := strconv.ParseUint(rawInfo[4], 10, 64)
	if err != nil {
		return err
	}
	foundation, err := strconv.ParseUint(rawInfo[5], 10, 64)
	if err != nil {
		return err
	}
	name := rawInfo[1]
	region := rawInfo[2]
	district := rawInfo[3]
	c.Create(uint16(id), uint16(foundation), uint32(population), name, region, district)
	return nil
}

func (c *City) PopulateUpdate(newPopulation uint32) {
	c.population = newPopulation
}

func (c *City) GetRegion() string {
	return c.region
}

func (c *City) GetDistrict() string {
	return c.district
}

func (c *City) GetFoundation() uint16 {
	return c.foundation
}

func (c *City) GetPopulation() uint32 {
	return c.population
}

func (c *City) ToMap() map[string]interface{} {
	cityMap := map[string]interface{}{
		"id":         c.ID,
		"name":       c.name,
		"region":     c.region,
		"district":   c.district,
		"population": c.population,
		"foundation": c.foundation,
	}
	return cityMap
}

func (c *City) ToList() []string {
	id := strconv.FormatUint(uint64(c.ID), 10)
	population := strconv.FormatUint(uint64(c.population), 10)
	foundation := strconv.FormatUint(uint64(c.foundation), 10)
	cityLists := []string{
		id,
		c.name,
		c.region,
		c.district,
		population,
		foundation,
	}
	return cityLists
}
