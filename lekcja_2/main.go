package main

import "fmt"

const DAWID_WIEK = 25

func main(){
	var imię string
	fmt.Println("Jak masz na imię?")
	fmt.Scanln(&imię)
	fmt.Printf("Cześć %s!\n", imię)

	var wiek int
	fmt.Println("A ile masz lat?")
	fmt.Scanln(&wiek)
	fmt.Printf("Super, to znaczy, że ja mam %d a ty %s masz %d\n", DAWID_WIEK, imię, wiek)
}