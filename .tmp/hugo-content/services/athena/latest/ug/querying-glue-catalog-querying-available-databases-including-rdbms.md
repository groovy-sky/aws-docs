# List databases and searching a specified database

The examples in this section show how to list the databases in metadata by schema
name.

###### Example– Listing databases

The following example query lists the databases from the
`information_schema.schemata` table.

```sql

SELECT schema_name
FROM   information_schema.schemata
LIMIT  10;
```

The following table shows sample results.

6alb-databas17alb\_original\_cust8alblogsdatabase9athena\_db\_test10athena\_ddl\_db

###### Example– Searching a specified database

In the following example query, `rdspostgresql` is a sample
database.

```sql

SELECT schema_name
FROM   information_schema.schemata
WHERE  schema_name = 'rdspostgresql'
```

The following table shows sample results.

schema\_name1rdspostgresql

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Query the AWS Glue Data Catalog

List tables

All content copied from https://docs.aws.amazon.com/.
