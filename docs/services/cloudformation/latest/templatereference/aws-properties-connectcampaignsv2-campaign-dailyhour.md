This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ConnectCampaignsV2::Campaign DailyHour

The daily hours configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Value" : [ TimeRange, ... ]
}

```

### YAML

```yaml

  Key: String
  Value:
    - TimeRange

```

## Properties

`Key`

The key for DailyHour.

_Required_: No

_Type_: String

_Allowed values_: `MONDAY | TUESDAY | WEDNESDAY | THURSDAY | FRIDAY | SATURDAY | SUNDAY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value for DailyHour.

_Required_: No

_Type_: Array of [TimeRange](aws-properties-connectcampaignsv2-campaign-timerange.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CommunicationTimeConfig

EmailChannelSubtypeConfig

All content copied from https://docs.aws.amazon.com/.
