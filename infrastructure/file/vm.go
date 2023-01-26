package file

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/CharVstack/CharV-backend/entity"
	"github.com/CharVstack/CharV-backend/infrastructure/system"
	"github.com/CharVstack/CharV-backend/interfaces/errors"
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

	path := filepath.Join(v.path.Guests, id.String()+".json")
	if _, err := os.Stat(path); err != nil {
		return entity.Vm{}, errors.NotFound.Wrap(err, fmt.Sprintf("%s is not found", path))
	}
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

func (v vmDataAccess) Edit(id uuid.UUID, vm entity.VmCore) (entity.Vm, error) {
	target, err := v.Read(id)
	if err != nil {
		return entity.Vm{}, err
	}
	path := filepath.Join(v.path.Guests, id.String()+".json")
	f, err := os.Open(path)
	if err != nil {
		return entity.Vm{}, err
	}
	defer func() {
		_ = f.Close()
	}()
	target.Cpu = vm.Cpu
	target.Memory = vm.Memory
	target.Name = vm.Name
	d, err := json.Marshal(target)
	if err != nil {
		return entity.Vm{}, err
	}
	if _, err := f.Write(d); err != nil {
		return entity.Vm{}, err
	}
	return target, nil
}

func (v vmDataAccess) Add(vm entity.Vm) error {
	path := filepath.Join(v.path.Guests, vm.ID.String()+".json")

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		f.Close()
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
	path := filepath.Join(v.path.Guests, id.String()+".json")
	if _, err := os.Stat(path); err != nil {
		return errors.NotFound.Wrap(err, fmt.Sprintf("%s is not found", path))
	}
	return os.Remove(path)
}
