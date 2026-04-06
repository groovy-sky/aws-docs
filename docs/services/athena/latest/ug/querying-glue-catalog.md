# Query the AWS Glue Data Catalog

Because AWS Glue Data Catalog is used by many AWS services as their central metadata repository,
you might want to query Data Catalog metadata. To do so, you can use SQL queries in Athena. You
can use Athena to query AWS Glue catalog metadata like databases, tables, partitions, and
columns.

To obtain AWS Glue Catalog metadata, you query the `information_schema` database
on the Athena backend. The example queries in this topic show how to use Athena to query AWS Glue
Catalog metadata for common use cases.

## Considerations and limitations

- Instead of querying the `information_schema` database, it is
possible to use individual Apache Hive [DDL commands](ddl-reference.md) to extract metadata
information for specific databases, tables, views, partitions, and columns from
Athena. However, the output is in a non-tabular format.

- Querying `information_schema` is most performant if you have a
small to moderate amount of AWS Glue metadata. If you have a large amount of
metadata, errors can occur.

- You cannot use `CREATE VIEW` to create a view on the
`information_schema` database.

###### Topics

- [List databases and searching a specified database](https://docs.aws.amazon.com/athena/latest/ug/querying-glue-catalog-querying-available-databases-including-rdbms.html)

- [List tables in a specified database and searching for a table by name](https://docs.aws.amazon.com/athena/latest/ug/querying-glue-catalog-listing-tables.html)

- [List partitions for a specific table](https://docs.aws.amazon.com/athena/latest/ug/querying-glue-catalog-listing-partitions.html)

- [List or search columns for a specified table or view](https://docs.aws.amazon.com/athena/latest/ug/querying-glue-catalog-listing-columns.html)

- [List the columns that specific tables have in common](https://docs.aws.amazon.com/athena/latest/ug/querying-glue-catalog-listing-columns-in-common.html)

- [List all columns for all tables](https://docs.aws.amazon.com/athena/latest/ug/querying-glue-catalog-listing-all-columns-for-all-tables.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Query across regions

List databases
