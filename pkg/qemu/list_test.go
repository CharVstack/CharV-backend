package qemu

import "testing"

func TestConvertToStruct(t *testing.T) {
	vms, err := ConvertToStruct("../../test/resources/machines/")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%#v", vms)
}
