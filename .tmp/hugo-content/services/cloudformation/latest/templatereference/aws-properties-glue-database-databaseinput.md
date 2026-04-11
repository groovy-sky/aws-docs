This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Database DatabaseInput

The structure used to create or update a database.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CreateTableDefaultPermissions" : [ PrincipalPrivileges, ... ],
  "Description" : String,
  "FederatedDatabase" : FederatedDatabase,
  "LocationUri" : String,
  "Name" : String,
  "Parameters" : Json,
  "TargetDatabase" : DatabaseIdentifier
}

```

### YAML

```yaml

  CreateTableDefaultPermissions:
    - PrincipalPrivileges
  Description: String
  FederatedDatabase:
    FederatedDatabase
  LocationUri: String
  Name: String
  Parameters: Json
  TargetDatabase:
    DatabaseIdentifier

```

## Properties

`CreateTableDefaultPermissions`

Creates a set of default permissions on the table for principals. Used by AWS Lake Formation. Not used in the normal course of AWS Glue operations.

_Required_: No

_Type_: Array of [PrincipalPrivileges](aws-properties-glue-database-principalprivileges.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the database.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `0`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FederatedDatabase`

A `FederatedDatabase` structure that references an entity outside the AWS Glue Data Catalog.

_Required_: No

_Type_: [FederatedDatabase](aws-properties-glue-database-federateddatabase.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LocationUri`

The location of the database (for example, an HDFS path).

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the database. For Hive compatibility, this is folded to lowercase when it is
stored.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Parameters`

These key-value pairs define parameters and properties
of the database.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetDatabase`

A `DatabaseIdentifier` structure that describes a target database for resource linking.

_Required_: No

_Type_: [DatabaseIdentifier](aws-properties-glue-database-databaseidentifier.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DatabaseIdentifier

DataLakePrincipal

All content copied from https://docs.aws.amazon.com/.
