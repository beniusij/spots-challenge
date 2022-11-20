package spots

import (
	"errors"
	"fmt"
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

func GetSpotsInCircle(longitude float64, latitude float64, radius int) ([]Spot, error) {
	var spots []Spot
	db := config.GetDb()

	if db == nil {
		return nil, errors.New("database is not available")
	}

	query := fmt.Sprintf(`
		SELECT * FROM "MY_TABLE" s 
		WHERE ST_DWithin(
			s.coordinates::geography,
			ST_MakePoint(%f,%f)::geography,
			%d
		)`,
		longitude,
		latitude,
		radius,
	)

	db.Raw(query).Scan(&spots)

	return spots, nil
}

func GetSpotsInSquare(longitude float64, latitude float64, radius int) ([]Spot, error) {
	var spots []Spot
	db := config.GetDb()

	if db == nil {
		return nil, errors.New("database is not available")
	}

	return spots, nil
}
