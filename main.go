package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"struct_example/note"
)

func main() {

	title, content := getNoteData()
	Usernote, err := note.New(title, content)

	if err != nil {
		fmt.Println("Error creating note:", err)
		return
	}

	Usernote.Display()
	err = Usernote.Save()

	if err != nil {
		fmt.Println("Error saving note:", err)
		return
	}

	fmt.Println("Note saved successfully!")

}

func getUserInput(prompt string) string {
	fmt.Println(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, error := reader.ReadString('\n')
	if error != nil {
		fmt.Println("Error reading input:", error)
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

func getNoteData() (string, string) {
	title := getUserInput("Note:Title")
	content := getUserInput("Note:Content")
	return title, content
}
