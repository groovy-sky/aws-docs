This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTAnalytics::Datastore Partition

A single dimension to partition a data store. The dimension must be an `AttributePartition` or a `TimestampPartition`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AttributeName" : String
}

```

### YAML

```yaml

  AttributeName: String

```

## Properties

`AttributeName`

The name of the attribute that defines a partition dimension.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9_]+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ParquetConfiguration

RetentionPeriod

All content copied from https://docs.aws.amazon.com/.
