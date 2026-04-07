# Register federated catalogs in Athena

After you create connections to federated data sources, you can register them as
federated data catalogs for simplified data discovery and manage data access with
fine-grained permissions using Lake Formation. For more information, see [Register your connection as a Glue Data Catalog](register-connection-as-gdc.md).

## Considerations and limitations

- DDL operations are not supported on federated catalogs.

- You can register the following connectors to integrate with AWS Glue for fine-grained access
control:

- [Azure Data Lake\
Storage](https://docs.aws.amazon.com/athena/latest/ug/connectors-adls-gen2.html)

- [Azure\
Synapse](https://docs.aws.amazon.com/athena/latest/ug/connectors-azure-synapse.html)

- [BigQuery](https://docs.aws.amazon.com/athena/latest/ug/connectors-bigquery.html)

- [CMDB](https://docs.aws.amazon.com/athena/latest/ug/connectors-cmdb.html)

- [Db2](https://docs.aws.amazon.com/athena/latest/ug/connectors-ibm-db2.html)

- [Db2 iSeries](https://docs.aws.amazon.com/athena/latest/ug/connectors-ibm-db2-as400.html)

- [DocumentDB](https://docs.aws.amazon.com/athena/latest/ug/connectors-docdb.html)

- [DynamoDB](https://docs.aws.amazon.com/athena/latest/ug/connectors-dynamodb.html)

- [Google Cloud Storage](https://docs.aws.amazon.com/athena/latest/ug/connectors-gcs.html)

- [HBase](https://docs.aws.amazon.com/athena/latest/ug/connectors-hbase.html)

- [MySQL](https://docs.aws.amazon.com/athena/latest/ug/connectors-mysql.html)

- [OpenSearch](https://docs.aws.amazon.com/athena/latest/ug/connectors-opensearch.html)

- [Oracle](https://docs.aws.amazon.com/athena/latest/ug/connectors-oracle.html)

- [PostgreSQL](https://docs.aws.amazon.com/athena/latest/ug/connectors-postgresql.html)

- [Redshift](https://docs.aws.amazon.com/athena/latest/ug/connectors-redshift.html)

- [SAP HANA](https://docs.aws.amazon.com/athena/latest/ug/connectors-sap-hana.html)

- [Snowflake](https://docs.aws.amazon.com/athena/latest/ug/connectors-snowflake.html)

- [SQL\
Server](https://docs.aws.amazon.com/athena/latest/ug/connectors-microsoft-sql-server.html)

- [Timestream](https://docs.aws.amazon.com/athena/latest/ug/connectors-timestream.html)

- [TPC-DS](https://docs.aws.amazon.com/athena/latest/ug/connectors-tpcds.html)

- When you create a resource link for Glue connection federation, the name of [resource link](https://docs.aws.amazon.com/lake-formation/latest/dg/create-resource-link-database.html) must be same as the database name of the
producer.

- Currently, only lowercase table and column names are recognized even if
the data source is case insensitive.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Register Redshift data catalogs

Register S3 table bucket catalogs
