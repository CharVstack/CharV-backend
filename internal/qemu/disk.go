package qemu

import (
	"bytes"
	"text/template"
)

func createDisk(name string) (string, error) {
	name = "/var/lib/charVstack/images/" + name + "." + "qcow2"
	tmpl, err := template.New("create").Parse(`qemu-img create -f qcow2 {{.}} 16G`)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, name)
	if err != nil {
		return "", err
	}
	cmd := buf.String()

	return name, run(cmd)
}
