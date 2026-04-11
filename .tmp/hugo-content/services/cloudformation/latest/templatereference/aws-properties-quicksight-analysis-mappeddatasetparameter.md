This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis MappedDataSetParameter

A dataset parameter that is mapped to an analysis parameter.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataSetIdentifier" : String,
  "DataSetParameterName" : String
}

```

### YAML

```yaml

  DataSetIdentifier: String
  DataSetParameterName: String

```

## Properties

`DataSetIdentifier`

A unique name that identifies a dataset within the analysis or dashboard.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataSetParameterName`

The name of the dataset parameter.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9]+$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LongFormatText

MaximumLabelType

All content copied from https://docs.aws.amazon.com/.
