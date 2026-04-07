This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ConnectCampaignsV2::Campaign Source

Contains source configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomerProfilesSegmentArn" : String,
  "EventTrigger" : EventTrigger
}

```

### YAML

```yaml

  CustomerProfilesSegmentArn: String
  EventTrigger:
    EventTrigger

```

## Properties

`CustomerProfilesSegmentArn`

The Amazon Resource Name (ARN) of the Customer Profiles segment.

_Required_: No

_Type_: String

_Pattern_: `^arn:.*$`

_Minimum_: `20`

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EventTrigger`

The event trigger of the campaign.

_Required_: No

_Type_: [EventTrigger](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connectcampaignsv2-campaign-eventtrigger.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SmsOutboundMode

Tag
