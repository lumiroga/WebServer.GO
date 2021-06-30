package main

import "fmt"

type animal interface {
	move() string
}

type dog struct {
}

type fish struct {
}

type bird struct {
}

func (dog) move() string {
	return "Mover al perro"
}
func (fish) move() string {
	return "Hacer nada al pez"
}

func moveAnimal(a animal) {
	fmt.Println(a.move())
}

func main() {
	p := dog{}
	f := fish{}
	moveAnimal(p)
	moveAnimal(f)

}
