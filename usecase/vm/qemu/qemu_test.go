package qemu

import (
	"errors"
	"reflect"
	"testing"

	"github.com/CharVstack/CharV-backend/entity"
	mock_models "github.com/CharVstack/CharV-backend/usecase/models/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

var (
	testVM = entity.Vm{
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
	}
	id = uuid.Must(uuid.Parse("0bfb8def-86ed-4b9d-8826-66a6ab1c1491"))
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
					m.EXPECT().Add(testVM).Return(nil)
				},
				id: func(m *mock_models.MockID) {
					m.EXPECT().GenID().Return(id, nil)
				},
				cmd: func(m *mock_models.MockCommand) {
					m.EXPECT().Run("qemu-system-x86_64", []string{"-name", "ubuntu", "-m", "1024", "-smp", "2", "-drive", "file=/var/lib/charv/images/0bfb8def-86ed-4b9d-8826-66a6ab1c1491.qcow2,format=qcow2", "-cdrom", "/var/lib/charv/images/ubuntu-20.04.5-live-server-amd64.iso", "-boot", "order=d", "-qmp", "unix:/0bfb8def-86ed-4b9d-8826-66a6ab1c1491.sock,server,nowait", "-vnc", "unix:/0bfb8def-86ed-4b9d-8826-66a6ab1c1491.sock", "-accel", "kvm", "-daemonize"}).Return(nil)
				},
				disk: func(m *mock_models.MockDisk) {
					m.EXPECT().Create("default.json", id).Return(nil)
				},
			},
			args: args{req: entity.VmCore{
				Cpu:    2,
				Memory: 1024,
				Name:   "ubuntu",
			}},
			want:    testVM,
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

func Test_qemuUseCase_ReadAll(t *testing.T) {
	type fields struct {
		da func(m *mock_models.MockVmDataAccess)
	}
	tests := []struct {
		name    string
		fields  fields
		want    []entity.Vm
		wantErr bool
	}{
		{
			name: "PASS01",
			fields: fields{da: func(m *mock_models.MockVmDataAccess) {
				m.EXPECT().Browse().Return([]entity.Vm{
					testVM,
				}, nil)
			}},
			want: []entity.Vm{
				testVM,
			},
			wantErr: false,
		},
		{
			name: "FAIL01",
			fields: fields{da: func(m *mock_models.MockVmDataAccess) {
				m.EXPECT().Browse().Return([]entity.Vm{}, errors.New("something wrong"))
			}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			da := mock_models.NewMockVmDataAccess(ctrl)
			tt.fields.da(da)

			q := qemuUseCase{
				da: da,
			}
			got, err := q.ReadAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_qemuUseCase_ReadById(t *testing.T) {
	type fields struct {
		da func(m *mock_models.MockVmDataAccess)
	}
	type args struct {
		id uuid.UUID
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
			fields: fields{da: func(m *mock_models.MockVmDataAccess) {
				m.EXPECT().Read(id).Return(testVM, nil)
			}},
			args:    args{id: id},
			want:    testVM,
			wantErr: false,
		},
		{
			name: "FAIL01",
			fields: fields{da: func(m *mock_models.MockVmDataAccess) {
				m.EXPECT().Read(id).Return(entity.Vm{ID: id}, errors.New("something wrong"))
			}},
			args:    args{id: id},
			want:    entity.Vm{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			da := mock_models.NewMockVmDataAccess(ctrl)
			tt.fields.da(da)

			q := qemuUseCase{
				da: da,
			}
			got, err := q.ReadById(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_qemuUseCase_Update(t *testing.T) {
	type fields struct {
		da  func(m *mock_models.MockVmDataAccess)
		vnc func(m *mock_models.MockSocket)
	}
	type args struct {
		id uuid.UUID
		vm entity.VmCore
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
					m.EXPECT().Edit(id, entity.VmCore{
						Cpu:    2,
						Memory: 1024,
						Name:   "ubuntu",
					}).Return(testVM, nil)
				},
				vnc: func(m *mock_models.MockSocket) {
					m.EXPECT().SearchFor(id).Return(false)
				},
			},
			args: args{
				id: id,
				vm: entity.VmCore{
					Cpu:    2,
					Memory: 1024,
					Name:   "ubuntu",
				},
			},
			want:    testVM,
			wantErr: false,
		},
		{
			name: "FAIL01",
			fields: fields{
				da: func(m *mock_models.MockVmDataAccess) {},
				vnc: func(m *mock_models.MockSocket) {
					m.EXPECT().SearchFor(id).Return(true)
				},
			},
			args: args{
				id: id,
				vm: entity.VmCore{
					Cpu:    2,
					Memory: 1024,
					Name:   "ubuntu",
				},
			},
			want:    entity.Vm{},
			wantErr: true,
		},
		{
			name: "FAIL02",
			fields: fields{
				da: func(m *mock_models.MockVmDataAccess) {
					m.EXPECT().Edit(id, entity.VmCore{
						Cpu:    2,
						Memory: 1024,
						Name:   "ubuntu",
					}).Return(entity.Vm{}, errors.New("something wrong"))
				},
				vnc: func(m *mock_models.MockSocket) {
					m.EXPECT().SearchFor(id).Return(false)
				},
			},
			args: args{
				id: id,
				vm: entity.VmCore{
					Cpu:    2,
					Memory: 1024,
					Name:   "ubuntu",
				},
			},
			want:    entity.Vm{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			da := mock_models.NewMockVmDataAccess(ctrl)
			vnc := mock_models.NewMockSocket(ctrl)
			tt.fields.da(da)
			tt.fields.vnc(vnc)

			q := qemuUseCase{
				da:  da,
				vnc: vnc,
			}
			got, err := q.Update(tt.args.id, tt.args.vm)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Update() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_qemuUseCase_Delete(t *testing.T) {
	type fields struct {
		da   func(m *mock_models.MockVmDataAccess)
		disk func(m *mock_models.MockDisk)
		vnc  func(m *mock_models.MockSocket)
		qmp  func(m *mock_models.MockSocket)
	}
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "PASS01",
			fields: fields{
				da: func(m *mock_models.MockVmDataAccess) {
					m.EXPECT().Delete(id).Return(nil)
				},
				disk: func(m *mock_models.MockDisk) {
					m.EXPECT().Delete("default.json", id).Return(nil)
				},
				vnc: func(m *mock_models.MockSocket) {
					m.EXPECT().SearchFor(id).Return(false)
					m.EXPECT().Delete(id).Return(nil)
				},
				qmp: func(m *mock_models.MockSocket) {
					m.EXPECT().Delete(id).Return(nil)
				},
			},
			args:    args{id: id},
			wantErr: false,
		},
		{
			name: "FAIL01",
			fields: fields{
				da:   func(m *mock_models.MockVmDataAccess) {},
				disk: func(m *mock_models.MockDisk) {},
				vnc: func(m *mock_models.MockSocket) {
					m.EXPECT().SearchFor(id).Return(true)
				},
				qmp: func(m *mock_models.MockSocket) {},
			},
			args:    args{id: id},
			wantErr: true,
		},
		{
			name: "FAIL02",
			fields: fields{
				da: func(m *mock_models.MockVmDataAccess) {
					m.EXPECT().Delete(id).Return(errors.New("something wrong"))
				},
				disk: func(m *mock_models.MockDisk) {},
				vnc: func(m *mock_models.MockSocket) {
					m.EXPECT().SearchFor(id).Return(false)
				},
				qmp: func(m *mock_models.MockSocket) {},
			},
			args:    args{id: id},
			wantErr: true,
		},
		{
			name: "FAIL03",
			fields: fields{
				da: func(m *mock_models.MockVmDataAccess) {
					m.EXPECT().Delete(id).Return(nil)
				},
				disk: func(m *mock_models.MockDisk) {
					m.EXPECT().Delete("default.json", id).Return(errors.New("something wrong"))
				},
				vnc: func(m *mock_models.MockSocket) {
					m.EXPECT().SearchFor(id).Return(false)
				},
				qmp: func(m *mock_models.MockSocket) {},
			},
			args:    args{id: id},
			wantErr: true,
		},
		{
			name: "FAIL04",
			fields: fields{
				da: func(m *mock_models.MockVmDataAccess) {
					m.EXPECT().Delete(id).Return(nil)
				},
				disk: func(m *mock_models.MockDisk) {
					m.EXPECT().Delete("default.json", id).Return(nil)
				},
				vnc: func(m *mock_models.MockSocket) {
					m.EXPECT().SearchFor(id).Return(false)
					m.EXPECT().Delete(id).Return(errors.New("something wrong"))
				},
				qmp: func(m *mock_models.MockSocket) {},
			},
			args:    args{id: id},
			wantErr: true,
		},
		{
			name: "FAIL05",
			fields: fields{
				da: func(m *mock_models.MockVmDataAccess) {
					m.EXPECT().Delete(id).Return(nil)
				},
				disk: func(m *mock_models.MockDisk) {
					m.EXPECT().Delete("default.json", id).Return(nil)
				},
				vnc: func(m *mock_models.MockSocket) {
					m.EXPECT().SearchFor(id).Return(false)
					m.EXPECT().Delete(id).Return(nil)
				},
				qmp: func(m *mock_models.MockSocket) {
					m.EXPECT().Delete(id).Return(errors.New("something wrong"))
				},
			},
			args:    args{id: id},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			da := mock_models.NewMockVmDataAccess(ctrl)
			disk := mock_models.NewMockDisk(ctrl)
			vnc := mock_models.NewMockSocket(ctrl)
			qmp := mock_models.NewMockSocket(ctrl)
			tt.fields.da(da)
			tt.fields.disk(disk)
			tt.fields.vnc(vnc)
			tt.fields.qmp(qmp)

			q := qemuUseCase{
				da:   da,
				disk: disk,
				vnc:  vnc,
				qmp:  qmp,
			}
			if err := q.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
