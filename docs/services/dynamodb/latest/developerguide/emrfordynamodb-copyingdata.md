# Copying data to and from Amazon DynamoDB

In the [Tutorial: Working with Amazon DynamoDB and Apache Hive](emrfordynamodb-tutorial.md), you copied data from a native Hive table into an external DynamoDB table, and then
queried the external DynamoDB table. The table is external because it exists outside of
Hive. Even if you drop the Hive table that maps to it, the table in DynamoDB is not
affected.

Hive is an excellent solution for copying data among DynamoDB tables, Amazon S3 buckets,
native Hive tables, and Hadoop Distributed File System (HDFS). This section provides
examples of these operations.

###### Topics

- [Copying data between DynamoDB and a native Hive table](emrfordynamodb-copyingdata-nativehive.md)

- [Copying data between DynamoDB and Amazon S3](emrfordynamodb-copyingdata-s3.md)

- [Copying data between DynamoDB and HDFS](emrfordynamodb-copyingdata-hdfs.md)

- [Using data compression](emrfordynamodb-copyingdata-compression.md)

- [Reading non-printable UTF-8 character data](emrfordynamodb-copyingdata-nonprintabledata.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Querying data in DynamoDB

Copying data between DynamoDB and a native Hive table

All content copied from https://docs.aws.amazon.com/.
