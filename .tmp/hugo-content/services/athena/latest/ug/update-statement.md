# UPDATE

Updates rows in an Apache Iceberg table. `UPDATE` is transactional and is
supported only for Apache Iceberg tables. The statement works only on existing rows and
cannot be used to insert or append a row.

## Synopsis

To update the rows in an Iceberg table, use the following syntax.

```sql

UPDATE [db_name.]table_name SET xx=yy[,...] [WHERE predicate]
```

For more information and examples, see the `UPDATE` section of [Update Iceberg table data](querying-iceberg-updating-iceberg-table-data.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DELETE

MERGE INTO

All content copied from https://docs.aws.amazon.com/.
