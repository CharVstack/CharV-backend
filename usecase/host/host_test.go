package host

import (
	"github.com/CharVstack/CharV-backend/entity"
	"github.com/CharVstack/CharV-backend/usecase/models"
	"reflect"
	"testing"
)

type testHostStatAccess struct{}

func (t testHostStatAccess) GetCpu() (entity.Cpu, error) {
	return entity.Cpu{
		Counts:  2,
		Percent: 33.4,
	}, nil
}

func (t testHostStatAccess) GetMem() (entity.Memory, error) {
	return entity.Memory{
		Free:    16,
		Percent: 50,
		Total:   32,
		Used:    16,
	}, nil
}

func (t testHostStatAccess) GetStoragePools() ([]entity.StoragePool, error) {
	return []entity.StoragePool{
		{
			Name:      "default",
			Path:      "/var/lig/libvirt/images",
			Status:    entity.ACTIVE,
			TotalSize: 1006530654208,
			UsedSize:  371360915456,
		},
	}, nil
}

func Test_hostUseCase_Get(t *testing.T) {
	type fields struct {
		hostStatAccess models.HostStatAccess
	}
	tests := []struct {
		name    string
		fields  fields
		want    entity.Host
		wantErr bool
	}{
		{
			name: "PASS_01",
			fields: fields{
				hostStatAccess: testHostStatAccess{},
			},
			want: entity.Host{
				Cpu: entity.Cpu{
					Counts:  2,
					Percent: 33.4,
				},
				Memory: entity.Memory{
					Free:    16,
					Percent: 50,
					Total:   32,
					Used:    16,
				},
				StoragePools: []entity.StoragePool{
					{
						Name:      "default",
						Path:      "/var/lig/libvirt/images",
						Status:    entity.ACTIVE,
						TotalSize: 1006530654208,
						UsedSize:  371360915456,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := hostUseCase{
				hostStatAccess: tt.fields.hostStatAccess,
			}
			got, err := h.Get()
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}
