This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet DataSetDateComparisonFilterCondition

A filter condition that compares date values using operators like `BEFORE`, `AFTER`, or
their inclusive variants.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Operator" : String,
  "Value" : DataSetDateFilterValue
}

```

### YAML

```yaml

  Operator: String
  Value:
    DataSetDateFilterValue

```

## Properties

`Operator`

The comparison operator to use, such as `BEFORE`, `BEFORE_OR_EQUALS_TO`, `AFTER`,
or `AFTER_OR_EQUALS_TO`.

_Required_: Yes

_Type_: String

_Allowed values_: `BEFORE | BEFORE_OR_EQUALS_TO | AFTER | AFTER_OR_EQUALS_TO`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The date value to compare against.

_Required_: No

_Type_: [DataSetDateFilterValue](aws-properties-quicksight-dataset-datasetdatefiltervalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataSetColumnIdMapping

DataSetDateFilterCondition

All content copied from https://docs.aws.amazon.com/.
