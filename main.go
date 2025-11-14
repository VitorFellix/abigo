package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func insertTransaction(t Transaction) bool {
	id := t.ID
	for _, transaction := range Transactions {
		if transaction.ID == id {
			return false
		}
	}
	Transactions = append(Transactions, t)
	return true
}

func getTransactions(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Transactions)
}

func postTransaction(c *gin.Context) {
	var newTransations Transaction
	c.BindJSON(&newTransations)

	if !insertTransaction(newTransations) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "transaction id already exists", "transaction": newTransations})
		return
	}
	c.IndentedJSON(http.StatusOK, newTransations)
}

func getTransactionByID(c *gin.Context) {
	id := c.Param("id")

	for _, transaction := range Transactions {
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

	router.Run("localhost:8080")
}

// curl localhost:8080/transactions
// curl localhost:8080/transactions/1
// curl localhost:8080/transactions --include --request "POST" --header "Content-Type: application/json" --data '{"id":"3","cat":"Transporte","acc":"BB","val": 200,"desc":"Gasolina"}'
