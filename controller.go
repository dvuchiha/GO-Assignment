package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getLatestPrices(c *gin.Context) {

	finalResponse, err := retrieve()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve latest prices from cache", "error": err.Error()})
		return
	}

	if finalResponse == nil {
		response, err := fetchData(COINDESK_API_URL)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Unable to fetch data from 3rd party", "error": err.Error()})
			return
		}
		pricesData := extractPrices(response)
		finalResponse = structureResponse(pricesData)
		err = store(finalResponse)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Failed to store latest prices in cache", "error": err.Error()})
			return
		}
	}

	c.IndentedJSON(http.StatusOK, finalResponse)
}
