This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Topic TopicNamedEntity

A structure that represents a named entity.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Definition" : [ NamedEntityDefinition, ... ],
  "EntityDescription" : String,
  "EntityName" : String,
  "EntitySynonyms" : [ String, ... ],
  "SemanticEntityType" : SemanticEntityType
}

```

### YAML

```yaml

  Definition:
    - NamedEntityDefinition
  EntityDescription: String
  EntityName: String
  EntitySynonyms:
    - String
  SemanticEntityType:
    SemanticEntityType

```

## Properties

`Definition`

The definition of a named entity.

_Required_: No

_Type_: Array of [NamedEntityDefinition](aws-properties-quicksight-topic-namedentitydefinition.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EntityDescription`

The description of the named entity.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EntityName`

The name of the named entity.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EntitySynonyms`

The other
names or aliases for the named entity.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SemanticEntityType`

The type of named entity that a topic represents.

_Required_: No

_Type_: [SemanticEntityType](aws-properties-quicksight-topic-semanticentitytype.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TopicFilter

TopicNumericEqualityFilter

All content copied from https://docs.aws.amazon.com/.
