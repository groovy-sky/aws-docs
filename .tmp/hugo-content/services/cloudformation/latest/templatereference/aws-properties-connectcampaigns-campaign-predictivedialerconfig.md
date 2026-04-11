This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ConnectCampaigns::Campaign PredictiveDialerConfig

Contains predictive dialer configuration for an outbound campaign.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BandwidthAllocation" : Number,
  "DialingCapacity" : Number
}

```

### YAML

```yaml

  BandwidthAllocation: Number
  DialingCapacity: Number

```

## Properties

`BandwidthAllocation`

Bandwidth allocation for the predictive dialer.

_Required_: Yes

_Type_: Number

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DialingCapacity`

The allocation of dialing capacity between multiple active campaigns.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OutboundCallConfig

ProgressiveDialerConfig

All content copied from https://docs.aws.amazon.com/.
