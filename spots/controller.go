package spots

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"strconv"
)

type Response struct {
	Status  string
	Results []Spot
}

func GetSpotsByRadius(c *gin.Context) {
	var resp Response
	var err error
	var searchType string
	var longitude, latitude float64
	var radius int

	if !validateQueryParams(c) {
		return
	}

	searchType = c.Query("type")
	longitude, _ = strconv.ParseFloat(c.Query("longitude"), 32)
	latitude, _ = strconv.ParseFloat(c.Query("latitude"), 32)
	radius, _ = strconv.Atoi(c.Query("radius"))

	switch searchType {
	case "circle":
		resp.Results, err = GetSpotsInCircle(longitude, latitude, radius)
		if err != nil {
			abortRequest(c, err)
			return
		}
	}

	resp.Status = "ok"
	c.JSON(http.StatusOK, resp)
}

func validateQueryParams(c *gin.Context) bool {
	var val string

	safe := true
	errJson := gin.H{"message": "Invalid payload"}
	expectedParams := [4]string{"longitude", "latitude", "radius", "type"}

	// Check if all required parameters are set and valid
	for _, p := range expectedParams {
		switch p {
		case expectedParams[0]:
			val = c.Query(expectedParams[0])

			if _, err := strconv.ParseFloat(val, 32); err != nil {
				safe = false
			}
		case expectedParams[1]:
			val = c.Query(expectedParams[1])

			if _, err := strconv.ParseFloat(val, 32); err != nil {
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

func abortRequest(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"message": "failed retrieving data",
		"details": err.Error(),
	})
	c.Abort()
}
