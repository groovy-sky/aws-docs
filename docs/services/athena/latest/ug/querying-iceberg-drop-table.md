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

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DESCRIBE

SHOW CREATE TABLE
