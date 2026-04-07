This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cassandra::Table Column

The name and data type of an individual column in a table. In addition to the data type, you can also use the
following two keywords:

- `STATIC` if the table has a clustering column. Static columns store values that are
shared by all rows in the same partition.

- `FROZEN` for collection data types. In frozen collections the values of the collection
are serialized into a single immutable value, and Amazon Keyspaces treats them like a `BLOB`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ColumnName" : String,
  "ColumnType" : String
}

```

### YAML

```yaml

  ColumnName: String
  ColumnType: String

```

## Properties

`ColumnName`

The name of the column. For more information, see [Identifiers](https://docs.aws.amazon.com/keyspaces/latest/devguide/cql.elements.html#cql.elements.identifier) in the
_Amazon Keyspaces Developer Guide_.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9][a-zA-Z0-9_]{1,47}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ColumnType`

The data type of the column. For more information, see [Data types](https://docs.aws.amazon.com/keyspaces/latest/devguide/cql.elements.html#cql.data-types)
in the _Amazon Keyspaces Developer Guide_.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ClusteringKeyColumn

EncryptionSpecification
