package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

func main() {
	title, content := getNoteData()

	noteData, err := note.New(title, content)

	if err != nil {
		fmt.Print(err)
		return
	}
	note.Note.Display(noteData)
}

func getNoteData() (string, string) {
	title := getUserInput("Enter the title:")
	content := getUserInput("Enter the content:")
	return title, content
}

func getUserInput(text string) string {

	fmt.Print(text)
	// var value string

	// fmt.Scanln(&value)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}
