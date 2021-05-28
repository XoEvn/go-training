package sturctInterface

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	t.Run("", func(t *testing.T) {
		catDog := &CatDog{
			"Lucy",
		}

		var cat Cat
		cat = catDog
		cat.CatchMouse()

		var dog Dog
		dog = catDog
		dog.Bark()
	})

	t.Run("", func(t *testing.T) {
		type Wheel struct {
			shape string
		}

		type Car struct {
			Wheel
			Name string
		}

		car := &Car{
			Wheel{
				"圆形的",
			},
			"福特",
		}
		fmt.Println(car.Name, car.shape, " ")
	})

	t.Run("", func(t *testing.T) {

		wild := WildDuck{}
		wild.fly()
		wild.swim()

		domestic := DomesticDuck{}
		domestic.swim()
	})

}

type Swimming struct {
}

func (swim *Swimming) swim() {
	fmt.Println("swimming is my ability")
}

type Flying struct {
}

func (fly *Flying) fly() {
	fmt.Println("flying is my ability")
}

type WildDuck struct {
	Swimming
	Flying
}

type DomesticDuck struct {
	Swimming
}
