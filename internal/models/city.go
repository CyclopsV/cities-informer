package models

import (
	"fmt"
	"strconv"
)

type City struct {
	id         uint16
	name       string
	region     string
	district   string
	population uint32
	foundation uint16
}

func (c *City) String() string {
	return fmt.Sprintf("%#v\n", c)
}

func (c *City) Create(rawInfo []string) error {
	buf, err := strconv.ParseUint(rawInfo[0], 10, 16)
	if err != nil {
		return err
	}
	c.id = uint16(buf)

	buf, err = strconv.ParseUint(rawInfo[4], 10, 32)
	if err != nil {
		return err
	}
	c.population = uint32(buf)

	buf, err = strconv.ParseUint(rawInfo[4], 10, 16)
	if err != nil {
		return err
	}
	c.foundation = uint16(buf)

	c.name = rawInfo[1]
	c.region = rawInfo[2]
	c.district = rawInfo[3]
	return nil
}
