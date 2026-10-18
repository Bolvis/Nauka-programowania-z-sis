package main

import "fmt"

func main() {
	fmt.Println("Witaj w kalkulatorze!!!")
	fmt.Println("Chcesz dodać czy odjąć? (+/-)")
	var znak string
	fmt.Scanln(&znak)
	if znak == "+" {
		fmt.Println("Wybrałeś dodawanie")
		var a, b int
		fmt.Println("Podaj pierwszą liczbę:")
		fmt.Scanln(&a)
		fmt.Println("Podaj drugą liczbę:")
		fmt.Scanln(&b)
		fmt.Printf("Wynik dodawania: %df\n", a+b)
	} else if znak == "-" {
		fmt.Println("Wybrałeś odejmowanie")
		var a, b int
		fmt.Println("Podaj pierwszą liczbę:")
		fmt.Scanln(&a)
		fmt.Println("Podaj drugą liczbę:")
		fmt.Scanln(&b)
		fmt.Printf("Wynik odejmowania: %d\n", a-b)
	} else {
		fmt.Println("Niepoprawny znak")
	}

}
