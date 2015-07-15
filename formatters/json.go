package formatters

import (
	"encoding/json"
)

type JsonFormatter struct{}

func (f *JsonFormatter) FormatOutput(model interface{}) (res string, err error) {
	byteRes, err := json.MarshalIndent(model, "", "    ")
	return string(byteRes), err
}