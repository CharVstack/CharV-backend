package utils

import (
	"os/exec"

	"github.com/CharVstack/CharV-backend/usecase/models"
)

type cmd struct{}

func NewCommand() models.Command {
	return &cmd{}
}

func (c cmd) Run(name string, args []string) error {
	cmd := exec.Command(name, args...)
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
