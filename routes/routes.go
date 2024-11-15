package routes

import (
	"customer-management/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/customers", controllers.GetCustomers)
	router.POST("/customers", controllers.AddCustomer)
	router.PUT("/customers/:id", controllers.UpdateCustomer)
	router.DELETE("/customers/:id", controllers.DeleteCustomer)

	return router
}
