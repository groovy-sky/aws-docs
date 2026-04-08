# Aurora MySQL database engine updates: 2016-06-01 (version 1.6.5) (Deprecated)

**Version:** 1.6.5

## New features

- **Efficient storage of Binary Logs** –
Efficient storage of binary logs is now enabled by default for all Aurora MySQL
DB clusters, and is not configurable. Efficient storage of binary logs was
introduced in the April 2016 update. For more information, see [Aurora MySQL database engine updates: 2016-04-06 (version 1.6) (Deprecated)](auroramysql-updates-20160406.md).

## Improvements

- Improved stability for Aurora Replicas when the primary instance is
encountering a heavy workload.

- Improved stability for Aurora Replicas when running queries on partitioned
tables and tables with special characters in the table name.

- Fixed connection issues when using secure connections.

## Integration of MySQL bug fixes

- SLAVE CAN'T CONTINUE REPLICATION AFTER MASTER'S CRASH RECOVERY (Port Bug
#17632285)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Aurora MySQL updates: 2016-08-30 (version 1.7.0) (Deprecated)

Aurora MySQL updates: 2016-04-06 (version 1.6) (Deprecated)

All content copied from https://docs.aws.amazon.com/.
