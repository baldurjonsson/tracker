package action

import (
	"bufio"
	"fmt"
	"strings"
)

func prompt(b *bufio.Reader, prompt string) string {
	fmt.Print(prompt, ": ")
	return readString(b)
}

func readString(b *bufio.Reader) string {
	str, _ := b.ReadString('\n')
	str = strings.Trim(str, "\r\n ")
	return str
}
