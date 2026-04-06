# ALTER TABLE SET LOCATION

Changes the location for the table named `table_name`, and optionally a partition with `partition_spec`.

## Synopsis

```sql

ALTER TABLE table_name [ PARTITION (partition_spec) ] SET LOCATION 'new location'
```

## Parameters

**PARTITION (partition\_spec)**

Specifies the partition with parameters `partition_spec` whose location you want to change. The `partition_spec` specifies a column name/value combination in the form `partition_col_name = partition_col_value`.

**SET LOCATION 'new location'**

Specifies the new location, which must be an Amazon S3 location. For information
about syntax, see [Table Location in\
Amazon S3](https://docs.aws.amazon.com/athena/latest/ug/tables-location-format.html).

## Examples

```sql

ALTER TABLE customers PARTITION (zip='98040', state='WA') SET LOCATION 's3://amzn-s3-demo-bucket/custdata/';
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ALTER TABLE REPLACE COLUMNS

ALTER TABLE SET TBLPROPERTIES
