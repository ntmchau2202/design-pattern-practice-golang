package hand

import "fmt"

type NormalHand struct {
	Hand
}

func (nh NormalHand) RaiseUp() {
	fmt.Println("This normal hand is raising up")
}

func (nh NormalHand) PutDown() {
	fmt.Println("This normal hand is putting down")
}

func (nh NormalHand) Wave() {
	fmt.Println("This normal hand is waving")
}
