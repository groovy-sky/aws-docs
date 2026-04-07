# List tables in a specified database and searching for a table by name

To list metadata for tables, you can query by table schema or by table name.

###### Example– Listing tables by schema

The following query lists tables that use the `rdspostgresql` table
schema.

```sql

SELECT table_schema,
       table_name,
       table_type
FROM   information_schema.tables
WHERE  table_schema = 'rdspostgresql'
```

The following table shows a sample result.

table\_schematable\_nametable\_type1rdspostgresqlrdspostgresqldb1\_public\_accountBASE TABLE

###### Example– Searching for a table by name

The following query obtains metadata information for the table
`athena1`.

```sql

SELECT table_schema,
       table_name,
       table_type
FROM   information_schema.tables
WHERE  table_name = 'athena1'
```

The following table shows a sample result.

table\_schematable\_nametable\_type1defaultathena1BASE TABLE

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

List databases

List partitions
