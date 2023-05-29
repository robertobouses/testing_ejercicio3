package user_test

import (
	"testing"

	"github.com/robertobouses/testing_ejercicio3/user"
)

func TestIva4(t *testing.T) {
	var expected float32
	calc := user.Iva4{}
	result := calc.Calcular(100.0)
	expected = 4.0
	if result != expected {
		t.Errorf("Cálculo de Iva incorrecto. Se esperaba %f, se obtuvo %f", expected, result)
	}
}
