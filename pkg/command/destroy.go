package cmd

import (
	"cloudlab/pkg/resource"
	"fmt"
)

func DestroyCloudLabResources() {
	ro := resource.New()
	ro.FindAll()
	ro.DestroyCloudLabResources()
	fmt.Println("deleted all cloudlab resources")
}
