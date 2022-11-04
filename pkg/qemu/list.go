package qemu

import (
	"encoding/json"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/CharVstack/CharV-backend/domain/models"
)

func GetRunningList() ([]string, error) {
	out, err := exec.Command("ps", "-o", "cmd", "axh").Output()
	if err != nil {
		return nil, err
	}
	sout := string(out)
	processes := strings.Split(sout, "\n")
	var guests []string
	for _, v := range processes {
		executable := strings.Split(v, " ")[0]
		r := regexp.MustCompile(`qemu-system-x86_64`)
		if r.MatchString(executable) {
			guests = append(guests, executable)
		}
	}
	return guests, nil
}

func ConvertToStruct(directoryPath string) ([]models.Vm, error) {
	var resJSONList []models.Vm
	dir, err := os.ReadDir(directoryPath)
	if err != nil {
		return []models.Vm{}, err
	}

	var resJSON models.Vm
	var raw []byte

	for _, file := range dir {
		raw, err = os.ReadFile(directoryPath + file.Name())
		if err != nil {
			return []models.Vm{}, err
		}

		err = json.Unmarshal(raw, &resJSON)
		if err != nil {
			return []models.Vm{}, err
		}

		resJSONList = append(resJSONList, resJSON)

	}
	return resJSONList, err
}
