package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	var newTransactions transaction
	c.BindJSON(&newTransactions)

	cat := c.DefaultQuery("cat", "sem categoria")
	acc := c.DefaultQuery("acc", "bb")

	newTransactions.Category = cat
	newTransactions.Account = acc

	if !insertTransaction(newTransactions) {
		c.IndentedJSON(http.StatusNotFound,
			gin.H{
				"message": "transaction id already exists",
				"transaction": newTransactions,
		})
		return
	}
	c.IndentedJSON(http.StatusOK, newTransactions)
}

func getTransactionByID(c *gin.Context) {
	id := c.Param("id")
	for _, transaction := range transactions {
		if transaction.ID == id {
			c.IndentedJSON(http.StatusOK, transaction)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "transaction not found."})
}

func main() {
	uri := os.Getenv("MONGODB_URI")

	log.Printf("using uri for mongo connection: %s\n", uri)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("sample_mflix").Collection("comments")
	name := "Mercedes Tyler"

	log.Printf("Querying for %s", name)

	var result bson.M
	err = coll.FindOne(context.TODO(), bson.D{{"name", name}}).
		Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("could not find document %s\n", name)
		return
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
	
	fmt.Printf("%s\n", queryMongo())

	// router := gin.Default()
	// router.GET("/transactions", getTransactions)
	// router.GET("/transactions/:id", getTransactionByID)
	// router.POST("/transactions", postTransaction)
	//
	// router.Run("localhost:8080")
}

func queryMongo() string {
	return "you queried me"
}

// curl "localhost:8080/transactions"
// curl "localhost:8080/transactions/1"
// curl "localhost:8080/transactions" --include --request "PORT" --header "content-type: application/json" --data '{"id":"3","cat":"transporte","acc":"bb","val": 200,"desc":"gasolina"}'
