package main

type carro struct {
	nome     string
	potencia int
	consumo  int
}

type ordenaPorPotencia []carro
type ordenaPorConsumo []carro
type ordenaPorLucro []carro

// Por Potência
func (x ordenaPorPotencia) Len() int {
	return len(x)
}

func (x ordenaPorPotencia) Less(i, j int) bool {
	return x[i].potencia < x[j].potencia
}

func (x ordenaPorPotencia) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

// Por Consumo
func (x ordenaPorConsumo) Len() int {
	return len(x)
}

func (x ordenaPorConsumo) Less(i, j int) bool {
	return x[i].consumo > x[j].consumo
}

func (x ordenaPorConsumo) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

// Por Lucro
func (x ordenaPorLucro) Len() int {
	return len(x)
}

func (x ordenaPorLucro) Less(i, j int) bool {
	return x[i].consumo < x[j].consumo
}

func (x ordenaPorLucro) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}
