package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	//rand.Seed(time.Now().UTC().UnixNano())
	//
	//const points = 1_000_000
	//
	//fmt.Printf("Estimating pi with %d point(s).\n\n", points)
	//
	//var sucess int
	//
	//for i := 0; i < points; i++ {
	//	x, y := genRandomPoint()
	//	if x*x+y*y < 1 {
	//		sucess++
	//	}
	//}
	//
	//piApprox := 4.0 * (float64(sucess) / float64(points))
	//errorPct := 100.0 * math.Abs(piApprox-math.Pi) / math.Pi
	//
	//fmt.Printf("Estimated pi: %9f \n", piApprox)
	//fmt.Printf("pi: %9f \n", math.Pi)
	//fmt.Printf("Error: %9f%%\n", errorPct)

	////===============================================
	//rand.Seed(time.Now().UTC().UnixNano())
	//
	//const (
	//	e           = math.E
	//	experiments = 1_000_000
	//)
	//
	//fmt.Printf("Estimating e with %d experiment(s).\n\n", experiments)
	//
	//var acc int
	//for i := 0; i < experiments; i++ {
	//	var (
	//		sum         float64
	//		num2Success int
	//	)
	//	for sum <= 1 {
	//		n := rand.Float64()
	//		sum += n
	//		num2Success++
	//	}
	//
	//	acc += num2Success
	//}
	//
	//expected := float64(acc) / float64(experiments)
	//
	//errorPct := 100.0 * math.Abs(expected-e) / e
	//
	//fmt.Printf("Expected vale: %9f \n", expected)
	//fmt.Printf("e: %9f \n", e)
	//fmt.Printf("Error: %9f%%\n", errorPct)

	//=================================================
	//rand.Seed(time.Now().UTC().UnixNano())
	//
	//const trials = 1_000_000
	//var (
	//	numPeople = 23
	//	success   = 0
	//)
	//for i := 0; i < trials; i++ {
	//	bdays :=genBdayList(numPeople)
	//	uniques :=uniqueSlice(bdays)
	//
	//	if len(uniques)!=numPeople{
	//		success++
	//	}
	//}
	//
	//probability :=float64(success)/float64(trials)
	//fmt.Printf("The probability of at least 2 persons in a group of %d people sharing a brithday is %.2f%%\n",numPeople,100.0*probability)
	//===============================================
	//rand.Seed(time.Now().UTC().UnixNano())
	//
	//const games = 1_000_000
	//
	//fmt.Printf("Estimating the probability of winning by switching doors with %d games(s).\n\n", games)
	//
	//var sucess int
	//
	//for i := 0; i < games; i++ {
	//	newDoor, prizeDoor := setMontyHallGame()
	//	if newDoor == prizeDoor {
	//		sucess++
	//	}
	//}
	//probability := float64(sucess) / float64(games)
	//const theoreticalValue = 2. / 3.
	//errorPct := 100. * math.Abs(probability-theoreticalValue) / theoreticalValue
	//
	//fmt.Printf("Estimated probability: %9f\n", probability)
	//fmt.Printf("Theoretical value: %9f\n", theoreticalValue)
	//fmt.Printf("Error: %9f%%\n", errorPct)




	var n uint64 = math.MaxUint64 // line A
	//n := math.MaxUint64 // line B
	fmt.Println(n)
}

func setMontyHallGame() (int, int) {
	var (
		montysChoice int
		prizeDoor    int
		goat1Door    int
		goat2Dorr    int
		newDoor      int
	)

	guestDoor := rand.Intn(3)

	areDoorSelected := false
	for !areDoorSelected {
		prizeDoor = rand.Intn(3)
		goat1Door = rand.Intn(3)
		goat2Dorr = rand.Intn(3)
		if prizeDoor != goat1Door && prizeDoor != goat2Dorr && goat1Door != goat2Dorr {
			areDoorSelected = true
		}
	}
	showGoat := false
	for !showGoat {
		montysChoice = rand.Intn(3)
		if montysChoice != prizeDoor && montysChoice != guestDoor {
			showGoat = true
		}
	}

	madeSwitch := false

	for !madeSwitch {
		newDoor = rand.Intn(3)
		if newDoor != guestDoor && newDoor != montysChoice {
			madeSwitch = true
		}
	}

	return newDoor, prizeDoor

}

func uniqueSlice(s []int) []int {
	keys := make(map[int]struct{})
	list := []int{}
	for _, entry := range s {
		if _, ok := keys[entry]; !ok {
			keys[entry] = struct{}{}
			list = append(list, entry)
		}
	}
	return list
}

func genBdayList(n int) []int {
	var bdays = make([]int, n)
	for i := 0; i < n; i++ {
		bdays[i] = rand.Intn(365)
	}
	return bdays
}

func genRandomPoint() (x, y float64) {
	x = 2.0*rand.Float64() - 1.0
	y = 2.0*rand.Float64() - 1.0
	return x, y
}
