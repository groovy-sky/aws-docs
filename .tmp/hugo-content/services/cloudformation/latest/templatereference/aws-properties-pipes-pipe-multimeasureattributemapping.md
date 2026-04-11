This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pipes::Pipe MultiMeasureAttributeMapping

A mapping of a source event data field to a measure in a Timestream for
LiveAnalytics record.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MeasureValue" : String,
  "MeasureValueType" : String,
  "MultiMeasureAttributeName" : String
}

```

### YAML

```yaml

  MeasureValue: String
  MeasureValueType: String
  MultiMeasureAttributeName: String

```

## Properties

`MeasureValue`

Dynamic path to the measurement attribute in the source event.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MeasureValueType`

Data type of the measurement attribute in the source event.

_Required_: Yes

_Type_: String

_Allowed values_: `DOUBLE | BIGINT | VARCHAR | BOOLEAN | TIMESTAMP`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MultiMeasureAttributeName`

Target measure name to be used.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MSKAccessCredentials

MultiMeasureMapping

All content copied from https://docs.aws.amazon.com/.
