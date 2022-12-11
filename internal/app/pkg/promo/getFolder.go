package promo

import "os"

func getFolder(path string) ([]string, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var domains []string

	for _, file := range files {
		if file.IsDir() {
			domains = append(domains, file.Name())
		}

	}
	return domains, nil
}
