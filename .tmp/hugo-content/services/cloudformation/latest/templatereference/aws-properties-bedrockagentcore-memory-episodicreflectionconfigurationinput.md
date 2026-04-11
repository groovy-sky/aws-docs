This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Memory EpisodicReflectionConfigurationInput

An episodic reflection configuration input.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Namespaces" : [ String, ... ],
  "NamespaceTemplates" : [ String, ... ]
}

```

### YAML

```yaml

  Namespaces:
    - String
  NamespaceTemplates:
    - String

```

## Properties

`Namespaces`

The namespaces over which to create reflections. Can be less nested than episode namespaces.

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

EpisodicOverrideReflectionConfigurationInput

InvocationConfigurationInput

All content copied from https://docs.aws.amazon.com/.
