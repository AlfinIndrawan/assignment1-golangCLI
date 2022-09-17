package controllers

import (
	"net/http"

	"assignment-2/model"

	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetOrders(c *gin.Context) {
	orders, err := idb.Repository.GetOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}

func (idb *InDB) CreateOrder(c *gin.Context) {
	var order model.Order
	c.BindJSON(&order)
	newOrder, err := idb.Repository.Create(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, newOrder)
}

func (idb *InDB) GetOrder(c *gin.Context) {
	id := c.Param("id")
	order, err := idb.Repository.GetOrder(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, order)
}

func (idb *InDB) UpdateOrder(c *gin.Context) {
	var order model.Order
	c.BindJSON(&order)
	newOrder, err := idb.Repository.Update(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, newOrder)
}

func (idb *InDB) DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	err := idb.Repository.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
