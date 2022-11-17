package qemu

import (
	"reflect"
	"testing"

	"github.com/CharVstack/CharV-backend/domain/models"
	"github.com/google/uuid"
)

func TestParse(t *testing.T) {
	testVm := models.Vm{
		Name:   "ubuntu",
		Memory: 1024,
		Cpu:    1,
		Devices: models.Devices{
			Disk: []models.Disk{
				{
					Name:   "ubuntu",
					Pool:   "default",
					Type:   "qcow2",
					Device: "disk",
				},
			},
		},
		Metadata: models.Metadata{
			Id:         uuid.Must(uuid.Parse("0bfb8def-86ed-4b9d-8826-66a6ab1c1491")),
			ApiVersion: "v1",
		},
	}
	machine, err := parse("../../test/resources/machines/ubuntu.json")
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(testVm, machine) {
		t.Fatalf("expected: %#v\nactually: %#v", testVm, machine)
	}
}

func TestGetAllVms(t *testing.T) {
	_, err := getAllVms("../../test/resources/machines/")
	if err != nil {
		t.Fatal(err)
	}
}
