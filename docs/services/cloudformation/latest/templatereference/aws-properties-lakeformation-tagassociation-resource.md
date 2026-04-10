This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::LakeFormation::TagAssociation Resource

A structure for the resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Catalog" : Json,
  "Database" : DatabaseResource,
  "Table" : TableResource,
  "TableWithColumns" : TableWithColumnsResource
}

```

### YAML

```yaml

  Catalog: Json
  Database:
    DatabaseResource
  Table:
    TableResource
  TableWithColumns:
    TableWithColumnsResource

```

## Properties

`Catalog`

The identifier for the Data Catalog. By default, the account ID. The Data Catalog is the persistent metadata store. It contains database definitions, table definitions, and other control information to manage your AWS Lake Formation environment.

_Required_: No

_Type_: Json

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Database`

The database for the resource. Unique to the Data Catalog. A database is a set of associated table definitions organized into a logical group. You can Grant and Revoke database permissions to a principal.

_Required_: No

_Type_: [DatabaseResource](aws-properties-lakeformation-tagassociation-databaseresource.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Table`

The table for the resource. A table is a metadata definition that represents your data. You can Grant and Revoke table privileges to a principal.

_Required_: No

_Type_: [TableResource](aws-properties-lakeformation-tagassociation-tableresource.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TableWithColumns`

The table with columns for the resource. A principal with permissions to this resource can select metadata from the columns of a table in the Data Catalog and the underlying data in Amazon S3.

_Required_: No

_Type_: [TableWithColumnsResource](aws-properties-lakeformation-tagassociation-tablewithcolumnsresource.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LFTagPair

TableResource

All content copied from https://docs.aws.amazon.com/.
