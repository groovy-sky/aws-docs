This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Topic DataAggregation

The definition of a data aggregation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DatasetRowDateGranularity" : String,
  "DefaultDateColumnName" : String
}

```

### YAML

```yaml

  DatasetRowDateGranularity: String
  DefaultDateColumnName: String

```

## Properties

`DatasetRowDateGranularity`

The level of time precision that is used to aggregate `DateTime` values.

_Required_: No

_Type_: String

_Allowed values_: `SECOND | MINUTE | HOUR | DAY | WEEK | MONTH | QUARTER | YEAR`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefaultDateColumnName`

The column name for the default date.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomInstructions

DatasetMetadata

All content copied from https://docs.aws.amazon.com/.
