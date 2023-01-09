package classA

import "cycleImport/cycleImport/classC"

type ClassA struct {
	data int
}

func (a *ClassA) Init(data int) {
	a.data = data
}

func (a ClassA) GetData() int { return a.data }

/* 造成 Cycle Import 的初始化 function */
func (c *ClassA) InitByClassC(classC classC.ClassC) {
	c.data = classC.GetMyClassB().GetMyClassA().GetData()
}
