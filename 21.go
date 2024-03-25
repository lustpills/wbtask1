package main

import (
	"fmt"
)

// OldPrinter представляет старый принтер с методом Print
type OldPrinter struct{}

func (op *OldPrinter) Print(s string) {
	fmt.Println("Printing:", s)
}

// Printer интерфейс определяет метод PrintMessage
type Printer interface {
	PrintMessage(string)
}

// Adapter структура адаптера, который преобразует OldPrinter в Printer
type Adapter struct {
	OldPrinter *OldPrinter
}

func (a *Adapter) PrintMessage(msg string) {
	a.OldPrinter.Print(msg)
}

func main() {
	oldPrinter := &OldPrinter{}

	adapter := &Adapter{
		OldPrinter: oldPrinter,
	}

	adapter.PrintMessage("Hello, world!")
}
