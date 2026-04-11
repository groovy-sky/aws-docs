# Aurora MySQL database engine updates 2019-05-02 (version 2.04.2) (Deprecated)

**Version:** 2.04.2

Aurora MySQL 2.04.2 is generally available. Aurora MySQL 2.x versions are compatible with MySQL 5.7
and Aurora MySQL 1.x versions are compatible with MySQL 5.6.

Currently supported Aurora MySQL releases are 1.19.5, 1.19.6, 1.22.\*, 1.23.\*, 2.04.\*, 2.07.\*, 2.08.\*, 2.09.\*, 2.10.\*, 3.01.\* and 3.02.\*.

When creating a new Aurora MySQL DB cluster (including restoring a snapshot), you have the
option of choosing compatibility with either MySQL 5.7 or MySQL 5.6. We do not allow in-place
upgrade of Aurora MySQL 1.\* clusters or restore of Aurora MySQL 1.\* clusters from an Amazon S3 backup into Aurora MySQL 2.04.2.
We plan to remove these restrictions in a later Aurora MySQL 2.\* release.

You can restore snapshots of Aurora MySQL 1.14.\*, 1.15.\*, 1.16.\*, 1.17.\*, 1.18.\*, 1.19.\*,
2.01.\*, 2.02.\*, 2.03.\*, 2.04.0, and 2.04.1 into Aurora MySQL 2.04.2.

If you have any questions or concerns, AWS Support is available on the community forums and through
[AWS Support](https://aws.amazon.com/support). For more information, see
[Maintaining an Amazon Aurora DB cluster](../aurorauserguide/user-upgradedbinstance-maintenance.md) in the _Amazon Aurora User Guide_.

###### Note

This version is currently not available in the AWS GovCloud (US-West) \[us-gov-west-1\]
and China (Ningxia) \[cn-northwest-1\] AWS Regions. There will be a separate announcement once
it is made available.

###### Note

For information on how to upgrade your Aurora MySQL database cluster, see [Upgrading the minor version or patch level of an Aurora MySQL DB cluster](../aurorauserguide/auroramysql-updates-patching.md) in the
_Amazon Aurora User Guide_.

## Improvements

- Added support for SSL binlog replication using custom certificates. For information on using SSL binlog replication in Aurora MySQL,
see [mysql\_rds\_import\_binlog\_ssl\_material](../userguide/mysql-rds-import-binlog-ssl-material.md).

- Fixed a deadlatch on the Aurora primary instance that occurs when a table with a Full Text Search index is being optimized.

- Fixed an issue on the Aurora Replicas where performance of certain queries using
`SELECT(*)` could be impacted on tables that have secondary indexes.

- Fixed a condition that resulted in Error 1032 being posted.

- Improved the stability of Aurora Replicas by fixing multiple deadlatches.

## Integration of MySQL bug fixes

- Bug #24829050 - INDEX\_MERGE\_INTERSECTION OPTIMIZATION CAUSES WRONG QUERY RESULTS

## Comparison with Aurora MySQL version 1

The following Amazon Aurora MySQL features are supported in Aurora MySQL Version 1 (compatible with
MySQL 5.6), but these features are currently not supported in Aurora MySQL Version 2 (compatible
with MySQL 5.7).

- Asynchronous key prefetch (AKP). For more
information, see [Optimizing Aurora indexed join queries with asynchronous key prefetch](../aurorauserguide/auroramysql-bestpractices.md#Aurora.BestPractices.AKP) in the
_Amazon Aurora User Guide_.

- Hash joins. For more information, see [Optimizing large Aurora MySQL join queries with hash joins](../aurorauserguide/auroramysql-bestpractices.md#Aurora.BestPractices.HashJoin) in the
_Amazon Aurora User Guide_.

- Native functions for synchronously invoking AWS Lambda functions. For more
information, see [Invoking a Lambda function with an Aurora MySQL native function](../aurorauserguide/auroramysql-integrating-lambda.md#AuroraMySQL.Integrating.NativeLambda) in the
_Amazon Aurora User Guide_.

- Scan batching. For more information, see [Aurora MySQL database engine updates 2017-12-11 (version 1.16) (Deprecated)](auroramysql-updates-20171211.md).

- Migrating data from MySQL using an Amazon S3 bucket. For more information, see [Migrating data from MySQL by using an Amazon S3 bucket](../aurorauserguide/auroramysql-migrating-extmysql.md#AuroraMySQL.Migrating.ExtMySQL.S3) in the
_Amazon Aurora User Guide_.

## MySQL 5.7 compatibility

Aurora MySQL 2.04.2 is wire-compatible with MySQL 5.7 and includes features such as JSON support, spatial indexes,
and generated columns. Aurora MySQL uses a native implementation of spatial indexing using z-order curves to deliver
>20x better write performance and >10x better read performance than MySQL 5.7 for spatial datasets.

Aurora MySQL 2.04.2 does not currently support the following MySQL 5.7 features:

- Group replication plugin

- Increased page size

- InnoDB buffer pool loading at startup

- InnoDB full-text parser plugin

- Multisource replication

- Online buffer pool resizing

- Password validation plugin

- Query rewrite plugins

- Replication filtering

- The `CREATE TABLESPACE` SQL statement

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Aurora MySQL updates: 2019-05-09 (version 2.04.3) (Deprecated)

Aurora MySQL updates: 2019-03-25 (version 2.04.1) (Deprecated)

All content copied from https://docs.aws.amazon.com/.
