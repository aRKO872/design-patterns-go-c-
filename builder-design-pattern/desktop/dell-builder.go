package desktop

type DellDesktopBuilder struct {
	desktop Desktop
}

func (d *DellDesktopBuilder) BuildRAM() {
	d.desktop.SetRAM("16 GB DDR4 RAM")
}

func (d *DellDesktopBuilder) BuildProcessor() {
	d.desktop.SetProcessor("AMD Ryzen")
}

func (d *DellDesktopBuilder) BuildMotherboard() {
	d.desktop.SetMotherboard("Dell Motherboard")
}

func (d *DellDesktopBuilder) BuildStorage() {
	d.desktop.SetStorage("1TB SSD")
}

func (d *DellDesktopBuilder) BuildGraphics() {
	d.desktop.SetGraphics("RTX 5090")
}

func (a *DellDesktopBuilder) GetDesktop() Desktop {
	return a.desktop
}
