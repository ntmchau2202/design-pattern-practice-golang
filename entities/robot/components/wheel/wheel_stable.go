package wheel

import "fmt"

type StableWheel struct {
	Wheel
}

func (sw StableWheel) Run() {
	fmt.Println("The stable wheel runs slowly...")
}
