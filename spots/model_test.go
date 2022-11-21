package spots_test

import (
	"reflect"
	"spotlas-challenge/config"
	"spotlas-challenge/spots"
	"testing"
)

func Test_GetSpotsInCircle(t *testing.T) {
	t.Run("should return error if database connection is not available", func(t *testing.T) {
		config.ResetDb()
		_, err := spots.GetSpotsInCircle(10.0, 10.0, 10)

		expected := "database is not available"
		if err.Error() != expected {
			t.Errorf("expected %q, got %q", expected, err.Error())
		}
	})

	t.Run("should return array of Spots", func(t *testing.T) {
		config.InitTestDb()
		results, err := spots.GetSpotsInCircle(-0.1387165, 51.5125127, 100)

		if err != nil {
			t.Errorf(err.Error())
		}

		if len(results) == 0 {
			t.Errorf("expected array with Spot, got empty array")
		}

		if reflect.TypeOf(results[0]) != reflect.TypeOf(spots.Spot{}) {
			t.Errorf("expected type to be Spot, got %q", reflect.TypeOf(results[0]))
		}

		if results[0].Rating > results[1].Rating {
			t.Errorf("expected to be sorted in descending order by rating when distance <50 meters")
		}
	})
}

func Test_GetSpotsInSquare(t *testing.T) {
	t.Run("should return error if database connection is not available", func(t *testing.T) {
		config.ResetDb()
		_, err := spots.GetSpotsInSquare(10.0, 10.0, 10)

		expected := "database is not available"
		if err.Error() != expected {
			t.Errorf("expected %q, got %q", expected, err.Error())
		}
	})

	t.Run("should return array of Spots", func(t *testing.T) {
		config.InitTestDb()
		results, err := spots.GetSpotsInSquare(-0.1387165, 51.5125127, 100)

		if err != nil {
			t.Errorf(err.Error())
		}

		if len(results) == 0 {
			t.Errorf("expected array with Spot, got empty array")
		}

		if reflect.TypeOf(results[0]) != reflect.TypeOf(spots.Spot{}) {
			t.Errorf("expected type to be Spot, got %q", reflect.TypeOf(results[0]))
		}

		if results[0].Rating > results[1].Rating {
			t.Errorf("expected to be sorted in descending order by rating when distance <50 meters")
		}
	})
}
