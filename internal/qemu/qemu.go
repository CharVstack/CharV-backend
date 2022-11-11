package qemu

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"

	"github.com/CharVstack/CharV-backend/domain/models"

	"github.com/google/uuid"
)

func getVmInfo(opts InstallOpts, filePath string) (vmInfo models.Vm, err error) {
	uuidObj, err := uuid.NewRandom()
	if err != nil {
		return models.Vm{}, err
	}

	var diskType models.DiskType
	diskType, err = CheckFileType(filePath)
	if err != nil {
		return models.Vm{}, err
	}
	typeMap := map[models.DiskType]string{models.DiskTypeQcow2: ".qcow2"}

	vmInfo = models.Vm{
		Devices: models.Devices{
			Disk: []models.Disk{
				{
					Type:   diskType,
					Device: models.DiskDeviceDisk,
					Path:   "/var/lib/charVstack/image/" + opts.Disk + typeMap[diskType],
				},
			},
		},
		Memory: opts.Memory,
		Metadata: models.Metadata{
			ApiVersion: "v1",
			Id:         uuidObj,
		},
		Name: opts.Name,
		Vcpu: opts.VCpu,
	}

	rawVmInfo, err := json.Marshal(vmInfo)
	if err != nil {
		return models.Vm{}, err
	}

	createJSONPath := "/var/lib/charVstack/machines/"

	fileName := createJSONPath + vmInfo.Name + "-" + vmInfo.Metadata.Id.String() + ".json"

	createFile, err := os.Create(fileName)
	if err != nil {
		return models.Vm{}, err
	}
	defer func() {
		err = createFile.Close()
	}()

	_, err = createFile.Write(rawVmInfo)
	if err != nil {
		return models.Vm{}, err
	}

	return vmInfo, nil
}

func CheckFileType(filePath string) (models.DiskType, error) {
	buf, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	if !bytes.Equal(buf[:4], []byte("QFI\xfb")) {
		return "", errors.New("Not QEMU QCOW Image (v3) ")
	}

	return models.DiskTypeQcow2, nil
}

func getAllVms(directoryPath string) ([]models.Vm, error) {
	var vmList []models.Vm
	dir, err := os.ReadDir(directoryPath)
	if err != nil {
		return []models.Vm{}, err
	}

	for _, file := range dir {
		raw, err := os.ReadFile(directoryPath + file.Name())
		if err != nil {
			return []models.Vm{}, err
		}

		var vm models.Vm
		err = json.Unmarshal(raw, &vm)
		if err != nil {
			return []models.Vm{}, err
		}

		vmList = append(vmList, vm)

	}
	return vmList, err
}
