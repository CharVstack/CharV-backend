package qemu

import (
	"errors"
	"path/filepath"
	"strconv"
	"time"

	"github.com/CharVstack/CharV-backend/entity"
	"github.com/CharVstack/CharV-backend/infrastructure/sockets"
	"github.com/CharVstack/CharV-backend/infrastructure/system"
	"github.com/CharVstack/CharV-backend/infrastructure/utils"
	usecase "github.com/CharVstack/CharV-backend/usecase/models"
	"github.com/digitalocean/go-qemu/qmp"
	"github.com/google/uuid"
)

type qemuUseCase struct {
	da    usecase.VmDataAccess
	id    usecase.ID
	cmd   usecase.Command
	disk  usecase.Disk
	vnc   usecase.Socket
	qmp   usecase.Socket
	paths system.Paths
}

func NewQemuUseCase(v *usecase.VmDataAccess, q *usecase.Disk, p system.Paths) usecase.VmUseCase {
	cmd := utils.NewCommand()
	id := utils.NewID()
	vs := sockets.NewVNCSocket(p)
	qs := sockets.NewQMPSocket(p)

	return &qemuUseCase{
		da:    *v,
		id:    id,
		cmd:   cmd,
		disk:  *q,
		vnc:   vs,
		qmp:   qs,
		paths: p,
	}
}

func (q qemuUseCase) Create(req entity.VmCore) (entity.Vm, error) {
	// --- create vm struct ---

	var vm entity.Vm

	// VmCore
	// Name, Cpu and Memory
	vm.VmCore = req

	// ID
	id, err := q.id.GenID() // UUID Version 4
	if err != nil {
		return entity.Vm{}, err
	}
	vm.ID = id

	// ISO
	vm.Devices.OS = entity.Disk{
		Device: entity.DiskDeviceCdrom,
		Name:   "ubuntu-20.04.5-live-server-amd64.iso",
		Pool:   "default",
		Type:   entity.DiskTypeIso,
	}

	// hard disk
	err = q.disk.Create("default.json", id)
	if err != nil {
		return entity.Vm{}, err
	}

	vm.Devices.Disk = append(vm.Devices.Disk, entity.Disk{
		Device: entity.DiskDeviceDisk,
		Name:   id.String(),
		Pool:   "default",
		Type:   entity.DiskTypeQcow2,
	})

	// boot order
	vm.Boot = entity.BootDeviceCdrom

	// virtualization
	vm.Virtualization = entity.VirtualizationTypeKvm

	// daemonize
	vm.Daemonize = true

	// --- run cmd-system-x86_64 command ---

	args := q.genArgs(&vm)
	if err := q.cmd.Run("qemu-system-x86_64", args); err != nil {
		_ = q.disk.Delete("default.json", id)
		return entity.Vm{}, err
	}

	// --- store data ---

	vm.Boot = entity.BootDeviceDisk

	if err := q.da.Add(vm); err != nil {
		return entity.Vm{}, err
	}

	return vm, nil
}

func (q qemuUseCase) ReadAll() ([]entity.Vm, error) {
	vms, err := q.da.Browse()
	if err != nil {
		return nil, err
	}

	return vms, nil
}

func (q qemuUseCase) ReadById(id uuid.UUID) (entity.Vm, error) {
	vm, err := q.da.Read(id)
	if err != nil {
		return entity.Vm{}, err
	}

	return vm, nil
}

func (q qemuUseCase) Update(id uuid.UUID, vm entity.Vm) (entity.Vm, error) {
	vm, err := q.da.Edit(id, vm)
	if err != nil {
		return entity.Vm{}, err
	}

	return vm, nil
}

func (q qemuUseCase) Delete(id uuid.UUID) error {
	// ToDo: 先にファイルがあるか確認してから削除する

	// Delete the guest constructor file
	err := q.da.Delete(id)
	if err != nil {
		return err
	}

	// Delete disk images
	err = q.disk.Delete("default.json", id)
	if err != nil {
		return err
	}

	// Delete the VNC socket
	err = q.vnc.Delete(id)
	if err != nil {
		return err
	}

	// Delete the QMP socket
	err = q.qmp.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (q qemuUseCase) GetPower(id uuid.UUID) usecase.State {
	s := q.vnc.SearchFor(id)
	if s {
		return usecase.RUNNING
	} else {
		return usecase.SHUTDOWN
	}
}

func (q qemuUseCase) Start(id uuid.UUID) error {
	vm, err := q.da.Read(id)
	if err != nil {
		return err
	}

	args := q.genArgs(&vm)
	if err := q.cmd.Run("qemu-system-x86_64", args); err != nil {
		return err
	}

	return nil
}

func (q qemuUseCase) Restart(id uuid.UUID) error {
	err := q.Shutdown(id)
	if err != nil {
		return err
	}

	/*
		シャットダウンが完了しているかの判定処理
		50秒以内に関数が完了しない場合はtimeout判定
	*/
	for cnt := 0; cnt <= 10; cnt++ {
		if q.GetPower(id) == usecase.SHUTDOWN {
			err = q.Start(id)
			if err != nil {
				return err
			}
			return nil
		}
		time.Sleep(time.Second * 5)
	}
	return errors.New("reboot request timed out")
}

func (q qemuUseCase) Shutdown(id uuid.UUID) error {
	addr := filepath.Join(q.paths.QMP, id.String()+".sock")

	sock, err := qmp.NewSocketMonitor("unix", addr, 2*time.Second)
	if err != nil {
		return err
	}

	err = sock.Connect()
	if err != nil {
		return err
	}

	defer func() {
		err = sock.Disconnect()
	}()

	cmd := []byte(`{ "execute": "system_powerdown" }`)
	_, err = sock.Run(cmd)
	if err != nil {
		return err
	}

	cmd = []byte(`{ "execute": "quit" }`)
	_, err = sock.Run(cmd)
	if err != nil {
		return err
	}

	err = q.vnc.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

// genArgs returns qemu command's args
// Detail: https://www.qemu.org/docs/master/system/index.html
func (q qemuUseCase) genArgs(vm *entity.Vm) []string {
	var args []string

	// `-name`
	// set the name of the guest
	name := []string{"-name", vm.Name}
	args = append(args, name...)

	// `-m`
	// configure guest RAM
	mem := []string{"-m", strconv.FormatUint(vm.Memory, 10)}
	args = append(args, mem...)

	// `-smp`
	// set the number of initial CPUs
	cpu := []string{"-smp", strconv.Itoa(vm.Cpu)}
	args = append(args, cpu...)

	// `-drive`
	// use 'file' as a drive image
	// disk
	d := []string{"-drive", "file=" + filepath.Join("/var/lib/charv/images/"+vm.ID.String()+".qcow2") + ",format=qcow2"}
	args = append(args, d...)

	// `-cdrom`
	// use 'file' as IDE cdrom image
	cdrom := []string{"-cdrom", filepath.Join("/var/lib/charv/images/" + vm.Devices.OS.Name)}
	args = append(args, cdrom...)

	// `-boot`
	// configure guest boot settings
	// 'drives': floppy (a), hard disk (c), CD-ROM (d), network (n)
	switch vm.Boot {
	case entity.BootDeviceCdrom:
		boot := []string{"-boot", "order=d"}
		args = append(args, boot...)
	case entity.BootDeviceDisk:
		boot := []string{"-boot", "order=c"}
		args = append(args, boot...)
	default:
		boot := []string{"-boot", "order=c"}
		args = append(args, boot...)
	}

	// `-qmp`
	// like -monitor but opens in 'control' mode
	p := []string{"-qmp", "unix:/" + filepath.Join(q.paths.QMP, vm.ID.String()+".sock") + ",server,nowait"}
	args = append(args, p...)

	// `-vnc`
	// select display VNC type
	vnc := []string{"-vnc", "unix:/" + filepath.Join(q.paths.VNC, vm.ID.String()+".sock")}
	args = append(args, vnc...)

	// select accelerator
	accel := []string{"-accel", string(vm.Virtualization)}
	args = append(args, accel...)

	// daemonize QEMU after initializing
	if vm.Daemonize {
		daemonize := "-daemonize"
		args = append(args, daemonize)
	}

	return args
}
