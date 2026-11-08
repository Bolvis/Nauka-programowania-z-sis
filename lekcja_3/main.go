package main

import "fmt"

func main() {
	fmt.Println("Witaj w kalkulatorze!!!")
	var a, b int
	fmt.Println("Podaj pierwszą liczbę:")
	fmt.Scanln(&a)
	fmt.Println("Podaj drugą liczbę:")
	fmt.Scanln(&b)
	fmt.Println("Chcesz dodać czy odjąć? (+/-)")
	var znak string
	fmt.Scanln(&znak)
	switch znak {
	case "+":
		fmt.Printf("Wynik dodawania: %df\n", a+b)
	case "-":
		fmt.Printf("Wynik odejmowania: %d\n", a-b)
	default:
		fmt.Println("Niepoprawny znak")
	}

}
