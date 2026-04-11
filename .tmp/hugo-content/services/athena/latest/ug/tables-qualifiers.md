# Understand federated table name qualifiers

Athena uses the following terms to refer to hierarchies of data objects:

- Data source – a group of databases

- Database – a group of tables

- Table – data organized as a group of rows or
columns

Sometimes these objects are also referred to with alternate but equivalent names such as
the following:

- A data source is sometimes referred to as a _catalog_.

- A database is sometimes referred to as a _schema_.

## Terms in federated data sources

When you query federated data sources, note that the underlying data source might not
use the same terminology as Athena. Keep this distinction in mind when you write your
federated queries. The following sections describe how data object terms in Athena
correspond to those in federated data sources.

### Amazon Redshift

An Amazon Redshift _database_ is a group of Redshift
_schemas_ that contains a group of Redshift
_tables_.

AthenaRedshiftRedshift data sourceA Redshift connector Lambda function configured to point to a
Redshift `database`.`data_source.database.table``database.schema.table`

Example query

```sql

SELECT * FROM
Athena_Redshift_connector_data_source.Redshift_schema_name.Redshift_table_name
```

For more information about this connector, see [Amazon Athena Redshift connector](connectors-redshift.md).

### Cloudera Hive

An Cloudera Hive _server_ or _cluster_ is a group of Cloudera Hive _databases_ that contains a group of Cloudera Hive
_tables_.

AthenaHiveCloudera Hive data sourceCloudera Hive connector Lambda function configured to point to a
Cloudera Hive `server`.`data_source.database.table``server.database.table`

Example query

```sql

SELECT * FROM
Athena_Cloudera_Hive_connector_data_source.Cloudera_Hive_database_name.Cloudera_Hive_table_name
```

For more information about this connector, see [Amazon Athena Cloudera Hive connector](connectors-cloudera-hive.md).

### Cloudera Impala

An Impala _server_ or _cluster_ is a group of Impala _databases_ that contains a group of Impala _tables_.

AthenaImpalaImpala data sourceImpala connector Lambda function configured to point to an Impala
`server`.`data_source.database.table``server.database.table`

Example query

```sql

SELECT * FROM
Athena_Impala_connector_data_source.Impala_database_name.Impala_table_name
```

For more information about this connector, see [Amazon Athena Cloudera Impala connector](connectors-cloudera-impala.md).

### MySQL

A MySQL _server_ is a group of MySQL _databases_ that contains a group of MySQL _tables_.

AthenaMySQLMySQL data sourceMySQL connector Lambda function configured to point to a MySQL
`server`.`data_source.database.table``server.database.table`

Example query

```sql

SELECT * FROM
Athena_MySQL_connector_data source.MySQL_database_name.MySQL_table_name
```

For more information about this connector, see [Amazon Athena MySQL connector](connectors-mysql.md).

### Oracle

An Oracle _server_ (or _database_) is a group of Oracle _schemas_ that contains a group of Oracle _tables_.

AthenaOracleOracle data sourceOracle connector Lambda function configured to point to an Oracle
`server`.`data_source.database.table``server.schema.table`

Example query

```sql

SELECT * FROM
Athena_Oracle_connector_data_source.Oracle_schema_name.Oracle_table_name
```

For more information about this connector, see [Amazon Athena Oracle connector](connectors-oracle.md).

### Postgres

A Postgres _server_ (or _cluster_) is a group of Postgres _databases_. A Postgres _database_ is
a group of Postgres _schemas_ that contains a group
of Postgres _tables_.

AthenaPostgresPostgres data sourcePostgres connector Lambda function configured to point to a
Postgres `server` and `database`.`data_source.database.table``server.database.schema.table`

Example query

```sql

SELECT * FROM
Athena_Postgres_connector_data_source.Postgres_schema_name.Postgres_table_name
```

For more information about this connector, see [Amazon Athena PostgreSQL connector](connectors-postgresql.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Use passthrough queries

Develop a data source connector

All content copied from https://docs.aws.amazon.com/.
