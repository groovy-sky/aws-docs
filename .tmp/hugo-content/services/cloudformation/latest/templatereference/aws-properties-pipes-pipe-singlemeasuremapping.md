This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pipes::Pipe SingleMeasureMapping

Maps a single source data field to a single record in the specified Timestream
for LiveAnalytics table.

For more information, see [Amazon Timestream for LiveAnalytics concepts](../../../timestream/latest/developerguide/concepts.md)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MeasureName" : String,
  "MeasureValue" : String,
  "MeasureValueType" : String
}

```

### YAML

```yaml

  MeasureName: String
  MeasureValue: String
  MeasureValueType: String

```

## Properties

`MeasureName`

Target measure name for the measurement attribute in the Timestream table.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MeasureValue`

Dynamic path of the source field to map to the measure in the record.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MeasureValueType`

Data type of the source field.

_Required_: Yes

_Type_: String

_Allowed values_: `DOUBLE | BIGINT | VARCHAR | BOOLEAN | TIMESTAMP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SelfManagedKafkaAccessConfigurationVpc

Tag

All content copied from https://docs.aws.amazon.com/.
