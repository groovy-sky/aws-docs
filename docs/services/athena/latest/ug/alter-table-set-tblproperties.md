# ALTER TABLE SET TBLPROPERTIES

Adds custom or predefined metadata properties to a table and sets their assigned values.
To see the properties in a table, use the [SHOW TBLPROPERTIES](https://docs.aws.amazon.com/athena/latest/ug/show-tblproperties.html) command.

Apache Hive [Managed tables](https://cwiki.apache.org/confluence/display/Hive/Managed+vs.+External+Tables) are not supported, so setting `'EXTERNAL'='FALSE'`
has no effect.

## Synopsis

```sql

ALTER TABLE table_name SET TBLPROPERTIES ('property_name' = 'property_value' [ , ... ])
```

## Parameters

**SET TBLPROPERTIES ('property\_name' = 'property\_value' \[ , ... \])**

Specifies the metadata properties to add as `property_name` and
the value for each as `property value`. If
`property_name` already exists, its value is set to the newly
specified `property_value`.

The following predefined table properties have special uses.

Predefined propertyDescription`classification`Indicates the data type for AWS Glue. Possible values are
`csv`, `parquet`,
`orc`, `avro`, or `json`.
Tables created for Athena in the CloudTrail console add
`cloudtrail` as a value for the
`classification` property. For more
information, see the TBLPROPERTIES section of [CREATE TABLE](https://docs.aws.amazon.com/athena/latest/ug/create-table.html).`has_encrypted_data`Indicates whether the dataset specified by
`LOCATION` is CSE-KMS encrypted. For more
information, see the TBLPROPERTIES section of [CREATE TABLE](https://docs.aws.amazon.com/athena/latest/ug/create-table.html)
and [Create tables based on encrypted datasets in Amazon S3](https://docs.aws.amazon.com/athena/latest/ug/creating-tables-based-on-encrypted-datasets-in-s3.html).`encryption_option`Indicates the highest level of encryption used in the underlying
dataset specified by `LOCATION`. For more
information, see the TBLPROPERTIES section of [CREATE TABLE](https://docs.aws.amazon.com/athena/latest/ug/create-table.html)
and [Create tables based on encrypted datasets in Amazon S3](https://docs.aws.amazon.com/athena/latest/ug/creating-tables-based-on-encrypted-datasets-in-s3.html).`kms_key`Indicates the AWS KMS key ARN used to encrypt and decrypt SSE-KMS or CSE-KMS
data files in the underlying dataset specified by `LOCATION`. For more
information, see the TBLPROPERTIES section of [CREATE TABLE](https://docs.aws.amazon.com/athena/latest/ug/create-table.html)
and [Create tables based on encrypted datasets in Amazon S3](https://docs.aws.amazon.com/athena/latest/ug/creating-tables-based-on-encrypted-datasets-in-s3.html).`orc.compress`Specifies a compression format for data in ORC format.
For more information, see [ORC SerDe](https://docs.aws.amazon.com/athena/latest/ug/orc-serde.html).`parquet.compression`Specifies a compression format for data in Parquet
format. For more information, see [Parquet SerDe](https://docs.aws.amazon.com/athena/latest/ug/parquet-serde.html).`write.compression`Specifies a compression format for data in the text file
or JSON formats. For the Parquet and ORC formats, use the
`parquet.compression` and
`orc.compress` properties
respectively.`compression_level`Specifies a compression level to use. This property
applies only to ZSTD compression. Possible values are from 1
to 22. The default value is 3. For more information, see
[Use ZSTD compression levels](https://docs.aws.amazon.com/athena/latest/ug/compression-support-zstd-levels.html).`projection.*`Custom properties used in partition projection that allow
Athena to know what partition patterns to expect when it runs
a query on a table. For more information, see [Use partition projection with Amazon Athena](https://docs.aws.amazon.com/athena/latest/ug/partition-projection.html).`skip.header.line.count`Ignores headers in data when you define a table. For more
information, see [Ignoring headers](https://docs.aws.amazon.com/athena/latest/ug/lazy-simple-serde.html#lazy-simple-serde-ignoring-headers).`storage.location.template`Specifies a custom Amazon S3 path template for projected
partitions. For more information, see [Set up partition projection](https://docs.aws.amazon.com/athena/latest/ug/partition-projection-setting-up.html).

## Examples

The following example adds a comment note to table properties.

```sql

ALTER TABLE orders
SET TBLPROPERTIES ('notes'="Please don't drop this table.");
```

The following example modifies the table `existing_table` to use Parquet
file format with ZSTD compression and ZSTD compression level 4.

```sql

ALTER TABLE existing_table
SET TBLPROPERTIES ('parquet.compression' = 'ZSTD', 'compression_level' = 4)
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ALTER TABLE SET LOCATION

ALTER VIEW DIALECT
