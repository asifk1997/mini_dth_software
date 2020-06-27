package input_output

import (
	"bufio"
	"os"
	"runtime"
	"strings"
)

//input from takeinput function consists of newline and whitespace
func removeWhiteSpace(input string) string {
	if runtime.GOOS == "windows" {
		input = strings.TrimRight(input, "\r\n")
	} else {
		input = strings.TrimRight(input, "\n")
	}
	return input
}

// Take input function used bufio reader for console
// input
func TakeInput() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = removeWhiteSpace(text)
	return text
}
