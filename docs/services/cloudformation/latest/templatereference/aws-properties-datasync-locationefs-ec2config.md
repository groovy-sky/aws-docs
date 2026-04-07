This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataSync::LocationEFS Ec2Config

The subnet and security groups that AWS DataSync uses to connect to one of your
Amazon EFS file system's [mount targets](https://docs.aws.amazon.com/efs/latest/ug/accessing-fs.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SecurityGroupArns" : [ String, ... ],
  "SubnetArn" : String
}

```

### YAML

```yaml

  SecurityGroupArns:
    - String
  SubnetArn: String

```

## Properties

`SecurityGroupArns`

Specifies the Amazon Resource Names (ARNs) of the security groups associated with an
Amazon EFS file system's mount target.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `128 | 5`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SubnetArn`

Specifies the ARN of a subnet where DataSync creates the [network interfaces](https://docs.aws.amazon.com/datasync/latest/userguide/datasync-network.html#required-network-interfaces.html) for managing traffic during your transfer.

The subnet must be located:

- In the same virtual private cloud (VPC) as the Amazon EFS file system.

- In the same Availability Zone as at least one mount target for the Amazon EFS
file system.

###### Note

You don't need to specify a subnet that includes a file system mount
target.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws|aws-cn|aws-us-gov|aws-eusc|aws-iso|aws-iso-b):ec2:[a-z\-0-9]*:[0-9]{12}:subnet/.*$`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::DataSync::LocationEFS

Tag
