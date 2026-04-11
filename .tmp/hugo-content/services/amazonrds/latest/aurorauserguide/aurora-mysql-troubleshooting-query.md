# Troubleshooting query performance for Aurora MySQL databases

MySQL provides [query optimizer control](https://dev.mysql.com/doc/refman/8.0/en/controlling-optimizer.html)
through system variables that affect how query plans are evaluated, switchable optimizations, optimizer and index hints, and the
optimizer cost model. These data points can be helpful not only while comparing different MySQL environments, but also to
compare previous query execution plans with current execution plans, and to understand the overall execution of a MySQL query at
any point.

Query performance depends on many factors, including the execution plan, table schema and size, statistics, resources,
indexes, and parameter configuration. Query tuning requires identifying bottlenecks and optimizing the execution path.

- Find the execution plan for the query and check whether the query is using appropriate indexes. You can optimize your
query by using `EXPLAIN` and reviewing the details of each plan.

- Aurora MySQL version 3 (compatible with MySQL 8.0 Community Edition) uses an `EXPLAIN ANALYZE` statement. The
`EXPLAIN ANALYZE` statement is a profiling tool that shows where MySQL spends time on your query and why.
With `EXPLAIN ANALYZE`, Aurora MySQL plans, prepares, and runs the query while counting rows and measuring the
time spent at various points of the execution plan. When the query completes, `EXPLAIN ANALYZE` prints the
plan and its measurements instead of the query result.

- Keep your schema statistics updated by using the `ANALYZE` statement. The query optimizer can sometimes
choose poor execution plans because of outdated statistics. This can lead to poor performance of a query because of
inaccurate cardinality estimates of both tables and indexes. The `last_update` column of the [innodb\_table\_stats](https://dev.mysql.com/doc/refman/8.0/en/innodb-persistent-stats.html) table shows the last time your schema statistics were updated, which is a good indicator
of "staleness."

- Other issues can occur, such as distribution skew of data, that aren't taken into account for table cardinality. For
more information, see [Estimating ANALYZE TABLE complexity for InnoDB tables](https://dev.mysql.com/doc/refman/8.0/en/innodb-analyze-table-complexity.html) and [Histogram statistics in MySQL](https://dev.mysql.com/blog-archive/histogram-statistics-in-mysql) in the
MySQL documentation.

## Understanding the time spent by queries

The following are ways to determine the time spent by queries:

- [Profiling](https://dev.mysql.com/doc/refman/8.0/en/show-profile.html)

- [Performance Schema](https://dev.mysql.com/doc/refman/8.0/en/performance-schema.html)

- [Query optimizer](https://dev.mysql.com/doc/refman/8.0/en/controlling-optimizer.html)

**Profiling**

By default, profiling is disabled. Enable profiling, then run the slow query and review its profile.

```nohighlight

SET profiling = 1;
Run your query.
SHOW PROFILE;
```

1. Identify the stage where the most time is spent. According to [General thread\
    states](https://dev.mysql.com/doc/refman/8.0/en/general-thread-states.html) in the MySQL documentation, reading and processing rows for a `SELECT`
    statement is often the longest-running state over the lifetime of a given query. You can use the `EXPLAIN` statement to understand how MySQL runs this query.

2. Review the slow query log to evaluate `rows_examined` and `rows_sent` to make
    sure that the workload is similar in each environment. For more information, see [Logging for Aurora MySQL databases](aurora-mysql-troubleshooting-logging.md).

3. Run the following command for tables that are part of the identified query:

```nohighlight

SHOW TABLE STATUS\G;
```

4. Capture the following outputs before and after running the query on each environment:

```nohighlight

SHOW GLOBAL STATUS;
```

5. Run the following commands on each environment to see if there are any other query/session influencing
    the performance of this sample query.

```nohighlight

SHOW FULL PROCESSLIST;

SHOW ENGINE INNODB STATUS\G;
```

Sometimes, when resources on the server are busy, it impacts every other operation on the server,
    including queries. You can also capture information periodically when queries are run or set up a
    `cron` job to capture information at useful intervals.

**Performance Schema**

The Performance Schema provides useful information about server runtime performance, while having minimal
impact on that performance. This is different from the `information_schema`, which provides schema
information about the DB instance. For more information, see [Overview of the Performance Schema for Performance Insights on Aurora MySQL](user-perfinsights-enablemysql.md).

**Query optimizer trace**

To understand why a particular [query plan was chosen for\
execution](https://dev.mysql.com/doc/refman/8.0/en/execution-plan-information.html), you can set up `optimizer_trace` to access the MySQL query optimizer.

Run an optimizer trace to show extensive information on all the paths available to the optimizer and its
choice.

```nohighlight

SET SESSION OPTIMIZER_TRACE="enabled=on";
SET optimizer_trace_offset=-5, optimizer_trace_limit=5;

-- Run your query.
SELECT * FROM table WHERE x = 1 AND y = 'A';

-- After the query completes:
SELECT * FROM information_schema.OPTIMIZER_TRACE;
SET SESSION OPTIMIZER_TRACE="enabled=off";
```

## Reviewing query optimizer settings

Aurora MySQL version 3 (compatible with MySQL 8.0 Community Edition) has many optimizer-related changes compared with
Aurora MySQL version 2 (compatible with MySQL 5.7 Community Edition). If you have some custom values for the
`optimizer_switch`, we recommend that you review the differences in the defaults and set
`optimizer_switch` values that work best for your workload. We also recommend that you test the options
available for Aurora MySQL version 3 to examine how your queries perform.

###### Note

Aurora MySQL version 3 uses the community default value of 20 for the [innodb\_stats\_persistent\_sample\_pages](https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html) parameter.

You can use the following command to show the `optimizer_switch` values:

```nohighlight

SELECT @@optimizer_switch\G;
```

The following table shows the default `optimizer_switch` values for Aurora MySQL versions 2 and 3.

SettingAurora MySQL version 2Aurora MySQL version 3`batched_key_access`offoff`block_nested_loop`onon`condition_fanout_filter`onon`derived_condition_pushdown`–on`derived_merge`onon`duplicateweedout`onon`engine_condition_pushdown`onon`firstmatch`onon`hash_join`offon`hash_join_cost_based`on–`hypergraph_optimizer`–off`index_condition_pushdown`onon`index_merge`onon`index_merge_intersection`onon`index_merge_sort_union`onon`index_merge_union`onon`loosescan`onon`materialization`onon`mrr`onon`mrr_cost_based`onon`prefer_ordering_index`onon`semijoin`onon`skip_scan`–on`subquery_materialization_cost_based`onon`subquery_to_derived`–off`use_index_extensions`onon`use_invisible_indexes`–off

For more information, see [Switchable\
optimizations (MySQL 5.7)](https://dev.mysql.com/doc/refman/5.7/en/switchable-optimizations.html) and [Switchable optimizations (MySQL 8.0)](https://dev.mysql.com/doc/refman/8.0/en/switchable-optimizations.html) in the MySQL documentation.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Troubleshooting database connection issues

Aurora MySQL reference

All content copied from https://docs.aws.amazon.com/.
