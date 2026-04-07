This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DeviceFarm::Project VpcConfig

The VPC security groups and subnets that are attached to a project.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SecurityGroupIds" : [ String, ... ],
  "SubnetIds" : [ String, ... ],
  "VpcId" : String
}

```

### YAML

```yaml

  SecurityGroupIds:
    - String
  SubnetIds:
    - String
  VpcId: String

```

## Properties

`SecurityGroupIds`

A list of VPC security group IDs.

A security group allows inbound traffic from network interfaces (and their associated instances) that are
assigned to the same security group. See [Security groups](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_SecurityGroups.html) in the _Amazon_
_Virtual Private Cloud user guide_.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `4096 | 5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubnetIds`

A subnet is a range of IP addresses in your VPC. You can launch Amazon resources, such as EC2 instances,
into a specific subnet. When you create a subnet, you specify the IPv4 CIDR block for the subnet, which is a
subset of the VPC CIDR block. See [VPCs and subnets](https://docs.aws.amazon.com/vpc/latest/userguide/VPC_Subnets.html) in the _Amazon Virtual Private Cloud user guide_.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `4096 | 8`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcId`

A list of VPC IDs.

Each VPC is given a unique ID upon creation.

_Required_: Yes

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::DeviceFarm::TestGridProject
