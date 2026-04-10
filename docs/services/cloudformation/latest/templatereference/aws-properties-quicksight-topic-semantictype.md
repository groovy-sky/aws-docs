This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Topic SemanticType

A structure that represents a semantic type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FalseyCellValue" : String,
  "FalseyCellValueSynonyms" : [ String, ... ],
  "SubTypeName" : String,
  "TruthyCellValue" : String,
  "TruthyCellValueSynonyms" : [ String, ... ],
  "TypeName" : String,
  "TypeParameters" : {Key: Value, ...}
}

```

### YAML

```yaml

  FalseyCellValue: String
  FalseyCellValueSynonyms:
    - String
  SubTypeName: String
  TruthyCellValue: String
  TruthyCellValueSynonyms:
    - String
  TypeName: String
  TypeParameters:
    Key: Value

```

## Properties

`FalseyCellValue`

The semantic type falsey cell value.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FalseyCellValueSynonyms`

The other names or aliases for the false cell value.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubTypeName`

The semantic type sub type name.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TruthyCellValue`

The semantic type truthy cell value.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TruthyCellValueSynonyms`

The other names or aliases for the true cell value.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TypeName`

The semantic type name.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TypeParameters`

The semantic type parameters.

_Required_: No

_Type_: Object of String

_Pattern_: `.+`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SemanticEntityType

Tag

All content copied from https://docs.aws.amazon.com/.
