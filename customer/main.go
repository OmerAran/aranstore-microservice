package main

import (
	"customer-service/db"
	service "customer-service/service"

	"github.com/gin-gonic/gin"
)

func main() {

	db := db.GetDBConnection()
	a := service.GetApp(db)

	r := gin.Default()

	r.GET("/customers/:customerId", a.GetHandler)

	r.POST("/customers", a.PostHandler)

	r.PUT("/customers/:customerId", a.PutHandler)

	r.DELETE("/customers/:customerId", a.DeleteHandler)

	r.Run("localhost:8080")
}
