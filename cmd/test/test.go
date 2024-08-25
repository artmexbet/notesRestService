package main

import (
	"fmt"
	"notesRestService/internal/textValidator"
)

func main() {
	validator := textValidator.New()
	texts, err := validator.ValidateTexts("превед", "медвед")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(texts)
}
