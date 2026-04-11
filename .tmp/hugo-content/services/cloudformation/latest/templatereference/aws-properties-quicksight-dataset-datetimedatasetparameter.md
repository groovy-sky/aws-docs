This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet DateTimeDatasetParameter

A date time parameter that is created in the dataset.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DefaultValues" : DateTimeDatasetParameterDefaultValues,
  "Id" : String,
  "Name" : String,
  "TimeGranularity" : String,
  "ValueType" : String
}

```

### YAML

```yaml

  DefaultValues:
    DateTimeDatasetParameterDefaultValues
  Id: String
  Name: String
  TimeGranularity: String
  ValueType: String

```

## Properties

`DefaultValues`

A list of default values for a given date time parameter. This structure only accepts static values.

_Required_: No

_Type_: [DateTimeDatasetParameterDefaultValues](aws-properties-quicksight-dataset-datetimedatasetparameterdefaultvalues.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Id`

An identifier for the parameter that is created in the dataset.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9-]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the date time parameter that is created in the dataset.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9]+$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeGranularity`

The time granularity of the date time parameter.

_Required_: No

_Type_: String

_Allowed values_: `YEAR | QUARTER | MONTH | WEEK | DAY | HOUR | MINUTE | SECOND | MILLISECOND`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ValueType`

The value type of the dataset parameter. Valid values are `single value` or `multi value`.

_Required_: Yes

_Type_: String

_Allowed values_: `MULTI_VALUED | SINGLE_VALUED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataSetUsageConfiguration

DateTimeDatasetParameterDefaultValues

All content copied from https://docs.aws.amazon.com/.
