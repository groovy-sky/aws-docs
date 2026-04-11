# Aurora MySQL database engine updates: 2017-01-12 (version 1.10.1) (Deprecated)

**Version:** 1.10.1

Version 1.10.1 of Aurora MySQL is an opt-in version and is not used to patch your
database instances. It is available for creating new Aurora instances and for upgrading
existing instances. You can apply the patch by choosing a cluster in the [Amazon RDS console](https://console.aws.amazon.com/rds), choosing **Cluster**
**Actions**, and then choosing **Upgrade Now**. Patching
requires a database restart with downtime typically lasting 5-30 seconds, after which
you can resume using your DB clusters. This patch is using a cluster patching model
where all nodes in an Aurora cluster are patched at the same time.

## New features

- **Advanced Auditing** – Aurora MySQL
provides a high-performance Advanced Auditing feature, which you can use to
audit database activity. For more information about enabling and using
Advanced Auditing, see [Using Advanced Auditing with an Aurora MySQL DB cluster](../aurorauserguide/auroramysql-auditing.md) in the _Amazon Aurora User Guide_.

## Improvements

- Fixed an issue with spatial indexing when creating a column and adding an
index on it in the same statement.

- Fixed an issue where spatial statistics aren't persisted across DB cluster
restart.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Aurora MySQL updates: 2017-02-23 (version 1.11) (Deprecated)

Aurora MySQL updates: 2016-12-14 (version 1.10) (Deprecated)

All content copied from https://docs.aws.amazon.com/.
