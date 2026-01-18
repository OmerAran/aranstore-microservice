package service

import (
	"customer-service/db"
	"customer-service/service/model"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func createCustomer(db *db.PostgresDb, c *gin.Context) (int, *model.Customer, error) {
	var customer model.Customer

	if err := c.ShouldBindJSON(&customer); err != nil {
		return http.StatusBadRequest, nil, err
	}

	if len(customer.Email) == 0 {
		return http.StatusBadRequest, nil, fmt.Errorf("email is required")
	}

	query := `
	INSERT INTO customers (name, email, address, status)
	VALUES ($1, $2, $3, $4)
	RETURNING id, status
	`
	err := db.DB.QueryRow(query, customer.Name, customer.Email, customer.Address, customer.Status).Scan(&customer.Id, &customer.Status)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, &customer, nil
}

func getCustomer(db *db.PostgresDb, c *gin.Context) (int, *model.Customer, error) {
	customerId := c.Param("customerId")
	id, err := strconv.Atoi(customerId)
	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	var customer model.Customer
	stmt := `SELECT id, name, email, address, status FROM customers WHERE id = $1 and status != -1`
	err = db.DB.QueryRow(stmt, id).Scan(&customer.Id, &customer.Name, &customer.Email, &customer.Address, &customer.Status)

	if err == sql.ErrNoRows {
		return http.StatusNotFound, nil, err
	}

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, &customer, nil
}

func updateCustomer(db *db.PostgresDb, c *gin.Context) (int, *model.Customer, error) {
	//I will implement this later
	return http.StatusNotImplemented, nil, nil
}

func deleteCustomer(db *db.PostgresDb, c *gin.Context) (int, error) {
	customerID := c.Param("customerId")
	id, err := strconv.Atoi(customerID)
	if err != nil {
		return http.StatusBadRequest, err
	}

	stmt := `UPDATE customers SET status = -1 WHERE id = $1`
	_, err = db.DB.Exec(stmt, id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusNoContent, nil
}
