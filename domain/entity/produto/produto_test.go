package produto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var categoria = Categoria{Id: "123", Name: "Filmadora"}
var fabricante = Fabricante{Id: "123", Name: "TecnoMania"}

func TestIfProductIsSell(t *testing.T) {
	var product = NewProduct("Filmadora",
		"TekPix, a câmera mais vendida do Brasil",
		10, 1000, categoria, fabricante)
	product.Sell()
	assert.Equal(t, 9, product.Quant)
}
func TestIfProductIsNotSell(t *testing.T) {
	var product = NewProduct("Filmadora",
		"TekPix, a câmera mais vendida do Brasil",
		0, 1000, categoria, fabricante)
	err := product.Sell()
	assert.NotNil(t, err)
	assert.Equal(t, "Quantidade insuficiente para ser vendida",
		err.Error())
}

func TestIfProductWasBuy(t *testing.T) {
	var product = NewProduct("Filmadora",
		"TekPix, a câmera mais vendida do Brasil",
		66, 1000, categoria, fabricante)
	product.Buy(1000)
	assert.Equal(t, 1066, product.Quant)
}
