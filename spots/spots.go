package spots

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"strconv"
)

type errorResponse struct {
	message string
	details []string
}

func GetSpotsByRadius(c *gin.Context) {
	if !validateQueryParams(c) {
		return
	}
}

func validateQueryParams(c *gin.Context) bool {
	var val string
	var safe bool

	errJson := gin.H{"message": "Invalid payload"}

	expectedParams := [4]string{"longitude", "latitude", "radius", "type"}
	pattern := "([\\-0-1]{0,3}\\.?[0-9]{0,5})"

	// Check if all required parameters are set and valid
	for _, p := range expectedParams {
		switch p {
		case expectedParams[0]:
			val = c.Query(expectedParams[0])
			match, err := regexp.MatchString(pattern, val)

			if err != nil || !match || val == "" {
				safe = false
			}
		case expectedParams[1]:
			val = c.Query(expectedParams[1])
			match, err := regexp.MatchString(pattern, val)

			if err != nil || !match || val == "" {
				safe = false
			}
		case expectedParams[2]:
			val = c.Query(expectedParams[2])
			intVal, err := strconv.Atoi(val)

			if err != nil || intVal < 0 {
				safe = false
			}
		case expectedParams[3]:
			val = c.Query(expectedParams[3])
			match, err := regexp.MatchString("([circle|square])", val)

			if err != nil || !match {
				safe = false
			}
		}
	}

	if !safe {
		c.JSON(http.StatusNotAcceptable, errJson)
		c.Abort()
	}

	return safe
}
