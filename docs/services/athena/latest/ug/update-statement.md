# UPDATE

Updates rows in an Apache Iceberg table. `UPDATE` is transactional and is
supported only for Apache Iceberg tables. The statement works only on existing rows and
cannot be used to insert or append a row.

## Synopsis

To update the rows in an Iceberg table, use the following syntax.

```sql

UPDATE [db_name.]table_name SET xx=yy[,...] [WHERE predicate]
```

For more information and examples, see the `UPDATE` section of [Update Iceberg table data](https://docs.aws.amazon.com/athena/latest/ug/querying-iceberg-updating-iceberg-table-data.html).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DELETE

MERGE INTO
