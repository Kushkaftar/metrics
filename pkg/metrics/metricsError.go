package metrics

import (
	"fmt"
	"go.uber.org/zap"
)

const (
	clientErr = "client return error"
	urlErr    = "url is broken"
	jsonErr   = "error unmarshal json"
)

func (m *Metrics) mErr(str string, err error) error {
	m.lg.Error(str, zap.Error(err))

	return fmt.Errorf("%s", str)
}
