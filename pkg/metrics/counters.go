package metrics

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"metrics/internal/models"
	"net/http"
)

const (
	counterPath                 = "/management/v1/counter/"
	countersPath                = "/management/v1/counters"
	searchNameCounterQueryParam = "search_string"
)

func (m *Metrics) GetCounter(counter *models.Counter) error {

	queryParam := make(map[string]string)

	queryParam[searchNameCounterQueryParam] = counter.MetricName

	url, err := m.url.AddQueryParams(countersPath, queryParam)
	if err != nil {
		return m.mErr(urlErr, err)
	}

	resp, err := m.client.Get(url)
	if err != nil {
		return m.mErr(clientErr, err)
	}

	var cResp CounterRESP

	if err = json.Unmarshal(resp.Body, &cResp); err != nil {
		return m.mErr(jsonErr, err)
	}

	// todo refactor хуйня!!!
	if cResp.Rows == 0 {
		str := "counter not found"
		err := fmt.Errorf("%s", str)
		return m.mErr(str, err)
	}
	if cResp.Rows != 1 {
		m.lg.Info("cResp.Rows != 1",
			zap.String("response metric", fmt.Sprintf("%+v", cResp)))
		str := "metric return more one counters"
		err := fmt.Errorf("%s", str)
		return m.mErr(str, err)
	}

	// todo refactor add label id
	counter.MetricID = cResp.Counters[0].Id

	return nil
}

func (m *Metrics) CreateCounter(counter *models.Counter) error {
	url, err := m.url.GetUrl(countersPath)
	if err != nil {
		return m.mErr(urlErr, err)
	}

	createCounter := newCounter(counter.MetricName)
	jsonData, err := json.Marshal(createCounter)
	if err != nil {
		return m.mErr(jsonErr, err)
	}

	resp, err := m.client.Post(url, jsonData)
	if err != nil {
		return m.mErr(clientErr, err)
	}

	var cResp CreateCounter

	if err = json.Unmarshal(resp.Body, &cResp); err != nil {
		return m.mErr(jsonErr, err)
	}
	counter.MetricID = cResp.Counter.Id

	return nil
}

func (m *Metrics) GetCounters(countersName string) ([]models.Counter, error) {

	queryParam := make(map[string]string)

	queryParam[searchNameCounterQueryParam] = countersName

	url, err := m.url.AddQueryParams(countersPath, queryParam)
	if err != nil {
		return nil, m.mErr(urlErr, err)
	}

	resp, err := m.client.Get(url)
	if err != nil {
		return nil, m.mErr(clientErr, err)
	}

	var cResp CounterRESP

	if err = json.Unmarshal(resp.Body, &cResp); err != nil {
		return nil, m.mErr(jsonErr, err)
	}

	var counters []models.Counter

	for _, counter := range cResp.Counters {
		ctr := models.Counter{
			MetricID:   counter.Id,
			MetricName: counter.Name,
		}
		counters = append(counters, ctr)
	}

	return counters, nil
}

func (m *Metrics) DelCounter(counter *models.Counter) error {

	path := fmt.Sprintf("%s%d", counterPath, counter.MetricID)
	url, err := m.url.GetUrl(path)
	if err != nil {
		return m.mErr(urlErr, err)
	}

	resp, err := m.client.Delete(url)
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
	return nil
}
