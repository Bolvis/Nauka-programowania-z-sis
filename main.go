package main

import (
	"fmt"
	"log"
)

const DAWID_WIEK = 25

func main() {
	var imię string
	fmt.Println("Jak masz na imię?")
	fmt.Scanln(&imię)
	fmt.Printf("Cześć %s!\n", imię)

	var wiek int
	fmt.Println("A ile masz lat?")
	fmt.Scanln(&wiek)

	var dawidMłodszy bool = czyDawidMłodszy(wiek)
	var dawidWTymSamymWieku bool = czyDawidWTymSamymWieku(wiek)

	if dawidWTymSamymWieku {
		fmt.Printf("Super, mamy tyle samo lat, czyli %d\n", wiek)
	} else {
		var różnicaWieku int = różnicaWieku(wiek, dawidMłodszy)
		var lat string = odmianaGramatycznaLat(różnicaWieku)
		if dawidMłodszy {
			fmt.Printf("Oho, staruszek nam się trafił, masz o %d %s więcej ode mnie!\n", różnicaWieku, lat)
		} else {
			fmt.Printf("Super,to znaczy, że masz o %d %s mniej ode mnie :)\n", różnicaWieku, lat)
		}
	}

}

func czyDawidMłodszy(wiekUżytkownika int) bool {
	return DAWID_WIEK < wiekUżytkownika
}

func czyDawidWTymSamymWieku(wiekUżytkownika int) bool {
	return DAWID_WIEK == wiekUżytkownika
}

func różnicaWieku(wiekUżytkownika int, dawidMłodszy bool) int {
	if dawidMłodszy {
		return wiekUżytkownika - DAWID_WIEK
	}

	return DAWID_WIEK - wiekUżytkownika
}

func odmianaGramatycznaLat(ilośćLat int) string {
	if ilośćLat <= 0 || ilośćLat > 999 {
		log.Printf("Niewspierana liczba: %d | dozwolony zakres jest między 1 a 999", ilośćLat)
		return "PODANA BŁĘDNA LICZBA"
	}

	if ilośćLat == 1 {
		return "rok"
	}

	var ostatniaCyfra int
	var nast_ostatnia int
	if ilośćLat < 100 {
		ostatniaCyfra = ilośćLat % 10
		nast_ostatnia = (ilośćLat / 10) % 10
	} else if ilośćLat >= 100 {
		ostatniaCyfra = ilośćLat % 100 % 10
		nast_ostatnia = (ilośćLat / 100) % 10
	}

	if nast_ostatnia != 1 && (ostatniaCyfra >= 2 && ostatniaCyfra <= 4) {
		return "lata"
	}

	return "lat"
}
