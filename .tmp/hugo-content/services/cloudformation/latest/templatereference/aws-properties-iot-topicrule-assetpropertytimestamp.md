This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::TopicRule AssetPropertyTimestamp

An asset property timestamp entry containing the following information.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OffsetInNanos" : String,
  "TimeInSeconds" : String
}

```

### YAML

```yaml

  OffsetInNanos: String
  TimeInSeconds: String

```

## Properties

`OffsetInNanos`

Optional. A string that contains the nanosecond time offset. Accepts substitution
templates.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeInSeconds`

A string that contains the time in seconds since epoch. Accepts substitution
templates.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Action

AssetPropertyValue

All content copied from https://docs.aws.amazon.com/.
