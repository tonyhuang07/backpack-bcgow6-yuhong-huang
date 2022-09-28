package main

import "fmt"

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

type Service struct {
	Name   string
	Price  float64
	Minute int
}

type Maintenance struct {
	Name  string
	Price float64
}

func calculatePriceProduct(products []Product) float64 {
	var total float64
	for _, product := range products {
		fmt.Println("process product: ", product.Name)
		total += product.Price * float64(product.Quantity)
	}
	return total
}

func calculatePriceService(services []Service) float64 {
	var total float64
	for _, service := range services {
		fmt.Println("process service: ", service.Name)
		total += service.Price * float64(service.Minute)
	}
	return total
}

func calculatePriceMaintenance(maintenances []Maintenance) float64 {
	var total float64
	for _, maintenance := range maintenances {
		fmt.Println("process maintenance: ", maintenance.Name)
		total += maintenance.Price
	}
	return total
}

func calculateTotalPrice(products []Product, services []Service, maintenances []Maintenance) float64 {
	priceProducts := make(chan float64)
	priceServices := make(chan float64)
	priceMaintenances := make(chan float64)

	go func() {
		priceProducts <- calculatePriceProduct(products)
	}()
	go func() {
		priceServices <- calculatePriceService(services)
	}()
	go func() {
		priceMaintenances <- calculatePriceMaintenance(maintenances)
	}()

	return <-priceProducts + <-priceServices + <-priceMaintenances
}

func main() {
	products := []Product{
		{Name: "000001", Price: 9000, Quantity: 1},
		{Name: "000002", Price: 3000, Quantity: 2},
		{Name: "000003", Price: 4000, Quantity: 4},
	}

	services := []Service{
		{Name: "000001", Price: 100, Minute: 30},
		{Name: "000002", Price: 200, Minute: 50},
		{Name: "000003", Price: 300, Minute: 60},
	}

	maintenances := []Maintenance{
		{Name: "000001", Price: 1000},
		{Name: "000002", Price: 2000},
	}

	total := calculateTotalPrice(products, services, maintenances)

	fmt.Println("Total price: ", total)
}
