package unload

import (
	"fmt"
	"html/template"
	"metrics/internal/models"
	"os"
)

const (
	pathTemplate = "./templates/counter.txt"
	pathFile     = "./tmp/"
	fileExt      = ".js"
)

func Unload(name string, counters []models.Counter) (string, error) {

	tmpl, err := template.ParseFiles(pathTemplate)
	if err != nil {
		return "", err
	}

	fileName := fmt.Sprintf("%s%s%s", pathFile, name, fileExt)

	f, err := os.Create(fileName)
	if err != nil {
		return "", err
	}

	if err = tmpl.Execute(f, counters); err != nil {
		return "", err
	}

	return fileName, nil
}
