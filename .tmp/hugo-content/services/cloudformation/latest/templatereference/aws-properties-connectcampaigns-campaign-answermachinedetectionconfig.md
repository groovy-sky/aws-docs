This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ConnectCampaigns::Campaign AnswerMachineDetectionConfig

Contains information about answering machine detection.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AwaitAnswerMachinePrompt" : Boolean,
  "EnableAnswerMachineDetection" : Boolean
}

```

### YAML

```yaml

  AwaitAnswerMachinePrompt: Boolean
  EnableAnswerMachineDetection: Boolean

```

## Properties

`AwaitAnswerMachinePrompt`

Whether waiting for answer machine prompt is enabled.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableAnswerMachineDetection`

Whether answering machine detection is enabled.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AgentlessDialerConfig

DialerConfig

All content copied from https://docs.aws.amazon.com/.
