package sturctInterface

import "fmt"

type Cat interface {
	CatchMouse()
}

type Dog interface {
	Bark()
}

type CatDog struct {
	Name string
}

func (cateDog *CatDog) CatchMouse() {
	fmt.Printf("%v caught the mouse and ate it!\n", cateDog.Name)
}

func (cateDog *CatDog) Bark() {
	fmt.Printf("%v barked loudly!\n", cateDog.Name)
}
