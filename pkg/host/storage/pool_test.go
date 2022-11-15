package storage

import (
	"reflect"
	"testing"

	"github.com/CharVstack/CharV-backend/domain/models"
)

func TestGetPoolFiles(t *testing.T) {
	files := []string{"default.json"}
	poolfile, _ := GetPoolFiles("../../../test/resources/storage")
	if !reflect.DeepEqual(poolfile, files) {
		t.Error(`miss`)
	}
}

func TestGetPoolInfo(t *testing.T) {
	storagePool := models.StoragePool{
		Name: "default",
		Path: "/var/lib/libvirt/default",
	}
	storagePoolInfo, _ := GetPoolInfo("default.json", "../../../test/resources/storage")
	if !reflect.DeepEqual(storagePoolInfo, storagePool) {
		t.Error(`miss`)
	}
}
