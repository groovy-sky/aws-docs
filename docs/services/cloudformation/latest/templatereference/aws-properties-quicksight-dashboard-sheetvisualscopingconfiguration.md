This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard SheetVisualScopingConfiguration

The filter that is applied to the options.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Scope" : String,
  "SheetId" : String,
  "VisualIds" : [ String, ... ]
}

```

### YAML

```yaml

  Scope: String
  SheetId: String
  VisualIds:
    - String

```

## Properties

`Scope`

The scope of the applied entities. Choose one of the following options:

- `ALL_VISUALS`

- `SELECTED_VISUALS`

_Required_: Yes

_Type_: String

_Allowed values_: `ALL_VISUALS | SELECTED_VISUALS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SheetId`

The selected sheet that the filter is applied to.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\-]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VisualIds`

The selected visuals that the filter is applied to.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 0`

_Maximum_: `512 | 50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SheetTextBox

ShortFormatText

All content copied from https://docs.aws.amazon.com/.
