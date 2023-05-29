package user_test

import (
	"testing"

	"github.com/robertobouses/testing_ejercicio3/user"
)

func TestIva21(t *testing.T) {
	var expected float32
	expected = 21.0
	calc := user.Iva21{}
	result := calc.Calcular(100)

	if expected != result {
		t.Errorf("Error al calcular IVA General. Se esperaba %f, se recibe %f", expected, result)
	}
}
