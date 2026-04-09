# ALTER TABLE DROP PARTITION

Drops one or more specified partitions for the named table.

## Synopsis

```sql

ALTER TABLE table_name DROP [IF EXISTS] PARTITION (partition_spec) [, PARTITION (partition_spec)]
```

## Parameters

**\[IF EXISTS\]**

Suppresses the error message if the partition specified does not
exist.

**PARTITION (partition\_spec)**

Each `partition_spec` specifies a column name/value combination
in the form `partition_col_name = partition_col_value
                        [,...]`.

## Examples

```sql

ALTER TABLE orders
DROP PARTITION (dt = '2014-05-14', country = 'IN');
```

```sql

ALTER TABLE orders
DROP PARTITION (dt = '2014-05-14', country = 'IN'), PARTITION (dt = '2014-05-15', country = 'IN');
```

## Notes

The `ALTER TABLE DROP PARTITION` statement does not provide a single syntax
for dropping all partitions at once or support filtering criteria to specify a range of
partitions to drop.

As a workaround, you can use the AWS Glue API [GetPartitions](../../../glue/latest/dg/aws-glue-api-catalog-partitions.md#aws-glue-api-catalog-partitions-GetPartitions) and [BatchDeletePartition](../../../glue/latest/dg/aws-glue-api-catalog-partitions.md#aws-glue-api-catalog-partitions-BatchDeletePartition) actions in scripting. The `GetPartitions`
action supports complex filter expressions like those in a SQL `WHERE`
expression. After you use `GetPartitions` to create a filtered list of
partitions to delete, you can use the `BatchDeletePartition` action to delete
the partitions in batches of 25.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ALTER TABLE CHANGE COLUMN

ALTER TABLE RENAME PARTITION

All content copied from https://docs.aws.amazon.com/.
