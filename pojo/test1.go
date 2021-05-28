package pojo

type Person struct {
	Name  string
	Birth string
	ID    int64
}

func test1() {
	var p1 Person
	p1.Name = "王小二"
	p1.Birth = " 1990-12-11"

	p2 := new(Person)

	p2.Name = "王小二"
	p2.Birth = "1990-12-23"

	p3 := &Person{}
	p3.Name = "王小二"
	p3.Birth = "1990-12-23"

	p4 := Person{
		Name:  "王雄武",
		Birth: "1990-12-20",
	}

	p5 := &Person{
		"wangwu",
		"1990-12-20",
		5,
	}

}
