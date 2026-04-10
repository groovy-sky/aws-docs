This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Memory SelfManagedConfiguration

The self managed configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HistoricalContextWindowSize" : Integer,
  "InvocationConfiguration" : InvocationConfigurationInput,
  "TriggerConditions" : [ TriggerConditionInput, ... ]
}

```

### YAML

```yaml

  HistoricalContextWindowSize: Integer
  InvocationConfiguration:
    InvocationConfigurationInput
  TriggerConditions:
    - TriggerConditionInput

```

## Properties

`HistoricalContextWindowSize`

The memory configuration for self managed.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InvocationConfiguration`

The self managed configuration.

_Required_: No

_Type_: [InvocationConfigurationInput](aws-properties-bedrockagentcore-memory-invocationconfigurationinput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TriggerConditions`

A list of conditions that trigger memory processing.

_Required_: No

_Type_: Array of [TriggerConditionInput](aws-properties-bedrockagentcore-memory-triggerconditioninput.md)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MessageBasedTriggerInput

SemanticMemoryStrategy

All content copied from https://docs.aws.amazon.com/.
