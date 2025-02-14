package deepseek

import "os"

type config struct {
	url    string
	apiKey string
}

func getConfig() config {
	return config{
		url:    os.Getenv("DEEPSEEK_URL"),
		apiKey: os.Getenv("DEEPSEEK_API_KEY"),
	}
}
