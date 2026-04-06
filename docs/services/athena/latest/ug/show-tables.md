# SHOW TABLES

Lists all the base tables and views in a database.

###### Note

The [StatementType](https://docs.aws.amazon.com/athena/latest/APIReference/API_QueryExecution.html#athena-Type-QueryExecution-StatementType) parameter for `SHOW TABLES` in [GetQueryExecution](https://docs.aws.amazon.com/athena/latest/APIReference/API_GetQueryExecution.html) API operation is categorized as `UTILITY`, not
`DDL`.

## Synopsis

```sql

SHOW TABLES [IN database_name] ['regular_expression']
```

## Parameters

**\[IN database\_name\]**

Specifies the `database_name` from which tables will be listed.
If omitted, the database from the current context is assumed.

###### Note

`SHOW TABLES` may fail if `database_name` uses
an [unsupported\
character](https://docs.aws.amazon.com/athena/latest/ug/tables-databases-columns-names.html) such as a hyphen. As a workaround, try enclosing
the database name in backticks.

**\['regular\_expression'\]**

Filters the list of tables to those that match the
`regular_expression` you specify. To indicate any character
in `AWSDataCatalog` tables, you can use the `*` or
`.*` wildcard expression. For Apache Hive databases, use the
`.*` wildcard expression. To indicate a choice between
characters, use the `|` character.

## Examples

###### Example– Show all of the tables in the database `sampledb`

```

SHOW TABLES IN sampledb
```

`Results`

```nohighlight

alb_logs
cloudfront_logs
elb_logs
flights_2016
flights_parquet
view_2016_flights_dfw
```

###### Example– Show the names of all tables in `sampledb` that include the word "flights"

```sql

SHOW TABLES IN sampledb '*flights*'
```

`Results`

```nohighlight

flights_2016
flights_parquet
view_2016_flights_dfw
```

###### Example– Show the names of all tables in `sampledb` that end in the word "logs"

```sql

SHOW TABLES IN sampledb '*logs'
```

`Results`

```nohighlight

alb_logs
cloudfront_logs
elb_logs
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SHOW PARTITIONS

SHOW TBLPROPERTIES
