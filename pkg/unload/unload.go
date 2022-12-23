package unload

import (
	"fmt"
	"html/template"
	"metrics/internal/models"
	"os"
	"strings"
)

const (
	pathTemplate = "./templates/counter.txt"
	pathFile     = "./tmp/"
	fileExt      = ".js"
)

func Unload(name string, counters []models.Counter) (string, error) {

	//d := strings.Replace(name, "_", "/", -1)
	labelNames := strings.Split(name, "_")
	domain := labelNames[0]

	for i, counter := range counters {
		promo := strings.Replace(counter.MetricName, domain, "", -1)
		counters[i].MetricName = promo
	}

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
