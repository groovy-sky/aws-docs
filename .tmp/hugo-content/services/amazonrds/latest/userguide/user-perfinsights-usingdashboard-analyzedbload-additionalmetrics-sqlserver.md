# SQL statistics for Amazon RDS for SQL Server

Amazon RDS for SQL Server collects SQL statistics both at the statement and digest level. At the statement level, the ID column represents the
value of `sql_handle`. At the digest level, the ID column shows the value of `query_hash`.

SQL Server returns NULL values for `query_hash` for a few statements. For example, ALTER INDEX, CHECKPOINT,
UPDATE STATISTICS, COMMIT TRANSACTION, FETCH NEXT FROM Cursor, and a few INSERT statements, SELECT @<variable>, conditional statements,
and executable stored procedures. In this case, the `sql_handle` value is displayed as the ID at the digest level for that statement.

###### Topics

- [Per-second statistics for SQL Server](#USER_PerfInsights.UsingDashboard.AnalyzeDBLoad.AdditionalMetrics.SQLServer.per-second)

- [Per-call statistics for SQL Server](#USER_PerfInsights.UsingDashboard.AnalyzeDBLoad.AdditionalMetrics.SQLServer.per-call)

- [Primary statistics for SQL Server](#USER_PerfInsights.UsingDashboard.AnalyzeDBLoad.AdditionalMetrics.SQLServer.primary)

## Per-second statistics for SQL Server

The following metrics provide per-second statistics for a SQL Server SQL query.

MetricUnitdb.sql.stats.execution\_count\_per\_secNumber of executions per seconddb.sql.stats.total\_elapsed\_time\_per\_secTotal elapsed time per seconddb.sql.stats.total\_rows\_per\_secTotal rows processed per seconddb.sql.stats.total\_logical\_reads\_per\_secTotal logical reads per seconddb.sql.stats.total\_logical\_writes\_per\_secTotal logical writes per seconddb.sql.stats.total\_physical\_reads\_per\_secTotal physical reads per seconddb.sql.stats.total\_worker\_time\_per\_secTotal CPU time (in ms)

The following metrics provide per-second statistics for a SQL Server SQL digest
query.

MetricUnitdb.sql\_tokenized.stats.execution\_count\_per\_secNumber of execution per seconddb.sql\_tokenized.stats.total\_elapsed\_time\_per\_secTotal elapsed time per seconddb.sql\_tokenized.stats.total\_rows\_per\_secTotal rows processed per seconddb.sql\_tokenized.stats.total\_logical\_reads\_per\_secTotal logical reads per seconddb.sql\_tokenized.stats.total\_logical\_writes\_per\_secTotal logical writes per seconddb.sql\_tokenized.stats.total\_physical\_reads\_per\_secTotal physical reads per seconddb.sql\_tokenized.stats.total\_worker\_time\_per\_secTotal CPU time (in ms)

## Per-call statistics for SQL Server

The following metrics provide per-call statistics for a SQL Server SQL statement.

MetricUnitdb.sql.stats.total\_elapsed\_time\_per\_callTotal elapsed time per execution (in ms)db.sql.stats.total\_rows\_per\_callTotal rows processed per executiondb.sql.stats.total\_logical\_reads\_per\_callTotal logical reads per executiondb.sql.stats.total\_logical\_writes\_per\_callTotal logical writes per executiondb.sql.stats.total\_physical\_reads\_per\_callTotal physical reads per executiondb.sql.stats.total\_worker\_time\_per\_callTotal CPU time per execution (in ms)

The following metrics provide per-call statistics for a SQL Server SQL digest query.

MetricUnitdb.sql\_tokenized.stats.total\_elapsed\_time\_per\_callTotal elapsed time per executiondb.sql\_tokenized.stats.total\_rows\_per\_callTotal rows processed per executiondb.sql\_tokenized.stats.total\_logical\_reads\_per\_callTotal logical reads per executiondb.sql\_tokenized.stats.total\_logical\_writes\_per\_callTotal logical writes per executiondb.sql\_tokenized.stats.total\_physical\_reads\_per\_callTotal physical reads per execution db.sql\_tokenized.stats.total\_worker\_time\_per\_callTotal CPU time per execution (in ms)

## Primary statistics for SQL Server

The following metrics provide primary statistics for a SQL Server SQL query.

MetricUnitdb.sql.stats.execution\_countNumber of executionsdb.sql.stats.total\_elapsed\_timeTotal elapsed time (in ms)db.sql.stats.total\_rowsTotal rows processeddb.sql.stats.total\_logical\_readsTotal logical readsdb.sql.stats.total\_logical\_writesTotal logical writesdb.sql.stats.total\_physical\_readsTotal physical readsdb.sql.stats.total\_worker\_timeTotal CPU time (in ms)

The following metrics provide primary statistics for a SQL Server SQL digest query.

MetricUnitdb.sql\_tokenized.stats.execution\_countNumber of executiondb.sql\_tokenized.stats.total\_elapsed\_timeTotal elapsed time (in ms)db.sql\_tokenized.stats.total\_rowsTotal rows processeddb.sql\_tokenized.stats.total\_logical\_readsTotal logical readsdb.sql\_tokenized.stats.total\_logical\_writesTotal logical writesdb.sql\_tokenized.stats.total\_physical\_readsTotal physical readsdb.sql\_tokenized.stats.total\_worker\_timeTotal CPU time (in ms)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SQL statistics for Oracle

SQL statistics for RDS PostgreSQL

All content copied from https://docs.aws.amazon.com/.
