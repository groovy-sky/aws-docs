This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTAnalytics::Dataset DeltaTime

Used to limit data to that which has arrived since the last execution of the
action.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OffsetSeconds" : Integer,
  "TimeExpression" : String
}

```

### YAML

```yaml

  OffsetSeconds: Integer
  TimeExpression: String

```

## Properties

`OffsetSeconds`

The number of seconds of estimated in-flight lag time of message data. When you create
dataset contents using message data from a specified timeframe, some message data might still
be in flight when processing begins, and so do not arrive in time to be processed. Use this
field to make allowances for the in flight time of your message data, so that data not
processed from a previous timeframe is included with the next timeframe. Otherwise, missed
message data would be excluded from processing during the next timeframe too, because its
timestamp places it within the previous timeframe.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeExpression`

An expression by which the time of the message data might be determined. This can be the
name of a timestamp field or a SQL expression that is used to derive the time the message data
was generated.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DatasetContentVersionValue

DeltaTimeSessionWindowConfiguration

All content copied from https://docs.aws.amazon.com/.
