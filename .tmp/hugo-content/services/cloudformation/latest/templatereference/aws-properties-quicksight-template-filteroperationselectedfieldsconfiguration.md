This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template FilterOperationSelectedFieldsConfiguration

The configuration of selected fields in the `CustomActionFilterOperation`.

This is a union type structure. For this structure to be valid, only one of the attributes can be defined.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SelectedColumns" : [ ColumnIdentifier, ... ],
  "SelectedFieldOptions" : String,
  "SelectedFields" : [ String, ... ]
}

```

### YAML

```yaml

  SelectedColumns:
    - ColumnIdentifier
  SelectedFieldOptions: String
  SelectedFields:
    - String

```

## Properties

`SelectedColumns`

The selected columns of a dataset.

_Required_: No

_Type_: Array of [ColumnIdentifier](aws-properties-quicksight-template-columnidentifier.md)

_Minimum_: `0`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SelectedFieldOptions`

A structure that contains the options that choose which fields are filtered in the `CustomActionFilterOperation`.

Valid values are defined as follows:

- `ALL_FIELDS`: Applies the filter operation to all fields.

_Required_: No

_Type_: String

_Allowed values_: `ALL_FIELDS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SelectedFields`

Chooses the fields that are filtered in `CustomActionFilterOperation`.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `512 | 20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FilterListControl

FilterOperationTargetVisualsConfiguration

All content copied from https://docs.aws.amazon.com/.
