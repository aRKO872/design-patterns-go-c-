package main

import guifactoryelements "github.com/abstract-factory-design-pattern/gui-factory-elements"

func main() {
	macFctry := guifactoryelements.NewGuiFactory("MacOS")
	winFctry := guifactoryelements.NewGuiFactory("Windows")

	if macFctry != nil {
		macBtn := macFctry.CreateButton()
		macBtn.Press()
		macTBox := macFctry.CreateTextBox()
		macTBox.SetText()
	}

	if winFctry != nil {
		winBtn := winFctry.CreateButton()
		winBtn.Press()
		winTBox := winFctry.CreateTextBox()
		winTBox.SetText()
	}
}