# Create a table from query results (CTAS)

A `CREATE TABLE AS SELECT` (CTAS) query creates a new table in Athena from the
results of a `SELECT` statement from another query. Athena stores data files
created by the CTAS statement in a specified location in Amazon S3. For syntax, see [CREATE TABLE AS](create-table-as.md).

`CREATE TABLE AS` combines a `CREATE TABLE` DDL statement with a
`SELECT` DML statement and therefore technically contains both DDL and DML.
However, note that for Service Quotas purposes, CTAS queries in Athena are treated as DML.
For information about Service Quotas in Athena, see [Service Quotas](service-limits.md).

Use CTAS queries to:

- Create tables from query results in one step, without repeatedly querying raw data
sets. This makes it easier to work with raw data sets.

- Transform query results and migrate tables into other table formats such as Apache
Iceberg. This improves query performance and reduces query costs in Athena. For
information, see [Create Iceberg tables](querying-iceberg-creating-tables.md).

- Transform query results into storage formats such as Parquet and ORC. This
improves query performance and reduces query costs in Athena. For information, see
[Use columnar storage formats](https://docs.aws.amazon.com/athena/latest/ug/columnar-storage.html).

- Create copies of existing tables that contain only the data you need.

###### Topics

- [Considerations and limitations for CTAS queries](https://docs.aws.amazon.com/athena/latest/ug/ctas-considerations-limitations.html)

- [Create CTAS queries](https://docs.aws.amazon.com/athena/latest/ug/ctas-console.html)

- [CTAS examples](https://docs.aws.amazon.com/athena/latest/ug/ctas-examples.html)

- [Use CTAS and INSERT INTO for ETL](https://docs.aws.amazon.com/athena/latest/ug/ctas-insert-into-etl.html)

- [Work around the 100 partition limit](ctas-insert-into.md)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Escape reserved keywords

Considerations and limitations for CTAS queries
