This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard SheetControlInfoIconLabelOptions

A control to display info icons for filters and parameters.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InfoIconText" : String,
  "Visibility" : String
}

```

### YAML

```yaml

  InfoIconText: String
  Visibility: String

```

## Properties

`InfoIconText`

The text content of info icon.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Visibility`

The visibility configuration of info icon label options.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Sheet

SheetControlLayout

All content copied from https://docs.aws.amazon.com/.
