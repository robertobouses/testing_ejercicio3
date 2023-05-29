package user_test

import (
	"testing"

	"github.com/robertobouses/testing_ejercicio3/user"
)

func TestIva10(t *testing.T) {
	var expected float32
	var result float32
	var calc user.Iva10
	result = calc.Calcular(100.0)
	expected = 10
	if result != expected {
		t.Errorf("Error, cálculo Iva: se esperaba %f, se recibe %f", expected, result)
	}

}
