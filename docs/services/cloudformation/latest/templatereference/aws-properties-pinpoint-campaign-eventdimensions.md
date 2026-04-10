This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::Campaign EventDimensions

Specifies the dimensions for an event filter that determines when a campaign
is sent or a journey activity is performed.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Attributes" : Json,
  "EventType" : SetDimension,
  "Metrics" : Json
}

```

### YAML

```yaml

  Attributes: Json
  EventType:
    SetDimension
  Metrics: Json

```

## Properties

`Attributes`

One or more custom attributes that your application reports to Amazon
Pinpoint. You can use these attributes as selection criteria when you create an
event filter.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EventType`

The name of the event that causes the campaign to be sent or the journey
activity to be performed. This can be a standard event that Amazon Pinpoint
generates, such as `_email.delivered` or
`_custom.delivered`. For campaigns, this can also be a custom
event that's specific to your application. For information about standard
events, see [Streaming Amazon\
Pinpoint Events](../../../pinpoint/latest/developerguide/event-streams.md) in the _Amazon Pinpoint Developer_
_Guide_.

_Required_: No

_Type_: [SetDimension](aws-properties-pinpoint-campaign-setdimension.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Metrics`

One or more custom metrics that your application reports to Amazon Pinpoint. You can use these metrics as selection criteria when you create an event
filter.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DefaultButtonConfiguration

InAppMessageBodyConfig

All content copied from https://docs.aws.amazon.com/.
