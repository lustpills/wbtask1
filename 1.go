package main

import "fmt"

// Определяем структуру Human
type Human struct {
	Name string
	Age  int
}

// Метод для структуры Human
func (h *Human) Speak() {
	fmt.Printf("Меня зовут %s, мне %d лет.\n", h.Name, h.Age)
}

// Структура Action встраивает структуру Human
type Action struct {
	Human  // Встраивание структуры Human
	Action string
}

// Метод для структуры Action
func (a *Action) DoAction() {
	fmt.Printf("%s действует: %s.\n", a.Name, a.Action) // Обратите внимание: мы можем напрямую обращаться к полям Human
}

func main() {
	// Создаем экземпляр Action, автоматически "наследуя" поля и методы Human
	a := Action{
		Human:  Human{Name: "Алексей", Age: 30},
		Action: "бегает",
	}

	// Вызываем метод Speak от "родительской" структуры Human
	a.Speak()
	// Вызываем метод DoAction от структуры Action
	a.DoAction()
}
