package internal

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

// FindAnswer finds the most relevant answer in data.md using simple word overlap + cosine similarity
func FindAnswer(question string) (string, bool) {
	data, err := os.ReadFile("data.md")
	if err != nil {
		fmt.Printf("error reading data.md: %v\n", err)
		return "", false
	}

	lines := strings.Split(string(data), "\n")
	fmt.Printf("The type of lines is %T\n", lines)
	cleaned := []string{}
	for _, line := range lines {
		line = strings.Trim(line, " |")
		if line != "" {
			cleaned = append(cleaned, line)
		}
	}

	if len(cleaned) == 0 {
		return "", false
	}

	// Convert question into vector (word counts)
	qVec := textToVector(question)

	bestScore := -1.0
	bestLine := ""

	// Compare question against each line
	for _, line := range cleaned {
		lineVec := textToVector(line)
		score := cosineSimilarity(qVec, lineVec)
		if score > bestScore {
			bestScore = score
			bestLine = line
		}
	}

	// Require a minimum similarity to avoid nonsense matches
	if bestScore > 0.1 {
		return bestLine, true
	}
	return "", false
}

// textToVector converts a string into a bag-of-words vector
func textToVector(s string) map[string]float64 {
	words := tokenize(s)
	vec := make(map[string]float64)
	for _, w := range words {
		vec[w]++
	}
	return vec
}

// tokenize lowercases, strips punctuation, and splits on spaces
func tokenize(s string) []string {
	s = strings.ToLower(s)
	var b strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			b.WriteRune(r)
		} else {
			b.WriteRune(' ')
		}
	}
	return strings.Fields(b.String())
}

// cosineSimilarity between two bag-of-words vectors
func cosineSimilarity(a, b map[string]float64) float64 {
	var dot, normA, normB float64
	for k, av := range a {
		bv := b[k]
		dot += av * bv
		normA += av * av
	}
	for _, bv := range b {
		normB += bv * bv
	}
	if normA == 0 || normB == 0 {
		return 0
	}
	return dot / (sqrt(normA) * sqrt(normB))
}

// simple sqrt (Newton’s method)
func sqrt(x float64) float64 {
	z := x
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}
