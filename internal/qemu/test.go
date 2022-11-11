package qemu

import (
	"reflect"
	"testing"

	"github.com/CharVstack/CharV-backend/domain/models"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	testVm := models.Vm{
		Name:   "ubuntu",
		Memory: 1024,
		Vcpu:   1,
		Devices: models.Devices{
			Disk: []models.Disk{
				{
					Type: "file",
					Path: "/path/to/ubuntu.qcow2",
				},
			},
		},
	}
	machine, err := parse("../../test/resources/machines/ubuntu.json")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(testVm, machine) {
		t.Fail()
	}
}

func TestGetAllVms(t *testing.T) {
	vms, err := getAllVms("../../test/resources/machines/")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", vms)
}

func TestCheckFileType(t *testing.T) {
	var err error
	_, err = CheckFileType("../../test/resources/image/ok.qcow2")
	if !assert.NoError(t, err) {
		t.Fatal(err)
	}

	_, err = CheckFileType("../../test/resources/image/bad.qcow2")
	if !assert.Error(t, err) {
		t.Fail()
	}
}
