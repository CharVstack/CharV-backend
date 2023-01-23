package disk

import (
	"github.com/CharVstack/CharV-backend/infrastructure/system"
	"github.com/CharVstack/CharV-backend/usecase/models"
	"github.com/google/uuid"
	"os"
	"os/exec"
	"path/filepath"
)

type qcow2 struct {
	storageAccess models.StorageAccess
	sys           system.Paths
}

func NewQCOW2Disk(q *models.StorageAccess, c system.Paths) models.Disk {
	return &qcow2{
		storageAccess: *q,
		sys:           c,
	}
}

func (q qcow2) Create(id uuid.UUID) error {
	// ToDo: `/storage_pools` のファイルから読み取る
	path := filepath.Join("/var/lib/charv/images/", id.String()+".qcow2")

	args := []string{"create", "-f", "qcow2", path, "16G"}

	cmd := exec.Command("qemu-img", args...)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func (q qcow2) Delete(id uuid.UUID) error {
	// ToDo: `/storage_pools` のファイルから読み取る
	err := os.Remove(filepath.Join("/var/lib/charv/images/", id.String()+".qcow2"))

	return err
}
