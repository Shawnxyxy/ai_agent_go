package agent

import (
	"fmt"
	"ai_agent/model"
)

func HandleTask(task model.Task) (string, error) {
	switch task.Type {
		case "echo":
			return handleEcho(task)

		case "init_memory":
			return "Memory initialized", nil

		case "warmup":
			return "Embedding model warmed", nil

		case "summarize":
			return handleSummarize(task)

		default:
			return "", fmt.Errorf("unknown task type: %s", task.Type)
	}
}

func handleEcho(task model.Task) (string, error) {
	return "Echo:" + task.Payload, nil
}

func handleSummarize(task model.Task) (string, error) {
	return "Summary: " + task.Payload, nil
}