package agent

import (
	"ai_agent/llm"
	"ai_agent/model"
	"fmt"
)

func HandleTask(task model.Task) (string, error) {
	memories, _ := LoadRecentMemory(0, 5)
	memoryContext := buildMemoryContext(memories)
	switch task.Type {
		case "echo":
			result := memoryContext + "\nUser says: " + task.Payload
			SaveMemory(0, "Echoed: " + task.Payload)
			return result, nil

		case "init_memory":
			return "Memory initialized", nil

		case "warmup":
			return "Embedding model warmed", nil

		case "summarize":
			return handleSummarizeWithMemory(task, memoryContext)

		case "chat":
			return handleChatWithMemory(task, memoryContext)
		
		case "recycle_memory":
			RecycleMemory(0)
			CompressMemory(0)
			return "Memory recycled and compressed", nil

		default:
			return "", fmt.Errorf("unknown task type: %s", task.Type)
	}
}

func handleSummarizeWithMemory(task model.Task, mem string) (string, error) {
	prompt := mem +
		"\n\nNow summarize the following text:\n" +
		task.Payload

	answer, err := llm.AskDeepSeek(prompt)
	if err != nil {
		return "", err
	}
	SaveMemory(0, "Summarized: " + task.Payload)
	SaveMemory(0, "AI said: " + answer)
	return answer, nil
}

func handleChatWithMemory(task model.Task, mem string) (string, error) {
	prompt := mem +
		"\n\nUser: " + task.Payload +
		"\nAI: (reply naturally based on memories)"

	answer, err := llm.AskDeepSeek(prompt)
	if err != nil {
		return "", err
	}
	SaveMemory(0, "User: " + task.Payload)
	SaveMemory(0, "AI said: " + answer)
	return answer, nil
}