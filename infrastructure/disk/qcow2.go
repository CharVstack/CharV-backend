package disk

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/CharVstack/CharV-backend/infrastructure/system"
	"github.com/CharVstack/CharV-backend/usecase/models"
	"github.com/google/uuid"
)

type qcow2 struct {
	da  models.StorageAccess
	sys system.Paths
}

func NewQCOW2Disk(q *models.StorageAccess, c system.Paths) models.Disk {
	return qcow2{
		da:  *q,
		sys: c,
	}
}

func (q qcow2) Create(pool string, id uuid.UUID) error {
	poolPath, err := q.da.Read(pool)
	if err != nil {
		return err
	}

	path := filepath.Join(poolPath.Path, id.String()+".qcow2")

	args := []string{"create", "-f", "qcow2", path, "16G"}

	cmd := exec.Command("qemu-img", args...)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func (q qcow2) Delete(pool string, id uuid.UUID) error {
	poolPath, err := q.da.Read(pool)
	if err != nil {
		return err
	}

	return os.Remove(filepath.Join(poolPath.Path, id.String()+".qcow2"))
}
