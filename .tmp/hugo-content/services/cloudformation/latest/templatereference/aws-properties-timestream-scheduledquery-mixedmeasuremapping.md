This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Timestream::ScheduledQuery MixedMeasureMapping

MixedMeasureMappings are mappings that can be used to ingest data into a mixture of narrow
and multi measures in the derived table.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MeasureName" : String,
  "MeasureValueType" : String,
  "MultiMeasureAttributeMappings" : [ MultiMeasureAttributeMapping, ... ],
  "SourceColumn" : String,
  "TargetMeasureName" : String
}

```

### YAML

```yaml

  MeasureName: String
  MeasureValueType: String
  MultiMeasureAttributeMappings:
    - MultiMeasureAttributeMapping
  SourceColumn: String
  TargetMeasureName: String

```

## Properties

`MeasureName`

Refers to the value of measure\_name in a result row. This field is required if
MeasureNameColumn is provided.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MeasureValueType`

Type of the value that is to be read from sourceColumn. If the mapping is for MULTI, use
MeasureValueType.MULTI.

_Required_: Yes

_Type_: String

_Allowed values_: `BIGINT | BOOLEAN | DOUBLE | VARCHAR | MULTI`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MultiMeasureAttributeMappings`

Required when measureValueType is MULTI. Attribute mappings for MULTI value
measures.

_Required_: No

_Type_: Array of [MultiMeasureAttributeMapping](aws-properties-timestream-scheduledquery-multimeasureattributemapping.md)

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceColumn`

This field refers to the source column from which measure-value is to be read for result
materialization.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TargetMeasureName`

Target measure name to be used. If not provided, the target measure name by default would
be measure-name if provided, or sourceColumn otherwise.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ErrorReportConfiguration

MultiMeasureAttributeMapping

All content copied from https://docs.aws.amazon.com/.
