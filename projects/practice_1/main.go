package main

import (
	"fmt"

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
	var value string

	fmt.Scanln(&value)
	return value
}
