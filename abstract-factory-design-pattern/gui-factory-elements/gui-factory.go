package guifactoryelements

type GUIFactory interface {
	CreateButton() IButton
	CreateTextBox() ITextBox
}

func NewGuiFactory(typeOfOS string) GUIFactory {
	switch typeOfOS {
		case "Windows" :
			return new(WinFactory)
		case "MacOS" :
			return new(MacFactory)
		case "Linux" :
			return new(LinuxFactory)
		default :
			return nil
	}
}

type MacFactory struct {
}

func (m *MacFactory) CreateButton() IButton {
	return new(MacButton)
}

func (m *MacFactory) CreateTextBox() ITextBox {
	return new(MacTextBox)
}

type WinFactory struct {
}

func (m *WinFactory) CreateButton() IButton {
	return new(WinButton)
}

func (m *WinFactory) CreateTextBox() ITextBox {
	return new(WinTextBox)
}

type LinuxFactory struct {
}

func (m *LinuxFactory) CreateButton() IButton {
	return new(LinuxButton)
}

func (m *LinuxFactory) CreateTextBox() ITextBox {
	return new(LinuxTextBox)
}