# Working with parameters on your RDS for PostgreSQL DB instance

In some cases, you might create an RDS for PostgreSQL DB instance without specifying a custom
parameter group. If so, your DB instance is created using the default parameter group for the
version of PostgreSQL that you choose. For example, suppose that you create an RDS for PostgreSQL
DB instance using PostgreSQL 13.3. In this case, the DB instance is created using the values
in the parameter group for PostgreSQL 13 releases, `default.postgres13`.

You can also create your own custom DB parameter group. You need to do this if you want to
modify any settings for the RDS for PostgreSQL DB instance from their default values. To learn how,
see [Parameter groups for Amazon RDS](user-workingwithparamgroups.md).

You can track the settings on your RDS for PostgreSQL DB instance in several different ways.
You can use the AWS Management Console, the AWS CLI, or the Amazon RDS API. You can also query the values from the
PostgreSQL `pg_settings` table of your instance, as shown following.

```sql

SELECT name, setting, boot_val, reset_val, unit
 FROM pg_settings
 ORDER BY name;
```

To learn more about the values returned from this query, see [`pg_settings`](https://www.postgresql.org/docs/current/view-pg-settings.html) in the PostgreSQL documentation.

Be especially careful when changing the settings for `max_connections` and
`shared_buffers` on your RDS for PostgreSQL DB instance. For example, suppose that you
modify settings for `max_connections` or `shared_buffers` and you use
values that are too high for your actual workload. In this case, your RDS for PostgreSQL DB
instance won't start. If this happens, you see an error such as the following in the
`postgres.log`.

```nohighlight

2018-09-18 21:13:15 UTC::@:[8097]:FATAL:  could not map anonymous shared memory: Cannot allocate memory
2018-09-18 21:13:15 UTC::@:[8097]:HINT:  This error usually means that PostgreSQL's request for a shared memory segment
exceeded available memory or swap space. To reduce the request size (currently 3514134274048 bytes), reduce
PostgreSQL's shared memory usage, perhaps by reducing shared_buffers or max_connections.
```

However, you can't change any values of the settings contained in the default
RDS for PostgreSQL DB parameter groups. To change settings for any parameters, first create a
custom DB parameter group. Then change the settings in that custom group, and then apply the
custom parameter group to your RDS for PostgreSQL DB instance. To learn more, see [Parameter groups for Amazon RDS](user-workingwithparamgroups.md).

There are two types of parameters in RDS for PostgreSQL.

- **Static parameters** – Static parameters require
that the RDS for PostgreSQL DB instance be rebooted after a change so that the new value can
take effect.

- **Dynamic parameters** – Dynamic parameters
don't require a reboot after changing their settings.

###### Note

If your RDS for PostgreSQL DB instance is using your own custom DB parameter group, you can
change the values of dynamic parameters on the running DB instance. You can do this by using
the AWS Management Console, the AWS CLI, or the Amazon RDS API.

If you have privileges to do so, you can also change parameter values by using the
`ALTER DATABASE`, `ALTER ROLE`, and `SET` commands.

## RDS for PostgreSQL DB instance parameter list

The following table lists some (but not all) parameters available in an RDS for PostgreSQL DB
instance. To view all available parameters, you use the [describe-db-parameters](https://docs.aws.amazon.com/cli/latest/reference/rds/describe-db-parameters.html) AWS CLI
command. For example, to get the list of all parameters available in the default parameter
group for RDS for PostgreSQL version 13, run the following.

```nohighlight

aws rds describe-db-parameters --db-parameter-group-name default.postgres13
```

You can also use the Console. Choose **Parameter groups** from the
Amazon RDS menu, and then choose the parameter group from those available in your
AWS Region.

Parameter name

Apply\_Type

Description

`application_name`

DynamicSets the application name to be reported in statistics and logs.

`archive_command`

DynamicSets the shell command that will be called to archive a WAL file.

`array_nulls`

DynamicEnables input of NULL elements in arrays.

`authentication_timeout`

DynamicSets the maximum allowed time to complete client authentication.

`autovacuum`

DynamicStarts the autovacuum subprocess.

`autovacuum_analyze_scale_factor`

DynamicNumber of tuple inserts, updates, or deletes before analyze as a fraction of
reltuples.

`autovacuum_analyze_threshold`

DynamicMinimum number of tuple inserts, updates, or deletes before analyze.

`autovacuum_freeze_max_age`

StaticAge at which to autovacuum a table to prevent transaction ID wraparound.

`autovacuum_naptime`

DynamicTime to sleep between autovacuum runs.

`autovacuum_max_workers`

StaticSets the maximum number of simultaneously running autovacuum worker
processes.

`autovacuum_vacuum_cost_delay`

DynamicVacuum cost delay, in milliseconds, for autovacuum.

`autovacuum_vacuum_cost_limit`

DynamicVacuum cost amount available before napping, for autovacuum.

`autovacuum_vacuum_scale_factor`

DynamicNumber of tuple updates or deletes before vacuum as a fraction of
reltuples.

`autovacuum_vacuum_threshold`

DynamicMinimum number of tuple updates or deletes before vacuum.

`backslash_quote`

DynamicSets whether a backslash (\\) is allowed in string literals.

`bgwriter_delay`

DynamicBackground writer sleep time between rounds.

`bgwriter_lru_maxpages`

DynamicBackground writer maximum number of LRU pages to flush per round.

`bgwriter_lru_multiplier`

DynamicMultiple of the average buffer usage to free per round.

`bytea_output`

DynamicSets the output format for bytes.

`check_function_bodies`

DynamicChecks function bodies during CREATE FUNCTION.

`checkpoint_completion_target`

DynamicTime spent flushing dirty buffers during checkpoint, as a fraction of the
checkpoint interval.

`checkpoint_segments`

DynamicSets the maximum distance in log segments between automatic write-ahead log
(WAL) checkpoints.

`checkpoint_timeout`

DynamicSets the maximum time between automatic WAL checkpoints.

`checkpoint_warning`

DynamicEnables warnings if checkpoint segments are filled more frequently than
this.

`client_connection_check_interval`

Dynamic Sets the time interval between checks for disconnection while running
queries.

`client_encoding`

DynamicSets the client's character set encoding.

`client_min_messages`

DynamicSets the message levels that are sent to the client.

`commit_delay`

DynamicSets the delay in microseconds between transaction commit and flushing WAL to
disk.

`commit_siblings`

DynamicSets the minimum concurrent open transactions before performing
commit\_delay.

`constraint_exclusion`

DynamicEnables the planner to use constraints to optimize queries.

`cpu_index_tuple_cost`

DynamicSets the planner's estimate of the cost of processing each index entry during
an index scan.

`cpu_operator_cost`

DynamicSets the planner's estimate of the cost of processing each operator or function
call.

`cpu_tuple_cost`

DynamicSets the planner's estimate of the cost of processing each tuple (row).

`cursor_tuple_fraction`

DynamicSets the planner's estimate of the fraction of a cursor's rows that will be
retrieved.

`datestyle`

DynamicSets the display format for date and time values.

`deadlock_timeout`

DynamicSets the time to wait on a lock before checking for deadlock.

`debug_pretty_print`

DynamicIndents parse and plan tree displays.

`debug_print_parse`

DynamicLogs each query's parse tree.

`debug_print_plan`

DynamicLogs each query's execution plan.

`debug_print_rewritten`

DynamicLogs each query's rewritten parse tree.

`default_statistics_target`

DynamicSets the default statistics target.

`default_tablespace`

DynamicSets the default tablespace to create tables and indexes in.

`default_transaction_deferrable`

DynamicSets the default deferrable status of new transactions.

`default_transaction_isolation`

DynamicSets the transaction isolation level of each new transaction.

`default_transaction_read_only`

DynamicSets the default read-only status of new transactions.

`default_with_oids`

DynamicCreates new tables with object IDs (OIDs) by default.

`effective_cache_size`

DynamicSets the planner's assumption about the size of the disk cache.

`effective_io_concurrency`

DynamicNumber of simultaneous requests that can be handled efficiently by the disk
subsystem.

`enable_bitmapscan`

DynamicEnables the planner's use of bitmap-scan plans.

`enable_hashagg`

DynamicEnables the planner's use of hashed aggregation plans.

`enable_hashjoin`

DynamicEnables the planner's use of hash join plans.

`enable_indexscan`

DynamicEnables the planner's use of index-scan plans.

`enable_material`

DynamicEnables the planner's use of materialization.

`enable_mergejoin`

DynamicEnables the planner's use of merge join plans.

`enable_nestloop`

DynamicEnables the planner's use of nested-loop join plans.

`enable_seqscan`

DynamicEnables the planner's use of sequential-scan plans.

`enable_sort`

DynamicEnables the planner's use of explicit sort steps.

`enable_tidscan`

DynamicEnables the planner's use of TID scan plans.

`escape_string_warning`

DynamicWarns about backslash (\\) escapes in ordinary string literals.

`extra_float_digits`

DynamicSets the number of digits displayed for floating-point values.

`from_collapse_limit`

DynamicSets the FROM-list size beyond which subqueries are not collapsed.

`fsync`

DynamicForces synchronization of updates to disk.

`full_page_writes`

DynamicWrites full pages to WAL when first modified after a checkpoint.

`geqo`

DynamicEnables genetic query optimization.

`geqo_effort`

DynamicGEQO: effort is used to set the default for other GEQO parameters.

`geqo_generations`

DynamicGEQO: number of iterations of the algorithm.

`geqo_pool_size`

DynamicGEQO: number of individuals in the population.

`geqo_seed`

DynamicGEQO: seed for random path selection.

`geqo_selection_bias`

DynamicGEQO: selective pressure within the population.

`geqo_threshold`

DynamicSets the threshold of FROM items beyond which GEQO is used.

`gin_fuzzy_search_limit`

DynamicSets the maximum allowed result for exact search by GIN.

`hot_standby_feedback`

DynamicDetermines whether a hot standby sends feedback messages to the primary or
upstream standby.

`intervalstyle`

DynamicSets the display format for interval values.

`join_collapse_limit`

DynamicSets the FROM-list size beyond which JOIN constructs are not flattened.

`lc_messages`

DynamicSets the language in which messages are displayed.

`lc_monetary`

DynamicSets the locale for formatting monetary amounts.

`lc_numeric`

DynamicSets the locale for formatting numbers.

`lc_time`

DynamicSets the locale for formatting date and time values.

`log_autovacuum_min_duration`

DynamicSets the minimum running time above which autovacuum actions will be
logged.

`log_checkpoints`

DynamicLogs each checkpoint.

`log_connections`

DynamicLogs each successful connection.

`log_disconnections`

DynamicLogs end of a session, including duration.

`log_duration`

DynamicLogs the duration of each completed SQL statement.

`log_error_verbosity`

DynamicSets the verbosity of logged messages.

`log_executor_stats`

DynamicWrites executor performance statistics to the server log.

`log_filename`

DynamicSets the file name pattern for log files.

`log_file_mode`

DynamicSets file permissions for log files. Default value is 0644.

`log_hostname`

DynamicLogs the host name in the connection logs. As of PostgreSQL 12 and later
versions, this parameter is 'off' by default. When turned on, the
connection uses DNS reverse-lookup to get the hostname that gets captured to the
connection logs. If you turn on this parameter, you should monitor the impact that
it has on the time it takes to establish connections.

`log_line_prefix `

DynamicControls information prefixed to each log line.

`log_lock_waits`

DynamicLogs long lock waits.

`log_min_duration_statement`

DynamicSets the minimum running time above which statements will be logged.

`log_min_error_statement`

DynamicCauses all statements generating an error at or above this level to be
logged.

`log_min_messages`

DynamicSets the message levels that are logged.

`log_parser_stats`

DynamicWrites parser performance statistics to the server log.

`log_planner_stats`

DynamicWrites planner performance statistics to the server log.

`log_rotation_age`

DynamicAutomatic log file rotation will occur after N minutes.

`log_rotation_size`

DynamicAutomatic log file rotation will occur after N kilobytes.

`log_statement`

DynamicSets the type of statements logged.

`log_statement_stats`

DynamicWrites cumulative performance statistics to the server log.

`log_temp_files`

DynamicLogs the use of temporary files larger than this number of kilobytes.

`log_timezone`

DynamicSets the time zone to use in log messages.

`log_truncate_on_rotation`

DynamicTruncate existing log files of same name during log rotation.

`logging_collector`

StaticStart a subprocess to capture stderr output and/or csvlogs into log
files.

`maintenance_work_mem`

DynamicSets the maximum memory to be used for maintenance operations.

`max_connections`

StaticSets the maximum number of concurrent connections.

`max_files_per_process`

StaticSets the maximum number of simultaneously open files for each server
process.

`max_locks_per_transaction`

StaticSets the maximum number of locks per transaction.

`max_pred_locks_per_transaction`

StaticSets the maximum number of predicate locks per transaction.

`max_prepared_transactions`

StaticSets the maximum number of simultaneously prepared transactions.

`max_stack_depth`

DynamicSets the maximum stack depth, in kilobytes.

`max_standby_archive_delay`

DynamicSets the maximum delay before canceling queries when a hot standby server is
processing archived WAL data.

`max_standby_streaming_delay`

DynamicSets the maximum delay before canceling queries when a hot standby server is
processing streamed WAL data.`max_wal_size`DynamicSets the WAL size (MB) that triggers a checkpoint.

- For RDS for PostgreSQL 15 and earlier versions, it is 2 GB.

- For RDS for PostgreSQL 16 and later versions, RDS for PostgreSQL automatically
configures based on the allocated storage size of your instance:

- 6 GB for instances with allocated storage 100 GB and more.

- 2 GB for instances with allocated storage less than 100 GB.

Use the following command on your Amazon RDS for PostgreSQL DB instance to see
its current value:

```

SHOW max_wal_size;
```

`min_wal_size`DynamicSets the minimum size to shrink the WAL to. For PostgreSQL version 9.6 and
earlier, `min_wal_size` is in units of 16 MB. For PostgreSQL version 10
and later, `min_wal_size` is in units of 1 MB.

`quote_all_identifiers`

DynamicAdds quotes (") to all identifiers when generating SQL fragments.

`random_page_cost`

DynamicSets the planner's estimate of the cost of a non-sequentially fetched disk
page. This parameter has no value unless query plan management (QPM) is turned on.
When QPM is on, the default value for this parameter `4`. `rds.adaptive_autovacuum`DynamicAutomatically tunes the autovacuum parameters whenever the transaction ID
thresholds are exceeded.`rds.force_ssl`DynamicRequires the use of SSL connections. The default value is set to 1 (on) for
RDS for PostgreSQL version 15. All other RDS for PostgreSQL major version 14 and older have
the default value set to 0 (off).

`rds.local_volume_spill_enabled`

StaticEnables writing logical spill files to the local volume.

`rds.log_retention_period`

DynamicSets log retention such that Amazon RDS deletes PostgreSQL logs that are older than
_n_ minutes.`rds.rds_superuser_reserved_connections`Static

Sets the number of connection slots reserved for rds\_superusers. This
parameter is only available in versions 15 and earlier. For more information, see
the PostgreSQL documentation [reserved\_connections](https://www.postgresql.org/docs/current/runtime-config-connection.html).

`rds.replica_identity_full`

Dynamic

When you set this parameter to `on`, it overrides the replica
identity setting to `FULL` for all database tables. This means all
column values are written to the write ahead log (WAL), regardless of your
`REPLICA IDENTITY FULL` settings.

###### Note

Turning on this parameter may increase your database instance IOPS due to
the additional WAL logging.

`rds.restrict_password_commands`StaticRestricts who can manage passwords to users with the `rds_password`
role. Set this parameter to 1 to enable password restriction. The default is
0.

`search_path`

DynamicSets the schema search order for names that are not schema-qualified.

`seq_page_cost`

DynamicSets the planner's estimate of the cost of a sequentially fetched disk
page.

`session_replication_role`

DynamicSets the sessions behavior for triggers and rewrite rules.

`shared_buffers`

StaticSets the number of shared memory buffers used by the server.

`shared_preload_libraries `

StaticLists the shared libraries to preload into the RDS for PostgreSQL DB instance.
Supported values include auto\_explain, orafce, pgaudit, pglogical, pg\_bigm, pg\_cron,
pg\_hint\_plan, pg\_prewarm, pg\_similarity, pg\_stat\_statements, pg\_tle, pg\_transport,
plprofiler, and plrust.

`ssl`

DynamicEnables SSL connections.

`sql_inheritance`

DynamicCauses subtables to be included by default in various commands.

`ssl_renegotiation_limit`

DynamicSets the amount of traffic to send and receive before renegotiating the
encryption keys.

`standard_conforming_strings`

DynamicCauses ... strings to treat backslashes literally.

`statement_timeout`

DynamicSets the maximum allowed duration of any statement.

`synchronize_seqscans`

DynamicEnables synchronized sequential scans.

`synchronous_commit`

DynamicSets the current transactions synchronization level.

`tcp_keepalives_count`

DynamicMaximum number of TCP keepalive retransmits.

`tcp_keepalives_idle`

DynamicTime between issuing TCP keepalives.

`tcp_keepalives_interval`

DynamicTime between TCP keepalive retransmits.

`temp_buffers`

DynamicSets the maximum number of temporary buffers used by each session.`temp_file_limit`DynamicSets the maximum size in KB to which the temporary files can grow.

`temp_tablespaces`

DynamicSets the tablespaces to use for temporary tables and sort files.

`timezone`

Dynamic

Sets the time zone for displaying and interpreting time stamps.

The Internet Assigned Numbers Authority (IANA) publishes new time zones at
[https://www.iana.org/time-zones](https://www.iana.org/time-zones) several times a year. Every time RDS
releases a new minor maintenance release of PostgreSQL, it ships with the latest
time zone data at the time of the release. When you use the latest RDS for PostgreSQL
versions, you have recent time zone data from RDS. To ensure that your DB instance
has recent time zone data, we recommend upgrading to a higher DB engine version.
You can't modify the time zone tables in PostgreSQL DB instances manually. RDS
doesn't modify or reset the time zone data of running DB instances. New time zone
data is installed only when you perform a database engine version
upgrade.

`track_activities`

DynamicCollects information about running commands.

`track_activity_query_size`

StaticSets the size reserved for pg\_stat\_activity.current\_query, in bytes.

`track_counts`

DynamicCollects statistics on database activity.

`track_functions`

DynamicCollects function-level statistics on database activity.

`track_io_timing`

DynamicCollects timing statistics on database I/O activity.

`transaction_deferrable`

DynamicIndicates whether to defer a read-only serializable transaction until it can be
started with no possible serialization failures.

`transaction_isolation`

DynamicSets the current transactions isolation level.

`transaction_read_only`

DynamicSets the current transactions read-only status.

`transform_null_equals`

DynamicTreats expr=NULL as expr IS NULL.

`update_process_title`

DynamicUpdates the process title to show the active SQL command.

`vacuum_cost_delay`

DynamicVacuum cost delay in milliseconds.

`vacuum_cost_limit`

DynamicVacuum cost amount available before napping.

`vacuum_cost_page_dirty`

DynamicVacuum cost for a page dirtied by vacuum.

`vacuum_cost_page_hit`

DynamicVacuum cost for a page found in the buffer cache.

`vacuum_cost_page_miss`

DynamicVacuum cost for a page not found in the buffer cache.

`vacuum_defer_cleanup_age`

DynamicNumber of transactions by which vacuum and hot cleanup should be deferred, if
any.

`vacuum_freeze_min_age`

DynamicMinimum age at which vacuum should freeze a table row.

`vacuum_freeze_table_age`

DynamicAge at which vacuum should scan a whole table to freeze tuples.

`wal_buffers`

StaticSets the number of disk-page buffers in shared memory for WAL.

`wal_writer_delay`

DynamicWAL writer sleep time between WAL flushes.

`work_mem`

DynamicSets the maximum memory to be used for query workspaces.

`xmlbinary`

DynamicSets how binary values are to be encoded in XML.

`xmloption`

DynamicSets whether XML data in implicit parsing and serialization operations is to be
considered as documents or content fragments.

Amazon RDS uses the default PostgreSQL units for all parameters. The following table shows
the PostgreSQL default unit for each parameter.

Parameter name

Unit

`archive_timeout`

s

`authentication_timeout`

s

`autovacuum_naptime`

s

`autovacuum_vacuum_cost_delay`

ms

`bgwriter_delay`

ms

`checkpoint_timeout`

s

`checkpoint_warning`

s

`deadlock_timeout`

ms

`effective_cache_size`

8 KB

`lock_timeout`

ms

`log_autovacuum_min_duration`

ms

`log_min_duration_statement`

ms

`log_rotation_age`

minutes

`log_rotation_size`

KB

`log_temp_files`

KB

`maintenance_work_mem`

KB

`max_stack_depth`

KB

`max_standby_archive_delay`

ms

`max_standby_streaming_delay`

ms

`post_auth_delay`

s

`pre_auth_delay`

s

`segment_size`

8 KB

`shared_buffers`

8 KB

`statement_timeout`

ms

`ssl_renegotiation_limit`

KB

`tcp_keepalives_idle`

s

`tcp_keepalives_interval`

s

`temp_file_limit`

KB

`work_mem`

KB

`temp_buffers`

8 KB

`vacuum_cost_delay`

ms

`wal_buffers`

8 KB

`wal_receiver_timeout`

ms

`wal_segment_size`

B

`wal_sender_timeout`

ms

`wal_writer_delay`

ms

`wal_receiver_status_interval`

s

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Best Practices for Parallel Queries in RDS for PostgreSQL

Tuning with wait events for RDS for PostgreSQL
