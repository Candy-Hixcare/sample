package classB

import "cycleImport/cycleImport/classA"

type ClassB struct {
	data     int
	myClassA classA.ClassA
}

func (c *ClassB) Init(dataB int, instanceA classA.ClassA) {
	c.data = dataB
	c.myClassA = instanceA
}

func (c ClassB) GetData() int { return c.data }

func (c ClassB) GetMyClassA() classA.ClassA { return c.myClassA }
