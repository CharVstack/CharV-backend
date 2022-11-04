package qemu

import (
	"reflect"
	"testing"

	"github.com/CharVstack/CharV-backend/domain/models"
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

func TestConvertToStruct(t *testing.T) {
	vms, err := ConvertToStruct("../../test/resources/machines/")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", vms)
}
