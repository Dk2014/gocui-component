package main

import (
	"fmt"

	"github.com/jroimartin/gocui"
	component "github.com/skanehira/gocui-component"
)

type signup struct {
	*component.Form
}

func main() {
	gui, err := gocui.NewGui(gocui.Output256)

	if err != nil {
		panic(err)
	}
	defer gui.Close()

	if err := gui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		panic(err)
	}

	signup := &signup{
		component.NewForm(gui, "Sign Up", 0, 0, 0, 0),
	}

	signup.AddInputField("First Name", 11, 18).
		AddValidator("required input", requireValidator)
	signup.AddInputField("Last Name", 11, 18).
		AddValidator("required input", requireValidator)

	signup.AddInputField("Password", 11, 18).
		AddValidator("required input", requireValidator).
		SetMask().
		SetMaskKeybinding(gocui.KeyCtrlM)

	signup.AddButton("Regist", signup.regist)
	signup.AddButton("Cancel", quit)

	signup.Draw()

	if err := gui.MainLoop(); err != nil && err != gocui.ErrQuit {
		panic(err)
	}
}

func (s *signup) regist(g *gocui.Gui, v *gocui.View) error {
	if !s.Validate() {
		return nil
	}

	if v, err := g.SetView("registed", 0, 0, 30, 5); err != nil {
		if err != gocui.ErrUnknownView {
			panic(err)
		}

		v.Title = v.Name()
		v.Wrap = true

		for label, text := range s.GetFormData() {
			fmt.Fprintf(v, "%s: %s\n", label, text)
		}

		g.SetCurrentView(v.Name())
		g.SetKeybinding(v.Name(), gocui.KeyEnter, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
			g.DeleteView(v.Name())
			g.DeleteKeybindings(v.Name())
			s.SetCurrentItem(0)
			return nil
		})
	}

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func requireValidator(text string) bool {
	if text == "" {
		return false
	}
	return true
}