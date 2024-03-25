package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Создаем переменные типа big.Int для хранения больших чисел

	a := new(big.Int)
	a.SetString("12345678901234567890", 10)
	b := new(big.Int)
	b.SetString("12345678901234567890", 10)
	//b := big.NewInt(98765432109876543210)

	// Перемножение a и b
	mulResult := new(big.Int).Mul(a, b)
	fmt.Println("Умножение a и b:", mulResult)

	// Деление a на b
	divResult := new(big.Int).Div(a, b)
	fmt.Println("Деление a на b:", divResult)

	// Сложение a и b
	addResult := new(big.Int).Add(a, b)
	fmt.Println("Сложение a и b:", addResult)

	// Вычитание b из a
	subResult := new(big.Int).Sub(a, b)
	fmt.Println("Вычитание b из a:", subResult)
}
