package file

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/CharVstack/CharV-backend/entity"
	"github.com/CharVstack/CharV-backend/infrastructure/system"
	"github.com/CharVstack/CharV-backend/usecase/models"
	"github.com/google/uuid"
)

type vmDataAccess struct {
	path system.Paths
}

func NewVmDataAccess(c system.Paths) models.VmDataAccess {
	return &vmDataAccess{
		path: c,
	}
}

func (v vmDataAccess) Browse() ([]entity.Vm, error) {
	files, err := os.ReadDir(v.path.Guests)
	if err != nil {
		return []entity.Vm{}, err
	}

	vms := make([]entity.Vm, len(files))

	for i, file := range files {
		name := file.Name()
		if len(name) != 41 {
			continue
		}

		vm, err := v.Read(uuid.Must(uuid.Parse(name[:36])))
		if err != nil {
			return vms, err
		}

		vms[i] = vm
	}

	return vms, nil
}

func (v vmDataAccess) Read(id uuid.UUID) (entity.Vm, error) {
	var vm entity.Vm

	path := v.path.Guests + id.String() + ".json"
	data, err := os.ReadFile(path)
	if err != nil {
		return vm, err
	}

	err = json.Unmarshal(data, &vm)
	if err != nil {
		return vm, err
	}

	return vm, err
}

func (v vmDataAccess) Edit(id uuid.UUID, vm entity.Vm) (entity.Vm, error) {
	return entity.Vm{}, nil
}

func (v vmDataAccess) Add(vm entity.Vm) error {
	n := v.path.Guests + vm.ID.String() + ".json"

	f, err := os.Create(n)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	d, err := json.Marshal(vm)
	if err != nil {
		return err
	}

	_, err = f.Write(d)
	if err != nil {
		return err
	}

	return nil
}

func (v vmDataAccess) Delete(id uuid.UUID) error {
	return os.Remove(filepath.Join(v.path.Guests, id.String()+".json"))
}
