package cmd

// import (
// 	"cloudlab/pkg/amazon"
// 	"cloudlab/pkg/args"
// 	"cloudlab/pkg/resource"
// 	"cloudlab/pkg/util"
// 	"fmt"
// 	"strconv"

// 	"github.com/aws/aws-sdk-go/service/ec2"
// )

// func ValidatePort(portString string) (portInt int) {
// 	portInt64, err := strconv.ParseInt(portString, 10, 32)
// 	if err != nil {
// 		panic("invalid port")
// 	}
// 	portInt = int(portInt64)
// 	if portInt > 65535 || portInt < 1 {
// 		panic("invalid port")
// 	}
// 	return portInt
// }

func OpenPort() {
	// port := args.Poll()
	// _ = args.Poll()

	// co := resource.NewResourceOperator()

	// sg := co.GetSecurityGroupIdByNameOrNil(port)

	// if sg == nil {
	// 	portInt := ValidatePort(port)
	// 	resource.CreateSecurityGroup(co.Vpc, port, portInt)
	// }

	// // instance := co.GetInstanceByName(instnaceName)
	// // securityGroup := co.GetSecurityGroupIdByNameOrPanic(port)

	// // co.AssignSecurityGroup(instance, securityGroup)
	// // fmt.Println(fmt.Sprintf("opened port %s on node %s", port, instnaceName))
}

func ClosePort() {
	// co := resource.NewResourceOperator()

	// port := args.Poll()
	// instanceName := args.Poll()

	// _ = ValidatePort(port)

	// var instance *ec2.Instance = co.FindInstanceByName(instanceName)

	// var newSecurityGroups []*string
	// for _, groupIdentifier := range instance.SecurityGroups {
	// 	if groupIdentifier.GroupName != nil && *groupIdentifier.GroupName == port {
	// 		continue
	// 	}
	// 	newSecurityGroups = append(newSecurityGroups, groupIdentifier.GroupId)
	// }

	// _, err := amazon.EC2().ModifyInstanceAttribute(&ec2.ModifyInstanceAttributeInput{
	// 	InstanceId: instance.InstanceId,
	// 	Groups:     newSecurityGroups,
	// })
	// util.MustExec(err)

	// fmt.Println(fmt.Sprintf("closed port %s on node %s", port, instanceName))
}
