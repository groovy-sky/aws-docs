This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ARCRegionSwitch::Plan Ec2Ungraceful

Configuration for handling failures when performing operations on EC2 resources.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MinimumSuccessPercentage" : Number
}

```

### YAML

```yaml

  MinimumSuccessPercentage: Number

```

## Properties

`MinimumSuccessPercentage`

The minimum success percentage that you specify for EC2 Auto Scaling groups.

_Required_: Yes

_Type_: Number

_Minimum_: `0`

_Maximum_: `99`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Ec2AsgCapacityIncreaseConfiguration

EcsCapacityIncreaseConfiguration

All content copied from https://docs.aws.amazon.com/.
