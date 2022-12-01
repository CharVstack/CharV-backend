package qemu

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/CharVstack/CharV-backend/domain/models"
	"github.com/google/uuid"
)

func TestParse(t *testing.T) {
	testVm := models.Vm{
		Name:   "ubuntu",
		Memory: 1024,
		Cpu:    1,
		Devices: models.Devices{
			Disk: []models.Disk{
				{
					Name:   "ubuntu",
					Pool:   "default",
					Type:   "qcow2",
					Device: "disk",
				},
			},
		},
		Metadata: models.Metadata{
			Id:         uuid.Must(uuid.Parse("0bfb8def-86ed-4b9d-8826-66a6ab1c1491")),
			ApiVersion: "v1",
		},
	}
	machine, err := parse("../../test/resources/machines/ubuntu.json")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(testVm, machine) {
		t.Fatalf("expected: %#v\nactually: %#v", testVm, machine)
	}
}

func TestGetAllVms(t *testing.T) {
	_, err := getAllVms("../../test/resources/machines/")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetVmPower(t *testing.T) {
	type args struct {
		id   uuid.UUID
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    models.VmPowerInfo
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "pass01",
			args: args{
				id:   uuid.Must(uuid.Parse("0bfb8def-86ed-4b9d-8826-66a6ab1c1491")),
				path: "../../test/resources/socks",
			},
			want: models.VmPowerInfo{
				CleanPowerOff: true,
				State:         "RUNNING",
			},
			wantErr: assert.NoError,
		},
		{
			name: "fail01",
			args: args{
				id:   uuid.Must(uuid.Parse("99999999-0000-4444-1111-112233445566")),
				path: "../../test/resources/socks",
			},
			want: models.VmPowerInfo{
				CleanPowerOff: true,
				State:         "SHUTDOWN",
			},
			wantErr: assert.NoError,
		},
		{
			name: "fail02",
			args: args{
				id:   uuid.Must(uuid.Parse("99999999-0000-4444-1111-112233445566")),
				path: "/no/exists/path",
			},
			want:    models.VmPowerInfo{},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetVmPower(tt.args.id, tt.args.path)
			if !tt.wantErr(t, err, fmt.Sprintf("GetVmPower(%v, %v)", tt.args.id, tt.args.path)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetVmPower(%v, %v)", tt.args.id, tt.args.path)
		})
	}
}

func Test_run(t *testing.T) {
	type args struct {
		cmd string
	}
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "pass01",
			args: args{
				cmd: "ls",
			},
			wantErr: assert.NoError,
		},
		{
			name: "fail01",
			args: args{
				cmd: "Non-Existent-Command",
			},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantErr(t, run(tt.args.cmd), fmt.Sprintf("run(%v)", tt.args.cmd))
		})
	}
}
