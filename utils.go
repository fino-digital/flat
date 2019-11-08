package qjson

import (
	"encoding/json"
)

// Unmarshal unmarshals the qjson in data into v
func Unmarshal(data []byte) (interface{}, error) {
	var parsed interface{}
	err := json.Unmarshal(data, &parsed)
	if err != nil {
		return nil, err
	}

	return Unflatten(parsed, nil)
}
