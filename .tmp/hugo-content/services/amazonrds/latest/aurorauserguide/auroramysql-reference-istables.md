# Aurora MySQL–specific information\_schema tables

Aurora MySQL has certain `information_schema` tables that are specific to Aurora.

## information\_schema.aurora\_global\_db\_instance\_status

The `information_schema.aurora_global_db_instance_status` table contains information about the status of all DB instances in a global database's primary and secondary DB clusters.
The following table shows the columns that you can use. The remaining columns are for Aurora internal use only.

###### Note

This information schema table is only available with Aurora MySQL version 3.04.0 and higher
global databases.

ColumnData typeDescriptionSERVER\_IDvarchar(100)The identifier of the DB instance.SESSION\_IDvarchar(100)A unique identifier for the current session. A value of `MASTER_SESSION_ID` identifies the Writer (primary) DB instance.AWS\_REGIONvarchar(100)The AWS Region in which this global database instance runs. For a list of Regions, see [Region availability](concepts-regionsandavailabilityzones.md#Aurora.Overview.Availability).DURABLE\_LSNbigint unsignedThe log sequence number (LSN) made durable in storage.
A log sequence number (LSN) is a unique sequential number that identifies a record in the database transaction log.
LSNs are ordered such that a larger LSN represents a later transaction.HIGHEST\_LSN\_RCVDbigint unsignedThe highest LSN received by the DB instance from the writer DB instance.OLDEST\_READ\_VIEW\_TRX\_IDbigint unsignedThe ID of the oldest transaction that the writer DB instance can purge to.OLDEST\_READ\_VIEW\_LSNbigint unsignedThe oldest LSN used by the DB instance to read from storage.VISIBILITY\_LAG\_IN\_MSECfloat(10,0) unsignedFor readers in the primary DB cluster, how far this DB instance is lagging behind the writer DB instance in milliseconds.
For readers in a secondary DB cluster, how far this DB instance is lagging behind the secondary volume in milliseconds.

## information\_schema.aurora\_global\_db\_status

The `information_schema.aurora_global_db_status` table contains information about various aspects of Aurora global database lag, specifically, lag
of the underlying Aurora storage (so called durability lag) and lag between the recovery point objective (RPO).
The following table shows the columns that you can use. The remaining columns are for Aurora internal use only.

###### Note

This information schema table is only available with Aurora MySQL version 3.04.0 and higher
global databases.

ColumnData typeDescriptionAWS\_REGIONvarchar(100)The AWS Region in which this global database instance runs. For a list of Regions, see [Region availability](concepts-regionsandavailabilityzones.md#Aurora.Overview.Availability).HIGHEST\_LSN\_WRITTENbigint unsignedThe highest log sequence number (LSN) that currently exists
on this DB cluster. A log sequence number (LSN) is a unique sequential number that identifies a record in the database transaction log.
LSNs are ordered such that a larger LSN represents a later transaction.DURABILITY\_LAG\_IN\_MILLISECONDSfloat(10,0) unsignedThe difference in the timestamp values between the `HIGHEST_LSN_WRITTEN` on a
secondary DB cluster and the `HIGHEST_LSN_WRITTEN` on the primary DB cluster. This value is always 0 on the primary DB cluster of the Aurora global database.RPO\_LAG\_IN\_MILLISECONDSfloat(10,0) unsigned

The recovery point objective (RPO) lag. The RPO lag is the time
it takes for the most recent user transaction COMMIT to be stored on a secondary DB cluster after it's been stored
on the primary DB cluster of the Aurora global database. This value is always 0 on the primary DB cluster of the Aurora global database.

In simple terms, this metric calculates the recovery point objective for each Aurora MySQL DB cluster in the Aurora global database, that is, how much data might be lost
if there were an outage. As with lag, RPO is measured in time.

LAST\_LAG\_CALCULATION\_TIMESTAMPdatetimeThe timestamp that specifies when values were last calculated for
`DURABILITY_LAG_IN_MILLISECONDS` and `RPO_LAG_IN_MILLISECONDS`. A time value such as `1970-01-01 00:00:00+00` means
this is the primary DB cluster.OLDEST\_READ\_VIEW\_TRX\_IDbigint unsignedThe ID of the oldest transaction that the writer DB instance can purge to.

## information\_schema.replica\_host\_status

The `information_schema.replica_host_status` table contains replication information. The columns that you can
use are shown in the following table. The remaining columns are for Aurora internal use only.

ColumnData typeDescriptionCPUdoubleThe CPU percentage usage of the replica host.IS\_CURRENTtinyintWhether the replica is current.LAST\_UPDATE\_TIMESTAMPdatetime(6)The time the last update occurred. Used to determine whether a record is stale.REPLICA\_LAG\_IN\_MILLISECONDSdoubleThe replica lag in milliseconds.SERVER\_IDvarchar(100)The ID of the database server.SESSION\_IDvarchar(100)The ID of the database session. Used to determine whether a DB instance is a writer or reader instance.

###### Note

When a replica instance falls behind, the information queried from its `information_schema.replica_host_status`
table might be outdated. In this situation, we recommend that you query from the writer instance instead.

While the `mysql.ro_replica_status` table has similar information, we don't recommend that you use
it.

## information\_schema.aurora\_forwarding\_processlist

The `information_schema.aurora_forwarding_processlist` table contains information about processes involved in
write forwarding.

The contents of this table are visible only on the writer DB instance for a DB cluster with global or in-cluster write
forwarding turned on. An empty result set is returned on reader DB instances.

FieldData typeDescriptionIDbigintThe identifier of the connection on the writer DB instance. This
identifier is the same value displayed in the `Id` column
of the `SHOW PROCESSLIST` statement and returned by the
`CONNECTION_ID()` function within the thread.USERvarchar(32)The MySQL user that issued the statement.HOSTvarchar(255)The MySQL client that issued the statement. For forwarded
statements, this field shows the application client host address
that established the connection on the forwarding reader DB
instance.DBvarchar(64)The default database for the thread.COMMANDvarchar(16)The type of command the thread is executing on behalf of the client, or `Sleep` if the session
is idle. For descriptions of thread commands, see the MySQL documentation on [Thread Command Values](https://dev.mysql.com/doc/refman/8.0/en/thread-commands.html) in the
MySQL documentation.TIMEintThe time in seconds that the thread has been in its current state.STATEvarchar(64)An action, event, or state that indicates what the thread is doing. For descriptions of state values, see
[General Thread\
States](https://dev.mysql.com/doc/refman/8.0/en/general-thread-states.html) in the MySQL documentation.INFOlongtextThe statement that the thread is executing, or `NULL` if it isn't executing a statement. The
statement might be the one sent to the server, or an innermost statement if the statement executes other
statements.IS\_FORWARDEDbigintIndicates whether the thread is forwarded from a reader DB instance.REPLICA\_SESSION\_IDbigintThe connection identifier on the Aurora Replica. This identifier
is the same value displayed in the `Id` column of the
`SHOW PROCESSLIST` statement on the forwarding Aurora
reader DB instance.REPLICA\_INSTANCE\_IDENTIFIERvarchar(64)The DB instance identifier of the forwarding thread.REPLICA\_CLUSTER\_NAMEvarchar(64)The DB cluster identifier of the forwarding thread. For
in-cluster write forwarding, this identifier is the same DB cluster
as the writer DB instance.REPLICA\_REGIONvarchar(64)The AWS Region from which the forwarding thread originates. For
in-cluster write forwarding, this Region is the same
AWS Region as the writer DB instance.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Setting and showing binary log configuration

Aurora MySQL updates

All content copied from https://docs.aws.amazon.com/.
