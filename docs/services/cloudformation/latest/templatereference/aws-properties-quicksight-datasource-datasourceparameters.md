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

_Type_: [AmazonElasticsearchParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-datasource-amazonelasticsearchparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AmazonOpenSearchParameters`

The parameters for OpenSearch.

_Required_: No

_Type_: [AmazonOpenSearchParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-datasource-amazonopensearchparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AthenaParameters`

The parameters for Amazon Athena.

_Required_: No

_Type_: [AthenaParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-datasource-athenaparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuroraParameters`

The parameters for Amazon Aurora MySQL.

_Required_: No

_Type_: [AuroraParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-datasource-auroraparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuroraPostgreSqlParameters`

The parameters for Amazon Aurora.

_Required_: No

_Type_: [AuroraPostgreSqlParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-datasource-aurorapostgresqlparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatabricksParameters`

The required parameters that are needed to connect to a Databricks data source.

_Required_: No

_Type_: [DatabricksParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-datasource-databricksparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MariaDbParameters`

The parameters for MariaDB.

_Required_: No

_Type_: [MariaDbParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-datasource-mariadbparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MySqlParameters`

The parameters for MySQL.

_Required_: No

_Type_: [MySqlParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-datasource-mysqlparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OracleParameters`

Oracle parameters.

_Required_: No

_Type_: [OracleParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-datasource-oracleparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PostgreSqlParameters`

The parameters for PostgreSQL.

_Required_: No

_Type_: [PostgreSqlParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-datasource-postgresqlparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PrestoParameters`

The parameters for Presto.

_Required_: No

_Type_: [PrestoParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-datasource-prestoparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RdsParameters`

The parameters for Amazon RDS.

_Required_: No

_Type_: [RdsParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-datasource-rdsparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RedshiftParameters`

The parameters for Amazon Redshift.

_Required_: No

_Type_: [RedshiftParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-datasource-redshiftparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3Parameters`

The parameters for S3.

_Required_: No

_Type_: [S3Parameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-datasource-s3parameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3TablesParameters`

Property description not available.

_Required_: No

_Type_: [S3TablesParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-datasource-s3tablesparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SnowflakeParameters`

The parameters for Snowflake.

_Required_: No

_Type_: [SnowflakeParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-datasource-snowflakeparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SparkParameters`

The parameters for Spark.

_Required_: No

_Type_: [SparkParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-datasource-sparkparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SqlServerParameters`

The parameters for SQL Server.

_Required_: No

_Type_: [SqlServerParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-datasource-sqlserverparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StarburstParameters`

The parameters that are required to connect to a Starburst data source.

_Required_: No

_Type_: [StarburstParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-datasource-starburstparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TeradataParameters`

The parameters for Teradata.

_Required_: No

_Type_: [TeradataParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-datasource-teradataparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrinoParameters`

The parameters that are required to connect to a Trino data source.

_Required_: No

_Type_: [TrinoParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-quicksight-datasource-trinoparameters.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DataSourceErrorInfo

IdentityCenterConfiguration
