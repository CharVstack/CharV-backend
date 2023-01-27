package controller

import (
	"net/http"

	"github.com/CharVstack/CharV-backend/api"
	"github.com/CharVstack/CharV-backend/entity"
	"github.com/CharVstack/CharV-backend/interfaces/errors"
	"github.com/CharVstack/CharV-backend/interfaces/middleware"
	usecase "github.com/CharVstack/CharV-backend/usecase/models"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type V1Handler struct {
	vmUseCase   usecase.VmUseCase
	hostUseCase usecase.HostUseCase
}

func NewV1Handler(vu *usecase.VmUseCase, hu *usecase.HostUseCase) *V1Handler {
	return &V1Handler{
		vmUseCase:   *vu,
		hostUseCase: *hu,
	}
}

func (v V1Handler) GetApiV1Host(c *gin.Context) {
	h, err := v.hostUseCase.Get()

	if err != nil {
		middleware.GenericErrorHandler(c, err, http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, api.GetHost200Response{
		Host: hostTransformer(h),
	})
}

func (v V1Handler) GetApiV1Vms(c *gin.Context) {
	vms, err := v.vmUseCase.ReadAll()

	if err != nil {
		middleware.GenericErrorHandler(c, err, http.StatusInternalServerError)
		return
	}

	ret := make([]api.Vm, len(vms))
	for i, vm := range vms {
		ret[i] = vmTransformer(vm)
	}

	c.JSON(http.StatusOK, api.GetAllVMsList200Response{
		Vms: ret,
	})
}

func (v V1Handler) GetApiV1VmsVmId(c *gin.Context, vmId openapi_types.UUID) {
	vm, err := v.vmUseCase.ReadById(vmId)

	if err != nil {
		switch errors.GetType(err) {
		case errors.NotFound:
			middleware.GenericErrorHandler(c, err, http.StatusNotFound)
		default:
			middleware.GenericErrorHandler(c, err, http.StatusInternalServerError)
		}
		return
	}

	c.JSON(http.StatusOK, api.GetVMByVMId200Response{
		Vm: vmTransformer(vm),
	})
}

// PostApiV1Vms Vm作成時にフロントから情報を受取りステータスを返す
func (v V1Handler) PostApiV1Vms(c *gin.Context) {
	var req api.PostApiV1VmsJSONRequestBody
	if err := c.ShouldBindJSON(&req); err != nil {
		middleware.GenericErrorHandler(c, err, http.StatusBadRequest)
		return
	}

	vm, err := v.vmUseCase.Create(entity.VmCore(req))
	if err != nil {
		middleware.GenericErrorHandler(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, api.PostCreateNewVM200Response{
		Vm: vmTransformer(vm),
	})
}

func (v V1Handler) PatchApiV1VmsVmId(c *gin.Context, vmId openapi_types.UUID) {
	//TODO implement me
	middleware.GenericErrorHandler(c, errors.Unknown.New("implement me"), http.StatusInternalServerError)
}

func (v V1Handler) DeleteApiV1VmsVmId(c *gin.Context, vmId openapi_types.UUID) {
	err := v.vmUseCase.Delete(vmId)

	if err != nil {
		switch errors.GetType(err) {
		case errors.NotFound:
			middleware.GenericErrorHandler(c, err, http.StatusNotFound)
		default:
			middleware.GenericErrorHandler(c, err, http.StatusInternalServerError)
		}
		return
	}

	c.Status(http.StatusOK)
}

func (v V1Handler) GetApiV1VmsVmIdPower(c *gin.Context, vmId uuid.UUID) {
	power := v.vmUseCase.GetPower(vmId)

	c.JSON(http.StatusOK, api.GetVMPowerByVMId200Response{
		VmPower: api.VmPowerInfo{
			CleanPowerOff: true,
			State:         api.VmPowerInfoState(power),
		},
	})
}

func (v V1Handler) PostApiV1VmsVmIdPower(c *gin.Context, vmId openapi_types.UUID) {
	var req api.PostChangeVMsPowerStatusByVMIdRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		middleware.GenericErrorHandler(c, err, http.StatusBadRequest)
		return
	}

	switch req.Action {
	case api.Start:
		err := v.vmUseCase.Start(vmId)
		if err != nil {
			middleware.GenericErrorHandler(c, err, http.StatusInternalServerError)
			return
		}

		c.Status(http.StatusOK)
	case api.Reboot:
		err := v.vmUseCase.Restart(vmId)
		if err != nil {
			middleware.GenericErrorHandler(c, err, http.StatusInternalServerError)
			return
		}

		c.Status(http.StatusOK)
	case api.Shutdown:
		err := v.vmUseCase.Shutdown(vmId)
		if err != nil {
			middleware.GenericErrorHandler(c, err, http.StatusInternalServerError)
			return
		}

		c.Status(http.StatusOK)
	}
}

// hostTransformer transforms type of Host to models.Host
func hostTransformer(h entity.Host) api.Host {
	var pools []api.StoragePool
	for _, p := range h.StoragePools {
		pools = append(pools, api.StoragePool{
			Name:      p.Name,
			Path:      p.Path,
			Status:    api.StoragePoolStatus(p.Status),
			TotalSize: p.TotalSize,
			UsedSize:  p.UsedSize,
		})
	}

	return api.Host{
		Cpu: api.Cpu{
			Counts:  h.Cpu.Counts,
			Percent: h.Cpu.Percent,
		},
		Memory: api.Memory{
			Free:    h.Memory.Free,
			Percent: h.Memory.Percent,
			Total:   h.Memory.Total,
			Used:    h.Memory.Used,
		},
		StoragePools: pools,
	}
}

// vmTransformer transforms type of Vm to models.Vm
func vmTransformer(vm entity.Vm) api.Vm {
	disks := make([]api.Disk, len(vm.Devices.Disk))
	for i, d := range vm.Devices.Disk {
		disks[i] = api.Disk{
			Device: api.DiskDevice(d.Device),
			Name:   d.Name,
			Pool:   d.Pool,
			Type:   api.DiskType(d.Type),
		}
	}

	return api.Vm{
		Cpu: vm.Cpu,
		Devices: api.Devices{
			Disk: disks,
		},
		Memory: vm.Memory,
		Metadata: api.Metadata{
			Id:         vm.ID,
			ApiVersion: "v1",
		},
		Name: vm.Name,
	}
}
