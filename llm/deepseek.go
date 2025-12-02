package llm

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
	"fmt"
	"ai_agent/config"
)

type DeepSeekRequest struct {
	Model  		string           		`json:"model"`
	Messages	[]map[string]string		`json:"messages"`
}

type DeepSeekResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func AskDeepSeek(prompt string) (string, error) {
	cfg := config.Cfg.DeepSeek
	fmt.Println("Deepseek prompt:", prompt) // 打印发送的 prompt

	body := DeepSeekRequest{
		Model: cfg.Model,
		Messages: []map[string]string{
			{"role": "user", "content": prompt},
		},
	}
	jsonData, _ := json.Marshal(body)
	fmt.Println("Request JSON:", string(jsonData)) // 打印请求 JSON

	req, _ := http.NewRequest("POST", cfg.BaseURL+"/chat/completions", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer " + cfg.ApiKey)

	client := &http.Client{Timeout: 20 * time.Second}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("HTTP request error:", err)
		return "", err
	}
	defer resp.Body.Close()

	var res DeepSeekResponse
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		fmt.Println("Decode response error:", err)
		return "", err
	}

	if len(res.Choices) == 0 {
		fmt.Println("Empty choices in response")
		return "", nil
	}
	answer := res.Choices[0].Message.Content
	fmt.Println("Deepseek answer:", answer) // 打印解析后的回答
	return answer, nil
}