# Configuring parameter settings for the pgactive extension

You can use the following query to view all the parameters associated with
`pgactive` extension.

```nohighlight

app=> SELECT * FROM pg_settings WHERE name LIKE 'pgactive.%';
```

You can configure the `pgactive` extension using various parameters. These
parameters can be set through either the AWS Management Console or the AWS CLI interface.

## Main pgactive extension parameters

The following table provides a reference for the main parameters of the
`pgactive` extension:

Parameter

Unit

Default

Description

`pgactive.conflict_logging_include_tuples`

`boolean`

–

Logs complete tuple information for the `pgactive`
extension.

###### Note

A server restart is required for changes to take effect.

`pgactive.log_conflicts_to_table`

`boolean`

–

Determines whether the `pgactive` extension logs the detected
conflicts to the `pgactive.pgactive_conflict_history` table. For more
information, see Conflict logging for details.

###### Note

A server restart is required for changes to take effect.

`pgactive.log_conflicts_to_logfile`

`boolean`

–

Determines whether the `pgactive` extension logs the detected
conflicts to the PostgreSQL log file. For more information, see Conflict logging
for details.

###### Note

A server restart is required for changes to take effect.

`pgactive.synchronous_commit`

`boolean`

off

Determines the commit behavior for pgactive apply workers. When
disabled(off), apply workers perform asynchronous commits, which improves
PostgreSQL throughput during apply operations but delays replay confirmations to
the upstream. Setting it to `off` is always safe and won't cause
transaction loss or skipping. This setting only affects the timing of disk flushes
on the downstream node and when confirmations are sent upstream. The system delays
sending replay flush confirmations until commits are flushed to disk through
unrelated operations like checkpoints or periodic work. However, if the upstream
has the downstream listed in `synchronous_standby_names`, setting it to
`off` causes synchronous commits on the upstream to take longer to
report success to the client. In this case, set the parameter to
`on`.

###### Note

Even when this parameter is set to `on` with nodes listed in
`synchronous_standby_names`, replication conflicts can still occur
in active-active configurations. This is because the system lacks inter-node
locking and global snapshot management, allowing concurrent transactions on
different nodes to modify the same tuple. Additionally, transactions only begin
replication after committing on the upstream node. Enabling synchronous commit
doesn't transform the pgactive extension into an always-consistent
system.

`pgactive.temp_dump_directory`

`string`

–

Defines the temporary storage path required for database cloning
operations during initial setup. This directory must be writable by the postgres
user and have sufficient storage space to contain a complete database dump. The
system uses this location only during initial database setup with logical copy
operations. This parameter isn't used by the `pgactive_init_copy
                    command`.

`pgactive.max_ddl_lock_delay`

`milliseconds`

`-1`

Specifies the maximum wait time for DDL lock before forcibly aborting
concurrent write transactions. The default value is `-1`, which adopts
the value set in `max_standby_streaming_delay`. This parameter accepts
time units. For example, you can set it to 10s for 10 seconds. During this wait
period, the system attempts to acquire DDL locks while waiting for ongoing write
transactions to either commit or roll back. For more information, see the DDL
Locking.

`pgactive.ddl_lock_timeout`

`milliseconds`

`-1`

Specifies how long a DDL lock attempt waits to obtain the lock. The
default value is `-1`, which uses the value specified in lock\_timeout.
You can set this parameter using time units such as 10s for 10 seconds. This timer
only controls the waiting period for obtaining a DDL lock. Once the system obtains
the lock and begins the DDL operation, the timer stops. This parameter doesn't
limit the total duration a DDL lock can be held or the overall DDL operation time.
To control the total duration of the operation, use `statement_timeout`
instead. For more information, see DDL Locking.

`pgactive.debug_trace_ddl_locks_level`

`boolean`

–

Overrides the default debug log level for DDL locking operations in the
`pgactive` extension. When configured, this setting causes DDL
lock-related messages to be emitted at the LOG debug level instead of their
default level. Use this parameter to monitor DDL locking activity without enabling
the verbose `DEBUG1` or `DEBUG2` log levels across your
entire server.

Available log levels, in increasing order of verbosity:

- **none** \- DDL lock messages appear only at
DEBUG1 and lower server log levels.

- **statement** \- Adds LOG output for DDL lock
acquisition attempts.

- **acquire\_release** \- Records lock acquisition,
release, declination events, and peer node applications of remote DDL
locks.

- **peers** \- Provides additional details about DDL
lock negotiations between peer nodes.

- **debug** \- Logs all DDL lock-related activities
at LOG level.

For more information about monitoring options, see Monitoring global DDL
locks.

###### Note

Changes to this setting take effect when you reload the configuration. You
don't need to restart the server.

## Additional pgactive extension parameters

The following table presents less frequently used and internal configuration options
available for the `pgactive` extension.

Parameter

Unit

Default

Description

`pgactive.debug_apply_delay`

`integer`

–

Sets an apply delay (in milliseconds) for configured connections that don't
have an explicit apply delay in their `pgactive.pgactive_connections`
entry. This delay is set during node creation or join time, and pgactive won't
replay a transaction on peer nodes until at least the specified number of
milliseconds have elapsed since it was committed.

Primarily used to simulate high-latency networks in testing environments to
make it easier to create conflicts. For example, with a 500ms delay on nodes A and
B, you have at least 500ms to perform a conflicting insert on node B after
inserting a value on node A.

###### Note

Requires a server reload or restart of apply workers to take effect.

`pgactive.connectability_check_duration`

`integer`

–

Specifies the duration (in seconds) that a database worker attempts to
establish connections during failed attempts. The worker makes one connection
attempt per second until it succeeds or reaches this timeout value. This setting
is useful when the database engine starts before the worker is ready to establish
connections.

`pgactive.skip_ddl_replication`

`boolean`

`on`

Controls how DDL changes are replicated or handled in Amazon RDS with
`pgactive` enabled. When set to `on`, the node processes
DDL changes like a non-pgcctive node. The following requirements apply when
working with this parameter:

- New nodes can't join a pgactive group if their
`skip_ddl_replication` value differs from the upstream
node.

- Existing nodes can't start pgactive workers if their parameter value
doesn't match the upstream node.

- All pgactive members must use the same parameter value.

You can modify this parameter in two ways with super user privileges:
globally, locally (session level).

###### Note

Changing this parameter incorrectly can break your replication
setups.

`pgactive.do_not_replicate`

`boolean`

–

This parameter is for internal use only. When you set this parameter in a
transaction, the changes are not replicated to other nodes in your DB cluster.

###### Note

Changing this parameter incorrectly can break your replication
setups.

`pgactive.discard_mismatched_row_attributes`

`boolean`

–

This parameter is intended for specialist use only. We recommend using
this parameter only when troubleshooting specific replication issues. Use this
parameter when:

- The incoming replication stream contains rows with more columns than your
local table.

- These remote rows contain non-null values.

This setting overrides the following error message and allows data divergence
to arise to let replication continue: `cannot right-pad mismatched
                    attributes; attno %u is missing in local table and remote row has non-null,
                    non-dropped value for this attribute`

###### Note

Changing this parameter incorrectly can break your replication
setups.

`pgactive.debug_trace_replay`

`boolean`

–

When set to `on`, it emits a log message for each remote
action that downstream apply workers process. The logs include:

- Change type

- Affected table name

- Number of changes since transaction start

- Transaction commit LSN

- Commit timestamp

- Upstream node identifier

- Forwarding node identifier (if applicable)

The logs also capture queued DDL commands and table drops.

para>

By default, the logs do not include row field contents. To include row
values in the logs, you must recompile with the following flags enabled:

- VERBOSE\_INSERT

- VERBOSE\_UPDATE

- VERBOSE\_DELETE

###### Note

Enabling this logging setting can impact performance. We recommend enabling
it only when needed for troubleshooting. Changes to this setting take effect
when you reload the configuration. You don't need to restart the server.

`pgactive.extra_apply_connection_options`

–

You can configure connection parameters for all peer node connections
with pgactive nodes. These parameters control settings such as keepalives and SSL
modes. By default, pgactive uses the following connection parameters:

- connect\_timeout=30

- keepalives=1

- keepalives\_idle=20

- keepalives\_interval=20

- keepalives\_count=5

To override the default parameters, use the following similar command:

`pgactive.extra_apply_connection_options = 'keepalives=0'`

Individual node connection strings take precedence over both these settings
and pgactive's built-in connection options. For more information about connection
string formats, see [libpq connection strings](https://www.postgresql.org/docs/current/libpq-connect.html).

We recommend keeping the default keepalive settings enabled. Only disable
keepalives if you experience issues with large transactions completing over
unreliable networks.

###### Note

We recommend keeping the default keepalive settings enabled. Only disable
keepalives if you experience issues with large transactions completing over
unreliable networks. Changes to this setting take effect when you reload the
configuration. You don't need to restart the server.

`pgactive.init_node_parallel_jobs` ( `int`)

–

Specifies the number of parallel jobs that `pg_dump` and
`pg_restore` can use during logical node joins with the
`pgactive.pgactive_join_group` function.

Changes to this setting take effect when you reload the configuration. You
don't need to restart the server.

`pgactive.max_nodes`

`int`

4

Specifies the maximum number of nodes allowed in a pgactive extension group.
The default value is 4 nodes. You must consider the following when setting the
value of this parameter:

- All nodes in a pgactive extension group must use the same parameter
value.

- A new node can't join if its parameter value differs from the upstream
node.

- Existing nodes can't start pgactive extension workers if their parameter
value differs from the upstream node.

- Larger groups require additional monitoring and maintenance effort, so the
value of this parameter wisely.

You can set this parameter in two ways: in the configuration file, using the
`ALTER SYSTEM SET` command

Default value for this parameter is `4`, meaning, there can be
maximum of 4 nodes allowed in the `pgactive` extension group at any
point of time.

###### Note

The change takes effect after you restart the server.

`pgactive.permit_node_identifier_getter_function_creation`

`boolean`

–

This parameter is intended for internal use only. When enabled,
`pgactive` extension allows creation of pgactive node identifier
getter function.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Measuring
replication lag among pgactive members

Understanding active-active conflicts

All content copied from https://docs.aws.amazon.com/.
