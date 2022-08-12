package robot

import (
	"design-pattern/entities/robot/components"
	"design-pattern/entities/robot/components/cabin"
	"design-pattern/entities/robot/components/hand"
	"design-pattern/entities/robot/components/wheel"
)

type Robot struct {
	Name  string
	ID    string
	Parts struct {
		RobotHand  [2]hand.WorkableHand
		RobotCabin [2]cabin.WorkableCabin
		RobotWheel []wheel.WorkableWheel
		Others     []components.RobotComponent
	}
	WorkableRobot
}

type WorkableRobot interface {
	SetHands(left, right hand.WorkableHand)
	SetWheels(wheels ...wheel.WorkableWheel)
	SetCabins(mainCabin, subCabin cabin.WorkableCabin)
	Greet()
	Bye()
	DoMainTask()
	DoSubTask()
}

type ClonableRobot interface {
	Clone() WorkableRobot
}

func (wb *Robot) SetHands(left, right hand.WorkableHand) {
	if wb.Parts.RobotHand[0] == nil && left != nil {
		wb.Parts.RobotHand[0] = left
	}

	if wb.Parts.RobotHand[1] == nil && right != nil {
		wb.Parts.RobotHand[1] = right
	}
}

func (wb *Robot) SetWheels(wheels ...wheel.WorkableWheel) {
	wb.Parts.RobotWheel = append(wb.Parts.RobotWheel, wheels...)
}

func (wb *Robot) SetCabins(mainCabin, subCabin cabin.WorkableCabin) {
	if wb.Parts.RobotCabin[0] == nil && mainCabin != nil {
		wb.Parts.RobotCabin[0] = mainCabin
	}

	if wb.Parts.RobotCabin[1] == nil && subCabin != nil {
		wb.Parts.RobotCabin[1] = subCabin
	}
}
