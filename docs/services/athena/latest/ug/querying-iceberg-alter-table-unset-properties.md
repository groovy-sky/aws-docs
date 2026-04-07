# ALTER TABLE UNSET TBLPROPERTIES

Drops existing properties from an Iceberg table.

## Synopsis

```sql

ALTER TABLE [db_name.]table_name UNSET TBLPROPERTIES ('property_name' [ , ... ])
```

## Example

```sql

ALTER TABLE iceberg_table UNSET TBLPROPERTIES ('write_compression')
```

The following example removes the `write.data.path` property from
an Iceberg table.

```sql

ALTER TABLE iceberg_table UNSET TBLPROPERTIES ('write.data.path')
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ALTER TABLE SET TBLPROPERTIES

DESCRIBE
