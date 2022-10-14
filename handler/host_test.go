package handler

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/CharVstack/CharV-backend/domain"
	"github.com/CharVstack/CharV-backend/openapi/v1"
	backendHost "github.com/CharVstack/CharV-backend/usecase/host"
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

func TransStruct(getInfo host.Host) (openapi.Cpu, openapi.Memory, []openapi.StoragePool) {
	cpuStruct := backendHost.GetCpuInfo(getInfo)
	memoryStruct := backendHost.GetMemoryInfo(getInfo)
	storageStruct := backendHost.GetStorageInfo(getInfo)

	return cpuStruct, memoryStruct, storageStruct
}

func TestGetHostInfo(t *testing.T) {
	storageDirEnv := os.Getenv("STORAGE_DIR")
	storageDir := host.GetInfoOptions{
		StorageDir: storageDirEnv,
	}
	var getHostInfo = host.GetInfo(storageDir)
	var cpuInfo, memoryInfo, storageInfo = TransStruct(getHostInfo)
	type args struct {
		getInfo host.Host
	}
	tests := []struct {
		name string
		args args
		want domain.GetApiV1Host200Response
	}{
		{
			name: "テストデータ",
			args: args{
				getInfo: host.Host{
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
