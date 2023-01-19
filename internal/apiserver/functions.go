package apiserver

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func getNumbersOfDays(c *gin.Context) {
	now := time.Now()
	year, err := strconv.Atoi(c.Param("year"))
	if err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	target := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	days := calculateTime(now, target)
	switch {
	case now.Year() == target.Year():
		switch{
		case now.Month() == target.Month():
			switch{
			case now.Day() == target.Day():
				c.JSON(200, gin.H{"Days gone": days})
			case now.Day() > target.Day():
				c.JSON(200, gin.H{"Days gone": days})
			case now.Day() < target.Day():
				c.JSON(200, gin.H{"Days left": days})
			}
		case now.Month() > target.Month():
			c.JSON(200, gin.H{"Days gone": days})
			
		case now.Month() < target.Month():
			c.JSON(200, gin.H{"Days left": days})
		}
	case now.Year() > target.Year():
		c.JSON(200, gin.H{"Days gone": days})
			
	case now.Year() < target.Year():
		c.JSON(200, gin.H{"Days left": days})
	}
	}

func calculateTime(now, target time.Time) int {
	days := abs(int(now.Sub(target).Hours() / 24))
	return days
}

func abs(x int) int {
	if x < 0{
		return -x
	} 
	return x
}