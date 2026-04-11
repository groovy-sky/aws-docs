This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::LakeFormation::Permissions Resource

A structure for the resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DatabaseResource" : DatabaseResource,
  "DataLocationResource" : DataLocationResource,
  "TableResource" : TableResource,
  "TableWithColumnsResource" : TableWithColumnsResource
}

```

### YAML

```yaml

  DatabaseResource:
    DatabaseResource
  DataLocationResource:
    DataLocationResource
  TableResource:
    TableResource
  TableWithColumnsResource:
    TableWithColumnsResource

```

## Properties

`DatabaseResource`

A structure for the database object.

_Required_: No

_Type_: [DatabaseResource](aws-properties-lakeformation-permissions-databaseresource.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DataLocationResource`

A structure for a data location object where permissions are granted or revoked.

_Required_: No

_Type_: [DataLocationResource](aws-properties-lakeformation-permissions-datalocationresource.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TableResource`

A structure for the table object. A table is a metadata definition that represents your data. You can Grant and Revoke table privileges to a principal.

_Required_: No

_Type_: [TableResource](aws-properties-lakeformation-permissions-tableresource.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TableWithColumnsResource`

A structure for a table with columns object. This object is only used when granting a SELECT permission.

_Required_: No

_Type_: [TableWithColumnsResource](aws-properties-lakeformation-permissions-tablewithcolumnsresource.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataLocationResource

TableResource

All content copied from https://docs.aws.amazon.com/.
