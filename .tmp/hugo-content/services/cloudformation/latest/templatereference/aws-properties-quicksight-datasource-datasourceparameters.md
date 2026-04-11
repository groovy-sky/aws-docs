This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSource DataSourceParameters

The parameters that Quick Sight uses to connect to your underlying data source.
This is a variant type structure. For this structure to be valid, only one of the
attributes can be non-null.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AmazonElasticsearchParameters" : AmazonElasticsearchParameters,
  "AmazonOpenSearchParameters" : AmazonOpenSearchParameters,
  "AthenaParameters" : AthenaParameters,
  "AuroraParameters" : AuroraParameters,
  "AuroraPostgreSqlParameters" : AuroraPostgreSqlParameters,
  "DatabricksParameters" : DatabricksParameters,
  "MariaDbParameters" : MariaDbParameters,
  "MySqlParameters" : MySqlParameters,
  "OracleParameters" : OracleParameters,
  "PostgreSqlParameters" : PostgreSqlParameters,
  "PrestoParameters" : PrestoParameters,
  "RdsParameters" : RdsParameters,
  "RedshiftParameters" : RedshiftParameters,
  "S3Parameters" : S3Parameters,
  "S3TablesParameters" : S3TablesParameters,
  "SnowflakeParameters" : SnowflakeParameters,
  "SparkParameters" : SparkParameters,
  "SqlServerParameters" : SqlServerParameters,
  "StarburstParameters" : StarburstParameters,
  "TeradataParameters" : TeradataParameters,
  "TrinoParameters" : TrinoParameters
}

```

### YAML

```yaml

  AmazonElasticsearchParameters:
    AmazonElasticsearchParameters
  AmazonOpenSearchParameters:
    AmazonOpenSearchParameters
  AthenaParameters:
    AthenaParameters
  AuroraParameters:
    AuroraParameters
  AuroraPostgreSqlParameters:
    AuroraPostgreSqlParameters
  DatabricksParameters:
    DatabricksParameters
  MariaDbParameters:
    MariaDbParameters
  MySqlParameters:
    MySqlParameters
  OracleParameters:
    OracleParameters
  PostgreSqlParameters:
    PostgreSqlParameters
  PrestoParameters:
    PrestoParameters
  RdsParameters:
    RdsParameters
  RedshiftParameters:
    RedshiftParameters
  S3Parameters:
    S3Parameters
  S3TablesParameters:
    S3TablesParameters
  SnowflakeParameters:
    SnowflakeParameters
  SparkParameters:
    SparkParameters
  SqlServerParameters:
    SqlServerParameters
  StarburstParameters:
    StarburstParameters
  TeradataParameters:
    TeradataParameters
  TrinoParameters:
    TrinoParameters

```

## Properties

`AmazonElasticsearchParameters`

The parameters for OpenSearch.

_Required_: No

_Type_: [AmazonElasticsearchParameters](aws-properties-quicksight-datasource-amazonelasticsearchparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AmazonOpenSearchParameters`

The parameters for OpenSearch.

_Required_: No

_Type_: [AmazonOpenSearchParameters](aws-properties-quicksight-datasource-amazonopensearchparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AthenaParameters`

The parameters for Amazon Athena.

_Required_: No

_Type_: [AthenaParameters](aws-properties-quicksight-datasource-athenaparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuroraParameters`

The parameters for Amazon Aurora MySQL.

_Required_: No

_Type_: [AuroraParameters](aws-properties-quicksight-datasource-auroraparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuroraPostgreSqlParameters`

The parameters for Amazon Aurora.

_Required_: No

_Type_: [AuroraPostgreSqlParameters](aws-properties-quicksight-datasource-aurorapostgresqlparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatabricksParameters`

The required parameters that are needed to connect to a Databricks data source.

_Required_: No

_Type_: [DatabricksParameters](aws-properties-quicksight-datasource-databricksparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MariaDbParameters`

The parameters for MariaDB.

_Required_: No

_Type_: [MariaDbParameters](aws-properties-quicksight-datasource-mariadbparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MySqlParameters`

The parameters for MySQL.

_Required_: No

_Type_: [MySqlParameters](aws-properties-quicksight-datasource-mysqlparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OracleParameters`

Oracle parameters.

_Required_: No

_Type_: [OracleParameters](aws-properties-quicksight-datasource-oracleparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PostgreSqlParameters`

The parameters for PostgreSQL.

_Required_: No

_Type_: [PostgreSqlParameters](aws-properties-quicksight-datasource-postgresqlparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrestoParameters`

The parameters for Presto.

_Required_: No

_Type_: [PrestoParameters](aws-properties-quicksight-datasource-prestoparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RdsParameters`

The parameters for Amazon RDS.

_Required_: No

_Type_: [RdsParameters](aws-properties-quicksight-datasource-rdsparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RedshiftParameters`

The parameters for Amazon Redshift.

_Required_: No

_Type_: [RedshiftParameters](aws-properties-quicksight-datasource-redshiftparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3Parameters`

The parameters for S3.

_Required_: No

_Type_: [S3Parameters](aws-properties-quicksight-datasource-s3parameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3TablesParameters`

Property description not available.

_Required_: No

_Type_: [S3TablesParameters](aws-properties-quicksight-datasource-s3tablesparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SnowflakeParameters`

The parameters for Snowflake.

_Required_: No

_Type_: [SnowflakeParameters](aws-properties-quicksight-datasource-snowflakeparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SparkParameters`

The parameters for Spark.

_Required_: No

_Type_: [SparkParameters](aws-properties-quicksight-datasource-sparkparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SqlServerParameters`

The parameters for SQL Server.

_Required_: No

_Type_: [SqlServerParameters](aws-properties-quicksight-datasource-sqlserverparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StarburstParameters`

The parameters that are required to connect to a Starburst data source.

_Required_: No

_Type_: [StarburstParameters](aws-properties-quicksight-datasource-starburstparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TeradataParameters`

The parameters for Teradata.

_Required_: No

_Type_: [TeradataParameters](aws-properties-quicksight-datasource-teradataparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrinoParameters`

The parameters that are required to connect to a Trino data source.

_Required_: No

_Type_: [TrinoParameters](aws-properties-quicksight-datasource-trinoparameters.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataSourceErrorInfo

IdentityCenterConfiguration

All content copied from https://docs.aws.amazon.com/.
