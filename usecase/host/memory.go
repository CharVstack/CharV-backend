package host

import (
	backendModels "github.com/CharVstack/CharV-backend/domain/models"
	libModels "github.com/CharVstack/CharV-lib/domain/models"
)

func GetMemoryInfo(getInfo libModels.Host) backendModels.Memory {
	memoryInfo := backendModels.Memory{
		Total:       backendModels.TypeUint64(getInfo.Mem.Total),
		Used:        backendModels.TypeUint64(getInfo.Mem.Used),
		Free:        backendModels.TypeUint64(getInfo.Mem.Free),
		UsedPercent: backendModels.TypeFloat64(getInfo.Mem.UsedPercent),
	}

	return memoryInfo
}
