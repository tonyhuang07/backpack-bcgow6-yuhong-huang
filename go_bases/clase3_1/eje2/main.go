package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

const FilePath = "../products.csv"

func main() {
	file, err := os.Open(FilePath)
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(file)
	products, _ := reader.ReadAll()
	var total float64
	for _, product := range products {
		price, _ := strconv.ParseFloat(product[1], 64)
		cant, _ := strconv.Atoi(product[2])
		total = total + (price * float64(cant))
		fmt.Println(product[0] + "\t" + product[1] + "\t" + product[2] + "\t")
	}
	fmt.Println("Total: " + strconv.FormatFloat(total, 'f', 2, 64))
}
