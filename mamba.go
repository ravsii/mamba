package mamba

import (
	"os"

	"github.com/gdamore/tcell/v2"
)

type Screen struct {
	s tcell.Screen
	f *Frame
}

func NewScreen() (*Screen, error) {
	screen, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}
	if err := screen.Init(); err != nil {
		return nil, err
	}

	w, h := screen.Size()

	frame, err := NewFrame(screen, 0, 0, w, h)
	if err != nil {
		return nil, err
	}

	return &Screen{screen, frame}, nil
}

// Run is a blocking operation.
func (s *Screen) Run() {
	quit := func() {
		s.s.Fini()
		os.Exit(0)
	}
	for {
		s.s.Show()
		ev := s.s.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.s.Clear()
			w, h := ev.Size()
			s.f.Resize(w, h)
			s.f.Show()
			s.s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				quit()
			}
		}
	}
}

func (s *Screen) Frame() *Frame {
	return s.f
}
