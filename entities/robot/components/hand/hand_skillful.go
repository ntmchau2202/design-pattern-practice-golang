package hand

import "fmt"

type SkillfulHand struct {
	Hand
}

func (sh SkillfulHand) RaiseUp() {
	fmt.Println("This skillfull hand is raising up")
}

func (sh SkillfulHand) PutDown() {
	fmt.Println("This skillfull hand is putting down")
}

func (sh SkillfulHand) Wave() {
	fmt.Println("This skillfull hand is waving")
}
