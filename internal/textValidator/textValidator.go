package textValidator

import (
	"encoding/json"
	"net/http"
	"net/url"
	"notesRestService/internal/models"
)

type TextValidator struct {
}

func New() *TextValidator {
	return &TextValidator{}
}

func (t *TextValidator) ValidateTexts(text ...string) ([][]models.SpellerJSON, error) {
	formData := url.Values{
		"text":   text,
		"format": {"plain"},
	}

	resp, err := http.PostForm("https://speller.yandex.net/services/spellservice.json/checkTexts", formData)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var spellerJSON [][]models.SpellerJSON
	err = json.NewDecoder(resp.Body).Decode(&spellerJSON)
	if err != nil {
		return nil, err
	}

	return spellerJSON, nil
}
