This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ConnectCampaigns::Campaign OutboundCallConfig

Contains outbound call configuration for an outbound campaign.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AnswerMachineDetectionConfig" : AnswerMachineDetectionConfig,
  "ConnectContactFlowArn" : String,
  "ConnectQueueArn" : String,
  "ConnectSourcePhoneNumber" : String
}

```

### YAML

```yaml

  AnswerMachineDetectionConfig:
    AnswerMachineDetectionConfig
  ConnectContactFlowArn: String
  ConnectQueueArn: String
  ConnectSourcePhoneNumber: String

```

## Properties

`AnswerMachineDetectionConfig`

Whether answering machine detection has been enabled.

_Required_: No

_Type_: [AnswerMachineDetectionConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connectcampaigns-campaign-answermachinedetectionconfig.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConnectContactFlowArn`

The Amazon Resource Name (ARN) of the flow.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/contact-flow/[-a-zA-Z0-9]*$`

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConnectQueueArn`

The Amazon Resource Name (ARN) of the queue.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/queue/[-a-zA-Z0-9]*$`

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConnectSourcePhoneNumber`

The phone number associated with the outbound call. This is the caller ID that is
displayed to customers when an agent calls them.

_Required_: No

_Type_: String

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DialerConfig

PredictiveDialerConfig
