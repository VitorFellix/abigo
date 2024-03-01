package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type transaction struct {
	ID          string  `json:"id"`
	Category    string  `json:"cat"`
	Description string  `json:"desc"`
	Account     string  `json:"acc"`
	Value       float32 `json:"val"`
}

var transactions = []transaction{
	{ID: "1", Category: "Transporte", Description: "Gasolina", Account: "BB", Value: 200},
	{ID: "2", Category: "Encontro", Description: "Taisho", Account: "BB", Value: 140},
}

func insertTransaction(t transaction) bool {
	id := t.ID
	for _, transaction := range transactions {
		if transaction.ID == id {
			return false
		}
	}
	transactions = append(transactions, t)
	return true
}

func getTransactions(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, transactions)
}

func postTransactionFromFile(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"context": c})
}

func postTransaction(c *gin.Context) {
	var newTransations transaction
	c.BindJSON(&newTransations)

	if !insertTransaction(newTransations) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "transaction id already exists", "transaction": newTransations})
		return
	}
	c.IndentedJSON(http.StatusOK, newTransations)
}

func getTransactionByID(c *gin.Context) {
	id := c.Param("id")

	for _, transaction := range transactions {
		if transaction.ID == id {
			c.IndentedJSON(http.StatusOK, transaction)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "transaction not"})
}

func main() {
	router := gin.Default()
	router.GET("/transactions", getTransactions)
	router.GET("/transactions/:id", getTransactionByID)
	router.POST("/transactions", postTransaction)
	router.POST("/transactions/fromFile", postTransactionFromFile)

	router.Run("localhost:8080")
}

// curl localhost:8080/transactions
// curl localhost:8080/transactions/1
// curl localhost:8080/transactions --include --request "POST" --header "Content-Type: application/json" --data '{"id":"3","cat":"Transporte","acc":"BB","val": 200,"desc":"Gasolina"}'
