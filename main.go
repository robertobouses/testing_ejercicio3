package main

import (
	"fmt"

	"github.com/robertobouses/testing_ejercicio3/user"
)

type CalculadoraIva interface {
	Calcular(float32) float32
}

func Eleccion() {
	var Calc CalculadoraIva

	fmt.Println("Indique la base imponible")
	var baseImponible float32
	fmt.Scanln(&baseImponible)

	fmt.Println("Indique el tipo de IVA general, reducido o superreducido(GEN/RED/SUP)")
	var tipoIva string
	fmt.Scanln(&tipoIva)
	switch tipoIva {
	case "GEN":
		Calc = user.Iva21{}
	case "RED":
		Calc = user.Iva10{}
	case "SUP":
		Calc = user.Iva4{}
	}

	Resultado := Calc.Calcular(baseImponible)
	fmt.Println(Resultado)

}

func main() {
	Eleccion()
}
