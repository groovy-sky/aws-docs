This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Memory EpisodicMemoryStrategy

The configuration for an episodic memory strategy. Episodic memory stores and retrieves specific interaction episodes between agents and users.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CreatedAt" : String,
  "Description" : String,
  "Name" : String,
  "Namespaces" : [ String, ... ],
  "NamespaceTemplates" : [ String, ... ],
  "ReflectionConfiguration" : EpisodicReflectionConfigurationInput,
  "Status" : String,
  "StrategyId" : String,
  "Type" : String,
  "UpdatedAt" : String
}

```

### YAML

```yaml

  CreatedAt: String
  Description: String
  Name: String
  Namespaces:
    - String
  NamespaceTemplates:
    - String
  ReflectionConfiguration:
    EpisodicReflectionConfigurationInput
  Status: String
  StrategyId: String
  Type: String
  UpdatedAt: String

```

## Properties

`CreatedAt`

The timestamp when the episodic memory strategy was created.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the episodic memory strategy.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the episodic memory strategy.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z][a-zA-Z0-9_]{0,47}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Namespaces`

The namespaces for which to create episodes.

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

`ReflectionConfiguration`

The configuration for the reflections created with the episodic memory
strategy.

_Required_: No

_Type_: [EpisodicReflectionConfigurationInput](aws-properties-bedrockagentcore-memory-episodicreflectionconfigurationinput.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

The current status of the episodic memory strategy.

_Required_: No

_Type_: String

_Allowed values_: `CREATING | ACTIVE | DELETING | FAILED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StrategyId`

The unique identifier of the memory strategy.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z][a-zA-Z0-9-_]{0,99}-[a-zA-Z0-9]{10}$`

_Minimum_: `12`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of the memory strategy.

_Required_: No

_Type_: String

_Allowed values_: `SEMANTIC | SUMMARIZATION | USER_PREFERENCE | CUSTOM | EPISODIC`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UpdatedAt`

The timestamp when the episodic memory strategy was last updated.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomMemoryStrategy

EpisodicOverride

All content copied from https://docs.aws.amazon.com/.
