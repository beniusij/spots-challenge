package spots

import (
	"errors"
	"spotlas-challenge/config"
)

func GetSpotsInSquare(radius int, longitude float32, latitude float32) ([]any, error) {
	db := config.GetDb()

	if db == nil {
		return nil, errors.New("database is not available")
	}

	return nil, errors.New("test")
}
