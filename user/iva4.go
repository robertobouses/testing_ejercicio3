package user

type Iva4 struct {
}

func (I Iva4) Calcular(valor float32) float32 {
	return valor * 0.04
}
