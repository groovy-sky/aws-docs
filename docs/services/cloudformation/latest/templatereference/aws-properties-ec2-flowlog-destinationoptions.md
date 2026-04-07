This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::FlowLog DestinationOptions

Describes the destination options for a flow log.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FileFormat" : String,
  "HiveCompatiblePartitions" : Boolean,
  "PerHourPartition" : Boolean
}

```

### YAML

```yaml

  FileFormat: String
  HiveCompatiblePartitions: Boolean
  PerHourPartition: Boolean

```

## Properties

`FileFormat`

The format for the flow log. The default is `plain-text`.

_Required_: Yes

_Type_: String

_Allowed values_: `plain-text | parquet`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`HiveCompatiblePartitions`

Indicates whether to use Hive-compatible prefixes for flow logs stored in Amazon S3.
The default is `false`.

_Required_: Yes

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PerHourPartition`

Indicates whether to partition the flow log per hour. This reduces the cost and response
time for queries. The default is `false`.

_Required_: Yes

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::EC2::FlowLog

Tag
