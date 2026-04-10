This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cassandra::Table ClusteringKeyColumn

Defines an individual column within the clustering key.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Column" : Column,
  "OrderBy" : String
}

```

### YAML

```yaml

  Column:
    Column
  OrderBy: String

```

## Properties

`Column`

The name and data type of this clustering key column.

_Required_: Yes

_Type_: [Column](aws-properties-cassandra-table-column.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OrderBy`

The order in which this column's data is stored:

- `ASC` (default) - The column's data is stored in ascending
order.

- `DESC` \- The column's data is stored in descending order.

_Required_: No

_Type_: String

_Allowed values_: `ASC | DESC`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CdcSpecification

Column

All content copied from https://docs.aws.amazon.com/.
