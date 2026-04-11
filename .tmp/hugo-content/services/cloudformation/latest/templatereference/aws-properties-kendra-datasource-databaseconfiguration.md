This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::DataSource DatabaseConfiguration

Provides the configuration information to an [Amazon Kendra supported \
database](../../../kendra/latest/dg/data-source-database.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AclConfiguration" : AclConfiguration,
  "ColumnConfiguration" : ColumnConfiguration,
  "ConnectionConfiguration" : ConnectionConfiguration,
  "DatabaseEngineType" : String,
  "SqlConfiguration" : SqlConfiguration,
  "VpcConfiguration" : DataSourceVpcConfiguration
}

```

### YAML

```yaml

  AclConfiguration:
    AclConfiguration
  ColumnConfiguration:
    ColumnConfiguration
  ConnectionConfiguration:
    ConnectionConfiguration
  DatabaseEngineType: String
  SqlConfiguration:
    SqlConfiguration
  VpcConfiguration:
    DataSourceVpcConfiguration

```

## Properties

`AclConfiguration`

Information about the database column that provides information for user context
filtering.

_Required_: No

_Type_: [AclConfiguration](aws-properties-kendra-datasource-aclconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ColumnConfiguration`

Information about where the index should get the document information from the
database.

_Required_: Yes

_Type_: [ColumnConfiguration](aws-properties-kendra-datasource-columnconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConnectionConfiguration`

Configuration information that's required to connect to a database.

_Required_: Yes

_Type_: [ConnectionConfiguration](aws-properties-kendra-datasource-connectionconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatabaseEngineType`

The type of database engine that runs the database.

_Required_: Yes

_Type_: String

_Allowed values_: `RDS_AURORA_MYSQL | RDS_AURORA_POSTGRESQL | RDS_MYSQL | RDS_POSTGRESQL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SqlConfiguration`

Provides information about how Amazon Kendra uses quote marks around SQL identifiers
when querying a database data source.

_Required_: No

_Type_: [SqlConfiguration](aws-properties-kendra-datasource-sqlconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcConfiguration`

Provides information for connecting to an Amazon VPC.

_Required_: No

_Type_: [DataSourceVpcConfiguration](aws-properties-kendra-datasource-datasourcevpcconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomDocumentEnrichmentConfiguration

DataSourceConfiguration

All content copied from https://docs.aws.amazon.com/.
