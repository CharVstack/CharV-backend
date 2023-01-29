package qemu

import (
	"github.com/CharVstack/CharV-backend/entity"
	mock_models "github.com/CharVstack/CharV-backend/usecase/models/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"reflect"
	"testing"
)

func Test_qemuUseCase_Create(t *testing.T) {
	type fields struct {
		da   func(m *mock_models.MockVmDataAccess)
		id   func(m *mock_models.MockID)
		cmd  func(m *mock_models.MockCommand)
		disk func(m *mock_models.MockDisk)
	}
	type args struct {
		req entity.VmCore
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.Vm
		wantErr bool
	}{
		{
			name: "PASS01",
			fields: fields{
				da: func(m *mock_models.MockVmDataAccess) {
					m.EXPECT().Add(entity.Vm{
						VmCore: entity.VmCore{
							Cpu:    2,
							Memory: 1024,
							Name:   "ubuntu",
						},
						ID: uuid.Must(uuid.Parse("0bfb8def-86ed-4b9d-8826-66a6ab1c1491")),
						Devices: entity.Devices{
							OS: entity.Disk{
								Device: entity.DiskDeviceCdrom,
								Name:   "ubuntu-20.04.5-live-server-amd64.iso",
								Pool:   "default",
								Type:   entity.DiskTypeIso,
							},
							Disk: []entity.Disk{
								{
									Device: entity.DiskDeviceDisk,
									Name:   "0bfb8def-86ed-4b9d-8826-66a6ab1c1491",
									Pool:   "default",
									Type:   entity.DiskTypeQcow2,
								},
							},
						},
						Boot:           entity.BootDeviceDisk,
						Virtualization: entity.VirtualizationTypeKvm,
						Daemonize:      true,
					}).Return(nil)
				},
				id: func(m *mock_models.MockID) {
					m.EXPECT().GenID().Return(uuid.Must(uuid.Parse("0bfb8def-86ed-4b9d-8826-66a6ab1c1491")), nil)
				},
				cmd: func(m *mock_models.MockCommand) {
					m.EXPECT().Run("qemu-system-x86_64", []string{"-name", "ubuntu", "-m", "1024", "-smp", "2", "-drive", "file=/var/lib/charv/images/0bfb8def-86ed-4b9d-8826-66a6ab1c1491.qcow2,format=qcow2", "-cdrom", "/var/lib/charv/images/ubuntu-20.04.5-live-server-amd64.iso", "-boot", "order=d", "-qmp", "unix:/0bfb8def-86ed-4b9d-8826-66a6ab1c1491.sock,server,nowait", "-vnc", "unix:/0bfb8def-86ed-4b9d-8826-66a6ab1c1491.sock", "-accel", "kvm", "-daemonize"}).Return(nil)
				},
				disk: func(m *mock_models.MockDisk) {
					m.EXPECT().Create("default.json", uuid.Must(uuid.Parse("0bfb8def-86ed-4b9d-8826-66a6ab1c1491"))).Return(nil)
				},
			},
			args: args{req: entity.VmCore{
				Cpu:    2,
				Memory: 1024,
				Name:   "ubuntu",
			}},
			want: entity.Vm{
				VmCore: entity.VmCore{
					Cpu:    2,
					Memory: 1024,
					Name:   "ubuntu",
				},
				ID: uuid.Must(uuid.Parse("0bfb8def-86ed-4b9d-8826-66a6ab1c1491")),
				Devices: entity.Devices{
					OS: entity.Disk{
						Device: entity.DiskDeviceCdrom,
						Name:   "ubuntu-20.04.5-live-server-amd64.iso",
						Pool:   "default",
						Type:   entity.DiskTypeIso,
					},
					Disk: []entity.Disk{
						{
							Device: entity.DiskDeviceDisk,
							Name:   "0bfb8def-86ed-4b9d-8826-66a6ab1c1491",
							Pool:   "default",
							Type:   entity.DiskTypeQcow2,
						},
					},
				},
				Boot:           entity.BootDeviceDisk,
				Virtualization: entity.VirtualizationTypeKvm,
				Daemonize:      true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			id := mock_models.NewMockID(ctrl)
			da := mock_models.NewMockVmDataAccess(ctrl)
			cmd := mock_models.NewMockCommand(ctrl)
			disk := mock_models.NewMockDisk(ctrl)
			tt.fields.da(da)
			tt.fields.id(id)
			tt.fields.cmd(cmd)
			tt.fields.disk(disk)

			q := qemuUseCase{
				da:   da,
				id:   id,
				cmd:  cmd,
				disk: disk,
			}
			got, err := q.Create(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}
