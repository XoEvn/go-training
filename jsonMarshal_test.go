package training

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Stu struct {
	Name string `json:"name"`
	Age int
	HIgh bool
	sex string
	Class *Class `json:"class"`
}

type Class struct {
	Name string
	Grade int
}


func Test01(t *testing.T){
	stu := Stu{
		Name: "张三",
		Age :18,
		HIgh: true,
		sex: "男",
	}

	cla :=new(Class)
	cla.Name = "1班"
	cla.Grade = 3
	stu.Class=cla
	jsonStu,err :=json.Marshal(stu)
	if err!=nil {
		fmt.Println("生成json字符串错误")
	}

	//jsonStu是[]byte类型，转化成string类型便于查看
	fmt.Println(string(jsonStu))
}
