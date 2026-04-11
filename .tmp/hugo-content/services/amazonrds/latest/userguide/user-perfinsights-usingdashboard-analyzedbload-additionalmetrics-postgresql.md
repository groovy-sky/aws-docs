# SQL statistics for RDS PostgreSQL

For each SQL call and for each second that a query runs,
Performance Insights collects SQL statistics. RDS for PostgreSQL collect SQL statistics only at the digest–level. No statistics are shown at the
statement-level.

Following, you can find information about digest-level statistics for
RDS for PostgreSQL.

###### Topics

- [Digest statistics for RDS PostgreSQL](#USER_PerfInsights.UsingDashboard.AnalyzeDBLoad.AdditionalMetrics.PostgreSQL.digest)

- [Per-second digest statistics for RDS PostgreSQL](#USER_PerfInsights.UsingDashboard.AnalyzeDBLoad.AdditionalMetrics.PostgreSQL.per-second)

- [Per-call digest statistics for RDS PostgreSQL](#USER_PerfInsights.UsingDashboard.AnalyzeDBLoad.AdditionalMetrics.PostgreSQL.per-call)

- [Primary statistics for RDS PostgreSQL](#USER_PerfInsights.UsingDashboard.AnalyzeDBLoad.AdditionalMetrics.PostgreSQL.primary)

## Digest statistics for RDS PostgreSQL

To view SQL digest statistics, RDS PostgreSQL must load the `pg_stat_statements` library. For PostgreSQL DB instances that are compatible
with PostgreSQL 11 or later, the database loads this library by default. For PostgreSQL DB instances that are compatible with
PostgreSQL 10 or earlier, enable this library manually. To enable it manually, add `pg_stat_statements` to
`shared_preload_libraries` in the DB parameter group associated with the DB instance. Then reboot your DB instance. For
more information, see [Parameter groups for Amazon RDS](user-workingwithparamgroups.md).

###### Note

Performance Insights can only collect statistics for queries in `pg_stat_activity` that aren't truncated. By
default, PostgreSQL databases truncate queries longer than 1,024 bytes. To increase the query size, change the
`track_activity_query_size` parameter in the DB parameter group associated with your DB instance. When you change
this parameter, a DB instance reboot is required.

## Per-second digest statistics for RDS PostgreSQL

The following SQL digest statistics are available for PostgreSQL DB instances.

MetricUnitdb.sql\_tokenized.stats.calls\_per\_secCalls per seconddb.sql\_tokenized.stats.rows\_per\_secRows per seconddb.sql\_tokenized.stats.total\_time\_per\_secAverage active executions per second (AAE)db.sql\_tokenized.stats.shared\_blks\_hit\_per\_secBlock hits per seconddb.sql\_tokenized.stats.shared\_blks\_read\_per\_secBlock reads per seconddb.sql\_tokenized.stats.shared\_blks\_dirtied\_per\_secBlocks dirtied per seconddb.sql\_tokenized.stats.shared\_blks\_written\_per\_secBlock writes per seconddb.sql\_tokenized.stats.local\_blks\_hit\_per\_secLocal block hits per seconddb.sql\_tokenized.stats.local\_blks\_read\_per\_secLocal block reads per seconddb.sql\_tokenized.stats.local\_blks\_dirtied\_per\_secLocal block dirtied per seconddb.sql\_tokenized.stats.local\_blks\_written\_per\_secLocal block writes per seconddb.sql\_tokenized.stats.temp\_blks\_written\_per\_secTemporary writes per seconddb.sql\_tokenized.stats.temp\_blks\_read\_per\_secTemporary reads per seconddb.sql\_tokenized.stats.blk\_read\_time\_per\_secAverage concurrent reads per seconddb.sql\_tokenized.stats.blk\_write\_time\_per\_secAverage concurrent writes per second

## Per-call digest statistics for RDS PostgreSQL

The following metrics provide per call statistics for a SQL statement.

MetricUnitdb.sql\_tokenized.stats.rows\_per\_callRows per calldb.sql\_tokenized.stats.avg\_latency\_per\_callAverage latency per call (in ms)db.sql\_tokenized.stats.shared\_blks\_hit\_per\_callBlock hits per calldb.sql\_tokenized.stats.shared\_blks\_read\_per\_callBlock reads per calldb.sql\_tokenized.stats.shared\_blks\_written\_per\_callBlock writes per calldb.sql\_tokenized.stats.shared\_blks\_dirtied\_per\_callBlocks dirtied per calldb.sql\_tokenized.stats.local\_blks\_hit\_per\_callLocal block hits per calldb.sql\_tokenized.stats.local\_blks\_read\_per\_callLocal block reads per calldb.sql\_tokenized.stats.local\_blks\_dirtied\_per\_callLocal block dirtied per calldb.sql\_tokenized.stats.local\_blks\_written\_per\_callLocal block writes per calldb.sql\_tokenized.stats.temp\_blks\_written\_per\_callTemporary block writes per calldb.sql\_tokenized.stats.temp\_blks\_read\_per\_callTemporary block reads per calldb.sql\_tokenized.stats.blk\_read\_time\_per\_callRead time per call (in ms)db.sql\_tokenized.stats.blk\_write\_time\_per\_callWrite time per call (in ms)

## Primary statistics for RDS PostgreSQL

The following SQL statistics are available for PostgreSQL DB instances.

MetricUnitdb.sql\_tokenized.stats.callsCalls db.sql\_tokenized.stats.rowsRows db.sql\_tokenized.stats.total\_timeTotal time (in ms)db.sql\_tokenized.stats.shared\_blks\_hitBlock hits db.sql\_tokenized.stats.shared\_blks\_readBlock reads db.sql\_tokenized.stats.shared\_blks\_dirtiedBlocks dirtied db.sql\_tokenized.stats.shared\_blks\_writtenBlock writes db.sql\_tokenized.stats.local\_blks\_hitLocal block hits db.sql\_tokenized.stats.local\_blks\_readLocal block reads db.sql\_tokenized.stats.local\_blks\_dirtiedLocal blocks dirtieddb.sql\_tokenized.stats.local\_blks\_writtenLocal block writes db.sql\_tokenized.stats.temp\_blks\_writtenTemporary writes db.sql\_tokenized.stats.temp\_blks\_readTemporary reads db.sql\_tokenized.stats.blk\_read\_timeAverage concurrent reads (in ms)db.sql\_tokenized.stats.blk\_write\_timeAverage concurrent writes (in ms)

For more information about these metrics, see [pg\_stat\_statements](https://www.postgresql.org/docs/current/pgstatstatements.html) in the PostgreSQL documentation.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SQL statistics for SQL Server

OS metrics in Enhanced Monitoring

All content copied from https://docs.aws.amazon.com/.
