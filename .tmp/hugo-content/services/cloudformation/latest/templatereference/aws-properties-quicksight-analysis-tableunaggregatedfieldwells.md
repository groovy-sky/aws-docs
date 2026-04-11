This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis TableUnaggregatedFieldWells

The unaggregated field well for the table.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Values" : [ UnaggregatedField, ... ]
}

```

### YAML

```yaml

  Values:
    - UnaggregatedField

```

## Properties

`Values`

The values field well for a pivot table. Values are unaggregated for an unaggregated table.

_Required_: No

_Type_: Array of [UnaggregatedField](aws-properties-quicksight-analysis-unaggregatedfield.md)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TableStyleTarget

TableVisual

All content copied from https://docs.aws.amazon.com/.
