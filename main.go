package main

import (
	"design-pattern/uc/factory"
)

func main() {
	deliverFactory := factory.DeliverRobotFactory{}
	naiveBot := deliverFactory.CreateNaiveRobot()

	naiveBot.Greet()
	naiveBot.Parts.RobotWheel[3].Run()
	naiveBot.Parts.RobotCabin[0].Open()

	skillfulBot := deliverFactory.CreateAdvancedRobot()
	skillfulBot.Greet()
	skillfulBot.Parts.RobotWheel[3].Run()
	for _, cab := range skillfulBot.Parts.RobotCabin {
		cab.Open()
	}

}
