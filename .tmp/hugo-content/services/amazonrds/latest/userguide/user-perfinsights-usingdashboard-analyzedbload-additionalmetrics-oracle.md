# SQL statistics for Amazon RDS for Oracle

Amazon RDS for Oracle collects SQL statistics both at the statement and digest level. At the statement level, the ID column represents the
value of `V$SQL.SQL_ID`. At the digest level, the ID column shows the value of `V$SQL.FORCE_MATCHING_SIGNATURE`.

If the ID is `0` at the digest level, Oracle Database has determined that this statement is not suitable for reuse. In this
case, the child SQL statements could belong to different digests. However, the statements are grouped together under the
`digest_text` for the first SQL statement collected.

###### Topics

- [Per-second statistics for Oracle](#USER_PerfInsights.UsingDashboard.AnalyzeDBLoad.AdditionalMetrics.Oracle.per-second)

- [Per-call statistics for Oracle](#USER_PerfInsights.UsingDashboard.AnalyzeDBLoad.AdditionalMetrics.Oracle.per-call)

- [Primary statistics for Oracle](#USER_PerfInsights.UsingDashboard.AnalyzeDBLoad.AdditionalMetrics.Oracle.primary)

## Per-second statistics for Oracle

The following metrics provide per-second statistics for an Oracle SQL query.

MetricUnitdb.sql.stats.executions\_per\_secNumber of executions per seconddb.sql.stats.elapsed\_time\_per\_secAverage active executions (AAE)db.sql.stats.rows\_processed\_per\_secRows processed per seconddb.sql.stats.buffer\_gets\_per\_secBuffer gets per seconddb.sql.stats.physical\_read\_requests\_per\_secPhysical reads per seconddb.sql.stats.physical\_write\_requests\_per\_secPhysical writes per seconddb.sql.stats.total\_sharable\_mem\_per\_secTotal shareable memory per second (in bytes) db.sql.stats.cpu\_time\_per\_secCPU time per second (in ms)

The following metrics provide per-second statistics for an Oracle SQL digest query.

MetricUnitdb.sql\_tokenized.stats.executions\_per\_secNumber of executions per seconddb.sql\_tokenized.stats.elapsed\_time\_per\_secAverage active executions (AAE)db.sql\_tokenized.stats.rows\_processed\_per\_secRows processed per seconddb.sql\_tokenized.stats.buffer\_gets\_per\_secBuffer gets per seconddb.sql\_tokenized.stats.physical\_read\_requests\_per\_secPhysical reads per seconddb.sql\_tokenized.stats.physical\_write\_requests\_per\_secPhysical writes per seconddb.sql\_tokenized.stats.total\_sharable\_mem\_per\_secTotal shareable memory per second (in bytes) db.sql\_tokenized.stats.cpu\_time\_per\_secCPU time per second (in ms)

## Per-call statistics for Oracle

The following metrics provide per-call statistics for an Oracle SQL statement.

MetricUnitdb.sql.stats.elapsed\_time\_per\_execElapsed time per executions (in ms) db.sql.stats.rows\_processed\_per\_execRows processed per executiondb.sql.stats.buffer\_gets\_per\_execBuffer gets per executiondb.sql.stats.physical\_read\_requests\_per\_execPhysical reads per executiondb.sql.stats.physical\_write\_requests\_per\_execPhysical writes per executiondb.sql.stats.total\_sharable\_mem\_per\_execTotal shareable memory per execution (in bytes)db.sql.stats.cpu\_time\_per\_execCPU time per execution (in ms)

The following metrics provide per-call statistics for an Oracle SQL digest query.

MetricUnitdb.sql\_tokenized.stats.elapsed\_time\_per\_execElapsed time per executions (in ms) db.sql\_tokenized.stats.rows\_processed\_per\_execRows processed per executiondb.sql\_tokenized.stats.buffer\_gets\_per\_execBuffer gets per executiondb.sql\_tokenized.stats.physical\_read\_requests\_per\_execPhysical reads per executiondb.sql\_tokenized.stats.physical\_write\_requests\_per\_execPhysical writes per executiondb.sql\_tokenized.stats.total\_sharable\_mem\_per\_execTotal shareable memory per execution (in bytes)db.sql\_tokenized.stats.cpu\_time\_per\_execCPU time per execution (in ms)

## Primary statistics for Oracle

The following metrics provide primary statistics for an Oracle SQL query.

MetricUnitdb.sql.stats.executionsNumber of executions db.sql.stats.elapsed\_timeElapsed time (in ms)db.sql.stats.rows\_processedRows processed db.sql.stats.buffer\_getsBuffer gets db.sql.stats.physical\_read\_requestsPhysical reads db.sql.stats.physical\_write\_requestsPhysical writes db.sql.stats.total\_sharable\_memTotal shareable memory (in bytes) db.sql.stats.cpu\_timeCPU time (in ms)

The following metrics provide primary statistics for an Oracle SQL digest query.

MetricUnitdb.sql\_tokenized.stats.executionsNumber of executionsdb.sql\_tokenized.stats.elapsed\_timeElapsed time (in ms)db.sql\_tokenized.stats.rows\_processedRows processeddb.sql\_tokenized.stats.buffer\_getsBuffer getsdb.sql\_tokenized.stats.physical\_read\_requestsPhysical readsdb.sql\_tokenized.stats.physical\_write\_requestsPhysical writesdb.sql\_tokenized.stats.total\_sharable\_memTotal shareable memory (in bytes) db.sql\_tokenized.stats.cpu\_timeCPU time (in ms)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SQL statistics for MariaDB and MySQL

SQL statistics for SQL Server

All content copied from https://docs.aws.amazon.com/.
