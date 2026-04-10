This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Topic SemanticEntityType

A structure that represents a semantic entity type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SubTypeName" : String,
  "TypeName" : String,
  "TypeParameters" : {Key: Value, ...}
}

```

### YAML

```yaml

  SubTypeName: String
  TypeName: String
  TypeParameters:
    Key: Value

```

## Properties

`SubTypeName`

The semantic entity sub type name.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TypeName`

The semantic entity type name.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TypeParameters`

The semantic entity type parameters.

_Required_: No

_Type_: Object of String

_Pattern_: `.+`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RangeConstant

SemanticType

All content copied from https://docs.aws.amazon.com/.
