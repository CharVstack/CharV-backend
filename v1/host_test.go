package openapi

import (
	"fmt"
	"github.com/CharVstack/CharV-lib/host"
	"github.com/joho/godotenv"
	"os"
	"reflect"
	"testing"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
}

func TransStruct(getInfo host.Host) (Cpu, Memory, []StoragePool) {
	cpuStruct := getCpuInfo(getInfo)
	memoryStruct := getMemoryInfo(getInfo)
	storageStruct := getStorageInfo(getInfo)

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
		want GetApiV1Host200Response
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
			want: GetApiV1Host200Response{
				Cpu:          cpuInfo,
				Mem:          memoryInfo,
				StoragePools: storageInfo,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetHostInfo(tt.args.getInfo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetHostInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
