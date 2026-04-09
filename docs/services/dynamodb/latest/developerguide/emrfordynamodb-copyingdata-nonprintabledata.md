# Reading non-printable UTF-8 character data

To read and write non-printable UTF-8 character data, you can use the `STORED
                    AS SEQUENCEFILE` clause when you create a Hive table. A SequenceFile is a
Hadoop binary file format. You need to use Hadoop to read this file. The following
example shows how to export data from DynamoDB into Amazon S3. You can use this
functionality to handle non-printable UTF-8 encoded characters.

```nohighlight

CREATE EXTERNAL TABLE s3_export(a_col string, b_col bigint, c_col array<string>)
STORED AS SEQUENCEFILE
LOCATION 's3://bucketname/path/subpath/';

INSERT OVERWRITE TABLE s3_export SELECT *
FROM hiveTableName;
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Using data compression

Performance tuning

All content copied from https://docs.aws.amazon.com/.
