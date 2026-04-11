This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ConnectCampaignsV2::Campaign TimeWindow

Contains information about a time window.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OpenHours" : OpenHours,
  "RestrictedPeriods" : RestrictedPeriods
}

```

### YAML

```yaml

  OpenHours:
    OpenHours
  RestrictedPeriods:
    RestrictedPeriods

```

## Properties

`OpenHours`

The open hours configuration.

_Required_: Yes

_Type_: [OpenHours](aws-properties-connectcampaignsv2-campaign-openhours.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RestrictedPeriods`

The restricted periods configuration.

_Required_: No

_Type_: [RestrictedPeriods](aws-properties-connectcampaignsv2-campaign-restrictedperiods.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TimeRange

WhatsAppChannelSubtypeConfig

All content copied from https://docs.aws.amazon.com/.
