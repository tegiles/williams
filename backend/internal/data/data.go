package data

import (
	"encoding/json"
	"os"
)

func ReadJSON[T interface{}](path string) (result T, err error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return
	}
	err = json.Unmarshal(file, &result)
	if err != nil {
		return
	}
	return
}
