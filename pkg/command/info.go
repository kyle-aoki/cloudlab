package command

import "cloud/pkg/resource"

func Info() {
	ro := resource.New()
	ro.FindAll()
	ro.Info()
}
