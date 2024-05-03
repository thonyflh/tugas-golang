package main

import (
	"log"

	"github.com/thonyflh/tugas-golang/app"
)

func main() {
	calculator := &app.Calculator{}

	result := calculator.Multiply(4.5, 3)

	log.Printf("Hasil perkalian: %.2f", result)

	if isValid, err := calculator.IsInteger("1"); err != nil {
		log.Println(isValid, err.Error())
	} else {
		log.Println(isValid)
	}

	log.Println(calculator.IntToDayName(4))

	log.Println(calculator.DayNameToInt("selasa"))

	var kosong string
	calculator.IsiHello(&kosong)
	log.Println(kosong)
}
