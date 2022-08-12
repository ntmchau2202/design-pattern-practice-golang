package cabin

import "fmt"

type SubCabin struct {
	Cabin
}

func (sc SubCabin) Open() {
	fmt.Println("The sub cabin is openned")
}

func (sc SubCabin) Close() {
	fmt.Println("The sub cabin is closed")
}

func (sc SubCabin) Load() {
	fmt.Println("The sub cabin is loading")
}

func (sc SubCabin) Unload() {
	fmt.Println("The sub cabin is unloading")
}
