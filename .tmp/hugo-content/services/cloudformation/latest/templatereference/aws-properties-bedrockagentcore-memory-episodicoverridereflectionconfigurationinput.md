This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Memory EpisodicOverrideReflectionConfigurationInput

Configurations for overriding the reflection step of the episodic memory
strategy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AppendToPrompt" : String,
  "ModelId" : String,
  "Namespaces" : [ String, ... ],
  "NamespaceTemplates" : [ String, ... ]
}

```

### YAML

```yaml

  AppendToPrompt: String
  ModelId: String
  Namespaces:
    - String
  NamespaceTemplates:
    - String

```

## Properties

`AppendToPrompt`

The text to append to the prompt for reflection step of the episodic memory
strategy.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `30000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ModelId`

The model ID to use for the reflection step of the episodic memory strategy.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Namespaces`

The namespaces to use for episodic reflection. Can be less nested than the episodic
namespaces.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NamespaceTemplates`

Property description not available.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EpisodicOverrideExtractionConfigurationInput

EpisodicReflectionConfigurationInput

All content copied from https://docs.aws.amazon.com/.
