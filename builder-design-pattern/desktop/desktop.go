package desktop

import "fmt"

type Desktop struct {
	motherboard string
	ram string
	graphics string
	storage string
	processor string
}

// Getters
func (d *Desktop) GetMotherboard() string {
	return d.motherboard
}

func (d *Desktop) GetRAM() string {
	return d.ram
}

func (d *Desktop) GetGraphics() string {
	return d.graphics
}

func (d *Desktop) GetStorage() string {
	return d.storage
}

func (d *Desktop) GetProcessor() string {
	return d.processor
}

// Setters
func (d *Desktop) SetMotherboard(motherboard string) {
	d.motherboard = motherboard
}

func (d *Desktop) SetRAM(ram string) {
	d.ram = ram
}

func (d *Desktop) SetGraphics(graphics string) {
	d.graphics = graphics
}

func (d *Desktop) SetStorage(storage string) {
	d.storage = storage
}

func (d *Desktop) SetProcessor(processor string) {
	d.processor = processor
}

func (d *Desktop) Display() {
	fmt.Printf("RAM: %s,\nProcessor: %s,\nStorage: %s,\nGraphics: %s,\nMotherboard: %s,\n", 
	d.GetRAM(),
	d.GetProcessor(),
	d.GetStorage(),
	d.GetGraphics(),
	d.GetMotherboard(),
	)
}