This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ConnectCampaignsV2::Campaign Schedule

Contains the schedule configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EndTime" : String,
  "RefreshFrequency" : String,
  "StartTime" : String
}

```

### YAML

```yaml

  EndTime: String
  RefreshFrequency: String
  StartTime: String

```

## Properties

`EndTime`

The end time of the schedule in UTC.

_Required_: Yes

_Type_: String

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RefreshFrequency`

The refresh frequency of the campaign.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9.]*$`

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartTime`

The start time of the schedule in UTC.

_Required_: Yes

_Type_: String

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

RestrictedPeriods

SmsChannelSubtypeConfig
