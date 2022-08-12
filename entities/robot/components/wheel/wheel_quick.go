package wheel

import "fmt"

type QuickWheel struct {
	Wheel
}

func (qw QuickWheel) Run() {
	fmt.Println("The quick wheel runs quick...")
}
