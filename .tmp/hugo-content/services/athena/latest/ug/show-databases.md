# SHOW DATABASES

Lists all databases defined in the metastore. You can use `DATABASES` or
`SCHEMAS`. They mean the same thing.

The programmatic equivalent of `SHOW DATABASES` is the [ListDatabases](../../../../reference/athena/latest/apireference/api-listdatabases.md) Athena API action. The equivalent method in AWS SDK for Python (Boto3) is [list\_databases](https://boto3.amazonaws.com/v1/documentation/api/latest/reference/services/athena/client/list_databases.html).

## Synopsis

```sql

SHOW {DATABASES | SCHEMAS} [LIKE 'regular_expression']
```

## Parameters

**\[LIKE ' `regular_expression`'\]**

Filters the list of databases to those that match the
`regular_expression` that you
specify. For wildcard character matching, you can use the combination
`.*`, which matches any character zero to unlimited
times.

## Examples

```sql

SHOW SCHEMAS;
```

```sql

SHOW DATABASES LIKE '.*analytics';
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SHOW CREATE VIEW

SHOW PARTITIONS

All content copied from https://docs.aws.amazon.com/.
