package accpetStuct

import "testing"

func Test1(t *testing.T) {
	t.Run("", func(t *testing.T) {
		p1 := Person{
			Name:  "wangxiaoer",
			Birth: "1990-12-32",
			ID:    1,
		}

		p1.printMess()
		p1.changeName("网老热")
		p1.printMess()
	})
}
