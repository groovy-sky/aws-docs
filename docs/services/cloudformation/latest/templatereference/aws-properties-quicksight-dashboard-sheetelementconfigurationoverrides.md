This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard SheetElementConfigurationOverrides

The override configuration of the rendering rules of a sheet.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Visibility" : String
}

```

### YAML

```yaml

  Visibility: String

```

## Properties

`Visibility`

Determines whether or not the overrides are visible. Choose one of the following options:

- `VISIBLE`

- `HIDDEN`

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SheetDefinition

SheetElementRenderingRule

All content copied from https://docs.aws.amazon.com/.
