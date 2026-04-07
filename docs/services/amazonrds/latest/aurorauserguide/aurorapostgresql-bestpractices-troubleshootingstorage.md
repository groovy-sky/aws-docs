# Troubleshooting storage issues in Aurora PostgreSQL

If the amount of working memory needed for sort or index-creation operations exceeds
the amount allocated by the `work_mem` parameter, Aurora PostgreSQL writes the
excess data to temporary disk files. When it writes the data, Aurora PostgreSQL uses the
same storage space that it uses for storing error and message logs, that is,
_local storage_. Each instance in your Aurora PostgreSQL DB cluster
has an amount of local storage available. The amount of storage is based on its DB
instance class. To increase the amount of local storage, you need to modify the instance
to use a larger DB instance class. For DB instance class specifications, see [Hardware specifications for DB instance classesfor Aurora](concepts-dbinstanceclass-summary.md).

You can monitor your Aurora PostgreSQL DB cluster's local storage space by watching
the Amazon CloudWatch metric for `FreeLocalStorage`. This metric reports the amount of
storage available to each DB instance in the Aurora DB cluster for temporary tables and
logs. For more information, see [Monitoring Amazon Aurora metrics with Amazon CloudWatch](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/monitoring-cloudwatch.html).

Sorting, indexing, and grouping operations start in working memory but often must be
offloaded to local storage. If your Aurora PostgreSQL DB cluster runs out of local storage
because of these types of operations, you can resolve the issue by taking one of the
following actions.

- Increase the amount of working memory. This reduces the need to use local
storage. By default, PostgreSQL allocates 4 MB for each sort, group, and index
operation. To check the current working memory value for your Aurora PostgreSQL DB
cluster's writer instance, connect to the instance using `psql`
and run the following command.

```nohighlight

postgres=> SHOW work_mem;
work_mem
  ----------
4MB
(1 row)
```

You can increase the working memory at the session level before sort, group,
and other operations, as follows.

```nohighlight

SET work_mem TO '1 GB';
```

For more information about working memory, see [Resource Consumption](https://www.postgresql.org/docs/current/runtime-config-resource.html) in the PostgreSQL documentation.

- Change the log retention period so that logs are stored for shorter
timeframes. To learn how, see [Aurora PostgreSQL database log files](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_LogAccess.Concepts.PostgreSQL.html).

For Aurora PostgreSQL DB clusters larger than 40 TB, don't use db.t2, db.t3, or
db.t4g instance classes. We recommend using the T DB instance classes only for
development and test servers, or other non-production servers. For more information, see
[DB instance class types](concepts-dbinstanceclass-types.md).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Best Practices for Parallel Queries in Aurora PostgreSQL

Replication with Aurora PostgreSQL
