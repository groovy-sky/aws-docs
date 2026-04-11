# Best practices with Amazon Aurora MySQL

This topic includes information on best practices and options for using or migrating data
to an Amazon Aurora MySQL DB cluster. The information in this topic summarizes and reiterates
some of the guidelines and procedures that you can find in
[Managing an Amazon Aurora DB cluster](chap-aurora.md).

###### Contents

- [Determining which DB instance you are connected to](auroramysql-bestpractices.md#AuroraMySQL.BestPractices.DeterminePrimaryInstanceConnection)

- [Best practices for Aurora MySQL performance and scaling](auroramysql-bestpractices-performance.md)

  - [Using T instance classes for development and testing](auroramysql-bestpractices-performance.md#AuroraMySQL.BestPractices.T2Medium)

  - [Optimizing Aurora MySQL indexed join queries with asynchronous key prefetch](auroramysql-bestpractices-performance.md#Aurora.BestPractices.AKP)

    - [Enabling asynchronous key prefetch](auroramysql-bestpractices-performance.md#Aurora.BestPractices.AKP.Enabling)

    - [Optimizing queries for asynchronous key prefetch](auroramysql-bestpractices-performance.md#Aurora.BestPractices.AKP.Optimizing)
  - [Optimizing large Aurora MySQL join queries with hash joins](auroramysql-bestpractices-performance.md#Aurora.BestPractices.HashJoin)

    - [Enabling hash joins](auroramysql-bestpractices-performance.md#Aurora.BestPractices.HashJoin.Enabling)

    - [Optimizing queries for hash joins](auroramysql-bestpractices-performance.md#Aurora.BestPractices.HashJoin.Optimizing)
  - [Using Amazon Aurora to scale reads for your MySQL database](auroramysql-bestpractices-performance.md#AuroraMySQL.BestPractices.ReadScaling)

  - [Optimizing timestamp operations](auroramysql-bestpractices-performance.md#AuroraMySQL.BestPractices.Performance.TimeZone)

  - [Virtual index ID overflow errors](auroramysql-bestpractices-performance.md#AuroraMySQL.BestPractices.Performance.VirtualIndexIDOverflow)
- [Best practices for Aurora MySQL high availability](auroramysql-bestpractices-ha.md)

  - [Using Amazon Aurora for Disaster Recovery with your MySQL databases](auroramysql-bestpractices-ha.md#AuroraMySQL.BestPractices.DisasterRecovery)

  - [Migrating from MySQL to Amazon Aurora MySQL with reduced downtime](auroramysql-bestpractices-ha.md#AuroraMySQL.BestPractices.Migrating)

  - [Avoiding slow performance, automatic restart, and failover for Aurora MySQL DB instances](auroramysql-bestpractices-ha.md#AuroraMySQL.BestPractices.Avoiding)
- [Recommendations for MySQL features in Aurora MySQL](auroramysql-bestpractices-featurerecommendations.md)

  - [Using multithreaded replication in Aurora MySQL](auroramysql-bestpractices-featurerecommendations.md#AuroraMySQL.BestPractices.MTReplica)

  - [Invoking AWS Lambda functions using native MySQL functions](auroramysql-bestpractices-featurerecommendations.md#AuroraMySQL.BestPractices.Lambda)

  - [Avoiding XA transactions with Amazon Aurora MySQL](auroramysql-bestpractices-featurerecommendations.md#AuroraMySQL.BestPractices.XA)

  - [Keeping foreign keys turned on during DML statements](auroramysql-bestpractices-featurerecommendations.md#Aurora.BestPractices.ForeignKeys)

  - [Configuring how frequently the log buffer is flushed](auroramysql-bestpractices-featurerecommendations.md#AuroraMySQL.BestPractices.Flush)

  - [Minimizing and troubleshooting Aurora MySQL deadlocks](auroramysql-bestpractices-featurerecommendations.md#AuroraMySQL.BestPractices.deadlocks)

    - [Minimizing InnoDB deadlocks](auroramysql-bestpractices-featurerecommendations.md#AuroraMySQL.BestPractices.deadlocks-minimize)

    - [Monitoring InnoDB deadlocks](auroramysql-bestpractices-featurerecommendations.md#AuroraMySQL.BestPractices.deadlocks-monitor)
- [Evaluating DB instance usage for Aurora MySQL with Amazon CloudWatch metrics](auroramysql-bestpractices-cw.md)

## Determining which DB instance you are connected to

To determine which DB instance in an Aurora MySQL DB cluster a connection is
connected to, check the `innodb_read_only` global variable, as shown in
the following example.

```sql

SHOW GLOBAL VARIABLES LIKE 'innodb_read_only';
```

The `innodb_read_only` variable is set to `ON` if you are
connected to a reader DB instance. This setting is `OFF` if you are connected to a writer DB
instance, such as primary instance in a provisioned cluster.

This approach can be helpful if you want to add logic to your application code to balance the workload or to ensure that a
write operation is using the correct connection.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Aurora MySQL lab mode

Best practices for performance and scaling

All content copied from https://docs.aws.amazon.com/.
