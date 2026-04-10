This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ConnectCampaignsV2::Campaign TelephonyOutboundConfig

The outbound configuration for telephony.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AnswerMachineDetectionConfig" : AnswerMachineDetectionConfig,
  "ConnectContactFlowId" : String,
  "ConnectSourcePhoneNumber" : String,
  "RingTimeout" : Integer
}

```

### YAML

```yaml

  AnswerMachineDetectionConfig:
    AnswerMachineDetectionConfig
  ConnectContactFlowId: String
  ConnectSourcePhoneNumber: String
  RingTimeout: Integer

```

## Properties

`AnswerMachineDetectionConfig`

The answering machine detection configuration.

_Required_: No

_Type_: [AnswerMachineDetectionConfig](aws-properties-connectcampaignsv2-campaign-answermachinedetectionconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConnectContactFlowId`

The identifier of the published Amazon Connect contact flow.

_Required_: Yes

_Type_: String

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConnectSourcePhoneNumber`

The Amazon Connect source phone number.

_Required_: No

_Type_: String

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RingTimeout`

The ring timeout configuration for outbound calls. Specifies how long to wait for the call to be answered before timing out.

_Required_: No

_Type_: Integer

_Minimum_: `15`

_Maximum_: `60`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TelephonyChannelSubtypeConfig

TelephonyOutboundMode

All content copied from https://docs.aws.amazon.com/.
