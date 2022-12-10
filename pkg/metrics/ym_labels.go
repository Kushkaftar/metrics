package metrics

import (
	"encoding/json"
	"fmt"
	"metrics/internal/models"
	"net/http"
)

const (
	labelsPath = "/management/v1/labels"
	labelPath  = "/label/"
)

func (m *Metrics) GetAllLabels() ([]models.Label, error) {
	url, err := m.url.GetUrl(labelsPath)
	if err != nil {
		return nil, m.mErr(urlErr, err)
	}

	resp, err := m.client.Get(url)
	if err != nil {
		return nil, m.mErr(clientErr, err)
	}

	var labels Labels

	if err = json.Unmarshal(resp.Body, &labels); err != nil {
		return nil, m.mErr(jsonErr, err)
	}

	var modeLabels []models.Label

	for _, label := range labels.Labels {
		modeLabel := models.Label{
			MetricID:   label.Id,
			MetricName: label.Name,
		}
		modeLabels = append(modeLabels, modeLabel)
	}
	return modeLabels, nil
}

func (m *Metrics) CreateLabel(label *models.Label) error {
	var createLabel CreateLabel
	createLabel.Label.Name = label.MetricName

	jsonData, err := json.Marshal(createLabel)
	if err != nil {
		return m.mErr(jsonErr, err)
	}

	url, err := m.url.GetUrl(labelsPath)
	if err != nil {
		return m.mErr(urlErr, err)
	}

	resp, err := m.client.Post(url, jsonData)
	if err != nil {
		return m.mErr(clientErr, err)
	}

	var respLabel CreateLabel

	if err = json.Unmarshal(resp.Body, &respLabel); err != nil {
		return m.mErr(jsonErr, err)
	}

	label.MetricID = respLabel.Label.Id

	return nil
}

func (m *Metrics) SetLabelInCounter(counter *models.Counter, label *models.Label) error {
	path := fmt.Sprintf("%s%d%s%d", counterPath, counter.MetricID, labelPath, label.MetricID)
	url, err := m.url.GetUrl(path)
	if err != nil {
		return m.mErr(urlErr, err)
	}

	resp, err := m.client.Post(url, nil)
	if err != nil {
		return m.mErr(clientErr, err)
	}

	if resp.StatusCode != http.StatusOK {
		str := "status code ot 200"
		err := fmt.Errorf("%s", str)
		return m.mErr(str, err)
	}
	type success struct {
		Success bool
	}
	s := success{}

	if err = json.Unmarshal(resp.Body, &s); err != nil {
		return m.mErr(jsonErr, err)
	}

	if !s.Success {
		str := "success not true"
		err := fmt.Errorf("%s", str)
		return m.mErr(string(resp.Body), err)
	}

	counter.LabelID = label.MetricID

	return nil
}
