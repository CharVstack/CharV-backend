package qemu

import (
	"os/exec"

	"github.com/mattn/go-shellwords"
)

func run(cmd string) error {
	c, err := shellwords.Parse(cmd)
	if err != nil {
		return err
	}
	switch len(c) {
	case 0:
		return nil
	case 1:
		err = exec.Command(c[0]).Run()
	default:
		err = exec.Command(c[0], c[1:]...).Run()
	}
	if err != nil {
		return err
	}
	return nil
}
