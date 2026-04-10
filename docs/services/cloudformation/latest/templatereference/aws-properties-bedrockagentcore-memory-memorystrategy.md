This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Memory MemoryStrategy

The memory strategy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomMemoryStrategy" : CustomMemoryStrategy,
  "EpisodicMemoryStrategy" : EpisodicMemoryStrategy,
  "SemanticMemoryStrategy" : SemanticMemoryStrategy,
  "SummaryMemoryStrategy" : SummaryMemoryStrategy,
  "UserPreferenceMemoryStrategy" : UserPreferenceMemoryStrategy
}

```

### YAML

```yaml

  CustomMemoryStrategy:
    CustomMemoryStrategy
  EpisodicMemoryStrategy:
    EpisodicMemoryStrategy
  SemanticMemoryStrategy:
    SemanticMemoryStrategy
  SummaryMemoryStrategy:
    SummaryMemoryStrategy
  UserPreferenceMemoryStrategy:
    UserPreferenceMemoryStrategy

```

## Properties

`CustomMemoryStrategy`

The memory strategy.

_Required_: No

_Type_: [CustomMemoryStrategy](aws-properties-bedrockagentcore-memory-custommemorystrategy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EpisodicMemoryStrategy`

The episodic memory strategy configuration.

_Required_: No

_Type_: [EpisodicMemoryStrategy](aws-properties-bedrockagentcore-memory-episodicmemorystrategy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SemanticMemoryStrategy`

The memory strategy.

_Required_: No

_Type_: [SemanticMemoryStrategy](aws-properties-bedrockagentcore-memory-semanticmemorystrategy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SummaryMemoryStrategy`

The memory strategy summary.

_Required_: No

_Type_: [SummaryMemoryStrategy](aws-properties-bedrockagentcore-memory-summarymemorystrategy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserPreferenceMemoryStrategy`

The memory strategy.

_Required_: No

_Type_: [UserPreferenceMemoryStrategy](aws-properties-bedrockagentcore-memory-userpreferencememorystrategy.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KinesisResource

MessageBasedTriggerInput

All content copied from https://docs.aws.amazon.com/.
