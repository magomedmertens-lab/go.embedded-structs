package main

import "fmt"

type Human struct {
    Name string
    Age  int
}

func (h Human) Speak() {
    fmt.Printf("Меня зовут %s, мне %d лет.\n", h.Name, h.Age)
}

func (h Human) Walk() {
    fmt.Printf("%s идёт пешком.\n", h.Name)
}

type Action struct {
    Human
    Power int
}

func (a Action) Fight() {
    fmt.Printf("%s атакует с силой %d!\n", a.Name, a.Power)
}

func main() {
    a := Action{
        Human: Human{Name: "Иван", Age: 25},
        Power: 100,
    }

    a.Speak()
    a.Walk()
    a.Fight()
}
