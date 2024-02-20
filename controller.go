package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// format the required response
func structureResponse(prices_data map[string]string) map[string]map[string]map[string]string {

	return map[string]map[string]map[string]string{
		"data": {
			"bitcoin": prices_data,
		},
	}

}

func getLatestPrices(c *gin.Context) {

	finalResponse, err := retrieve()

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve latest prices from cache", "error": err.Error()})
		return
	}

	if finalResponse == nil {
		response, err := fetchData(COINDESK_API_URL)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Unable to fetch and parse data", "error": err.Error()})
			return
		}
		finalResponse = structureResponse(response)
		err = store(finalResponse)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Failed to store latest prices in cache", "error": err.Error()})
			return
		}
	}

	c.IndentedJSON(http.StatusOK, finalResponse)
}
