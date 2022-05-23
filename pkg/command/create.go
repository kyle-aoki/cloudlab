package command

import (
	"cloud/pkg/amazon"
	"cloud/pkg/defaults"
	"cloud/pkg/names"
	"cloud/pkg/util"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func CreatePublicInstance() {
	cldo := &defaults.CloudLabDefaultsOperator{}
	cldo.FindAll()
	CreateInstanceInSubnet(cldo, cldo.PublicSubnet.SubnetId)
}

func CreatePrivateInstance() {
	cldo := &defaults.CloudLabDefaultsOperator{}
	cldo.FindAll()
	CreateInstanceInSubnet(cldo, cldo.PrivateSubnet.SubnetId)
}

func CreateInstanceInSubnet(cldo *defaults.CloudLabDefaultsOperator, subnetId *string) {
	rio, err := amazon.EC2().RunInstances(&ec2.RunInstancesInput{
		SubnetId: subnetId,
		BlockDeviceMappings: []*ec2.BlockDeviceMapping{{
			DeviceName: util.StrPtr("/dev/sda1"),
			Ebs:        &ec2.EbsBlockDevice{VolumeSize: util.IntToInt64Ptr(8)}},
		},
		ImageId:  util.StrPtr(amazon.GetAmi()),
		MinCount: util.IntToInt64Ptr(1),
		MaxCount: util.IntToInt64Ptr(1),
		KeyName:  util.StrPtr(cldo.GetCurrentCloudLabKeyPairName()),
		SecurityGroupIds: []*string{
			util.StrPtr(*cldo.GetSecurityGroupIdByNameOrPanic("22").GroupId),
		},
		InstanceType: util.StrPtr("t2.micro"),
		TagSpecifications: []*ec2.TagSpecification{
			{ResourceType: util.StrPtr("instance"), Tags: []*ec2.Tag{
				{Key: util.StrPtr("Name"), Value: util.StrPtr(names.GetRandomName(cldo.Instances))},
				{Key: util.StrPtr(defaults.CloudLabInstance), Value: util.StrPtr("true")},
			}},
		},
	})
	util.MustExec(err)
	util.SinglePrint("created node\t" + *rio.Instances[0].InstanceId)
}
