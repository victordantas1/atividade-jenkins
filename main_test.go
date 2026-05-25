package main

import "testing"

func TestCelsiusParaFahrenheit(t *testing.T) {
	resultado := CelsiusParaFahrenheit(0)
	esperado := 32.0

	if resultado != esperado {
		t.Errorf("Resultado incorreto. Esperado: %f, Recebido: %f", esperado, resultado)
	}
}

func TestFahrenheitParaCelsius(t *testing.T) {
	resultado := FahrenheitParaCelsius(32)
	esperado := 0.0

	if resultado != esperado {
		t.Errorf("Resultado incorreto. Esperado: %f, Recebido: %f", esperado, resultado)
	}
}
