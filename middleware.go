package main

import (
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
)

var lastTimeReceivedACall int64 = -1

func ipInputValidation(c *gin.Context){
	ip := c.Query("ip")
	if ip == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "IP is required"})
		return
	}
	regex := regexp.MustCompile(`\b\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\b`)
	if !regex.MatchString(ip) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid IP format"})
		return
	}

	c.Next()
}

func maxCallsCheck(keepGapOfSecondsBetweenCalls int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		if(lastTimeReceivedACall == -1){
			fmt.Println("First call",lastTimeReceivedACall)
			lastTimeReceivedACall = time.Now().UnixMilli()
		}else{
			timeGapInMilliseconds := time.Now().UnixMilli() - lastTimeReceivedACall
			fmt.Println("Time gap in milliseconds",timeGapInMilliseconds , "Keep gap of seconds between calls",keepGapOfSecondsBetweenCalls * 1000)


			if(timeGapInMilliseconds < keepGapOfSecondsBetweenCalls * 1000){
				c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
				return
			}else{
				lastTimeReceivedACall = time.Now().UnixMilli()
				fmt.Println("Updating last time received a call",lastTimeReceivedACall)
			}
		}
		c.Next()
	}
}