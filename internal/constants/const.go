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
	{Name: " mistral-7b-instruct-v0.1", Desc: "Instruct fine-tuned version of the Mistral-7b generative text model with 7 billion parameters", Endpoint: "@cf/mistral/mistral-7b-instruct-v0.1"},
	{Name: " llama-2-7b-chat-fp16 ", Desc: "Full precision (fp16) generative text model with 7 billion parameters from Meta", Endpoint: "@cf/meta/llama-2-7b-chat-fp16"},
	{Name: "qwen1.5-1.8b-chat (beta)", Desc: "Qwen1.5 is the improved version of Qwen, the large language model series developed by Alibaba Cloud.", Endpoint: "@cf/qwen/qwen1.5-1.8b-chat"},
}
