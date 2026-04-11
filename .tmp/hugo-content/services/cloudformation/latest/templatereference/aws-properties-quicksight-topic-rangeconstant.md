This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Topic RangeConstant

The value of the constant that is used to specify the endpoints of a range filter.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Maximum" : String,
  "Minimum" : String
}

```

### YAML

```yaml

  Maximum: String
  Minimum: String

```

## Properties

`Maximum`

The maximum value for a range constant.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Minimum`

The minimum value for a range constant.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NegativeFormat

SemanticEntityType

All content copied from https://docs.aws.amazon.com/.
