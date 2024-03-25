package main

import (
	"fmt"
	"reflect"
)

// Функция checkType принимает переменную типа interface{} и определяет ее тип
func checkType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("Это переменная типа int")
	case string:
		fmt.Println("Это переменная типа string")
	case bool:
		fmt.Println("Это переменная типа bool")
	case chan int:
		fmt.Println("Это переменная типа channel int")
	default:
		fmt.Println("Неизвестный тип")
	}
}

func main() {
	var i int = 10
	var s string = "Hello"
	var b bool = true
	var c = make(chan int)

	// Проверяем типы переменных i, s, b, c
	checkType(i)
	checkType(s)
	checkType(b)
	checkType(c)

	// Пример с использованием пакета reflect для определения типа переменной x
	var x interface{} = make(chan int)
	switch reflect.TypeOf(x).Kind() {
	case reflect.Int:
		fmt.Println("Переменная x - int")
	case reflect.String:
		fmt.Println("Переменная x - string")
	case reflect.Chan:
		fmt.Println("Переменная x - chan")
	default:
		fmt.Println("Неизвестный тип переменной x")
	}
}
