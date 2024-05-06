package types

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
)

// follows github.com/sirupsen/logrus@v1.9.3/json_formatter.go
type CustomJSONFormatter struct{}

// Format renders a single log entry
func (f *CustomJSONFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	data := make(map[string]interface{}, len(entry.Data)+4)
	for k, v := range entry.Data {
		switch v := v.(type) {
		case error:
			// Otherwise errors are ignored by `encoding/json`
			// https://github.com/sirupsen/logrus/issues/137
			data[k] = v.Error()
		default:
			data[k] = v
		}
	}

	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	encoder := json.NewEncoder(b)
	encoder.SetIndent("", "  ")
	err := encoder.Encode(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal fields to JSON, %w", err)
	}

	return b.Bytes(), nil
}
