This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard ThousandSeparatorOptions

The options that determine the thousands separator configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "GroupingStyle" : String,
  "Symbol" : String,
  "Visibility" : String
}

```

### YAML

```yaml

  GroupingStyle: String
  Symbol: String
  Visibility: String

```

## Properties

`GroupingStyle`

Determines the way numbers are styled to accommodate different readability standards. The `DEFAULT` value uses the standard international grouping system and groups numbers by the thousands. The `LAKHS` value uses the Indian numbering system and groups numbers by lakhs and crores.

_Required_: No

_Type_: String

_Allowed values_: `DEFAULT | LAKHS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Symbol`

Determines the thousands separator symbol.

_Required_: No

_Type_: String

_Allowed values_: `COMMA | DOT | SPACE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Visibility`

Determines the visibility of the thousands separator.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TextFieldControlDisplayOptions

TimeBasedForecastProperties

All content copied from https://docs.aws.amazon.com/.
