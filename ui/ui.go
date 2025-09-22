package ui

import "fmt"

type UserInterface struct {
	content string
}

func New(formularioHtml string) UserInterface {
	ui := UserInterface{
		content: formularioHtml,
	}
	return ui
}

func (u UserInterface) Render() string {
	fmt.Printf("LOG: imprimiendo formulario: <h1>%s</h1>", u.content)
	return fmt.Sprintf("<h1>%s</h1>", u.content)
}
