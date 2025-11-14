package main

type Transaction struct {
	ID          string  `json:"id"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
	Account     string  `json:"account"`
	Value       float32 `json:"value"`
}

var Transactions = []Transaction{
	{ID: "1", Category: "Transporte", Description: "Gasolina", Account: "BB", Value: 200},
	{ID: "2", Category: "Encontro", Description: "Taisho", Account: "BB", Value: 140},
}

