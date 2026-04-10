This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTAnalytics::Dataset ResourceConfiguration

The configuration of the resource used to execute the `containerAction`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ComputeType" : String,
  "VolumeSizeInGB" : Integer
}

```

### YAML

```yaml

  ComputeType: String
  VolumeSizeInGB: Integer

```

## Properties

`ComputeType`

The type of the compute resource used to execute the `containerAction`.
Possible values are: `ACU_1` (vCPU=4, memory=16 GiB) or `ACU_2` (vCPU=8,
memory=32 GiB).

_Required_: Yes

_Type_: String

_Allowed values_: `ACU_1 | ACU_2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VolumeSizeInGB`

The size, in GB, of the persistent storage available to the resource instance used to
execute the `containerAction` (min: 1, max: 50).

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

QueryAction

RetentionPeriod

All content copied from https://docs.aws.amazon.com/.
