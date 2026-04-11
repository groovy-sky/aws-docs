This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataBrew::Dataset DatasetParameter

Represents a dataset paramater that defines type and conditions for a parameter in the
Amazon S3 path of the dataset.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CreateColumn" : Boolean,
  "DatetimeOptions" : DatetimeOptions,
  "Filter" : FilterExpression,
  "Name" : String,
  "Type" : String
}

```

### YAML

```yaml

  CreateColumn: Boolean
  DatetimeOptions:
    DatetimeOptions
  Filter:
    FilterExpression
  Name: String
  Type: String

```

## Properties

`CreateColumn`

Optional boolean value that defines whether the captured value of this parameter
should be loaded as an additional column in the dataset.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatetimeOptions`

Additional parameter options such as a format and a timezone. Required for datetime
parameters.

_Required_: No

_Type_: [DatetimeOptions](aws-properties-databrew-dataset-datetimeoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Filter`

The optional filter expression structure to apply additional matching criteria to the
parameter.

_Required_: No

_Type_: [FilterExpression](aws-properties-databrew-dataset-filterexpression.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the parameter that is used in the dataset's Amazon S3
path.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of the dataset parameter, can be one of a 'String', 'Number' or
'Datetime'.

_Required_: Yes

_Type_: String

_Allowed values_: `String | Number | Datetime`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataCatalogInputDefinition

DatetimeOptions

All content copied from https://docs.aws.amazon.com/.
