package eval

import "strings"

func splitIntoTokens(expression string) []string {
	tokens := []string{}

	start := 0
	for i, char := range expression {
		switch char {
		case '+', '-', '*', '/', '^', '(', ')':
			token := strings.TrimSpace(expression[start:i])
			if token != "" {
				tokens = append(tokens, token, string(char))
			} else {
				tokens = append(tokens, string(char))
			}

			start = i + 1
		}
	}

	token := strings.TrimSpace(expression[start:])
	if token != "" {
		tokens = append(tokens, token)
	}

	return tokens
}
