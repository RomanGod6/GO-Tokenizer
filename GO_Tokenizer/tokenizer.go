package GO_Tokenizer

import (
	"log"
	"strings"
	"unicode"
)

// Settings holds the configuration for token limits
type Settings struct {
	MaxNewTokens  int
	ContextWindow int
}

// TokenizeInput checks the total token count of JSON, system, and user prompts, and identifies unnecessary spaces.
func TokenizeInput(jsonData string, systemPrompt string, userPrompt string, settings Settings) (int, int, map[string]interface{}, error) {
	// Detect and report unnecessary spaces
	reportUnnecessarySpaces(jsonData)

	// Token count for JSON, system, and user prompts
	jsonTokenCount := len(strings.Fields(jsonData))
	systemTokenCount := len(strings.Fields(systemPrompt))
	userTokenCount := len(strings.Fields(userPrompt))

	// Total token count
	totalTokenCount := jsonTokenCount + systemTokenCount + userTokenCount + settings.MaxNewTokens

	// Check if the total tokens exceed the context window
	if totalTokenCount > settings.ContextWindow {
		log.Printf("Warning: Total token count %d exceeds the context window of %d. Consider splitting the document.\n", totalTokenCount, settings.ContextWindow)
	}

	// Return useful info
	additionalInfo := map[string]interface{}{
		"json_token_count":   jsonTokenCount,
		"system_token_count": systemTokenCount,
		"user_token_count":   userTokenCount,
		"total_token_count":  totalTokenCount,
	}

	return totalTokenCount, settings.ContextWindow, additionalInfo, nil
}

// reportUnnecessarySpaces detects and reports unnecessary spaces in a string.
func reportUnnecessarySpaces(input string) {
	words := strings.Fields(input) // Split by any whitespace
	reconstructed := strings.Join(words, " ")

	if len(input) != len(reconstructed) {
		log.Printf("Unnecessary spaces detected. Original length: %d, Optimized length: %d", len(input), len(reconstructed))
	}

	// Additional detection logic
	if strings.HasPrefix(input, " ") || strings.HasSuffix(input, " ") {
		log.Println("Leading or trailing spaces detected.")
	}

	// Check for spaces before punctuation
	for i := 0; i < len(input)-1; i++ {
		if unicode.IsSpace(rune(input[i])) && unicode.IsPunct(rune(input[i+1])) {
			log.Printf("Unnecessary space before punctuation detected at position %d: '%s'", i, string(input[i:i+2]))
		}
	}
}
