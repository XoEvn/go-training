package closePackage

func createCounter(initial int) func() int {
	if initial < 0 {
		initial = 0
	}

	return func() int {
		initial++
		return initial
	}
}
