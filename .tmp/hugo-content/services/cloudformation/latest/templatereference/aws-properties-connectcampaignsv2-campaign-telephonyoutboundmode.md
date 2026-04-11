This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ConnectCampaignsV2::Campaign TelephonyOutboundMode

Contains information about telephony outbound mode.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AgentlessConfig" : Json,
  "PredictiveConfig" : PredictiveConfig,
  "PreviewConfig" : PreviewConfig,
  "ProgressiveConfig" : ProgressiveConfig
}

```

### YAML

```yaml

  AgentlessConfig: Json
  PredictiveConfig:
    PredictiveConfig
  PreviewConfig:
    PreviewConfig
  ProgressiveConfig:
    ProgressiveConfig

```

## Properties

`AgentlessConfig`

The agentless outbound mode configuration for telephony.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PredictiveConfig`

Contains predictive outbound mode configuration.

_Required_: No

_Type_: [PredictiveConfig](aws-properties-connectcampaignsv2-campaign-predictiveconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PreviewConfig`

Contains preview outbound mode configuration.

_Required_: No

_Type_: [PreviewConfig](aws-properties-connectcampaignsv2-campaign-previewconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProgressiveConfig`

Contains progressive telephony outbound mode configuration.

_Required_: No

_Type_: [ProgressiveConfig](aws-properties-connectcampaignsv2-campaign-progressiveconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TelephonyOutboundConfig

TimeoutConfig

All content copied from https://docs.aws.amazon.com/.
