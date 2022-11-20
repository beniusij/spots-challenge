package main

import (
	"github.com/gin-gonic/gin"
	"spotlas-challenge/config"
	"spotlas-challenge/spots"
)

func main() {
	config.InitDb()
	db := config.GetDb()
	router := gin.Default()

	defer db.Close()

	router.GET("/spots", spots.GetSpotsByRadius)

	_ = router.Run(":3000")
}
