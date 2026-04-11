This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet DataSetStringComparisonFilterCondition

A filter condition that compares string values using operators like `EQUALS`, `CONTAINS`,
or `STARTS_WITH`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Operator" : String,
  "Value" : DataSetStringFilterValue
}

```

### YAML

```yaml

  Operator: String
  Value:
    DataSetStringFilterValue

```

## Properties

`Operator`

The comparison operator to use, such as `EQUALS`, `CONTAINS`, `STARTS_WITH`,
`ENDS_WITH`, or their negations.

_Required_: Yes

_Type_: String

_Allowed values_: `EQUALS | DOES_NOT_EQUAL | CONTAINS | DOES_NOT_CONTAIN | STARTS_WITH | ENDS_WITH`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The string value to compare against.

_Required_: No

_Type_: [DataSetStringFilterValue](aws-properties-quicksight-dataset-datasetstringfiltervalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataSetRefreshProperties

DataSetStringFilterCondition

All content copied from https://docs.aws.amazon.com/.
