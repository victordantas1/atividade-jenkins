package main

import "fmt"

// CelsiusParaFahrenheit converte C para F
func CelsiusParaFahrenheit(c float64) float64 {
	return (c * 9 / 5) + 32
}

// FahrenheitParaCelsius converte F para C
func FahrenheitParaCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

func main() {
	fmt.Printf("25°C em Fahrenheit é: %.2f°F\n", CelsiusParaFahrenheit(25))
	fmt.Printf("77°F em Celsius é: %.2f°C\n", FahrenheitParaCelsius(77))
}
