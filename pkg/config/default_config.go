package config

const DEFAULT_CONFIG = 
`{
	"ShowTerminatedNodes": true,
	"NodeConfigurations": [
	  {
		"ConfigName": "public-node",
		"SubnetNameTag": "PUBLIC",
		"SecurityGroupNames": ["SSH", "UDP_500", "UDP_4500"],
		"AMI": "ami-091130e4e91d7bb45",
		"KeyPair": "kp1",
		"InstanceType": "t2.nano",
		"StorageSize": "8gb"
	  },
	  {
		"ConfigName": "private-node",
		"SubnetNameTag": "private",
		"SecurityGroupNames": ["ALL_TRAFFIC"],
		"AMI": "ami-091130e4e91d7bb45",
		"KeyPair": "kp1",
		"InstanceType": "t2.nano",
		"StorageSize": "8gb"
	  }
	]
}
`