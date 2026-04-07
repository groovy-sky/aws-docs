# SHOW TBLPROPERTIES

Shows one or more table properties of an Iceberg table. Only Athena-supported table
properties are shown.

## Synopsis

```sql

SHOW TBLPROPERTIES [db_name.]table_name [('property_name')]
```

## Example

```sql

SHOW TBLPROPERTIES iceberg_table
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SHOW CREATE TABLE

Evolve Iceberg table schema
