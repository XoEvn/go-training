package accpetStuct

import "fmt"

type Person struct {
	Name  string
	Birth string
	ID    int64
}

func (person *Person) changeName(name string) {
	person.Name = name
}

func (person Person) printMess() {
	fmt.Printf("My name is %v, and my birthday is %v, and my id is %v\n", person.Name, person.Birth, person.ID)
	person.ID = 3
}
