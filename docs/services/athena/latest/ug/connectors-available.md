# Available data source connectors

This section lists prebuilt Athena data source connectors that you can use to query a
variety of data sources external to Amazon S3. To use a connector in your Athena queries,
configure it and deploy it to your account.

## Considerations and limitations

- Some prebuilt connectors require that you create a VPC and a security group
before you can use the connector. For information about creating VPCs, see [Create a VPC for a data source connector or AWS Glue connection](https://docs.aws.amazon.com/athena/latest/ug/athena-connectors-vpc-creation.html).

- To use the Athena Federated Query feature with AWS Secrets Manager, you must configure an Amazon VPC
private endpoint for Secrets Manager. For more information, see [Create a Secrets Manager VPC private endpoint](https://docs.aws.amazon.com/secretsmanager/latest/userguide/vpc-endpoint-overview.html#vpc-endpoint-create) in the _AWS Secrets Manager User Guide_.

- For connectors that do not support predicate pushdown, queries that include a
predicate take longer to execute. For small datasets, very little data is
scanned, and queries take an average of about 2 minutes. However, for large
datasets, many queries can time out.

- Some federated data sources use terminology to refer data objects that is
different from Athena. For more information, see [Understand federated table name qualifiers](https://docs.aws.amazon.com/athena/latest/ug/tables-qualifiers.html).

- We update our connectors periodically based on upgrades from the database or
data source provider. We do not support data sources that are at end-of-life for
support.

- For connectors that do not support pagination when you list tables, the web
service can time out if your database has many tables and metadata. The
following connectors provide pagination support for listing tables:

- DocumentDB

- DynamoDB

- MySQL

- OpenSearch

- Oracle

- PostgreSQL

- Redshift

- SQL Server

## Case resolver modes in Federation SDK

The Federation SDK supports the following standardized case resolver modes for schema and table names:

- `NONE` – Does not change case of the given schema and table names.

- `LOWER` – Lower case all given schema and table names.

- `UPPER` – Upper case all given schema and table names.

- `ANNOTATION` – This mode is maintained for backward compatibility only and is supported exclusively by existing Snowflake and SAP HANA connectors.

- `CASE_INSENSITIVE_SEARCH` – Perform case insensitive searches against schema and tables names.

## Connector support for case resolver modes

### Basic mode support

All JDBC connectors support the following basic modes:

- `NONE`

- `LOWER`

- `UPPER`

### Annotation mode support

Only the following connectors support the `ANNOTATION` mode:

- Snowflake

- SAP HANA

###### Note

It is recommended to use CASE\_INSENSITIVE\_SEARCH instead of ANNOTATION.

### Case-insensitive search support

The following connectors support `CASE_INSENSITIVE_SEARCH`:

- DataLake Gen2

- Snowflake

- Oracle

- Synapse

- MySQL

- PostgreSQL

- Redshift

- ClickHouse

- SQL Server

- DB2

## Case resolver limitations

Be aware of the following limitations when using case resolver modes:

- When using `LOWER` mode, your schema name and all tables within the schema must be in lowercase.

- When using `UPPER` mode, your schema name and all tables within the schema must be in uppercase.

- When using `CASE_INSENSITIVE_SEARCH`:

- Schema names must be unique

- Table names within a schema must be unique (for example, you cannot have both "Apple" and "APPLE")

- Glue integration limitations:

- Glue only supports lowercase names

- Only `NONE` or `LOWER` modes will work when registering your Lambda function with GlueDataCatalog/LakeFormation

## Additional information

- For information about deploying an Athena data source connector, see [Use Amazon Athena Federated Query](federated-queries.md).

- For information about queries that use Athena data source connectors, see [Run federated queries](https://docs.aws.amazon.com/athena/latest/ug/running-federated-queries.html).

###### Athena data source connectors

- [Azure Data Lake Storage](https://docs.aws.amazon.com/athena/latest/ug/connectors-adls-gen2.html)

- [Azure Synapse](https://docs.aws.amazon.com/athena/latest/ug/connectors-azure-synapse.html)

- [Cloudera Hive](https://docs.aws.amazon.com/athena/latest/ug/connectors-cloudera-hive.html)

- [Cloudera Impala](https://docs.aws.amazon.com/athena/latest/ug/connectors-cloudera-impala.html)

- [CloudWatch](https://docs.aws.amazon.com/athena/latest/ug/connectors-cloudwatch.html)

- [CloudWatch metrics](https://docs.aws.amazon.com/athena/latest/ug/connectors-cwmetrics.html)

- [CMDB](https://docs.aws.amazon.com/athena/latest/ug/connectors-cmdb.html)

- [Db2](https://docs.aws.amazon.com/athena/latest/ug/connectors-ibm-db2.html)

- [Db2 iSeries](https://docs.aws.amazon.com/athena/latest/ug/connectors-ibm-db2-as400.html)

- [DocumentDB](https://docs.aws.amazon.com/athena/latest/ug/connectors-docdb.html)

- [DynamoDB](https://docs.aws.amazon.com/athena/latest/ug/connectors-dynamodb.html)

- [Google BigQuery](https://docs.aws.amazon.com/athena/latest/ug/connectors-bigquery.html)

- [Google Cloud Storage](https://docs.aws.amazon.com/athena/latest/ug/connectors-gcs.html)

- [HBase](https://docs.aws.amazon.com/athena/latest/ug/connectors-hbase.html)

- [Hortonworks](https://docs.aws.amazon.com/athena/latest/ug/connectors-hortonworks.html)

- [Kafka](https://docs.aws.amazon.com/athena/latest/ug/connectors-kafka.html)

- [MSK](https://docs.aws.amazon.com/athena/latest/ug/connectors-msk.html)

- [MySQL](https://docs.aws.amazon.com/athena/latest/ug/connectors-mysql.html)

- [Neptune](https://docs.aws.amazon.com/athena/latest/ug/connectors-neptune.html)

- [OpenSearch](https://docs.aws.amazon.com/athena/latest/ug/connectors-opensearch.html)

- [Oracle](https://docs.aws.amazon.com/athena/latest/ug/connectors-oracle.html)

- [PostgreSQL](https://docs.aws.amazon.com/athena/latest/ug/connectors-postgresql.html)

- [Redis OSS](https://docs.aws.amazon.com/athena/latest/ug/connectors-redis.html)

- [Redshift](https://docs.aws.amazon.com/athena/latest/ug/connectors-redshift.html)

- [SAP HANA](https://docs.aws.amazon.com/athena/latest/ug/connectors-sap-hana.html)

- [Snowflake](https://docs.aws.amazon.com/athena/latest/ug/connectors-snowflake.html)

- [SQL Server](https://docs.aws.amazon.com/athena/latest/ug/connectors-microsoft-sql-server.html)

- [Teradata](https://docs.aws.amazon.com/athena/latest/ug/connectors-teradata.html)

- [Timestream](https://docs.aws.amazon.com/athena/latest/ug/connectors-timestream.html)

- [TPC-DS](https://docs.aws.amazon.com/athena/latest/ug/connectors-tpcds.html)

- [Vertica](https://docs.aws.amazon.com/athena/latest/ug/connectors-vertica.html)

###### Note

The [AthenaJdbcConnector](https://serverlessrepo.aws.amazon.com/applications/us-east-1/292517598671/AthenaJdbcConnector) (latest version 2022.4.1) has been deprecated. Instead,
use a database-specific connector like those for [MySQL](https://docs.aws.amazon.com/athena/latest/ug/connectors-mysql.html), [Redshift](https://docs.aws.amazon.com/athena/latest/ug/connectors-redshift.html), or [PostgreSQL](https://docs.aws.amazon.com/athena/latest/ug/connectors-postgresql.html).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Use federated queries

Azure Data Lake Storage
