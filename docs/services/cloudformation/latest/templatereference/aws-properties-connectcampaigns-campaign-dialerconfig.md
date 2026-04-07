This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ConnectCampaigns::Campaign DialerConfig

Contains dialer configuration for an outbound campaign.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AgentlessDialerConfig" : AgentlessDialerConfig,
  "PredictiveDialerConfig" : PredictiveDialerConfig,
  "ProgressiveDialerConfig" : ProgressiveDialerConfig
}

```

### YAML

```yaml

  AgentlessDialerConfig:
    AgentlessDialerConfig
  PredictiveDialerConfig:
    PredictiveDialerConfig
  ProgressiveDialerConfig:
    ProgressiveDialerConfig

```

## Properties

`AgentlessDialerConfig`

The configuration of the agentless dialer.

_Required_: No

_Type_: [AgentlessDialerConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connectcampaigns-campaign-agentlessdialerconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PredictiveDialerConfig`

The configuration of the predictive dialer.

_Required_: No

_Type_: [PredictiveDialerConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connectcampaigns-campaign-predictivedialerconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProgressiveDialerConfig`

The configuration of the progressive dialer.

_Required_: No

_Type_: [ProgressiveDialerConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connectcampaigns-campaign-progressivedialerconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AnswerMachineDetectionConfig

OutboundCallConfig
