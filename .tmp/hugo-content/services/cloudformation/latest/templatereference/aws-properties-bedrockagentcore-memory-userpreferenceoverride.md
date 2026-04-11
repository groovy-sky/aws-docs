This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Memory UserPreferenceOverride

The memory user preference override.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Consolidation" : UserPreferenceOverrideConsolidationConfigurationInput,
  "Extraction" : UserPreferenceOverrideExtractionConfigurationInput
}

```

### YAML

```yaml

  Consolidation:
    UserPreferenceOverrideConsolidationConfigurationInput
  Extraction:
    UserPreferenceOverrideExtractionConfigurationInput

```

## Properties

`Consolidation`

The memory override consolidation information.

_Required_: No

_Type_: [UserPreferenceOverrideConsolidationConfigurationInput](aws-properties-bedrockagentcore-memory-userpreferenceoverrideconsolidationconfigurationinput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Extraction`

The memory user preferences for extraction.

_Required_: No

_Type_: [UserPreferenceOverrideExtractionConfigurationInput](aws-properties-bedrockagentcore-memory-userpreferenceoverrideextractionconfigurationinput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UserPreferenceMemoryStrategy

UserPreferenceOverrideConsolidationConfigurationInput

All content copied from https://docs.aws.amazon.com/.
