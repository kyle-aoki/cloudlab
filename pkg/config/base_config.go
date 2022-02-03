package config

const configFileCreationMessage = "Config file has been created at ~/.cloud/config.json.\nPlease open this file and update it with your security groups and subnets."

const baseConfig = `{
  "Variables": {
    "Subnets": [
      { "Name": "PRIVATE_SUBNET", "Value": "subnet-xxxxxxxxxxxxxxxxx" },
      { "Name": "PUBLIC_SUBNET", "Value": "subnet-xxxxxxxxxxxxxxxxx" }
    ],
    "SecurityGroups": [
      { "Name": "SSH_ONLY", "Value": "sg-xxxxxxxxxxxxxxxxx" },
      { "Name": "UDP_500", "Value": "sg-xxxxxxxxxxxxxxxxx" },
      { "Name": "UDP_4500", "Value": "sg-xxxxxxxxxxxxxxxxx" },
      { "Name": "ALL_TRAFFIC", "Value": "sg-xxxxxxxxxxxxxxxxx" }
    ],
    "AMIs": [
      { "Name": "Ubuntu_21_10_US_West_1", "Value": "ami-091130e4e91d7bb45" }
    ]
  },
  "InstanceConfigurations": [
    {
      "Name": "default",
      "Subnet": "PRIVATE_SUBNET",
      "SecurityGroups": ["ALL_TRAFFIC"],
      "KeyPair": "kp1",
      "InstanceType": "t2.nano",
      "StorageSize": "8gb",
      "AMI": "Ubuntu_21_10_US_West_1"
    },
    {
      "Name": "public",
      "Subnet": "PUBLIC",
      "SecurityGroups": ["SSH_ONLY", "UDP_500", "UDP_4500"],
      "KeyPair": "kp1",
      "InstanceType": "t2.nano",
      "StorageSize": "8gb",
      "AMI": "Ubuntu_21_10_US_West_1"
    },
    {
      "Name": "<name>",
      "Subnet": "PRIVATE_SUBNET",
      "SecurityGroups": ["SSH_ONLY"],
      "KeyPair": "kp1",
      "InstanceType": "t2.nano",
      "StorageSize": "8gb",
      "AMI": "Ubuntu_21_10_US_West_1"
    }
  ]
}
`
