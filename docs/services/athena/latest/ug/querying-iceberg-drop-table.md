# DROP TABLE

Drops an Iceberg table.

###### Warning

Because Iceberg tables are considered managed tables in Athena, dropping an
Iceberg table also removes all the data in the table.

## Synopsis

```sql

DROP TABLE [IF EXISTS] [db_name.]table_name
```

## Example

```sql

DROP TABLE iceberg_table
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DESCRIBE

SHOW CREATE TABLE

All content copied from https://docs.aws.amazon.com/.
