package main

import (
	"cycleImport/cycleImport/classA"
	"cycleImport/cycleImport/classB"
	"cycleImport/cycleImport/classC"
)

func main() {
	a := classA.ClassA{}
	a.Init(5)

	b := classB.ClassB{}
	b.Init(3, a)

	c := classC.ClassC{}
	c.Init(2, b)

	/* 期望實作一個函式，利用已存在的 ClassC 建立一個新的 ClassA
	 * 並解決 import cycle not allowed 的問題
	 */
	newA := classA.ClassA{}
	newA.InitByClassC(c)
}
