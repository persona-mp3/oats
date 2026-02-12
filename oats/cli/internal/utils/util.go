package utils

import (
	"encoding/json"
	"fmt"
)

func ToJson(data any) ([]byte, error) {
	content, err := json.Marshal(data)
	if err != nil {
		return []byte{}, fmt.Errorf("error in converting to_json: %w", err)
	}
	return content, nil
}
