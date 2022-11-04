package qemu

import "github.com/CharVstack/CharV-backend/domain/models"

func GetAllVmInfo() ([]models.Vm, error) {
	path := "/var/lib/charVstack/machines/"
	vms, err := ConvertToStruct(path)
	if err != nil {
		return []models.Vm{}, err
	}
	return vms, nil
}
