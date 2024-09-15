package mamba

import (
	"github.com/gdamore/tcell/v2"
)

type Shower interface {
	Show()
}

type Frame struct {
	startX, startY, width, height int
	s                             tcell.Screen
	borderChar                    rune
	children                      []Shower
}

func NewFrame(s tcell.Screen, startX, startY, width, height int) (*Frame, error) {
	return &Frame{
		s:          s,
		startX:     startX,
		startY:     startY,
		width:      width,
		height:     height,
		borderChar: '+',
	}, nil
}

func (f *Frame) Resize(w, h int) {
	f.width = w
	f.height = h
}

func (f *Frame) AddChild(c Shower) {
	f.children = append(f.children, c)
}

func (f *Frame) Show() {
	for i := range f.width {
		f.s.SetContent(i, f.startY, f.borderChar, nil, tcell.StyleDefault)
		f.s.SetContent(i, f.height-1, f.borderChar, nil, tcell.StyleDefault)
	}
	for i := range f.height {
		f.s.SetContent(f.startX, i, f.borderChar, nil, tcell.StyleDefault)
		f.s.SetContent(f.width-1, i, f.borderChar, nil, tcell.StyleDefault)
	}

	for i := range f.children {
		f.children[i].Show()
	}
}
