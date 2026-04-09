# Supported data types for Iceberg tables in Athena

Athena can query Iceberg tables that contain the following data types:

```nohighlight

binary
boolean
date
decimal
double
float
int
list
long
map
string
struct
timestamp without time zone
```

For more information about Iceberg table types, see the [schemas page for\
Iceberg](https://iceberg.apache.org/docs/latest/schemas) in the Apache documentation.

The following table shows the relationship between Athena data types and Iceberg table
data types.

Iceberg typeAthena typeNotes`boolean``boolean`-`tinyint`Not supported for Iceberg tables in Athena.-`smallint`Not supported for Iceberg tables in Athena.`int``int`In Athena DML statements, this type is `INTEGER`.`long``bigint``double``double``float``float``decimal(P, S)``decimal(P, S)``P` is precision, `S` is scale.-`char`Not supported for Iceberg tables in Athena.`string``string`In Athena DML statements, this type is `VARCHAR`.`binary``binary``date``date``time`-Only Iceberg timestamp (without time zone) is supported
for Athena Iceberg DDL statements like `CREATE TABLE`, but all
timestamp types can be queried through Athena.`timestamp``timestamp``timestamptz``timestamptz``list<E>``array``map<K,V>``map``struct<...>``struct``fixed(L)`-The `fixed(L)` type is not currently supported in
Athena.

For more information about data types in Athena, see [Data types in Amazon Athena](data-types.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Query AWS Glue Data Catalog materialized views

Additional resources

All content copied from https://docs.aws.amazon.com/.
