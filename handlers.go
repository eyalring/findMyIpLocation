package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func calculateCountryFromIP(c *gin.Context, dbConnector *DatabaseConnector) {

	ip := c.Query("ip")

	myLocation := dbConnector.Finder.FindLocation(ip)
	   if myLocation.found {
        c.IndentedJSON(http.StatusOK, gin.H{"country": myLocation.loc.country, "city": myLocation.loc.city})
    } else {
        c.IndentedJSON(http.StatusNotFound, gin.H{"error": "IP not found"})
    }
}