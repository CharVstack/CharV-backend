package host

import (
	"reflect"
	"testing"

	"github.com/CharVstack/CharV-backend/entity"
	mock_models "github.com/CharVstack/CharV-backend/usecase/models/mocks"
	"github.com/golang/mock/gomock"
)

func Test_hostUseCase_Get(t *testing.T) {
	tests := []struct {
		name    string
		mock    func(m *mock_models.MockHostStatAccess)
		want    entity.Host
		wantErr bool
	}{
		{
			name: "PASS01",
			mock: func(m *mock_models.MockHostStatAccess) {
				m.EXPECT().GetCpu().Return(entity.Cpu{
					Counts:  2,
					Percent: 33.4,
				}, nil)
				m.EXPECT().GetMem().Return(entity.Memory{
					Free:    16,
					Percent: 50,
					Total:   32,
					Used:    16,
				}, nil)
				m.EXPECT().GetStoragePools().Return([]entity.StoragePool{
					{
						Name:      "default",
						Path:      "/var/lig/libvirt/images",
						Status:    "ACTIVE",
						TotalSize: 1006530654208,
						UsedSize:  371360915456,
					},
				}, nil)
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
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			m := mock_models.NewMockHostStatAccess(ctrl)
			tt.mock(m)

			h := hostUseCase{
				hostStatAccess: m,
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
