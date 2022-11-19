package spots_test

import (
	"spotlas-challenge/spots"
	"testing"
)

func Test_GetSpotsInSquare(t *testing.T) {
	t.Run("should return error if database connection is not available", func(t *testing.T) {
		_, err := spots.GetSpotsInSquare(10, 10.0, 10.0)

		expected := "database is not available"
		if err.Error() != expected {
			t.Errorf("expected %q, got %q", expected, err.Error())
		}
	})
}
