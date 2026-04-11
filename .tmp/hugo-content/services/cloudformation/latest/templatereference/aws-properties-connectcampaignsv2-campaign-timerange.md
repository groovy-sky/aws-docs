This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ConnectCampaignsV2::Campaign TimeRange

Contains information about a time range.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EndTime" : String,
  "StartTime" : String
}

```

### YAML

```yaml

  EndTime: String
  StartTime: String

```

## Properties

`EndTime`

The end time of the time range.

_Required_: Yes

_Type_: String

_Pattern_: `^T\d{2}:\d{2}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartTime`

The start time of the time range.

_Required_: Yes

_Type_: String

_Pattern_: `^T\d{2}:\d{2}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TimeoutConfig

TimeWindow

All content copied from https://docs.aws.amazon.com/.
