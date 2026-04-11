This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::LakeFormation::Permissions TableWithColumnsResource

A structure for a table with columns object. This object is only used when granting a SELECT permission.

This object must take a value for at least one of `ColumnsNames`, `ColumnsIndexes`, or `ColumnsWildcard`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CatalogId" : String,
  "ColumnNames" : [ String, ... ],
  "ColumnWildcard" : ColumnWildcard,
  "DatabaseName" : String,
  "Name" : String
}

```

### YAML

```yaml

  CatalogId: String
  ColumnNames:
    - String
  ColumnWildcard:
    ColumnWildcard
  DatabaseName: String
  Name: String

```

## Properties

`CatalogId`

The identifier for the Data Catalog. By default, it is the account ID of the caller.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ColumnNames`

The list of column names for the table. At least one of `ColumnNames` or `ColumnWildcard` is required.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ColumnWildcard`

A wildcard specified by a `ColumnWildcard` object. At least one of `ColumnNames` or `ColumnWildcard` is required.

_Required_: No

_Type_: [ColumnWildcard](aws-properties-lakeformation-permissions-columnwildcard.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DatabaseName`

The name of the database for the table with columns resource. Unique to the Data Catalog. A database is a set of associated table definitions organized into a logical group. You can Grant and Revoke database privileges to a principal.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the table resource. A table is a metadata definition that represents your data. You can Grant and Revoke table privileges to a principal.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Examples

### Input format for TableWithColumnsResource

#### JSON

```json

{
  "CatalogId" : “123456789012”,
  "ColumnNames" : [ “col1”, “col2” ],
  "DatabaseName" : “my_database”,
  "Name" : “my_table”
}

```

#### YAML

```yaml

CatalogId: “123456789012”
  ColumnNames:
    - “col1”
  DatabaseName: “my_database”
  Name: “my_table”

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TableResource

AWS::LakeFormation::PrincipalPermissions

All content copied from https://docs.aws.amazon.com/.
