This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Memory EpisodicOverrideExtractionConfigurationInput

Configurations for overriding the extraction step of the episodic memory
strategy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AppendToPrompt" : String,
  "ModelId" : String
}

```

### YAML

```yaml

  AppendToPrompt: String
  ModelId: String

```

## Properties

`AppendToPrompt`

The text to append to the prompt for the extraction step of the episodic memory
strategy.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `30000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelId`

The model ID to use for the extraction step of the episodic memory strategy.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EpisodicOverrideConsolidationConfigurationInput

EpisodicOverrideReflectionConfigurationInput

All content copied from https://docs.aws.amazon.com/.
