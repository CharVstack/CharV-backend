package handler

import (
	"fmt"
	"github.com/CharVstack/CharV-backend/domain"
	"github.com/CharVstack/CharV-backend/openapi/v1"
	backendHost "github.com/CharVstack/CharV-backend/usecase/host"
	"os"
	"reflect"
	"testing"

	"github.com/CharVstack/CharV-lib/host"
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
	var getHostInfo = host.GetInfo()
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
