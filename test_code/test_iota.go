package testcode

import "fmt"

type BikeType int

const (
	Nakid BikeType = iota + 1
	SuperSport
	American
	Harley
)

type BikeOption uint64

const (
	Mafler BikeOption = 1 << iota
	BackStep
	Cruchi
	Breaki
)

func TestIota() {
	//Enum for BikeType
	fmt.Println("testIota")
	fmt.Println(Nakid)
	fmt.Println(SuperSport)
	fmt.Println(American)
	fmt.Println(Harley)
	t := SuperSport
	fmt.Println(t)

	// Enum for BikeOption
	var o = Cruchi | BackStep
	fmt.Println(o)

	if o&Breaki != 0 {
		fmt.Println("ハンドル周り")
	}
}
