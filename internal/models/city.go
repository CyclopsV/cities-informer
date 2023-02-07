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
	foundation, err := strconv.ParseUint(rawInfo[4], 10, 64)
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
