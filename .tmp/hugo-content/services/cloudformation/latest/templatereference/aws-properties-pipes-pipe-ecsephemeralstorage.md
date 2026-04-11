This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pipes::Pipe EcsEphemeralStorage

The amount of ephemeral storage to allocate for the task. This parameter is used to
expand the total amount of ephemeral storage available, beyond the default amount, for
tasks hosted on Fargate. For more information, see [Fargate task storage](../../../amazonecs/latest/userguide/using-data-volumes.md) in the _Amazon ECS User Guide_
_for Fargate_.

###### Note

This parameter is only supported for tasks hosted on Fargate using
Linux platform version `1.4.0` or later. This parameter is not supported for
Windows containers on Fargate.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SizeInGiB" : Integer
}

```

### YAML

```yaml

  SizeInGiB: Integer

```

## Properties

`SizeInGiB`

The total amount, in GiB, of ephemeral storage to set for the task. The minimum
supported value is `21` GiB and the maximum supported value is `200`
GiB.

_Required_: Yes

_Type_: Integer

_Minimum_: `21`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EcsEnvironmentVariable

EcsInferenceAcceleratorOverride

All content copied from https://docs.aws.amazon.com/.
