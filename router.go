package main

import "github.com/gin-gonic/gin"

func setupRouter() *gin.Engine {
	router := gin.New()
	router.Use(logRequest)

	router.GET("/coindesk_prices", getLatestPrices)

	return router
}
