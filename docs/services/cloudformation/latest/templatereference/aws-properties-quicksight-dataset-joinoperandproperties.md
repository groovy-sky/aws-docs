This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet JoinOperandProperties

Properties that control how columns are handled for a join operand, including column name overrides.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OutputColumnNameOverrides" : [ OutputColumnNameOverride, ... ]
}

```

### YAML

```yaml

  OutputColumnNameOverrides:
    - OutputColumnNameOverride

```

## Properties

`OutputColumnNameOverrides`

A list of column name overrides to apply to the join operand's output columns.

_Required_: Yes

_Type_: Array of [OutputColumnNameOverride](aws-properties-quicksight-dataset-outputcolumnnameoverride.md)

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

JoinKeyProperties

JoinOperation

All content copied from https://docs.aws.amazon.com/.
