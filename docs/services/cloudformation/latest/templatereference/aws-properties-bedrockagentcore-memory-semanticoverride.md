This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Memory SemanticOverride

The memory override.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Consolidation" : SemanticOverrideConsolidationConfigurationInput,
  "Extraction" : SemanticOverrideExtractionConfigurationInput
}

```

### YAML

```yaml

  Consolidation:
    SemanticOverrideConsolidationConfigurationInput
  Extraction:
    SemanticOverrideExtractionConfigurationInput

```

## Properties

`Consolidation`

The memory override consolidation.

_Required_: No

_Type_: [SemanticOverrideConsolidationConfigurationInput](aws-properties-bedrockagentcore-memory-semanticoverrideconsolidationconfigurationinput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Extraction`

The memory override extraction.

_Required_: No

_Type_: [SemanticOverrideExtractionConfigurationInput](aws-properties-bedrockagentcore-memory-semanticoverrideextractionconfigurationinput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SemanticMemoryStrategy

SemanticOverrideConsolidationConfigurationInput

All content copied from https://docs.aws.amazon.com/.
