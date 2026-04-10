This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Timestream::ScheduledQuery TimestreamConfiguration

Configuration to write data into Timestream database and table. This configuration allows
the user to map the query result select columns into the destination table columns.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DatabaseName" : String,
  "DimensionMappings" : [ DimensionMapping, ... ],
  "MeasureNameColumn" : String,
  "MixedMeasureMappings" : [ MixedMeasureMapping, ... ],
  "MultiMeasureMappings" : MultiMeasureMappings,
  "TableName" : String,
  "TimeColumn" : String
}

```

### YAML

```yaml

  DatabaseName: String
  DimensionMappings:
    - DimensionMapping
  MeasureNameColumn: String
  MixedMeasureMappings:
    - MixedMeasureMapping
  MultiMeasureMappings:
    MultiMeasureMappings
  TableName: String
  TimeColumn: String

```

## Properties

`DatabaseName`

Name of Timestream database to which the query result will be written.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DimensionMappings`

This is to allow mapping column(s) from the query result to the dimension in the
destination table.

_Required_: Yes

_Type_: Array of [DimensionMapping](aws-properties-timestream-scheduledquery-dimensionmapping.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MeasureNameColumn`

Name of the measure column. Also see `MultiMeasureMappings` and
`MixedMeasureMappings` for how measure name properties on those relate to
`MeasureNameColumn`.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MixedMeasureMappings`

Specifies how to map measures to multi-measure records.

_Required_: No

_Type_: Array of [MixedMeasureMapping](aws-properties-timestream-scheduledquery-mixedmeasuremapping.md)

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MultiMeasureMappings`

Multi-measure mappings.

_Required_: No

_Type_: [MultiMeasureMappings](aws-properties-timestream-scheduledquery-multimeasuremappings.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TableName`

Name of Timestream table that the query result will be written to. The table should be
within the same database that is provided in Timestream configuration.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TimeColumn`

Column from query result that should be used as the time column in destination table.
Column type for this should be TIMESTAMP.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TargetConfiguration

AWS::Timestream::Table

All content copied from https://docs.aws.amazon.com/.
