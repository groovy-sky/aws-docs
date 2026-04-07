This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::Service VpcLatticeConfiguration

The VPC Lattice configuration for your service that holds the information for the
target group(s) Amazon ECS tasks will be registered to.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PortName" : String,
  "RoleArn" : String,
  "TargetGroupArn" : String
}

```

### YAML

```yaml

  PortName: String
  RoleArn: String
  TargetGroupArn: String

```

## Properties

`PortName`

The name of the port mapping to register in the VPC Lattice target group. This is the
name of the `portMapping` you defined in your task definition.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The ARN of the IAM role to associate with this VPC Lattice configuration. This is the
Amazon ECS infrastructure IAM role that is used to manage your VPC Lattice
infrastructure.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetGroupArn`

The full Amazon Resource Name (ARN) of the target group or groups associated with the
VPC Lattice configuration that the Amazon ECS tasks will be registered to.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TimeoutConfiguration

AWS::ECS::TaskDefinition
