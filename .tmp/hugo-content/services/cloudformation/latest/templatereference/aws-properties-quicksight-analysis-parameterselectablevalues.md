This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis ParameterSelectableValues

A list of selectable values that are used in a control.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LinkToDataSetColumn" : ColumnIdentifier,
  "Values" : [ String, ... ]
}

```

### YAML

```yaml

  LinkToDataSetColumn:
    ColumnIdentifier
  Values:
    - String

```

## Properties

`LinkToDataSetColumn`

The column identifier that fetches values from the data set.

_Required_: No

_Type_: [ColumnIdentifier](aws-properties-quicksight-analysis-columnidentifier.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Values`

The values that are used in `ParameterSelectableValues`.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `50000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Parameters

ParameterSliderControl

All content copied from https://docs.aws.amazon.com/.
