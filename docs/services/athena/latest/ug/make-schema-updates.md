# Make schema updates

This topic describes some of the changes that you can make to the schema in `CREATE
            TABLE` statements without actually altering your data. To update a schema, you can
in some cases use an `ALTER TABLE` command, but in other cases you do not
actually modify an existing table. Instead, you create a table with a new name that modifies
the schema that you used in your original `CREATE TABLE` statement.

Depending on how you expect your schemas to evolve, to continue using Athena queries,
choose a compatible data format.

Consider an application that reads orders information from an `orders` table
that exists in two formats: CSV and Parquet.

The following example creates a table in Parquet:

```sql

CREATE EXTERNAL TABLE orders_parquet (
   `orderkey` int,
   `orderstatus` string,
   `totalprice` double,
   `orderdate` string,
   `orderpriority` string,
   `clerk` string,
   `shippriority` int
) STORED AS PARQUET
LOCATION 's3://amzn-s3-demo-bucket/orders_ parquet/';

```

The following example creates the same table in CSV:

```sql

CREATE EXTERNAL TABLE orders_csv (
   `orderkey` int,
   `orderstatus` string,
   `totalprice` double,
   `orderdate` string,
   `orderpriority` string,
   `clerk` string,
   `shippriority` int
)
ROW FORMAT DELIMITED FIELDS TERMINATED BY ','
LOCATION 's3://amzn-s3-demo-bucket/orders_csv/';

```

The following topics show how updates to these tables affect Athena queries.

###### Topics

- [Add columns at the beginning or in the middle of the table](https://docs.aws.amazon.com/athena/latest/ug/updates-add-columns-beginning-middle-of-table.html)

- [Add columns at the end of the table](https://docs.aws.amazon.com/athena/latest/ug/updates-add-columns-end-of-table.html)

- [Remove columns](https://docs.aws.amazon.com/athena/latest/ug/updates-removing-columns.html)

- [Rename columns](https://docs.aws.amazon.com/athena/latest/ug/updates-renaming-columns.html)

- [Reorder columns](https://docs.aws.amazon.com/athena/latest/ug/updates-reordering-columns.html)

- [Change a column data type](https://docs.aws.amazon.com/athena/latest/ug/updates-changing-column-type.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Handle schema updates

Add columns at the beginning or in the middle of the table
