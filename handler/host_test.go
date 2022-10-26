package handler

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/CharVstack/CharV-backend/domain"
	backendHost "github.com/CharVstack/CharV-backend/usecase/host"
	"github.com/CharVstack/CharV-lib/domain/models"
	"github.com/CharVstack/CharV-lib/pkg/host"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
}

func TransStruct(getInfo models.Host) (models.Cpu, models.Memory, []models.StoragePool) {
	cpuStruct := backendHost.GetCpuInfo(getInfo)
	memoryStruct := backendHost.GetMemoryInfo(getInfo)
	storageStruct := backendHost.GetStorageInfo(getInfo)

	return cpuStruct, memoryStruct, storageStruct
}

func TestGetHostInfo(t *testing.T) {
	storageDirEnv := os.Getenv("STORAGE_DIR")
	storageDir := models.GetInfoOptions{
		StorageDir: storageDirEnv,
	}
	getHostInfo, err := host.GetInfo(storageDir)
	if err != nil {
		t.Fatal(err)
	}
	var cpuInfo, memoryInfo, storageInfo = TransStruct(getHostInfo)
	type args struct {
		getInfo models.Host
	}
	tests := []struct {
		name string
		args args
		want domain.GetApiV1Host200Response
	}{
		{
			name: "テストデータ",
			args: args{
				getInfo: models.Host{
					Cpu:          getHostInfo.Cpu,
					Mem:          getHostInfo.Mem,
					StoragePools: getHostInfo.StoragePools,
				},
			},
			want: domain.GetApiV1Host200Response{
				Cpu:          cpuInfo,
				Mem:          memoryInfo,
				StoragePools: storageInfo,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backendHost.GetHostInfo(tt.args.getInfo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHostInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
