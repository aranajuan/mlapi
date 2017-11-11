package main

import (
	"github.com/gin-gonic/gin"
	"ml/itemCat"
	"ml/math"
)

func priceRoute(r *gin.Engine) {
	var a itemCat.ItemCat
	r.GET("/categories/:idCat/price", func(c *gin.Context) {
		if a.LoadItems(c.Param("idCat")) == 1 {
			c.JSON(504, gin.H{
				"error": "Could not connect to ML's api",
			})
			return
		}
		prices, err := a.GetPrices()
		if err != 0 {
			c.JSON(404, gin.H{
				"error": "Could not get prices for this categorie",
			})
			return
		}
		max, min, avg, _ := math.GetStats(prices)
		c.JSON(200, gin.H{
			"max":       max,
			"min":       min,
			"suggested": avg,
		})
	})
}
