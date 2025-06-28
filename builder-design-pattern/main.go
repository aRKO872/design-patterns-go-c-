package main

import (
	"github.com/builder-design-pattern/desktop"
	"github.com/builder-design-pattern/director"
)

func main() {
	desktopDirector := new(director.DesktopDirector)
	dellDesktopBuilder := new(desktop.DellDesktopBuilder)
	hpDesktopBuilder := new(desktop.HPDesktopBuilder)

	dellDesktop := desktopDirector.BuildDesktop(dellDesktopBuilder)
	dellDesktop.Display()

	hpDesktop := desktopDirector.BuildDesktop(hpDesktopBuilder)
	hpDesktop.Display()
}