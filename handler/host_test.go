package handler

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/CharVstack/CharV-backend/domain"
	backendModels "github.com/CharVstack/CharV-backend/domain/models"
	backendHost "github.com/CharVstack/CharV-backend/usecase/host"
	libModels "github.com/CharVstack/CharV-lib/domain/models"
	libHost "github.com/CharVstack/CharV-lib/pkg/host"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
}

func TransStruct(getInfo libModels.Host) (backendModels.Cpu, backendModels.Memory, []backendModels.StoragePool) {
	cpuStruct := backendHost.GetCpuInfo(getInfo)
	memoryStruct := backendHost.GetMemoryInfo(getInfo)
	storageStruct := backendHost.GetStorageInfo(getInfo)

	return cpuStruct, memoryStruct, storageStruct
}

func TestGetHostInfo(t *testing.T) {
	storageDirEnv := os.Getenv("STORAGE_DIR")
	storageDir := libModels.GetInfoOptions{
		StorageDir: storageDirEnv,
	}
	var getHostInfo, err = libHost.GetInfo(storageDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	var cpuInfo, memoryInfo, storageInfo = TransStruct(getHostInfo)
	type args struct {
		getInfo libModels.Host
	}
	tests := []struct {
		name string
		args args
		want domain.GetApiV1Host200Response
	}{
		{
			name: "テストデータ",
			args: args{
				getInfo: libModels.Host{
					Cpu:          getHostInfo.Cpu,
					Memory:       getHostInfo.Memory,
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
