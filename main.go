package main

import (
	"os"
	"os/exec"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/creack/pty"
)

func main() {
	a := app.New()
	w := a.NewWindow("goterm")

	ui := widget.NewTextGrid()       // Create a new TextGrid
	ui.SetText("I'm on a terminal!") // Set text to display

	c := exec.Command("/usr/bin/bash")
	p, err := pty.Start(c)

	if err != nil {
		fyne.LogError("Failed to open pty", err)
		os.Exit(1)
	}

	defer c.Process.Kill()

	onTypedKey := func(e *fyne.KeyEvent) {
		if e.Name == fyne.KeyEnter || e.Name == fyne.KeyReturn {
			_, _ = p.Write([]byte{'\r'})
		}
	}

	onTypedRune := func(r rune) {
		_, _ = p.WriteString(string(r))
	}

	w.Canvas().SetOnTypedKey(onTypedKey)
	w.Canvas().SetOnTypedRune(onTypedRune)

	go func() {
		for {
			time.Sleep(1 * time.Second)
			b := make([]byte, 256)
			_, err = p.Read(b)
			if err != nil {
				fyne.LogError("Failed to read pty", err)
			}
			w.SetContent(widget.NewLabel(string(b)))
		}
	}()

	w.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewGridWrapLayout(fyne.NewSize(420, 200)),
			ui,
		),
	)
	w.ShowAndRun()
}
