package classC

import (
	"cycleImport/cycleImport/classB"
)

type ClassC struct {
	data     int
	myClassB classB.ClassB
}

func (c *ClassC) Init(dataC int, instanceB classB.ClassB) {
	c.data = dataC
	c.myClassB = instanceB
}

func (c ClassC) GetData() int { return c.data }

func (c ClassC) GetMyClassB() classB.ClassB { return c.myClassB }
