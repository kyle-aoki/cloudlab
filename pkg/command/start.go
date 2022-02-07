package command

import (
	"cloud/pkg/amazon"
	"cloud/pkg/args"
	"cloud/pkg/util"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func Start() {
	names := args.Collect()
	client := amazon.EC2Client()

	nodeNames := GetNodeNames(GetNodes())
	var nodesToStart []NodeName

	for i := range names {
		for j := range nodeNames {
			if strings.Contains(nodeNames[j].Name, names[i]) {
				nodesToStart = append(nodesToStart, nodeNames[j])
			}
		}
	}

	ids := GetIds(nodesToStart)

	sio, err := client.StartInstances(&ec2.StartInstancesInput{
		InstanceIds: util.StrSlicePtr(ids),
	})
	util.MustExec(err)
	
	for i := range sio.StartingInstances {
		fmt.Println(*sio.StartingInstances[i].InstanceId)
	}
}

type NodeName struct {
	Name string
	Id   string
}

func GetIds(nns []NodeName) (ids []string) {
	for i := range nns {
		ids = append(ids, nns[i].Id)
	}
	return ids
}

func GetNodeNames(nodes []*ec2.Instance) (nodeNames []NodeName) {
	for i := 0; i < len(nodes); i++ {
		var nameTag string
		for j := range nodes[i].Tags {
			if *nodes[i].Tags[j].Key == "Name" {
				nameTag = *nodes[i].Tags[j].Value
			}
		}
		nodeNames = append(nodeNames, NodeName{
			Id:   *nodes[i].InstanceId,
			Name: nameTag,
		})
	}
	return nodeNames
}