package _interface

type Tank interface {
	Walk()
	Fire()
}
type Plane interface {
	Fly()
}

type PlaneTank interface {
	Tank
	Plane
}

type Printer interface {
	Print(interface{})
}

type FuncCaller func(p interface{})

func (funcCaller FuncCaller) Print(p interface{}) {
	funcCaller(p)
}
