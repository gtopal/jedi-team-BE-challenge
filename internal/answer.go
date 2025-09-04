package internal

import (
	"fmt"
	"os"
	"strings"
)

func FindAnswer(question string) (string, bool) {
	data, err := os.ReadFile("data.md")

	fmt.Printf("the error is %v\n", err)
	if err != nil {
		return "", false
	}
	lines := strings.Split(string(data), "\n")
	qLower := strings.ToLower(question)
	for _, line := range lines {
		if strings.Contains(strings.ToLower(line), qLower) {
			return line, true
		}
	}
	return "", false
}
