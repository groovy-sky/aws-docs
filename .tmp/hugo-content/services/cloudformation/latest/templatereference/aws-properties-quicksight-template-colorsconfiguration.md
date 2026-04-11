This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template ColorsConfiguration

The color configurations for a column.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomColors" : [ CustomColor, ... ]
}

```

### YAML

```yaml

  CustomColors:
    - CustomColor

```

## Properties

`CustomColors`

A list of up to 50 custom colors.

_Required_: No

_Type_: Array of [CustomColor](aws-properties-quicksight-template-customcolor.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ColorScale

ColumnConfiguration

All content copied from https://docs.aws.amazon.com/.
