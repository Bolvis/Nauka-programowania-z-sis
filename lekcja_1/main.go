package main

import "fmt"

func main(){
	var imięUżytkownika string
	fmt.Println("Jak masz na imię?")
	fmt.Scanln(&imięUżytkownika)

	var imięKolegi string
	fmt.Println("A jak na imię ma twój kolega?")
	fmt.Scan(&imięKolegi)

	fmt.Printf("Cześć %s!\nA kolega ma na imię %s\n", imięUżytkownika, imięKolegi)
}