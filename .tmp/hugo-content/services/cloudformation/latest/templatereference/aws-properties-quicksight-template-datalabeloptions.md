This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template DataLabelOptions

The options that determine the presentation of the data labels.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CategoryLabelVisibility" : String,
  "DataLabelTypes" : [ DataLabelType, ... ],
  "LabelColor" : String,
  "LabelContent" : String,
  "LabelFontConfiguration" : FontConfiguration,
  "MeasureLabelVisibility" : String,
  "Overlap" : String,
  "Position" : String,
  "TotalsVisibility" : String,
  "Visibility" : String
}

```

### YAML

```yaml

  CategoryLabelVisibility: String
  DataLabelTypes:
    - DataLabelType
  LabelColor: String
  LabelContent: String
  LabelFontConfiguration:
    FontConfiguration
  MeasureLabelVisibility: String
  Overlap: String
  Position: String
  TotalsVisibility: String
  Visibility: String

```

## Properties

`CategoryLabelVisibility`

Determines the visibility of the category field labels.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataLabelTypes`

The option that determines the data label type.

_Required_: No

_Type_: Array of [DataLabelType](aws-properties-quicksight-template-datalabeltype.md)

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LabelColor`

Determines the color of the data labels.

_Required_: No

_Type_: String

_Pattern_: `^#[A-F0-9]{6}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LabelContent`

Determines the content of the data labels.

_Required_: No

_Type_: String

_Allowed values_: `VALUE | PERCENT | VALUE_AND_PERCENT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LabelFontConfiguration`

Determines the font configuration of the data labels.

_Required_: No

_Type_: [FontConfiguration](aws-properties-quicksight-template-fontconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MeasureLabelVisibility`

Determines the visibility of the measure field labels.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Overlap`

Determines whether overlap is enabled or disabled for the data labels.

_Required_: No

_Type_: String

_Allowed values_: `DISABLE_OVERLAP | ENABLE_OVERLAP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Position`

Determines the position of the data labels.

_Required_: No

_Type_: String

_Allowed values_: `INSIDE | OUTSIDE | LEFT | TOP | BOTTOM | RIGHT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TotalsVisibility`

Determines the visibility of the total.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Visibility`

Determines the visibility of the data labels.

_Required_: No

_Type_: String

_Allowed values_: `HIDDEN | VISIBLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataFieldSeriesItem

DataLabelType

All content copied from https://docs.aws.amazon.com/.
