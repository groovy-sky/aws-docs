This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis CalculatedField

The calculated field of an analysis.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataSetIdentifier" : String,
  "Expression" : String,
  "Name" : String
}

```

### YAML

```yaml

  DataSetIdentifier: String
  Expression: String
  Name: String

```

## Properties

`DataSetIdentifier`

The data set that is used in this calculated field.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Expression`

The expression of the calculated field.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `32000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the calculated field.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BoxPlotVisual

CalculatedMeasureField

All content copied from https://docs.aws.amazon.com/.
