package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Объявляем переменные для хранения вводимых значений
	var a, i, OneZero int64

	// Считываем три целых числа: число a, позицию бита i и значение бита OneZero (0 или 1)
	fmt.Scanf("%d %d %d", &a, &i, &OneZero)

	// Выводим двоичное представление числа a
	fmt.Println(strconv.FormatInt(a, 2))

	// Создаем маску с единицей на позиции i-1
	mask1 := int64(1) << (i - 1)
	// Выводим маску в двоичном формате
	fmt.Println(strconv.FormatInt(mask1, 2))

	// Создаем обратную маску с нулем на позиции i-1
	mask2 := ^mask1
	// Выводим обратную маску в двоичном формате
	fmt.Println(strconv.FormatInt(mask2, 2))

	// Модифицируем число a, очищая бит на позиции i и устанавливая его в значение OneZero
	a = (a & mask2) | (mask1 * OneZero)

	// Выводим модифицированное число a в двоичной системе
	fmt.Println(strconv.FormatInt(a, 2))
}
