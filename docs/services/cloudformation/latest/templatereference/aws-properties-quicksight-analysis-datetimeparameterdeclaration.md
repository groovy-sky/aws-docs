This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis DateTimeParameterDeclaration

A parameter declaration for the `DateTime` data type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DefaultValues" : DateTimeDefaultValues,
  "MappedDataSetParameters" : [ MappedDataSetParameter, ... ],
  "Name" : String,
  "TimeGranularity" : String,
  "ValueWhenUnset" : DateTimeValueWhenUnsetConfiguration
}

```

### YAML

```yaml

  DefaultValues:
    DateTimeDefaultValues
  MappedDataSetParameters:
    - MappedDataSetParameter
  Name: String
  TimeGranularity: String
  ValueWhenUnset:
    DateTimeValueWhenUnsetConfiguration

```

## Properties

`DefaultValues`

The default values of a parameter. If the parameter is a single-value parameter, a maximum of one default value can be provided.

_Required_: No

_Type_: [DateTimeDefaultValues](aws-properties-quicksight-analysis-datetimedefaultvalues.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MappedDataSetParameters`

Property description not available.

_Required_: No

_Type_: Array of [MappedDataSetParameter](aws-properties-quicksight-analysis-mappeddatasetparameter.md)

_Minimum_: `0`

_Maximum_: `150`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the parameter that is being declared.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9]+$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeGranularity`

The level of time precision that is used to aggregate `DateTime` values.

_Required_: No

_Type_: String

_Allowed values_: `YEAR | QUARTER | MONTH | WEEK | DAY | HOUR | MINUTE | SECOND | MILLISECOND`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ValueWhenUnset`

The configuration that defines the default value of a `DateTime` parameter when a value has not been set.

_Required_: No

_Type_: [DateTimeValueWhenUnsetConfiguration](aws-properties-quicksight-analysis-datetimevaluewhenunsetconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DateTimeParameter

DateTimePickerControlDisplayOptions

All content copied from https://docs.aws.amazon.com/.
