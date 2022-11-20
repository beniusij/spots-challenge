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

func GetSpotsInCircle(long float64, lat float64, r int) ([]Spot, error) {
	var spots []Spot
	db := config.GetDb()

	if db == nil {
		return nil, errors.New("database is not available")
	}

	query := fmt.Sprintf(`
		SELECT *
		FROM "MY_TABLE" s
		WHERE ST_DWithin(s.coordinates::geography,ST_MakePoint(%f,%f)::geography,%d)
		ORDER BY
			CASE
				WHEN ST_Distance(s.coordinates::geography,ST_MakePoint(%f,%f)::geography) <= 50 THEN s.rating
				ELSE ST_Distance(s.coordinates::geography,ST_MakePoint(%f,%f)::geography)
			END;
		`,
		long, lat,
		r,
		long, lat,
		long, lat,
	)

	db.Raw(query).Scan(&spots)

	return spots, nil
}

func GetSpotsInSquare(long float64, lat float64, r int) ([]Spot, error) {
	var spots []Spot
	db := config.GetDb()

	if db == nil {
		return nil, errors.New("database is not available")
	}

	return spots, nil
}
