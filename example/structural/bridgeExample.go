package main

import (
	"fmt"

	"github.com/Zando74/generic-patterns/structural"
)

type Processor interface {
	structural.Bridgeable
	process() string
}

type Intel struct{}

func (i *Intel) process() string {
	return "Intel Processor"
}

type AMD struct{}

func (a *AMD) process() string {
	return "AMD Processor"
}

type OS interface {
	structural.Bridgeable
	boot() string
}

type WindowsOS struct{}

func (w *WindowsOS) boot() string {
	return "Windows OS"
}

type MacOS struct{}

func (m *MacOS) boot() string {
	return "Mac OS"
}

type Computer interface {
	Compute()
}

type Windows struct {
	Proc structural.Bridge[Processor]
	OS   structural.Bridge[OS]
}

func (w *Windows) Compute() {
	fmt.Println("WINDOWS Compute using : ", (*w.Proc.Impl).process(), " and ", (*w.OS.Impl).boot())
}

type Mac struct {
	Proc structural.Bridge[Processor]
	OS   structural.Bridge[OS]
}

func (m *Mac) Compute() {
	fmt.Println("MAC Compute using : ", (*m.Proc.Impl).process(), " and ", (*m.OS.Impl).boot())
}

func MainBridgeExample() {
	windows := &Windows{}
	mac := &Mac{}

	intel := &Intel{}
	amd := &AMD{}

	windowsOS := &WindowsOS{}
	macOS := &MacOS{}

	windows.Proc.SetImpl(intel)
	mac.Proc.SetImpl(amd)

	windows.OS.SetImpl(windowsOS)
	mac.OS.SetImpl(macOS)

	windows.Compute()
	mac.Compute()

}
