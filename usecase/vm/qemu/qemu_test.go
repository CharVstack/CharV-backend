package qemu

import (
	"github.com/CharVstack/CharV-backend/entity"
	"github.com/CharVstack/CharV-backend/infrastructure/system"
	"github.com/CharVstack/CharV-backend/usecase/models"
	"reflect"
	"testing"
)

func Test_qemuUseCase_genArgs(t *testing.T) {
	type fields struct {
		da   models.VmDataAccess
		id   models.ID
		cmd  models.Command
		disk models.Disk
		vnc  models.Socket
		sys  system.Paths
	}
	type args struct {
		vm *entity.Vm
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := qemuUseCase{
				da:    tt.fields.da,
				id:    tt.fields.id,
				cmd:   tt.fields.cmd,
				disk:  tt.fields.disk,
				vnc:   tt.fields.vnc,
				paths: tt.fields.sys,
			}
			if got := q.genArgs(tt.args.vm); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("genArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
