package scraper

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

/*
A small TUI that allows you to highlight text to save from Wikipedia.

1. Fetch the list of all of the Consuls from Wikipedia.
2. On each that has a link we need to visit that link and then create a cmd that pipes
   the text into this TUI.

This TUI app allows you to highlight text and save it into a csv file (for the moment).
The text should be from the opening paragraph of wikipedia.
*/

type Model struct {
	text     string
	sPos     int
	ePos     int
	vMod     bool
	selected string

	info string
	w, h int

	lPos  int
	links []Link
}

type Link struct {
	Name string
	Link string
}

func NewModel() *Model {
	// Should update lPos based on already saved consuls
	return &Model{
		text:     "",
		sPos:     0,
		ePos:     0,
		vMod:     false,
		selected: "",

		info: "",

		lPos:  0,
		links: GetConsulLinks(),
	}
}

func (m Model) Init() tea.Cmd {
	return NewCmd(m.links[m.lPos])
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.w = msg.Width
		m.h = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "p":
			m.selected = m.text[m.sPos:m.ePos]
		case "right", "l":
			if !m.vMod {
				if m.sPos < len(m.text) {
					m.sPos++
					m.ePos = m.sPos
				}
			} else {
				if m.ePos < len(m.text) {
					m.ePos++
				}
			}
		case "left", "h":
			if !m.vMod {
				if m.sPos > 0 {
					m.sPos--
					m.ePos = m.sPos
				}
			} else {
				if m.ePos > m.sPos {
					m.ePos--
				}
			}
		case "v":
			m.vMod = true
		case "esc":
			m.vMod = false
			m.selected = ""
			m.ePos = m.sPos
		case "enter":
			if m.selected != "" {
				m.info = ""
				if m.lPos < len(m.links)-1 {
					m.selected = ""
					m.sPos = 0
					m.ePos = 0
					m.vMod = false

					m.lPos++
					return m, NewCmd(m.links[m.lPos])
				} else {
					m.info = "No more links..."
				}
			} else {
				m.info = "You have not selected any text for this link...?"
			}

		}
	case wikiIntroMsg:
		m.text = string(msg)
		return m, nil
	case wikiErrorMsg:
		fmt.Println("error: ", msg)
		return nil, tea.Quit
	}
	return m, nil
}

func (m Model) View() string {
	if m.w == 0 {
		return "Loading..."
	}

	cursor := "|"
	s := "Move the cursor along the text\n\n"

	if m.sPos == m.ePos {
		b, a := m.text[:m.sPos], m.text[m.sPos:]
		s += b + cursor + a
	} else {
		// split the string in three parts
		// before sPos : inbetween sPos & ePos : ePos to end of string
		a, b, c := m.text[:m.sPos], m.text[m.sPos:m.ePos], m.text[m.ePos:]
		s += a + cursor + b + cursor + c
	}

	if m.selected != "" {
		s += fmt.Sprintf("\n\nSelected: %v", m.selected)
	}

	if m.info != "" {
		s += fmt.Sprintf("\n\nInfo: %v", m.info)
	}

	s += "\n\nPress q to quite.\n"
	return s
}

func NewCmd(l Link) tea.Cmd {
	return func() tea.Msg {
		s, err := GetConsulIntroText(l.Link)
		if err != nil {
			return wikiErrorMsg{err}
		}
		return wikiIntroMsg(s)
	}
}

type wikiIntroMsg string
type wikiErrorMsg struct{ error }

func (w wikiErrorMsg) Error() string { return w.error.Error() }
