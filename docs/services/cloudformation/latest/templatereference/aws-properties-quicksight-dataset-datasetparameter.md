This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet DatasetParameter

The parameter declarations of the dataset.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DateTimeDatasetParameter" : DateTimeDatasetParameter,
  "DecimalDatasetParameter" : DecimalDatasetParameter,
  "IntegerDatasetParameter" : IntegerDatasetParameter,
  "StringDatasetParameter" : StringDatasetParameter
}

```

### YAML

```yaml

  DateTimeDatasetParameter:
    DateTimeDatasetParameter
  DecimalDatasetParameter:
    DecimalDatasetParameter
  IntegerDatasetParameter:
    IntegerDatasetParameter
  StringDatasetParameter:
    StringDatasetParameter

```

## Properties

`DateTimeDatasetParameter`

A date time parameter that is created in the dataset.

_Required_: No

_Type_: [DateTimeDatasetParameter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-dataset-datetimedatasetparameter.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DecimalDatasetParameter`

A decimal parameter that is created in the dataset.

_Required_: No

_Type_: [DecimalDatasetParameter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-dataset-decimaldatasetparameter.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IntegerDatasetParameter`

An integer parameter that is created in the dataset.

_Required_: No

_Type_: [IntegerDatasetParameter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-dataset-integerdatasetparameter.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StringDatasetParameter`

A string parameter that is created in the dataset.

_Required_: No

_Type_: [StringDatasetParameter](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-dataset-stringdatasetparameter.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DataSetNumericRangeFilterCondition

DataSetRefreshProperties
