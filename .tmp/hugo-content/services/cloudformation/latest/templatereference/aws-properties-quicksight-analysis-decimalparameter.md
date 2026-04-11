This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis DecimalParameter

A decimal parameter.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "Values" : [ Number, ... ]
}

```

### YAML

```yaml

  Name: String
  Values:
    - Number

```

## Properties

`Name`

A display name for the decimal parameter.

_Required_: Yes

_Type_: String

_Pattern_: `\S`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Values`

The values for the decimal parameter.

_Required_: Yes

_Type_: Array of Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DecimalDefaultValues

DecimalParameterDeclaration

All content copied from https://docs.aws.amazon.com/.
