package agent

import (
	"ai_agent/llm"
	"strings"
)

// -------- Memory Classification --------

// 让 LLM 判断这段 memory 属于哪种类型
func classifyMemory(content string) string {
	prompt := `
	你是一名负责分类用户记忆的 AI 系统。
	请根据下面的内容判断 memory 类型。只返回以下类别之一：

	- interaction:普通对话
	- preference:偏好（喜欢什么/不喜欢什么）
	- plan:计划、打算做的事情
	- identity:用户的个人信息(角色、职业、自述)
	- goal:目标、长期想实现的事
	- fact:客观信息

	要分类的内容：
	` + content

		resp, _ := llm.AskDeepSeek(prompt)
		resp = strings.ToLower(resp)

		switch {
			case strings.Contains(resp, "preference"):
				return "preference"
			case strings.Contains(resp, "plan"):
				return "plan"
			case strings.Contains(resp, "identity"):
				return "identity"
			case strings.Contains(resp, "goal"):
				return "goal"
			case strings.Contains(resp, "fact"):
				return "fact"
			default:
				return "interaction"
		}
}

// -------- Memory Importance Scoring --------

// 让 LLM 对这段 memory 的重要程度打分（0~2）
func scoreImportance(content string) int {
	prompt := `
	你是一名负责判断用户记忆重要度的 AI。
	请根据内容给它一个重要度评分：
	0 = 普通对话，不重要
	1 = 有一定意义
	2 = 非常重要、需要长期记住

	只返回一个数字（0、1、2）。

	内容：
	` + content

	resp, _ := llm.AskDeepSeek(prompt)
	resp = strings.TrimSpace(resp)

	if strings.HasPrefix(resp, "2") {
		return 2
	}
	if strings.HasPrefix(resp, "1") {
		return 1
	}
	return 0
}