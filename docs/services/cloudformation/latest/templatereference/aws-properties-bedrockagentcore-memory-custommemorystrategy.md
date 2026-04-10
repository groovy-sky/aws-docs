This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Memory CustomMemoryStrategy

The memory strategy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Configuration" : CustomConfigurationInput,
  "CreatedAt" : String,
  "Description" : String,
  "Name" : String,
  "Namespaces" : [ String, ... ],
  "NamespaceTemplates" : [ String, ... ],
  "Status" : String,
  "StrategyId" : String,
  "Type" : String,
  "UpdatedAt" : String
}

```

### YAML

```yaml

  Configuration:
    CustomConfigurationInput
  CreatedAt: String
  Description: String
  Name: String
  Namespaces:
    - String
  NamespaceTemplates:
    - String
  Status: String
  StrategyId: String
  Type: String
  UpdatedAt: String

```

## Properties

`Configuration`

The memory strategy configuration.

_Required_: No

_Type_: [CustomConfigurationInput](aws-properties-bedrockagentcore-memory-customconfigurationinput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreatedAt`

The memory strategy creation date and time.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The memory strategy description.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The memory strategy name.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z][a-zA-Z0-9_]{0,47}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Namespaces`

The memory strategy namespaces.

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

`Status`

The memory strategy status.

_Required_: No

_Type_: String

_Allowed values_: `CREATING | ACTIVE | DELETING | FAILED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StrategyId`

The memory strategy ID.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z][a-zA-Z0-9-_]{0,99}-[a-zA-Z0-9]{10}$`

_Minimum_: `12`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The memory strategy type.

_Required_: No

_Type_: String

_Allowed values_: `SEMANTIC | SUMMARIZATION | USER_PREFERENCE | CUSTOM | EPISODIC`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UpdatedAt`

The memory strategy update date and time.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomConfigurationInput

EpisodicMemoryStrategy

All content copied from https://docs.aws.amazon.com/.
