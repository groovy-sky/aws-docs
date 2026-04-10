This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis Parameters

A list of Quick Sight parameters and the list's override values.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DateTimeParameters" : [ DateTimeParameter, ... ],
  "DecimalParameters" : [ DecimalParameter, ... ],
  "IntegerParameters" : [ IntegerParameter, ... ],
  "StringParameters" : [ StringParameter, ... ]
}

```

### YAML

```yaml

  DateTimeParameters:
    - DateTimeParameter
  DecimalParameters:
    - DecimalParameter
  IntegerParameters:
    - IntegerParameter
  StringParameters:
    - StringParameter

```

## Properties

`DateTimeParameters`

The parameters that have a data type of date-time.

_Required_: No

_Type_: Array of [DateTimeParameter](aws-properties-quicksight-analysis-datetimeparameter.md)

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DecimalParameters`

The parameters that have a data type of decimal.

_Required_: No

_Type_: Array of [DecimalParameter](aws-properties-quicksight-analysis-decimalparameter.md)

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IntegerParameters`

The parameters that have a data type of integer.

_Required_: No

_Type_: Array of [IntegerParameter](aws-properties-quicksight-analysis-integerparameter.md)

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StringParameters`

The parameters that have a data type of string.

_Required_: No

_Type_: Array of [StringParameter](aws-properties-quicksight-analysis-stringparameter.md)

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ParameterListControl

ParameterSelectableValues

All content copied from https://docs.aws.amazon.com/.
