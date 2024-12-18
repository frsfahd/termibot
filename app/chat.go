package app

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/frsfahd/termiBot/internal/chat"
	"github.com/frsfahd/termiBot/internal/constants"
	"github.com/frsfahd/termiBot/internal/llm"
)

// sessionState is used to track which model is focused
type sessionState uint

const (
	textareaView sessionState = iota
	viewPortView
)

type Chat_Model struct {
	spinner  spinner.Model
	viewport viewport.Model
	textarea textarea.Model
	chat     chat.Chat
	llm      llm.LLM
	state    sessionState
	quitting bool
}

func (m *Chat_Model) setViewportContents() {
	if len(m.chat.Messages) > 0 {
		m.viewport.SetContent(strings.Join(m.chat.Messages, "\n"))

	}
}

func initChat(llm llm.LLM) (tea.Model, tea.Cmd) {
	//reset the chat history
	chat.MsgHistory = make([]chat.Message, 0)
	chat.MsgHistory = append(chat.MsgHistory, chat.Message{Role: "system", Content: "You are a helpful assistant"})

	//setup the styling
	x, y := constants.DocStyle.GetFrameSize()

	constants.DocStyle = constants.DocStyle.Width(constants.WindowSize.Width - x).Height(constants.WindowSize.Height - y - 5)
	docX, docY := lipgloss.Size(constants.DocStyle.String())
	constants.MsgViewStyle = constants.MsgViewStyle.Width(docX - x - 20).Height(docY - y - 5)
	constants.MsgInputStyle = constants.MsgInputStyle.Width(docX - x - 20).Height(3)

	//initialize textarea chat input
	ta := textarea.New()
	ta.Placeholder = "Message AI..."
	ta.Focus()
	ta.Prompt = ""
	ta.CharLimit = 1000

	// ta.FocusedStyle.Base.Padding(1).BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("#FFF275"))

	msgInputX, msgInputY := lipgloss.Size(constants.MsgInputStyle.String())

	ta.SetWidth(msgInputX)
	ta.SetHeight(msgInputY)
	// ta.FocusedStyle.CursorLine = lipgloss.NewStyle()

	// ta.FocusedStyle.Base = constants.MsgInputStyle

	ta.ShowLineNumbers = false

	ta.KeyMap.InsertNewline.SetEnabled(false)

	//initialize viewport chat messages
	msgViewX, msgViewY := lipgloss.Size(constants.MsgViewStyle.String())

	vp := viewport.New(msgViewX, msgViewY)

	return Chat_Model{
		viewport: vp,
		textarea: ta,
		llm:      llm,
	}, nil
}

func (m Chat_Model) Init() tea.Cmd {
	return nil
}

func (m Chat_Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	// var taCmd tea.Cmd
	// var vpCmd tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		constants.WindowSize = msg

		x, y := constants.DocStyle.GetFrameSize()

		constants.DocStyle = constants.DocStyle.Width(constants.WindowSize.Width - x).Height(constants.WindowSize.Height - y - 5)
		docX, docY := lipgloss.Size(constants.DocStyle.String())

		constants.MsgViewStyle = constants.MsgViewStyle.Width(docX - x - 20).Height(docY - y - 5)
		constants.MsgInputStyle = constants.MsgInputStyle.Width(docX - x - 20).Height(3)

		msgInputX, msgInputY := lipgloss.Size(constants.MsgInputStyle.String())
		msgViewX, msgViewY := lipgloss.Size(constants.MsgViewStyle.String())

		m.viewport = viewport.New(msgViewX, msgViewY)

		m.textarea.SetWidth(msgInputX)
		m.textarea.SetHeight(msgInputY)

	case errMsg:
		log.Println(msg)
	case msgByte:
		//delete spinner
		m.chat.Messages = m.chat.Messages[:len(m.chat.Messages)-1]

		var res Response
		if err := json.Unmarshal(msg, &res); err != nil {
			m.chat.Messages = append(m.chat.Messages, constants.SenderStyle.Render("System: ")+err.Error())
			log.Println(err)
		}
		if res.Success && res.Result.Response != "" {

			//update model chat history
			chat.MsgHistory = append(chat.MsgHistory, chat.Message{Role: "assistant", Content: res.Result.Response})
			//update displayed chat histpry
			m.chat.Messages = append(m.chat.Messages, constants.SenderStyle.Render("Assistant: ")+res.Result.Response+"\n")
		} else {
			m.chat.Messages = append(m.chat.Messages, constants.SenderStyle.Render("System: something wrong..."))

		}
		m.viewport.GotoBottom()

	case spinner.TickMsg:
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	case tea.KeyMsg:

		switch {
		case key.Matches(msg, constants.Keymap.Quit):
			m.quitting = true
			return m, tea.Quit
		case key.Matches(msg, constants.Keymap.Back):
			return InitLLM()
		case key.Matches(msg, constants.Keymap.Focus):
			if m.state == textareaView {
				m.state = viewPortView
			} else {
				m.state = textareaView
			}
		}
		switch m.state {
		case textareaView:
			m.textarea.Focus()

			if key.Matches(msg, constants.Keymap.Enter) {
				//append to model chat history
				chat.MsgHistory = append(chat.MsgHistory, chat.Message{Role: "user", Content: m.textarea.Value()})
				cmds = append(cmds, sendMsg(chat.MsgHistory, m.llm.Endpoint))
				//append to displayed chat history
				m.chat.Messages = append(m.chat.Messages, constants.SenderStyle.Render("You: ")+m.textarea.Value())
				m.textarea.Reset()
				// m.viewport.GotoBottom()
				// init loading text
				m.spinner = spinner.New()
				m.spinner.Spinner = spinner.Line
				m.spinner.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("69"))
				m.chat.Messages = append(m.chat.Messages, constants.SenderStyle.Render("System: ⏳ waiting for respond..."))
			}

			m.textarea, cmd = m.textarea.Update(msg)
		case viewPortView:
			m.textarea.Blur()

			m.viewport, cmd = m.viewport.Update(msg)

		}
		cmds = append(cmds, cmd)

	}

	m.setViewportContents()

	return m, tea.Batch(cmds...)
}

func (m Chat_Model) View() string {
	if m.quitting {
		return ""
	}
	currentLLM := constants.LLMNameStyle.Render(m.llm.Name)
	title := fmt.Sprintf(
		"You're currently chatting with %s\nType a message and press [Enter] to send.\nUse [tab] to switch between input box and message view.\nUse ⬆️ (up/PgUp) & ⬇️ (down/PgDown) for scrolling.", currentLLM)

	var msgInput string
	var msgView string
	var formatted string

	msgInput = constants.MsgInputStyle.Render(m.textarea.View())
	msgView = constants.MsgViewStyle.Render(m.viewport.View())
	// formatted := fmt.Sprintf("%s\n\n%s\n\n%s", title, msgView, msgInput)
	formatted = lipgloss.JoinVertical(lipgloss.Center, title, msgView, msgInput)

	return constants.DocStyle.Render(formatted)
}
