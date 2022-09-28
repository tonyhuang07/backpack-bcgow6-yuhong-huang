package main

import "fmt"

type User struct {
	Name     string
	Age      int
	Email    string
	Password string
}

func updateName(user *User, name string) error {
	if user == nil {
		return fmt.Errorf("user is nil")
	}
	user.Name = name
	return nil
}

func updateAge(user *User, age int) error {
	if user == nil {
		return fmt.Errorf("user is nil")
	}
	user.Age = age
	return nil
}
func updateEmail(user *User, email string) error {
	if user == nil {
		return fmt.Errorf("user is nil")
	}
	user.Email = email
	return nil
}

func updatePassword(user *User, password string) error {
	if user == nil {
		return fmt.Errorf("user is nil")
	}
	user.Password = password
	return nil
}

func main() {
	var user User
	user.Name = "Tony"
	user.Age = 20
	user.Email = "tony@meli.com"
	user.Password = "123456"

	fmt.Println("user before update: ", user)

	updateName(&user, "Tony1")
	updateAge(&user, 21)
	updateEmail(&user, "tony1@meli.com")
	updatePassword(&user, "1234567")

	fmt.Println("user after update: ", user)

}
