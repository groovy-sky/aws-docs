This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet SaaSTable

A table from a Software-as-a-Service (SaaS) data source, including connection details and column definitions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataSourceArn" : String,
  "InputColumns" : [ InputColumn, ... ],
  "TablePath" : [ TablePathElement, ... ]
}

```

### YAML

```yaml

  DataSourceArn: String
  InputColumns:
    - InputColumn
  TablePath:
    - TablePathElement

```

## Properties

`DataSourceArn`

The Amazon Resource Name (ARN) of the SaaS data source.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputColumns`

The list of input columns available from the SaaS table.

_Required_: Yes

_Type_: Array of [InputColumn](aws-properties-quicksight-dataset-inputcolumn.md)

_Minimum_: `0`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TablePath`

The hierarchical path to the table within the SaaS data source.

_Required_: Yes

_Type_: Array of [TablePathElement](aws-properties-quicksight-dataset-tablepathelement.md)

_Minimum_: `1`

_Maximum_: `32`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3Source

SemanticModelConfiguration

All content copied from https://docs.aws.amazon.com/.
