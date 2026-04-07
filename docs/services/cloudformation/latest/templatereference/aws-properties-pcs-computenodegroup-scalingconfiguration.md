This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PCS::ComputeNodeGroup ScalingConfiguration

Specifies the boundaries of the compute node group auto scaling.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxInstanceCount" : Integer,
  "MinInstanceCount" : Integer
}

```

### YAML

```yaml

  MaxInstanceCount: Integer
  MinInstanceCount: Integer

```

## Properties

`MaxInstanceCount`

The upper bound of the number of instances allowed in the compute fleet.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinInstanceCount`

The lower bound of the number of instances allowed in the compute fleet.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

InstanceConfig

SlurmConfiguration
