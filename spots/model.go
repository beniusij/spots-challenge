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

func GetSpotsInCircle(radius int, longitude float32, latitude float32) ([]Spot, error) {
	var spots []Spot
	db := config.GetDb()

	if db == nil {
		return nil, errors.New("database is not available")
	}

	query := fmt.Sprintf(`SELECT * FROM "MY_TABLE" LIMIT %d`, 10)
	db.Raw(query).Scan(&spots)

	return spots, nil
}
