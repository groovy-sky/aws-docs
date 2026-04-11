This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ConnectCampaignsV2::Campaign CommunicationTimeConfig

Communication time configuration for an outbound campaign.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Email" : TimeWindow,
  "LocalTimeZoneConfig" : LocalTimeZoneConfig,
  "Sms" : TimeWindow,
  "Telephony" : TimeWindow,
  "WhatsApp" : TimeWindow
}

```

### YAML

```yaml

  Email:
    TimeWindow
  LocalTimeZoneConfig:
    LocalTimeZoneConfig
  Sms:
    TimeWindow
  Telephony:
    TimeWindow
  WhatsApp:
    TimeWindow

```

## Properties

`Email`

The communication time configuration for the email channel subtype.

_Required_: No

_Type_: [TimeWindow](aws-properties-connectcampaignsv2-campaign-timewindow.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LocalTimeZoneConfig`

The local timezone configuration.

_Required_: Yes

_Type_: [LocalTimeZoneConfig](aws-properties-connectcampaignsv2-campaign-localtimezoneconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Sms`

The communication time configuration for the SMS channel subtype.

_Required_: No

_Type_: [TimeWindow](aws-properties-connectcampaignsv2-campaign-timewindow.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Telephony`

The communication time configuration for the telephony channel subtype.

_Required_: No

_Type_: [TimeWindow](aws-properties-connectcampaignsv2-campaign-timewindow.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WhatsApp`

The communication time configuration for the WhatsApp channel subtype.

_Required_: No

_Type_: [TimeWindow](aws-properties-connectcampaignsv2-campaign-timewindow.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CommunicationLimitsConfig

DailyHour

All content copied from https://docs.aws.amazon.com/.
