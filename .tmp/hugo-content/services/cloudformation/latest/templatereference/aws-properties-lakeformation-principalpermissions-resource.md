This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::LakeFormation::PrincipalPermissions Resource

A structure for the resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Catalog" : Json,
  "Database" : DatabaseResource,
  "DataCellsFilter" : DataCellsFilterResource,
  "DataLocation" : DataLocationResource,
  "LFTag" : LFTagKeyResource,
  "LFTagPolicy" : LFTagPolicyResource,
  "Table" : TableResource,
  "TableWithColumns" : TableWithColumnsResource
}

```

### YAML

```yaml

  Catalog: Json
  Database:
    DatabaseResource
  DataCellsFilter:
    DataCellsFilterResource
  DataLocation:
    DataLocationResource
  LFTag:
    LFTagKeyResource
  LFTagPolicy:
    LFTagPolicyResource
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

_Type_: [DatabaseResource](aws-properties-lakeformation-principalpermissions-databaseresource.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DataCellsFilter`

A data cell filter.

_Required_: No

_Type_: [DataCellsFilterResource](aws-properties-lakeformation-principalpermissions-datacellsfilterresource.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DataLocation`

The location of an Amazon S3 path where permissions are granted or revoked.

_Required_: No

_Type_: [DataLocationResource](aws-properties-lakeformation-principalpermissions-datalocationresource.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LFTag`

The LF-tag key and values attached to a resource.

_Required_: No

_Type_: [LFTagKeyResource](aws-properties-lakeformation-principalpermissions-lftagkeyresource.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LFTagPolicy`

A list of LF-tag conditions that define a resource's LF-tag policy.

_Required_: No

_Type_: [LFTagPolicyResource](aws-properties-lakeformation-principalpermissions-lftagpolicyresource.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Table`

The table for the resource. A table is a metadata definition that represents your data. You can Grant and Revoke table privileges to a principal.

_Required_: No

_Type_: [TableResource](aws-properties-lakeformation-principalpermissions-tableresource.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TableWithColumns`

The table with columns for the resource. A principal with permissions to this resource can select metadata from the columns of a table in the Data Catalog and the underlying data in Amazon S3.

_Required_: No

_Type_: [TableWithColumnsResource](aws-properties-lakeformation-principalpermissions-tablewithcolumnsresource.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LFTagPolicyResource

TableResource

All content copied from https://docs.aws.amazon.com/.
