# SQL statistics for MariaDB and MySQL

MariaDB and MySQL collect SQL statistics only at
the digest level. No statistics are shown at the statement level.

###### Topics

- [Digest statistics for MariaDB and MySQL](#USER_PerfInsights.UsingDashboard.AnalyzeDBLoad.AdditionalMetrics.MySQL.truncation)

- [Per-second statistics for MariaDB and MySQL](#USER_PerfInsights.UsingDashboard.AnalyzeDBLoad.AdditionalMetrics.MySQL.per-second)

- [Per-call statistics for MariaDB and MySQL](#USER_PerfInsights.UsingDashboard.AnalyzeDBLoad.AdditionalMetrics.MySQL.truncation.per-call)

- [Primary statistics for MariaDB and MySQL](#USER_PerfInsights.UsingDashboard.AnalyzeDBLoad.AdditionalMetrics.MySQL.primary)

## Digest statistics for MariaDB and MySQL

Performance Insights collects SQL digest statistics from the `events_statements_summary_by_digest` table. The
`events_statements_summary_by_digest` table is managed by your database.

The digest table doesn't have an eviction policy. When the table is full, the AWS Management Console shows the following message:

```nohighlight

Performance Insights is unable to collect SQL Digest statistics on new queries because the table events_statements_summary_by_digest is full.
Please truncate events_statements_summary_by_digest table to clear the issue. Check the User Guide for more details.
```

In this situation, MariaDB and MySQL don't track SQL queries. To address this issue, Performance Insights automatically truncates the digest table when both
of the following conditions are met:

- The table is full.

- Performance Insights manages the Performance Schema automatically.

For automatic management, the `performance_schema` parameter must be set to `0` and the
**Source** must not be set to `user`. If Performance Insights isn't managing the Performance
Schema automatically, see [Overview of the Performance Schema for Performance Insights on Amazon RDS for MariaDB or MySQL](user-perfinsights-enablemysql.md).

In the AWS CLI, check the source of a parameter value by running the [describe-db-parameters](../../../cli/latest/reference/rds/describe-db-parameters.md) command.

## Per-second statistics for MariaDB and MySQL

The following SQL statistics are available for MariaDB and MySQL DB instances.

MetricUnitdb.sql\_tokenized.stats.count\_star\_per\_secCalls per seconddb.sql\_tokenized.stats.sum\_timer\_wait\_per\_secAverage latency per second (in ms)db.sql\_tokenized.stats.sum\_select\_full\_join\_per\_secSelect full join per seconddb.sql\_tokenized.stats.sum\_select\_range\_check\_per\_secSelect range check per seconddb.sql\_tokenized.stats.sum\_select\_scan\_per\_secSelect scan per seconddb.sql\_tokenized.stats.sum\_sort\_merge\_passes\_per\_secSort merge passes per seconddb.sql\_tokenized.stats.sum\_sort\_scan\_per\_secSort scans per seconddb.sql\_tokenized.stats.sum\_sort\_range\_per\_secSort ranges per seconddb.sql\_tokenized.stats.sum\_sort\_rows\_per\_secSort rows per seconddb.sql\_tokenized.stats.sum\_rows\_affected\_per\_secRows affected per seconddb.sql\_tokenized.stats.sum\_rows\_examined\_per\_secRows examined per seconddb.sql\_tokenized.stats.sum\_rows\_sent\_per\_secRows sent per seconddb.sql\_tokenized.stats.sum\_created\_tmp\_disk\_tables\_per\_secCreated temporary disk tables per seconddb.sql\_tokenized.stats.sum\_created\_tmp\_tables\_per\_secCreated temporary tables per seconddb.sql\_tokenized.stats.sum\_lock\_time\_per\_secLock time per second (in ms)

## Per-call statistics for MariaDB and MySQL

The following metrics provide per call statistics for a SQL statement.

MetricUnitdb.sql\_tokenized.stats.sum\_timer\_wait\_per\_callAverage latency per call (in ms) db.sql\_tokenized.stats.sum\_select\_full\_join\_per\_callSelect full joins per calldb.sql\_tokenized.stats.sum\_select\_range\_check\_per\_callSelect range check per calldb.sql\_tokenized.stats.sum\_select\_scan\_per\_callSelect scans per calldb.sql\_tokenized.stats.sum\_sort\_merge\_passes\_per\_callSort merge passes per calldb.sql\_tokenized.stats.sum\_sort\_scan\_per\_callSort scans per calldb.sql\_tokenized.stats.sum\_sort\_range\_per\_callSort ranges per calldb.sql\_tokenized.stats.sum\_sort\_rows\_per\_callSort rows per calldb.sql\_tokenized.stats.sum\_rows\_affected\_per\_callRows affected per calldb.sql\_tokenized.stats.sum\_rows\_examined\_per\_callRows examined per calldb.sql\_tokenized.stats.sum\_rows\_sent\_per\_callRows sent per calldb.sql\_tokenized.stats.sum\_created\_tmp\_disk\_tables\_per\_callCreated temporary disk tables per calldb.sql\_tokenized.stats.sum\_created\_tmp\_tables\_per\_callCreated temporary tables per calldb.sql\_tokenized.stats.sum\_lock\_time\_per\_callLock time per call (in ms)

## Primary statistics for MariaDB and MySQL

The following SQL statistics are available for MariaDB and MySQL DB instances.

MetricUnitdb.sql\_tokenized.stats.count\_starCallsdb.sql\_tokenized.stats.sum\_timer\_waitWait time (in ms)db.sql\_tokenized.stats.sum\_select\_full\_joinSelect full joindb.sql\_tokenized.stats.sum\_select\_range\_checkSelect range checksdb.sql\_tokenized.stats.sum\_select\_scanSelect scansdb.sql\_tokenized.stats.sum\_sort\_merge\_passesSort merge passesdb.sql\_tokenized.stats.sum\_sort\_scanSort scansdb.sql\_tokenized.stats.sum\_sort\_rangeSort rangesdb.sql\_tokenized.stats.sum\_sort\_rowsSort rowsdb.sql\_tokenized.stats.sum\_rows\_affectedRows affecteddb.sql\_tokenized.stats.sum\_rows\_examinedRows examineddb.sql\_tokenized.stats.sum\_rows\_sentRows sentdb.sql\_tokenized.stats.sum\_created\_tmp\_disk\_tablesCreated temporary disk tablesdb.sql\_tokenized.stats.sum\_created\_tmp\_tablesCreated temporary tablesdb.sql\_tokenized.stats.sum\_lock\_timeLock time (in ms)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SQL statistics for Performance Insights

SQL statistics for Oracle

All content copied from https://docs.aws.amazon.com/.
