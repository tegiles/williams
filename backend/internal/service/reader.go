package service

import (
	"encoding/json"
	"os"
)

// ReadJSON reads in a given json file and outputs it in a given type
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
