This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis Sheet

A _sheet_, which is an object that contains a set of visuals that
are viewed together on one page in Quick Sight. Every analysis and dashboard
contains at least one sheet. Each sheet contains at least one visualization widget, for
example a chart, pivot table, or narrative insight. Sheets can be associated with other
components, such as controls, filters, and so on.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "SheetId" : String
}

```

### YAML

```yaml

  Name: String
  SheetId: String

```

## Properties

`Name`

The name of a sheet. This name is displayed on the sheet's tab in the Quick Sight
console.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SheetId`

The unique identifier associated with a sheet.

_Required_: No

_Type_: String

_Pattern_: `^[\w\-]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ShapeConditionalFormat

SheetControlInfoIconLabelOptions
