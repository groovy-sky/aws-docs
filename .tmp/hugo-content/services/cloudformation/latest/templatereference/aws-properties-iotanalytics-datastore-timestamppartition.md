This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTAnalytics::Datastore TimestampPartition

A partition dimension defined by a timestamp attribute.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AttributeName" : String,
  "TimestampFormat" : String
}

```

### YAML

```yaml

  AttributeName: String
  TimestampFormat: String

```

## Properties

`AttributeName`

The attribute name of the partition defined by a timestamp.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9_]+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimestampFormat`

The timestamp format of a partition defined by a timestamp. The default format is seconds since epoch (January 1, 1970 at midnight UTC time).

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9\s\[\]_,.'/:-]*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::IoTAnalytics::Pipeline

All content copied from https://docs.aws.amazon.com/.
