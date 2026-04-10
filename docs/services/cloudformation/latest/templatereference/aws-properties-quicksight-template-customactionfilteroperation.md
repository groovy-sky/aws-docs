This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template CustomActionFilterOperation

The filter operation that filters data included in a visual or in an entire sheet.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SelectedFieldsConfiguration" : FilterOperationSelectedFieldsConfiguration,
  "TargetVisualsConfiguration" : FilterOperationTargetVisualsConfiguration
}

```

### YAML

```yaml

  SelectedFieldsConfiguration:
    FilterOperationSelectedFieldsConfiguration
  TargetVisualsConfiguration:
    FilterOperationTargetVisualsConfiguration

```

## Properties

`SelectedFieldsConfiguration`

The configuration that chooses the fields to be filtered.

_Required_: Yes

_Type_: [FilterOperationSelectedFieldsConfiguration](aws-properties-quicksight-template-filteroperationselectedfieldsconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetVisualsConfiguration`

The configuration that chooses the target visuals to be filtered.

_Required_: Yes

_Type_: [FilterOperationTargetVisualsConfiguration](aws-properties-quicksight-template-filteroperationtargetvisualsconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CurrencyDisplayFormatConfiguration

CustomActionNavigationOperation

All content copied from https://docs.aws.amazon.com/.
