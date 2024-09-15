package mamba

import "github.com/gdamore/tcell/v2"

type TextView struct {
	Frame
	text string
}

func NewTextView(f Frame, text string) *TextView {
	f.startX++
	f.width -= 2
	f.startY++
	f.height -= 2
	return &TextView{Frame: f, text: text}
}

func (t *TextView) Show() {
	for i, r := range t.text {
		t.s.SetContent(t.startX+i, t.startY, r, nil, tcell.StyleDefault)
	}
}
