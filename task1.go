package main

import (
	"fmt"
	"strings"
)

type Human struct {
	Name    string
	Surname string
	Age     int
}

func (p Human) FullName() string {
	x := strings.Builder{}
	x.WriteString(p.Name)
	x.WriteString(" ")
	x.WriteString(p.Surname)
	return x.String()
}

type Action struct {
	WakeUp bool
	Human  //наследуем Human с его методом FullName
}

func main() {
	stas := Action{
		WakeUp: true,
		Human:  Human{Name: "Stanislav", Surname: "Krokhalev", Age: 27},
	}
	fmt.Println(stas.FullName()) // Вызывается метод FullName структуры Human через экземпляр структуры Action
}
