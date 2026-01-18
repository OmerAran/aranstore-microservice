package service

import (
	"customer-service/db"

	"github.com/gin-gonic/gin"
)

type App struct {
	db *db.PostgresDb
}

func GetApp(db *db.PostgresDb) *App {
	return &App{db: db}
}

func (a *App) PostHandler(c *gin.Context) {
	status, customer, err := createCustomer(a.db, c)
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	c.JSON(status, customer)
}

func (a *App) PutHandler(c *gin.Context) {
	status, customer, err := updateCustomer(a.db, c)
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	c.JSON(status, customer)

}

func (a *App) DeleteHandler(c *gin.Context) {
	status, err := deleteCustomer(a.db, c)
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	c.JSON(status, gin.H{"message": "customer deleted successfully"})

}

func (a *App) GetHandler(c *gin.Context) {
	status, customer, err := getCustomer(a.db, c)
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	c.JSON(status, customer)
}
