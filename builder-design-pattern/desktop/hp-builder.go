package desktop

type HPDesktopBuilder struct {
	desktop Desktop
}

func (d *HPDesktopBuilder) BuildRAM() {
	d.desktop.SetRAM("16 GB DDR4 RAM")
}

func (d *HPDesktopBuilder) BuildProcessor() {
	d.desktop.SetProcessor("Intel I7 - 15th Gen")
}

func (d *HPDesktopBuilder) BuildMotherboard() {
	d.desktop.SetMotherboard("HP Motherboard")
}

func (d *HPDesktopBuilder) BuildStorage() {
	d.desktop.SetStorage("1.5TB SSD")
}

func (d *HPDesktopBuilder) BuildGraphics() {
	d.desktop.SetGraphics("RTX 6090")
}

func (a *HPDesktopBuilder) GetDesktop() Desktop {
	return a.desktop
}

