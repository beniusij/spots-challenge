package spots_test

import (
	"reflect"
	"spotlas-challenge/config"
	"spotlas-challenge/spots"
	"testing"
)

func Test_GetSpotsInCircle(t *testing.T) {
	config.InitTestDb()

	t.Run("should return error if database connection is not available", func(t *testing.T) {
		_, err := spots.GetSpotsInCircle(10.0, 10.0, 10)

		expected := "database is not available"
		if err.Error() != expected {
			t.Errorf("expected %q, got %q", expected, err.Error())
		}
	})

	t.Run("should return array of Spots", func(t *testing.T) {
		results, err := spots.GetSpotsInCircle(-8.473656, 51.899216, 100)
		// lat, long
		// 51.899216,-8.473656

		if err != nil {
			t.Errorf(err.Error())
		}

		if len(results) == 0 {
			t.Errorf("expected array with Spot, got empty array")
		}

		if reflect.TypeOf(results[0]) != reflect.TypeOf(spots.Spot{}) {
			t.Errorf("expected type to be Spot, got %q", reflect.TypeOf(results[0]))
		}
	})
}
