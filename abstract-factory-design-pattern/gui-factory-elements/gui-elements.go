package guifactoryelements
import "log"

type IButton interface {
	Press()
}

type ITextBox interface {
	SetText()
}

type MacButton struct {
}

func (m *MacButton) Press() {
	log.Println("Mac Button Pressed")
}

type WinButton struct {
}

func (m *WinButton) Press() {
	log.Println("Windows Button Pressed")
}

type LinuxButton struct {
}

func (m *LinuxButton) Press() {
	log.Println("Linux Button Pressed")
}

type MacTextBox struct {
}

func (m *MacTextBox) SetText() {
	log.Println("Mac TextBox Set Text")
}

type WinTextBox struct {
}

func (m *WinTextBox) SetText() {
	log.Println("Windows TextBox Set Text")
}

type LinuxTextBox struct {
}

func (m *LinuxTextBox) SetText() {
	log.Println("Linux TextBox Set Text")
}

