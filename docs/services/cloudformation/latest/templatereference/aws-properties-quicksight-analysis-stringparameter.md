This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis StringParameter

A string parameter.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "Values" : [ String, ... ]
}

```

### YAML

```yaml

  Name: String
  Values:
    - String

```

## Properties

`Name`

A display name for a string parameter.

_Required_: Yes

_Type_: String

_Pattern_: `\S`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Values`

The values of a string parameter.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StringFormatConfiguration

StringParameterDeclaration

All content copied from https://docs.aws.amazon.com/.
