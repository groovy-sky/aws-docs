This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Memory TriggerConditionInput

The memory trigger condition input.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MessageBasedTrigger" : MessageBasedTriggerInput,
  "TimeBasedTrigger" : TimeBasedTriggerInput,
  "TokenBasedTrigger" : TokenBasedTriggerInput
}

```

### YAML

```yaml

  MessageBasedTrigger:
    MessageBasedTriggerInput
  TimeBasedTrigger:
    TimeBasedTriggerInput
  TokenBasedTrigger:
    TokenBasedTriggerInput

```

## Properties

`MessageBasedTrigger`

The memory trigger condition input for the message based trigger.

_Required_: No

_Type_: [MessageBasedTriggerInput](aws-properties-bedrockagentcore-memory-messagebasedtriggerinput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeBasedTrigger`

The memory trigger condition input.

_Required_: No

_Type_: [TimeBasedTriggerInput](aws-properties-bedrockagentcore-memory-timebasedtriggerinput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TokenBasedTrigger`

The trigger condition information for a token based trigger.

_Required_: No

_Type_: [TokenBasedTriggerInput](aws-properties-bedrockagentcore-memory-tokenbasedtriggerinput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TokenBasedTriggerInput

UserPreferenceMemoryStrategy

All content copied from https://docs.aws.amazon.com/.
