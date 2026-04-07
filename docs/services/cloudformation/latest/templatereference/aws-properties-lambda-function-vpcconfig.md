This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::Function VpcConfig

The VPC security groups and subnets that are attached to a Lambda function. When you connect a function to a
VPC, Lambda creates an elastic network interface for each combination of security group and subnet in the
function's VPC configuration. The function can only access resources and the internet through that VPC. For more
information, see [VPC\
Settings](https://docs.aws.amazon.com/lambda/latest/dg/configuration-vpc.html).

###### Note

When you delete a function, CloudFormation monitors the state of its network interfaces and waits for
Lambda to delete them before proceeding. If the VPC is defined in the same stack, the network interfaces need to
be deleted by Lambda before CloudFormation can delete the VPC's resources.

To monitor network interfaces, CloudFormation needs the `ec2:DescribeNetworkInterfaces`
permission. It obtains this from the user or role that modifies the stack. If you don't provide this permission,
CloudFormation does not wait for network interfaces to be deleted.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Ipv6AllowedForDualStack" : Boolean,
  "SecurityGroupIds" : [ String, ... ],
  "SubnetIds" : [ String, ... ]
}

```

### YAML

```yaml

  Ipv6AllowedForDualStack: Boolean
  SecurityGroupIds:
    - String
  SubnetIds:
    - String

```

## Properties

`Ipv6AllowedForDualStack`

Allows outbound IPv6 traffic on VPC functions that are connected to dual-stack subnets.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityGroupIds`

A list of VPC security group IDs.

_Required_: No

_Type_: Array of String

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubnetIds`

A list of VPC subnet IDs.

_Required_: No

_Type_: Array of String

_Maximum_: `16`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### VPC Configuration

Connect a function to a VPC.

#### YAML

```yaml

      VpcConfig:
        SecurityGroupIds:
          - sg-085912345678492fb
        SubnetIds:
          - subnet-071f712345678e7c8
          - subnet-07fd123456788a036
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TracingConfig

AWS::Lambda::LayerVersion
