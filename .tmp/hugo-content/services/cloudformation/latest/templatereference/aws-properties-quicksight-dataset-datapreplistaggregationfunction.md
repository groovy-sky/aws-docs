This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet DataPrepListAggregationFunction

An aggregation function that concatenates values from multiple rows into a single string with a specified
separator.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Distinct" : Boolean,
  "InputColumnName" : String,
  "Separator" : String
}

```

### YAML

```yaml

  Distinct: Boolean
  InputColumnName: String
  Separator: String

```

## Properties

`Distinct`

Whether to include only distinct values in the concatenated result, removing duplicates.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputColumnName`

The name of the column containing values to be concatenated.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Separator`

The string used to separate values in the concatenated result.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataPrepConfiguration

DataPrepPercentileAggregationFunction

All content copied from https://docs.aws.amazon.com/.
