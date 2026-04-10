This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::LakeFormation::TagAssociation TableWithColumnsResource

A structure for a table with columns object. This object is only used when granting a SELECT permission.

This object must take a value for at least one of `ColumnsNames`, `ColumnsIndexes`, or `ColumnsWildcard`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CatalogId" : String,
  "ColumnNames" : [ String, ... ],
  "DatabaseName" : String,
  "Name" : String
}

```

### YAML

```yaml

  CatalogId: String
  ColumnNames:
    - String
  DatabaseName: String
  Name: String

```

## Properties

`CatalogId`

A wildcard object representing every table under a database.

At least one of TableResource$Name or TableResource$TableWildcard is required.

_Required_: Yes

_Type_: String

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ColumnNames`

The list of column names for the table. At least one of `ColumnNames` or `ColumnWildcard` is required.

_Required_: Yes

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DatabaseName`

The name of the database for the table with columns resource. Unique to the Data Catalog. A database is a set of associated table definitions organized into a logical group. You can Grant and Revoke database privileges to a principal.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the table resource. A table is a metadata definition that represents your data. You can Grant and Revoke table privileges to a principal.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TableResource

Next

All content copied from https://docs.aws.amazon.com/.
