package user

type Iva10 struct {
}

func (I Iva10) Calcular(valor float32) float32 {
	return valor * 0.10
}
