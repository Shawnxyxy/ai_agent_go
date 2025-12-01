package service

import (
	"fmt"
)

func SummarizeText(text string) string {
	return fmt.Sprintf("Summary: %s", text)
}