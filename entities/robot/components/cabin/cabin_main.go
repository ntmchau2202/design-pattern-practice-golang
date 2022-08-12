package cabin

import "fmt"

type MainCabin struct {
	Cabin
}

func (mc MainCabin) Open() {
	fmt.Println("The main cabin is openned")
}

func (mc MainCabin) Close() {
	fmt.Println("The main cabin is closed")
}

func (mc MainCabin) Load() {
	fmt.Println("The main cabin is loading")
}

func (mc MainCabin) Unload() {
	fmt.Println("The main cabin is unloading")
}
