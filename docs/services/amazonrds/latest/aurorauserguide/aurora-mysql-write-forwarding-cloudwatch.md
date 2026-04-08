# Amazon CloudWatch metrics and Aurora MySQL status variables for write forwarding

The following Amazon CloudWatch metrics and Aurora MySQL status variables apply when you use write forwarding for Aurora MySQL. For more information about metrics for Aurora MySQL writer and reader DB instances, see the following topics.

###### Topics

- [Metrics for write forwarding for Aurora MySQL writer DB instances](#aurora-mysql-write-forwarding-cloudwatch-writer-metrics)

- [Metrics for write forwarding for Aurora MySQL reader DB instances](#aurora-mysql-write-forwarding-cloudwatch-reader-metrics)

## Metrics for write forwarding for Aurora MySQL writer DB instances

The following Amazon CloudWatch metrics apply when you use write forwarding on one or more DB clusters. These metrics are all measured on the writer DB
instance.

CloudWatch metricUnitDescription

`ForwardingWriterDMLLatency`

Milliseconds

Average time to process each forwarded DML statement on the writer DB instance.

It doesn't include the time for the DB cluster to forward the write request, or the time to replicate changes back to
the writer.

`ForwardingWriterDMLThroughput`

Count per secondNumber of forwarded DML statements processed each second by this writer DB instance.

`ForwardingWriterOpenSessions`

CountNumber of forwarded sessions on the writer DB instance.

The following Aurora MySQL status variables apply when you use write forwarding on one or more DB clusters. These status variables are all
measured on the writer DB instance.

Aurora MySQL status variableUnitDescription`Aurora_fwd_writer_dml_stmt_count`CountTotal number of DML statements forwarded to this writer DB instance.`Aurora_fwd_writer_dml_stmt_duration`MicrosecondsTotal duration of DML statements forwarded to this writer DB instance.`Aurora_fwd_writer_open_sessions`CountNumber of forwarded sessions on the writer DB instance.`Aurora_fwd_writer_select_stmt_count`CountTotal number of `SELECT` statements forwarded to this writer DB instance.`Aurora_fwd_writer_select_stmt_duration`MicrosecondsTotal duration of `SELECT` statements forwarded to this writer DB instance.

## Metrics for write forwarding for Aurora MySQL reader DB instances

The following CloudWatch metrics are measured on each reader DB instance in a DB cluster with write forwarding enabled.

CloudWatch metricUnitDescription

`ForwardingReplicaDMLLatency`

MillisecondsAverage response time of forwarded DMLs on replica.

`ForwardingReplicaDMLThroughput`

Count per secondNumber of forwarded DML statements processed each second.

`ForwardingReplicaOpenSessions`

CountNumber of sessions that are using write forwarding on a reader DB instance.

`ForwardingReplicaReadWaitLatency`

Milliseconds

Average wait time that a `SELECT` statement on a reader DB instance waits to catch up to the writer.

The degree to which the reader DB instance waits before processing a query depends on the
`aurora_replica_read_consistency` setting.

`ForwardingReplicaReadWaitThroughput`

Count per secondTotal number of `SELECT` statements processed each second in all sessions that are forwarding writes.

`ForwardingReplicaSelectLatency`

MillisecondsForwarded `SELECT` latency, averaged over all forwarded `SELECT` statements within the monitoring
period.

`ForwardingReplicaSelectThroughput`

Count per secondForwarded `SELECT` throughput per second average within the monitoring period.

The following Aurora MySQL status variables are measured on each reader DB instance in a DB cluster with write forwarding enabled.

Aurora MySQL status variableUnitDescription`Aurora_fwd_replica_dml_stmt_count`CountTotal number of DML statements forwarded from this reader DB instance.`Aurora_fwd_replica_dml_stmt_duration`MicrosecondsTotal duration of all DML statements forwarded from this reader DB instance.`Aurora_fwd_replica_errors_session_limit`Count

Number of sessions rejected by the primary cluster due to one of the following error conditions:

- **`writer full`**

- **`Too many forwarded statements in progress`**.

`Aurora_fwd_replica_open_sessions`CountNumber of sessions that are using write forwarding on a reader DB instance.`Aurora_fwd_replica_read_wait_count`CountTotal number of read-after-write waits on this reader DB instance.`Aurora_fwd_replica_read_wait_duration`MicrosecondsTotal duration of waits due to the read consistency setting on this reader DB instance.`Aurora_fwd_replica_select_stmt_count`CountTotal number of `SELECT` statements forwarded from this reader DB instance.`Aurora_fwd_replica_select_stmt_duration`MicrosecondsTotal duration of `SELECT` statements forwarded from this reader DB instance.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Read consistency

Integrating Aurora MySQL with AWS services

All content copied from https://docs.aws.amazon.com/.
