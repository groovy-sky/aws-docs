This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Memory SemanticOverrideConsolidationConfigurationInput

The memory override configuration.

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

The override configuration.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `30000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelId`

The memory override model ID.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SemanticOverride

SemanticOverrideExtractionConfigurationInput

All content copied from https://docs.aws.amazon.com/.
