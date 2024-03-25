package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(15, 15)

	if total != 33 {
		t.Errorf("resultado da soma é invalido: Resultado %d e o valor esperado era:%d", total, 30)
	}
}
