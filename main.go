package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type item struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Price  float32 `json:"price"`
	Amount int     `json:"amount"`
}

type order struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	Items  []item `json:"items"`
}

var orders = []order{}

func addOrder(o order) {
	orders = append(orders, o)
}

func getOrders() []order {
	return orders
}

func getOrder(id string) *order {
	for _, o := range orders {
		if o.ID == id {
			return &o
		}
	}
	return nil
}

func deleteOrder(id string) {
	for i, o := range orders {
		if o.ID == id {
			orders = append(orders[:i], orders[i+1:]...)
			return
		}
	}
}

func updateOrder(id string, o order) {
	for i, order := range orders {
		if order.ID == id {
			orders[i] = o
			return
		}
	}
}

func main() {
	router := gin.Default()
	router.POST("/orders", func(c *gin.Context) {
		var o order
		if err := c.BindJSON(&o); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		addOrder(o)
		c.JSON(http.StatusCreated, o)
	})
	router.GET("/orders", func(c *gin.Context) {
		c.JSON(http.StatusOK, getOrders())
	})

	router.GET("/orders/:id", func(c *gin.Context) {
		id := c.Param("id")
		o := getOrder(id)
		if o == nil {
			c.JSON(http.StatusNotFound, gin.H{})
			return
		}
		c.JSON(http.StatusOK, o)
	})

	router.DELETE("/orders/:id", func(c *gin.Context) {
		id := c.Param("id")
		deleteOrder(id)
		c.JSON(http.StatusNoContent, gin.H{})
	})

	router.PUT("/orders/:id", func(c *gin.Context) {
		id := c.Param("id")
		var o order
		if err := c.BindJSON(&o); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		updateOrder(id, o)
		c.JSON(http.StatusOK, o)
	})

	router.Run(":3003")
}
