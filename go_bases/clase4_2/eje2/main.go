package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"reflect"
)

type Client struct {
	ID        string
	Name      string
	DNI       string
	Cellphone string
	Direction string
}

const filePath = "clients.csv"
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func panicHandler() {
	rec := recover()
	if rec != nil {
		fmt.Println("RECOVER", rec)
	}
}

func generateID(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func verifyClient(client Client) bool {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	} else {
		reader := csv.NewReader(file)
		clients, _ := reader.ReadAll()
		for _, c := range clients {
			if c[1] == client.Name {
				return false
			}
		}
	}
	return true
}

func addClient(client Client) {
	fields := reflect.VisibleFields(reflect.TypeOf(client))
	emptyfields := make([]string, 0)
	defer panicHandler()
	for _, field := range fields {
		if field.Name == "" {
			emptyfields = append(emptyfields, field.Name)
		}
	}
	if len(emptyfields) > 0 {
		var result string
		for _, field := range emptyfields {
			result = result + field + ", "
		}
		panic("Empty fields: " + result)

	}

	verified := verifyClient(client)
	if !verified {
		panic("Client already exists")
	}

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.Write([]string{client.ID, client.Name, string(client.DNI),
		string(client.Cellphone), client.Direction})

}

func main() {
	fmt.Println("starting...")
	client := Client{
		ID:        generateID(10),
		Name:      "Juan",
		DNI:       "12345678",
		Cellphone: "123456789",
		Direction: "Av. Siempre Viva 123"}

	addClient(client)

	client2 := Client{
		ID:        generateID(10),
		Name:      "Juan",
		DNI:       "",
		Cellphone: "123456789",
		Direction: "Av. Siempre Viva 123"}

	addClient(client2)

	fmt.Println("ending...")
	fmt.Println("have detected several errors, but the program has not stopped")
	fmt.Println("the program has recovered from the errors and didn't left any file open")

}
