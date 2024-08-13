package main

import "fmt"

type Human struct {
	Name    string
	Surname string
	Age     int
}

func (p Human) FullName() string {
	return p.Name + " " + p.Surname
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
