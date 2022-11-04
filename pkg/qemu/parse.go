package qemu

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/CharVstack/CharV-backend/domain/models"
)

func parse(path string) (models.Vm, error) {
	var machine models.Vm
	abspath, err := filepath.Abs(path)
	f, err := os.Open(abspath)
	if err != nil {
		return models.Vm{}, err
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return models.Vm{}, err
	}
	if err := json.Unmarshal(b, &machine); err != nil {
		return models.Vm{}, err
	}
	return machine, nil
}
