package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/creack/pty"
	"os"
	"os/exec"
	"time"
)

func main() {
	a := app.New()
	w := a.NewWindow("goterm")

	w.SetContent(widget.NewLabel("I am on goterm!"))

	c := exec.Command("bin/bash")
	p, err := pty.Start(c)

	if err != nil {
		fyne.LogError("Failed to open pty", err)
		os.Exit(1)
	}

	defer c.Process.Kill()

	p.Write([]byte("ls\r"))
	time.Sleep(1 * time.Second)
	b := make([]byte, 1024)
	_, err = p.Read(b)
	if err != nil {
		fyne.LogError("Failed to read pty", err)
	}

	w.Resize(fyne.NewSize(420, 200))
	w.ShowAndRun()
}
