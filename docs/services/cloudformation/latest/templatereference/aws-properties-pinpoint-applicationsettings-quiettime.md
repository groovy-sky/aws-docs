This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::ApplicationSettings QuietTime

Specifies the start and end times that define a time range when messages
aren't sent to endpoints.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "End" : String,
  "Start" : String
}

```

### YAML

```yaml

  End: String
  Start: String

```

## Properties

`End`

The specific time when quiet time ends. This value has to use 24-hour notation
and be in HH:MM format, where HH is the hour (with a leading zero, if
applicable) and MM is the minutes. For example, use `02:30` to
represent 2:30 AM, or `14:30` to represent 2:30 PM.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Start`

The specific time when quiet time begins. This value has to use 24-hour
notation and be in HH:MM format, where HH is the hour (with a leading zero, if
applicable) and MM is the minutes. For example, use `02:30` to
represent 2:30 AM, or `14:30` to represent 2:30 PM.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Limits

AWS::Pinpoint::BaiduChannel
