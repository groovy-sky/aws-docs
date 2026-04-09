# ALTER TABLE RENAME

Renames a table.

Because the table metadata of an Iceberg table is stored in Amazon S3, you can update
the database and table name of an Iceberg managed table without affecting underlying
table information.

## Synopsis

```sql

ALTER TABLE [db_name.]table_name RENAME TO [new_db_name.]new_table_name
```

## Example

```sql

ALTER TABLE my_db.my_table RENAME TO my_db2.my_table2
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Manage Iceberg tables

ALTER TABLE SET TBLPROPERTIES

All content copied from https://docs.aws.amazon.com/.
