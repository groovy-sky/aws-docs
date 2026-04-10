This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet DataSetStringListFilterCondition

A filter condition that includes or excludes string values from a specified list.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Operator" : String,
  "Values" : DataSetStringListFilterValue
}

```

### YAML

```yaml

  Operator: String
  Values:
    DataSetStringListFilterValue

```

## Properties

`Operator`

The list operator to use, either `INCLUDE` to match values in the list or `EXCLUDE` to
filter out values in the list.

_Required_: Yes

_Type_: String

_Allowed values_: `INCLUDE | EXCLUDE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Values`

The list of string values to include or exclude in the filter.

_Required_: No

_Type_: [DataSetStringListFilterValue](aws-properties-quicksight-dataset-datasetstringlistfiltervalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataSetStringFilterValue

DataSetStringListFilterValue

All content copied from https://docs.aws.amazon.com/.
