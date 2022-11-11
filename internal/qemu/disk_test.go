package qemu

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
