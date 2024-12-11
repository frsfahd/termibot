package app

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/frsfahd/termiBot/internal/constants"
	"github.com/frsfahd/termiBot/internal/llm"
)

type LLM_Model struct {
	llmList  list.Model
	quitting bool
}

func setupLLMList() []list.Item {
	items := make([]list.Item, len(constants.LLM_LIST))
	for i, llm := range constants.LLM_LIST {
		items[i] = list.Item(llm)
	}
	return items
}

func InitLLM() (tea.Model, tea.Cmd) {
	llmList := setupLLMList()
	model := LLM_Model{
		llmList: list.New(llmList, list.NewDefaultDelegate(), 8, 8),
	}

	x, y := constants.DocStyle.GetFrameSize()

	constants.DocStyle = constants.DocStyle.Width(constants.WindowSize.Width - x).Height(constants.WindowSize.Height - y - 5)
	docX, docY := lipgloss.Size(constants.DocStyle.String())
	model.llmList.SetSize(docX, docY)

	model.llmList.Title = "Available LLMs"
	return model, func() tea.Msg { return errMsg{nil} }
}

func (m LLM_Model) Init() tea.Cmd {
	return nil
}

func (m LLM_Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		constants.WindowSize = msg
		x, y := constants.DocStyle.GetFrameSize()

		constants.DocStyle = constants.DocStyle.Width(constants.WindowSize.Width - x).Height(constants.WindowSize.Height - y - 5)
		docX, docY := lipgloss.Size(constants.DocStyle.String())

		m.llmList.SetSize(docX, docY)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, constants.Keymap.Quit):
			m.quitting = true
			return m, tea.Quit
		case key.Matches(msg, constants.Keymap.Enter):
			activeLLM := m.llmList.SelectedItem().(llm.LLM)
			return initChat(activeLLM)
			// return chat.Update(constants.WindowSize)

		}
	}
	m.llmList, cmd = m.llmList.Update(msg)

	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m LLM_Model) View() string {
	return constants.DocStyle.Render(m.llmList.View())
}
