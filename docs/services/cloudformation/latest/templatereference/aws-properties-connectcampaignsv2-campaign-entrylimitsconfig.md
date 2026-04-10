This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ConnectCampaignsV2::Campaign EntryLimitsConfig

The `EntryLimitsConfig` property type specifies Property description not available. for an [AWS::ConnectCampaignsV2::Campaign](aws-resource-connectcampaignsv2-campaign.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxEntryCount" : Integer,
  "MinEntryInterval" : String
}

```

### YAML

```yaml

  MaxEntryCount: Integer
  MinEntryInterval: String

```

## Properties

`MaxEntryCount`

Property description not available.

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinEntryInterval`

Property description not available.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9.]*$`

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EmailOutboundMode

EventTrigger

All content copied from https://docs.aws.amazon.com/.
