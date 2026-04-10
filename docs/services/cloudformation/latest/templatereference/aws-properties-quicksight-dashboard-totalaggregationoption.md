This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard TotalAggregationOption

The total aggregation settings map of a field id.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FieldId" : String,
  "TotalAggregationFunction" : TotalAggregationFunction
}

```

### YAML

```yaml

  FieldId: String
  TotalAggregationFunction:
    TotalAggregationFunction

```

## Properties

`FieldId`

The field id that's associated with the total aggregation option.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TotalAggregationFunction`

The total aggregation function that you want to set for a specified field id.

_Required_: Yes

_Type_: [TotalAggregationFunction](aws-properties-quicksight-dashboard-totalaggregationfunction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TotalAggregationFunction

TotalOptions

All content copied from https://docs.aws.amazon.com/.
