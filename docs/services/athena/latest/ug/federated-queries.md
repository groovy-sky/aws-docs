---
title: "Use Amazon Athena Federated Query"
---

# Use Amazon Athena Federated Query

If you have data in sources other than Amazon S3, you can use Athena Federated Query to query the data in
place or build pipelines that extract data from multiple data sources and store the data in
Amazon S3. With Athena Federated Query, you can run SQL queries across data stored in relational,
non-relational, object, and custom data sources. For a full list of supported data sources,
see [Available data source connectors](connectors-available.md).

When you run a query against a data source, Athena invokes the connector to determine
which data to read, manages parallelism, and pushes down filter predicates. Connectors
can also restrict access to data based on the user who submits the query.

Athena uses _data source connectors_ to run federated queries on
underlying data. Athena supports two types of data source connectors with different
capabilities:

- AWS Glue Data Catalog federated connectors
– These connectors use an AWS Glue connection to connect to the data source. They can be used with fine-grained data governance control support through Lake Formation.
For more information, see [Federated catalog data connections](../../../lake-formation/latest/dg/federated-catalog-data-connection.md) in the _AWS Lake Formation Developer Guide_.

- Connectors associated with a Lambda can optionally be manually registered as an AWS Glue Data Catalog to be used with Lake Formation for fine-grained data governance

- Starting April 21, 2026, certain newly created connectors are automatically registered as Glue Data Catalogs and do not use a Lambda function in your AWS account

- Athena data catalog federated connectors
– These connectors are specific to Athena and cannot be registered as
federated catalogs with AWS Glue Data Catalog. They require a Lambda function in
your AWS account to query data. Custom connectors developed using the Athena Query Federation SDK
are Athena data catalog connectors. For more information, see
[Develop a data source connector using the Athena Query Federation SDK](../../../../reference/athena/latest/ug/connect-data-source-federation-sdk.md).

For a list of data sources compatible with each type, see [Connector type support by data source](#federated-queries-connector-support).

###### Note

Third party developers may have used the Athena Query Federation SDK to write data source connectors.
For support or licensing issues with these data source connectors, please work with your
connector provider. These connectors are not tested or supported by AWS.

## Considerations and limitations

- Views – You can create and query views
on federated data sources. Federated views are stored in AWS Glue, not the
underlying data source. For more information, see [Query federated views](running-federated-queries.md#running-federated-queries-federated-views).

- Delimited identifiers – Delimited
identifiers (also known as quoted identifiers) begin and end with double
quotation marks ("). Currently, delimited identifiers are not supported for
federated queries in Athena.

- Write operations – Write operations like
[INSERT INTO](insert-into.md) are not supported.
Attempting to do so may result in the error message **`This operation is**
**currently not supported for external catalogs`**.

- Pricing – For pricing information, see
[Amazon Athena\
pricing](https://aws.amazon.com/athena/pricing).

- JDBC driver – To use the JDBC driver
with federated queries or an [external Hive metastore](connect-to-data-source-hive.md),
include `MetadataRetrievalMethod=ProxyAPI` in your JDBC connection
string. For information about the JDBC driver, see [Connect to Amazon Athena with JDBC](connect-with-jdbc.md).

- Secrets Manager – To use the Athena Federated Query feature with
AWS Secrets Manager, you must configure an Amazon VPC private endpoint for Secrets Manager. For more
information, see [Create a Secrets Manager VPC private endpoint](../../../secretsmanager/latest/userguide/vpc-endpoint-overview.md#vpc-endpoint-create) in the _AWS Secrets Manager User Guide_.

- Passthrough queries – Passthrough queries are not supported after a data source is registered as an AWS Glue Data Catalog.

## Connector type support by data source

The following table shows the connector types that each data source supports.
Certain AWS Glue Data Catalog federated catalog connectors that you create on or after April 21, 2026,
do not require Lambda.

Data sourceAWS Glue Data Catalog federated connectorsAthena data catalog federated connectorsWithout LambdaWith Lambda[Amazon CloudWatch Logs](connectors-cloudwatch.md)YesYes[Amazon CloudWatch Metrics](connectors-cwmetrics.md)YesYes[Amazon DocumentDB](connectors-docdb.md)YesYesYes[Amazon DynamoDB](connectors-dynamodb.md)YesYesYes[Amazon MSK](connectors-msk.md)Yes[Amazon Neptune](connectors-neptune.md)Yes[Amazon OpenSearch](connectors-opensearch.md)YesYesYes[Amazon Redshift](connectors-redshift.md)YesYesYes[Amazon Timestream](connectors-timestream.md)YesYes[Azure Data Lake Storage](connectors-adls-gen2.md)YesYes[Azure Synapse](connectors-azure-synapse.md)YesYes[Cloudera Hive](connectors-cloudera-hive.md)YesYes[Cloudera Impala](connectors-cloudera-impala.md)YesYes[CMDB](connectors-cmdb.md)YesYes[Confluent](connectors-kafka.md)Yes[Custom](../../../../reference/athena/latest/ug/connect-data-source-federation-sdk.md)Yes[Db2](connectors-ibm-db2.md)YesYes[Db2 iSeries](connectors-ibm-db2-as400.md)YesYes[Google BigQuery](connectors-bigquery.md)YesYesYes[Google Cloud Storage](connectors-gcs.md)YesYes[HBase](connectors-hbase.md)YesYes[Hortonworks (Hive)](connectors-hortonworks.md)Yes[Kafka](connectors-kafka.md)Yes[MySQL](connectors-mysql.md)YesYesYes[Oracle](connectors-oracle.md)YesYesYes[PostgreSQL](connectors-postgresql.md)YesYesYes[Redis OSS](connectors-redis.md)Yes[SAP HANA](connectors-sap-hana.md)YesYesYes[Snowflake](connectors-snowflake.md)YesYesYes[SQL Server](connectors-microsoft-sql-server.md)YesYesYes[Teradata](connectors-teradata.md)YesYesYes[TPC-DS](connectors-tpcds.md)YesYes[Vertica](connectors-vertica.md)YesYes

## Videos

Watch the following videos to learn more about using Athena Federated Query.

###### Video: Analyze Results of Federated Query in Amazon Athena in Quick

The following video demonstrates how to analyze results of an Athena Federated Query
in Quick.

###### Video: Game Analytics Pipeline

The following video shows how to deploy a scalable serverless data pipeline to
ingest, store, and analyze telemetry data from games and services using Amazon Athena
federated queries.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Work with geospatial data

Available data source connectors

All content copied from https://docs.aws.amazon.com/.
