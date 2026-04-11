This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ConnectCampaignsV2::Campaign PreviewConfig

Contains preview outbound mode configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AgentActions" : [ String, ... ],
  "BandwidthAllocation" : Number,
  "TimeoutConfig" : TimeoutConfig
}

```

### YAML

```yaml

  AgentActions:
    - String
  BandwidthAllocation: Number
  TimeoutConfig:
    TimeoutConfig

```

## Properties

`AgentActions`

Agent actions for the preview outbound mode.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BandwidthAllocation`

Bandwidth allocation for the preview outbound mode.

_Required_: Yes

_Type_: Number

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeoutConfig`

Countdown timer configuration for preview outbound mode.

_Required_: Yes

_Type_: [TimeoutConfig](aws-properties-connectcampaignsv2-campaign-timeoutconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PredictiveConfig

ProgressiveConfig

All content copied from https://docs.aws.amazon.com/.
