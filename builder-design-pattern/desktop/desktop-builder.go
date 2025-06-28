package desktop

type IDesktopBuilder interface{
	BuildRAM()
	BuildProcessor()
	BuildStorage()
	BuildMotherboard()
	BuildGraphics()
	GetDesktop() Desktop
}