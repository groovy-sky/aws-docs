This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::Campaign CampaignEventFilter

Specifies the settings for events that cause a campaign to be sent.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Dimensions" : EventDimensions,
  "FilterType" : String
}

```

### YAML

```yaml

  Dimensions:
    EventDimensions
  FilterType: String

```

## Properties

`Dimensions`

The dimension settings of the event filter for the campaign.

_Required_: No

_Type_: [EventDimensions](aws-properties-pinpoint-campaign-eventdimensions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterType`

The type of event that causes the campaign to be sent. Valid values are:
`SYSTEM`, sends the campaign when a system event occurs; and,
`ENDPOINT`, sends the campaign when an endpoint event (Events resource)
occurs.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CampaignEmailMessage

CampaignHook

All content copied from https://docs.aws.amazon.com/.
