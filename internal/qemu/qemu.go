package qemu

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/CharVstack/CharV-backend/domain/models"

	"github.com/google/uuid"
)

func createInfoJSON(opts InstallOpts, filePath string) (models.Vm, error) {
	uuidObj, err := uuid.NewRandom()
	if err != nil {
		return models.Vm{}, err
	}

	var diskType models.DiskType
	diskType, err = CheckFileType(filePath)
	if err != nil {
		fmt.Println(err)
	}
	typeMap := map[models.DiskType]string{models.DiskTypeQcow2: ".qcow2"}

	vmInfo := models.Vm{
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

	var MarshalJSON []byte
	MarshalJSON, err = json.Marshal(vmInfo)
	if err != nil {
		return models.Vm{}, err
	}

	createJSONPath := "/var/lib/charVstack/machines/"

	fileName := createJSONPath + vmInfo.Name + "-" + vmInfo.Metadata.Id.String() + ".json"

	var createFile *os.File
	createFile, err = os.Create(fileName)
	if err != nil {
		return models.Vm{}, err
	}
	defer func() {
		err = createFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	_, err = createFile.Write(MarshalJSON)
	if err != nil {
		return models.Vm{}, err
	}

	return vmInfo, err
}

func CheckFileType(filePath string) (models.DiskType, error) {
	buf, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	readFile := bytes.NewReader(buf)
	if !bytes.Equal(ReadBinaryFile(readFile, 4), []byte("QFI\xfb")) {
		return "", errors.New("Not QEMU QCOW Image (v3) ")
	}

	return models.DiskTypeQcow2, nil
}

func ReadBinaryFile(readFile io.Reader, index int) []byte {
	buf := make([]byte, index)
	cnt, err := readFile.Read(buf)
	if err != nil || index != cnt {
		return []byte{}
	}
	return buf
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
