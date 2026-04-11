This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard DestinationParameterValueConfiguration

The configuration of destination parameter values.

This is a union type structure. For this structure to be valid, only one of the attributes can be defined.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomValuesConfiguration" : CustomValuesConfiguration,
  "SelectAllValueOptions" : String,
  "SourceColumn" : ColumnIdentifier,
  "SourceField" : String,
  "SourceParameterName" : String
}

```

### YAML

```yaml

  CustomValuesConfiguration:
    CustomValuesConfiguration
  SelectAllValueOptions: String
  SourceColumn:
    ColumnIdentifier
  SourceField: String
  SourceParameterName: String

```

## Properties

`CustomValuesConfiguration`

The configuration of custom values for destination parameter in `DestinationParameterValueConfiguration`.

_Required_: No

_Type_: [CustomValuesConfiguration](aws-properties-quicksight-dashboard-customvaluesconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SelectAllValueOptions`

The configuration that selects all options.

_Required_: No

_Type_: String

_Allowed values_: `ALL_VALUES`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceColumn`

A column of a data set.

_Required_: No

_Type_: [ColumnIdentifier](aws-properties-quicksight-dashboard-columnidentifier.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceField`

The source field ID of the destination parameter.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceParameterName`

The source parameter name of the destination parameter.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DefaultTextFieldControlOptions

DimensionField

All content copied from https://docs.aws.amazon.com/.
