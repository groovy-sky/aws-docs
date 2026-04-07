# DESCRIBE

Describes table information.

## Synopsis

```sql

DESCRIBE [FORMATTED] [db_name.]table_name
```

When the `FORMATTED` option is specified, the output displays
additional information such as table location and properties.

## Example

```sql

DESCRIBE iceberg_table
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ALTER TABLE UNSET TBLPROPERTIES

DROP TABLE
