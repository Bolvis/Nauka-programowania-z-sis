package main

import "fmt"

func main(){
	var imię string
	fmt.Println("Jak masz na imię?")
	fmt.Scanln(&imię)
	fmt.Printf("Cześć %s!\n", imię)
}