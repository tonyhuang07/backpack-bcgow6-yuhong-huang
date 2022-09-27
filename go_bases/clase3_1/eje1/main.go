package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

const FilePath = "../products.csv"

type Product struct {
	ID    string
	Price float64
	Cant  int
}

func registerProduct(products []Product) {
	file, err := os.OpenFile(FilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("failed to open file: ", err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	for _, product := range products {
		line := []string{product.ID, strconv.FormatFloat(product.Price, 'f', 2, 64), strconv.Itoa(product.Cant)}
		err := writer.Write(line)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	products := []Product{
		{ID: "000001", Price: 32234, Cant: 1},
		{ID: "000002", Price: 343.6, Cant: 2},
		{ID: "000003", Price: 132333, Cant: 1},
	}
	registerProduct(products)
}
