package spots

import (
	"errors"
	"spotlas-challenge/config"
)

type Spot struct {
	Id          string `gorm:"primary_key"`
	Name        string
	Website     string
	Coordinates string
	Description string
	Rating      float32
}

func GetSpotsInSquare(radius int, longitude float32, latitude float32) ([]Spot, error) {
	db := config.GetDb()

	if db == nil {
		return nil, errors.New("database is not available")
	}

	return []Spot{}, nil
}
