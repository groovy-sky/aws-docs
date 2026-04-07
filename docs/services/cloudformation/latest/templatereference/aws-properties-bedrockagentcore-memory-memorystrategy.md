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

_Type_: [CustomMemoryStrategy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrockagentcore-memory-custommemorystrategy.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EpisodicMemoryStrategy`

The episodic memory strategy configuration.

_Required_: No

_Type_: [EpisodicMemoryStrategy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrockagentcore-memory-episodicmemorystrategy.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SemanticMemoryStrategy`

The memory strategy.

_Required_: No

_Type_: [SemanticMemoryStrategy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrockagentcore-memory-semanticmemorystrategy.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SummaryMemoryStrategy`

The memory strategy summary.

_Required_: No

_Type_: [SummaryMemoryStrategy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrockagentcore-memory-summarymemorystrategy.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserPreferenceMemoryStrategy`

The memory strategy.

_Required_: No

_Type_: [UserPreferenceMemoryStrategy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrockagentcore-memory-userpreferencememorystrategy.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

KinesisResource

MessageBasedTriggerInput
