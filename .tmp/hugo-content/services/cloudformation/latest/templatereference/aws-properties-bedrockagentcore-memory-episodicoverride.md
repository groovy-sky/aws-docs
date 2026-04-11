This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Memory EpisodicOverride

The configuration to override the episodic memory strategy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Consolidation" : EpisodicOverrideConsolidationConfigurationInput,
  "Extraction" : EpisodicOverrideExtractionConfigurationInput,
  "Reflection" : EpisodicOverrideReflectionConfigurationInput
}

```

### YAML

```yaml

  Consolidation:
    EpisodicOverrideConsolidationConfigurationInput
  Extraction:
    EpisodicOverrideExtractionConfigurationInput
  Reflection:
    EpisodicOverrideReflectionConfigurationInput

```

## Properties

`Consolidation`

Contains configurations for overriding the consolidation step of the episodic memory
strategy.

_Required_: No

_Type_: [EpisodicOverrideConsolidationConfigurationInput](aws-properties-bedrockagentcore-memory-episodicoverrideconsolidationconfigurationinput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Extraction`

Contains configurations for overriding the extraction step of the episodic memory
strategy.

_Required_: No

_Type_: [EpisodicOverrideExtractionConfigurationInput](aws-properties-bedrockagentcore-memory-episodicoverrideextractionconfigurationinput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Reflection`

Contains configurations for overriding the reflection step of the episodic memory
strategy.

_Required_: No

_Type_: [EpisodicOverrideReflectionConfigurationInput](aws-properties-bedrockagentcore-memory-episodicoverridereflectionconfigurationinput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EpisodicMemoryStrategy

EpisodicOverrideConsolidationConfigurationInput

All content copied from https://docs.aws.amazon.com/.
