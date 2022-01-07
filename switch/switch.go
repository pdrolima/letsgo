package main

import "fmt"

func calculateIMC(value float64) string {
	switch {
	case value < 18.5:
		return "Underweight"
	case value < 25:
		return "Normal"
	case value < 30:
		return "Overweight"
	default:
		return "Obese"
	}
}

func main() {
	result := calculateIMC(29)
	fmt.Println("IMC:", result)
}
