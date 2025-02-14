package deepseek

type ChatRequestBody struct {
	Messages    []Message `json:"messages"`
	Model       string    `json:"model"`
	MaxTokens   int       `json:"max_tokens"`
	Temperature float32   `json:"temperature"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func buildRequestBody(messages []Message, model string, maxTokens int, temperature float32) ChatRequestBody {
	return ChatRequestBody{
		Messages:    messages,
		Model:       model,
		MaxTokens:   maxTokens,
		Temperature: temperature,
	}
}
