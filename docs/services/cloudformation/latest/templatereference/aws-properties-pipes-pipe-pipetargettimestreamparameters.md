This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pipes::Pipe PipeTargetTimestreamParameters

The parameters for using a Timestream for LiveAnalytics table as a
target.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DimensionMappings" : [ DimensionMapping, ... ],
  "EpochTimeUnit" : String,
  "MultiMeasureMappings" : [ MultiMeasureMapping, ... ],
  "SingleMeasureMappings" : [ SingleMeasureMapping, ... ],
  "TimeFieldType" : String,
  "TimestampFormat" : String,
  "TimeValue" : String,
  "VersionValue" : String
}

```

### YAML

```yaml

  DimensionMappings:
    - DimensionMapping
  EpochTimeUnit: String
  MultiMeasureMappings:
    - MultiMeasureMapping
  SingleMeasureMappings:
    - SingleMeasureMapping
  TimeFieldType: String
  TimestampFormat: String
  TimeValue: String
  VersionValue: String

```

## Properties

`DimensionMappings`

Map source data to dimensions in the target Timestream for LiveAnalytics
table.

For more information, see [Amazon Timestream for LiveAnalytics concepts](../../../timestream/latest/developerguide/concepts.md)

_Required_: Yes

_Type_: Array of [DimensionMapping](aws-properties-pipes-pipe-dimensionmapping.md)

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EpochTimeUnit`

The granularity of the time units used. Default is `MILLISECONDS`.

Required if `TimeFieldType` is specified as `EPOCH`.

_Required_: No

_Type_: String

_Allowed values_: `MILLISECONDS | SECONDS | MICROSECONDS | NANOSECONDS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MultiMeasureMappings`

Maps multiple measures from the source event to the same record in the specified Timestream for LiveAnalytics table.

_Required_: No

_Type_: Array of [MultiMeasureMapping](aws-properties-pipes-pipe-multimeasuremapping.md)

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SingleMeasureMappings`

Mappings of single source data fields to individual records in the specified Timestream for LiveAnalytics table.

_Required_: No

_Type_: Array of [SingleMeasureMapping](aws-properties-pipes-pipe-singlemeasuremapping.md)

_Minimum_: `0`

_Maximum_: `8192`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeFieldType`

The type of time value used.

The default is `EPOCH`.

_Required_: No

_Type_: String

_Allowed values_: `EPOCH | TIMESTAMP_FORMAT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimestampFormat`

How to format the timestamps. For example,
`yyyy-MM-dd'T'HH:mm:ss'Z'`.

Required if `TimeFieldType` is specified as
`TIMESTAMP_FORMAT`.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeValue`

Dynamic path to the source data field that represents the time value for your data.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VersionValue`

64 bit version value or source data field that represents the version value for your data.

Write requests with a higher version number will update the existing measure values of the record and version.
In cases where the measure value is the same, the version will still be updated.

Default value is 1.

Timestream for LiveAnalytics does not support updating partial measure values in a record.

Write requests for duplicate data with a
higher version number will update the existing measure value and version. In cases where
the measure value is the same, `Version` will still be updated. Default value is
`1`.

###### Note

`Version` must be `1` or greater, or you will receive a
`ValidationException` error.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PipeTargetStateMachineParameters

PlacementConstraint

All content copied from https://docs.aws.amazon.com/.
