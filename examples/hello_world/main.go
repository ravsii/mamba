package main

import "github.com/ravsii/mamba"

func main() {
	screen, err := mamba.NewScreen()
	if err != nil {
		panic(err)
	}

	text := mamba.NewTextView(*screen.Frame(), "Hello World!")
	screen.Frame().AddChild(text)

	screen.Run()
}
