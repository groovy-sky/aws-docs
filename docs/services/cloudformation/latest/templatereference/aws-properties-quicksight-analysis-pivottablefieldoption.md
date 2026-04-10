This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis PivotTableFieldOption

The selected field options for the pivot table field options.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomLabel" : String,
  "FieldId" : String,
  "Visibility" : String
}

```

### YAML

```yaml

  CustomLabel: String
  FieldId: String
  Visibility: String

```

## Properties

`CustomLabel`

The custom label of the pivot table field.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FieldId`

The field ID of the pivot table field.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Visibility`

The visibility of the pivot table field.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PivotTableFieldCollapseStateTarget

PivotTableFieldOptions

All content copied from https://docs.aws.amazon.com/.
