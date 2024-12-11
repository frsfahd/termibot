package constants

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/frsfahd/termiBot/internal/llm"
)

/* CONSTANTS*/

var (
	P          *tea.Program
	WindowSize tea.WindowSizeMsg
)

var LLM_LIST = []llm.LLM{
	{Name: " llama-3-8b-instruct", Desc: "Generation over generation, Meta Llama 3 demonstrates state-of-the-art performance on a wide range of industry benchmarks and offers new capabilities, including improved reasoning.", Endpoint: "@cf/meta/llama-3-8b-instruct"},
	{Name: "llm_b", Desc: "model for multimedia generation", Endpoint: "api/v1/llm_b"},
	{Name: "llm_c", Desc: "model for image classification", Endpoint: "api/v1/llm_c"},
	{Name: "llm_code", Desc: "model for code generation", Endpoint: ""},
}
