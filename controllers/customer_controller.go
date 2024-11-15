package controllers

import (
	"customer-management/config"
	"customer-management/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

var customerCollection = config.DB.Collection("customers")

// Get all customers
func GetCustomers(c *gin.Context) {
	var customers []models.Customer
	cursor, err := customerCollection.Find(c, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(c)
	for cursor.Next(c) {
		var customer models.Customer
		cursor.Decode(&customer)
		customers = append(customers, customer)
	}
	c.JSON(http.StatusOK, customers)
}

// Add customer
func AddCustomer(c *gin.Context) {
	var customer models.Customer
	if err := c.BindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	customerCollection.InsertOne(c, customer)
	c.JSON(http.StatusCreated, customer)
}

// Update customer
func UpdateCustomer(c *gin.Context) {
	id := c.Param("id")
	var customer models.Customer
	if err := c.BindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	customerCollection.UpdateOne(c, bson.M{"id": id}, bson.M{"$set": customer})
	c.JSON(http.StatusOK, gin.H{"message": "Customer updated"})
}

// Delete customer
func DeleteCustomer(c *gin.Context) {
	id := c.Param("id")
	customerCollection.DeleteOne(c, bson.M{"id": id})
	c.JSON(http.StatusOK, gin.H{"message": "Customer deleted"})
}
