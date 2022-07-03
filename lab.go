package main

import (
	"cloudlab/pkg/amazon"
	"cloudlab/pkg/args"
	"cloudlab/pkg/cmd"
	"cloudlab/pkg/util"
	"fmt"
)

func main() {
	defer util.MainRecover()
	args.Init()
	amazon.InitEC2Client()

	if command, found := Syntax[args.PollOrEmpty()]; found {
		command()
	} else {
		fmt.Print(HelpText)
	}
}

var Syntax = map[string]func(){
	"version":    cmd.Version,
	"info":       cmd.Info,
	"init":       cmd.InitializeCloudLabResources,
	"destroy":    cmd.DestroyCloudLabResources,
	"list":       cmd.ListInstances,
	"run":        cmd.CreateInstance,
	"delete":     cmd.DeleteInstances,
	"open-port":  cmd.OpenPort,
	"close-port": cmd.ClosePort,
}

const HelpText = `command structure:
  lab [...flags] <command> [...arguments]
general flags:
  -v, -verbose    verbose logging
commands:
  version          print version
  info             print cloudlab resource info

  init             create base cloudlab resources (vpc, subnets, etc.)
                   (base resources cost nothing)
  destroy          destroy base cloudlab resrouces
                   (must terminate all instances first)

  list             list active instances
  run              run a new instance
                       -name             instance name
                       -private, -p      create instance in private subnet
                       -type,            instance type (t2.nano, t2.micro, etc.)
                       -gigs             gigabytes of storage
                       -script           path to start up script file

  delete           <instance-name>           terminate an instance
  open-port        <port> <instance-name>    open a port on an instance (all protocols)
  close-port       <port> <instance-name>    close a port on an instance
`
