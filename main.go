package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

const (
	unicorn_text = `
		\
		\je
		/.(ss
		(,/"(y(__,--.
			\  ) _( /{
			!|| " :||
			!||   :||
			'''   '''

`
	line       = "----------------------------------------"
	lineLength = 40
)

func isInputFromPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode()&os.ModeCharDevice == 0
}

func readPipedData() string {
	builder := strings.Builder{}
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		builder.WriteString(scanner.Text())
	}

	return builder.String()
}

func spacePadSides(text string, padding int) string {
	leftPadded := fmt.Sprintf("%*s", padding+len(text), text)
	// Compromise here is that we're just going to pad the one
	// side in the event that it's too long ;)
	// Oh, and check if the result is too long as well... Then
	// we're cutting off some text... Centering is quite hard,
	// who'd have thunk?
	if len(leftPadded) >= lineLength - 2 {
		return leftPadded
	}
	result := fmt.Sprintf("%-*s", padding+len(leftPadded), leftPadded)
	if len(result) > lineLength - 2 {
		return result[:len(result)-1]
	}
	return result
}

func splitMessageIntoLines(message string) []string {
	result := []string{}
	words := strings.Split(message, " ")
	line := words[0]
	for i := 1; i < len(words); i++ {
		if (len(line) + len(words[i]) + 1) > (lineLength - 2) {
			result = append(result, line)
			line = words[i]
		} else {
			line += " " + words[i]
		}
	}
	if len(line) > 0 {
		result = append(result, line)
	}

	return result
}

// Assuming that the maxLength will never be less
// than the text
func findPadding(text string, maxLength int) int {
	return int(math.Ceil(float64(maxLength - len(text)) / 2))
}

func createMessageBox(message string) string {
	builder := strings.Builder{}
	builder.WriteString(line + "\n")

	lines := splitMessageIntoLines(message)
	for _, line := range lines {
		padding := findPadding(line, lineLength-2)
		paddedText := spacePadSides(line, padding)
		// Adding pipes on either side for box effect, can change
		builder.WriteString(fmt.Sprintf("|%s|\n", paddedText))
	}

	builder.WriteString(line + "\n")

	return builder.String()
}

func main() {
	say := ""
	if isInputFromPipe() {
		say = readPipedData()
	} else {
		say = strings.Join(os.Args[1:], " ")
	}
	if say == "" {
		say = "I have no words..."
	}

	messageBox := createMessageBox(say)
	fmt.Println(messageBox)

	// Unicorn is rendering differently?
	// Maybe it's a terminal thing?
	fmt.Print(unicorn_text)
}
