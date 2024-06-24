package main

import "fmt"

func main() {

	var viper = User{"viper", "viper@vv.com", 23, 98765432}
	fmt.Println(viper)
	(&viper).ChangeEmail("viper2@vv.com")
	fmt.Println(viper)
	fmt.Println(&viper.Email)
	fmt.Println(*&viper.Email)
}

type User struct {
	Name  string
	Email string
	Age   int
	Phone int64
}

func (email *User) ChangeEmail(mail string) {
	email.Email = mail
}
