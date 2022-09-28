package main

import "fmt"

type Product struct {
	Name  string
	Price float64
}

type User struct {
	Name     string
	Age      int
	Email    string
	Password string
	Procuts  []Product
}

func newProduct(name string, price float64) Product {
	return Product{
		Name:  name,
		Price: price,
	}
}

func addProduct(user *User, product Product) error {
	if user == nil {
		return fmt.Errorf("user is nil")
	}
	user.Procuts = append(user.Procuts, product)
	return nil
}

func deleteProducts(user *User) error {
	if user == nil {
		return fmt.Errorf("user is nil")
	}
	user.Procuts = nil
	return nil
}

// func updateName(user *User, name string) error {
// 	if user == nil {
// 		return fmt.Errorf("user is nil")
// 	}
// 	user.Name = name
// 	return nil
// }

// func updateAge(user *User, age int) error {
// 	if user == nil {
// 		return fmt.Errorf("user is nil")
// 	}
// 	user.Age = age
// 	return nil
// }
// func updateEmail(user *User, email string) error {
// 	if user == nil {
// 		return fmt.Errorf("user is nil")
// 	}
// 	user.Email = email
// 	return nil
// }

// func updatePassword(user *User, password string) error {
// 	if user == nil {
// 		return fmt.Errorf("user is nil")
// 	}
// 	user.Password = password
// 	return nil
// }

func main() {
	var user User
	user.Name = "Tony"
	user.Age = 20
	user.Email = "tony@meli.com"
	user.Password = "123456"

	fmt.Println("user before adding product: ", user)

	var product1 = newProduct("product1", 100)
	fmt.Println("new product: ", product1)

	addProduct(&user, product1)
	fmt.Println("user after adding product: ", user)

	var product2 = newProduct("product2", 200)
	fmt.Println("new product: ", product2)

	addProduct(&user, product2)
	fmt.Println("user after adding product: ", user)

	deleteProducts(&user)
	fmt.Println("user after deleting all products: ", user)

}
