This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::LakeFormation::TagAssociation TableResource

A structure for the table object. A table is a metadata definition that represents your data. You can Grant and Revoke table privileges to a principal.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CatalogId" : String,
  "DatabaseName" : String,
  "Name" : String,
  "TableWildcard" : Json
}

```

### YAML

```yaml

  CatalogId: String
  DatabaseName: String
  Name: String
  TableWildcard: Json

```

## Properties

`CatalogId`

The identifier for the Data Catalog. By default, it is the account ID of the caller.

_Required_: Yes

_Type_: String

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DatabaseName`

The name of the database for the table. Unique to a Data Catalog. A database is a set of associated table definitions organized into a logical group. You can Grant and Revoke database privileges to a principal.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the table.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TableWildcard`

A wildcard object representing every table under a database.This is an object with no properties that effectively behaves as a true or false depending on whether not it is passed as a parameter.
The valid inputs for a property with this type in either yaml or json is null or {}.

At least one of `TableResource$Name` or `TableResource$TableWildcard` is required.

_Required_: No

_Type_: Json

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Resource

TableWithColumnsResource

All content copied from https://docs.aws.amazon.com/.
