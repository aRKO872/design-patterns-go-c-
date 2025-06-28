package director

import "github.com/builder-design-pattern/desktop"

type DesktopDirector struct {
}

func (d *DesktopDirector) BuildDesktop(builder desktop.IDesktopBuilder) desktop.Desktop {
	builder.BuildStorage()
	builder.BuildRAM()
	builder.BuildProcessor()
	builder.BuildGraphics()
	builder.BuildMotherboard()
	return builder.GetDesktop()
}