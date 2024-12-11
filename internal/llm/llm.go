package llm

type LLM struct {
	Name, Desc, Endpoint string
}

func (i LLM) Title() string       { return i.Name }
func (i LLM) Description() string { return i.Desc }
func (i LLM) FilterValue() string { return i.Name }
