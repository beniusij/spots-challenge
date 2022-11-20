package main

import (
	"github.com/gin-gonic/gin"
	"spotlas-challenge/config"
)

func main() {
	config.InitDb()
	db := config.GetDb()
	router := gin.Default()

	defer db.Close()

	router.GET("/spots")

	_ = router.Run(":3000")
}
