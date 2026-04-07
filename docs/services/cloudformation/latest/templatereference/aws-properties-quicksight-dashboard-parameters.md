This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard Parameters

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

_Type_: Array of [DateTimeParameter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-dashboard-datetimeparameter.html)

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DecimalParameters`

The parameters that have a data type of decimal.

_Required_: No

_Type_: Array of [DecimalParameter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-dashboard-decimalparameter.html)

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IntegerParameters`

The parameters that have a data type of integer.

_Required_: No

_Type_: Array of [IntegerParameter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-dashboard-integerparameter.html)

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StringParameters`

The parameters that have a data type of string.

_Required_: No

_Type_: Array of [StringParameter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-dashboard-stringparameter.html)

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ParameterListControl

ParameterSelectableValues
