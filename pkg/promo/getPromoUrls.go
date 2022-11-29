package promo

import (
	"fmt"
	"go.uber.org/zap"
	"io/fs"
	"metrics/internal/models"
	"path/filepath"
	"strings"
)

const searchFile = "index.html"

func (p *Promo) GetPromoUrls(label *models.Label) ([]models.Counter, error) {
	var counters []models.Counter

	//path := fmt.Sprintf("%s%s", p.path, domain.Name)
	path := strings.ReplaceAll(label.MetricName, "_", "/")
	path = fmt.Sprintf("%s%s", p.path, path)

	err := filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			p.lg.Error("GetPromoUrls func()",
				zap.Error(err))
			return nil
		}

		if info.Name() == searchFile {
			file := cutPath(path, p.path)
			counter := models.Counter{
				MetricName: file,
			}
			counters = append(counters, counter)

			return nil
		}

		return nil
	})

	if err != nil {
		p.lg.Error("GetPromoUrls filepath.Walk",
			zap.Error(err))
		return nil, err
	}

	return counters, nil
}

func cutPath(path, pathToDir string) string {

	str := strings.TrimRight(path, searchFile)
	_, after, _ := strings.Cut(str, pathToDir)
	return after
}
