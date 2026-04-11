# Performance Insights counter metrics

Counter metrics are operating system and database performance metrics in the Performance
Insights dashboard. To help identify and analyze performance problems, you can correlate
counter metrics with DB load. You must append a statistic
function to the metric to get the metric values. For example, the supported functions for
`os.memory.active` metric are `.avg`, `.min`,
`.max`, `.sum`, and `.sample_count`.

The counter metrics are collected one time each minute. The OS metrics collection depends on whether
Enhanced Monitoring is turned on or off. If Enhanced Monitoring is turned off, the OS metrics are collected one time each minute. If Enhanced Monitoring is turned on, the OS metrics
are collected for the selected time period. For more information about turning Enhanced Monitoring on or off, see
[Turning Enhanced Monitoring on and off](user-monitoring-os-enabling.md#USER_Monitoring.OS.Enabling.Procedure).

###### Topics

- [Performance Insights operating system counters](#USER_PerfInsights_Counters.OS)

- [Performance Insights counters for Aurora MySQL](#USER_PerfInsights_Counters.Aurora_MySQL)

- [Performance Insights counters for Aurora PostgreSQL](#USER_PerfInsights_Counters.Aurora_PostgreSQL)

## Performance Insights operating system counters

The following operating system counters, which are prefixed with `os`, are
available with Performance Insights for
Aurora PostgreSQL and Aurora MySQL.

You can use `ListAvailableResourceMetrics` API
for the list of available counter metrics for your DB instance. For more information, see
[ListAvailableResourceMetrics](../../../../reference/performance-insights/latest/apireference/api-listavailableresourcemetrics.md) in the Amazon RDS Performance Insights API
Reference guide.

CounterTypeUnitMetricDescriptionActiveMemoryKilobytesos.memory.activeThe amount of assigned memory, in kilobytes.BuffersMemoryKilobytesos.memory.buffersThe amount of memory used for buffering I/O requests prior to writing to the storage device, in kilobytes.CachedMemoryKilobytesos.memory.cachedThe amount of memory used for caching file system–based I/O, in kilobytes.DB CacheMemoryBytesos.memory.db.cache

The amount of memory used for page cache by database process
including tmpfs (shmem), in bytes.

DB Resident Set SizeMemoryBytesos.memory.db.residentSetSize

The amount of memory used for anonymous and swap cache by database
process not including tmpfs (shmem), in bytes.

DB SwapMemoryBytesos.memory.db.swap

The amount of memory used for swap by database process, in
bytes.

DirtyMemoryKilobytesos.memory.dirtyThe amount of memory pages in RAM that have been modified but not written to their related data block in storage, in kilobytes.FreeMemoryKilobytesos.memory.freeThe amount of unassigned memory, in kilobytes.Huge Pages FreeMemoryPagesos.memory.hugePagesFreeThe number of free huge pages. Huge pages are a feature of the Linux kernel.Huge Pages RsvdMemoryPagesos.memory.hugePagesRsvdThe number of committed huge pages.Huge Pages SizeMemoryKilobytesos.memory.hugePagesSizeThe size for each huge pages unit, in kilobytes.Huge Pages SurpMemoryPagesos.memory.hugePagesSurpThe number of available surplus huge pages over the total.Huge Pages TotalMemoryPagesos.memory.hugePagesTotalThe total number of huge pages.InactiveMemoryKilobytesos.memory.inactiveThe amount of least-frequently used memory pages, in kilobytes.MappedMemoryKilobytesos.memory.mappedThe total amount of file-system contents that is memory mapped inside a process address space, in kilobytes.Out of Memory Kill CountMemoryKillsos.memory.outOfMemoryKillCount

The number of OOM kills that happened over the last collection
interval.

Page TablesMemoryKilobytesos.memory.pageTablesThe amount of memory used by page tables, in kilobytes.SlabMemoryKilobytesos.memory.slabThe amount of reusable kernel data structures, in kilobytes.TotalMemoryKilobytesos.memory.totalThe total amount of memory, in kilobytes.WritebackMemoryKilobytesos.memory.writebackThe amount of dirty pages in RAM that are still being written to the backing storage, in kilobytes.GuestCpu UtilizationPercentageos.cpuUtilization.guestThe percentage of CPU in use by guest programs.IdleCpu UtilizationPercentageos.cpuUtilization.idleThe percentage of CPU that is idle.IrqCpu UtilizationPercentageos.cpuUtilization.irqThe percentage of CPU in use by software interrupts.NiceCpu UtilizationPercentageos.cpuUtilization.niceThe percentage of CPU in use by programs running at lowest priority.StealCpu UtilizationPercentageos.cpuUtilization.stealThe percentage of CPU in use by other virtual machines.SystemCpu UtilizationPercentageos.cpuUtilization.systemThe percentage of CPU in use by the kernel.TotalCpu UtilizationPercentageos.cpuUtilization.totalThe total percentage of the CPU in use. This value includes the nice value.UserCpu UtilizationPercentageos.cpuUtilization.userThe percentage of CPU in use by user programs.WaitCpu UtilizationPercentageos.cpuUtilization.waitThe percentage of CPU unused while waiting for I/O access.
Aurora Storage Aurora Storage
Bytes Rx
Disk IOBytes per secondos.diskIO.auroraStorage.auroraStorageBytesRx

The number of bytes received from Aurora storage per second.

Aurora Storage Aurora Storage
Bytes Tx
Disk IOBytes per secondos.diskIO.auroraStorage.auroraStorageBytesTx

The number of bytes uploaded to aurora storage per second.

Aurora Storage Disk Queue Depth

Disk IORequests

os.diskIO.auroraStorage.diskQueueDepth

The length of Aurora storage disk queue.

Aurora Storage Read IOs PS

Disk IORequests per second

os.diskIO.auroraStorage.readIOsPS

The number of read operations per second.

Aurora Storage Read Latency

Disk IOMilliseconds

os.diskIO.auroraStorage.readLatency

The average latency of a read I/O request to Aurora storage, in milliseconds.

Aurora Storage Read Throughput

Disk IOBytes per second

os.diskIO.auroraStorage.readThroughput

The amount of network throughput used by requests to the DB cluster, in bytes per second.

Aurora Storage Write IOs PS

Disk IORequests per second

os.diskIO.auroraStorage.writeIOsPS

The number of write operations per second.

Aurora Storage Write Latency

Disk IOMilliseconds

os.diskIO.auroraStorage.writeLatency

The average latency of a write I/O request to Aurora storage, in milliseconds.

Aurora Storage Write Throughput

Disk IOBytes per second

os.diskIO.auroraStorage.writeThroughput

The amount of network throughput used by responses from the DB cluster, in bytes per second.

Rdstemp Avg Queue Len

Disk IORequests

os.diskIO.rdstemp.avgQueueLen

The number of requests waiting in the I/O device's queue.

Rdstemp Avg Req Sz

Disk IORequests

os.diskIO.rdstemp.avgReqSz

The number of requests waiting in the I/O device's queue.

Rdstemp Await

Disk IOMilliseconds

os.diskIO.rdstemp.await

The number of milliseconds required to respond to requests, including queue time and service time.

Rdstemp Read IOs PS

Disk IORequests

os.diskIO.rdstemp.readIOsPS

The number of read operations per second.

Rdstemp Read KB

Disk IOKilobytes

os.diskIO.rdstemp.readKb

The total number of kilobytes read.

Rdstemp Read KB PS

Disk IOKilobytes per second

os.diskIO.rdstemp.readKbPS

The number of kilobytes read per second.

Rdstemp Rrqm PS

Disk IORequests per second

os.diskIO.rdstemp.rrqmPS

The number of merged read requests queued per second.

Rdstemp TPS

Disk IOTransactions per second

os.diskIO.rdstemp.tps

The number of I/O transactions per second.

Rdstemp Util

Disk IOPercentage

os.diskIO.rdstemp.util

The percentage of CPU time during which requests were issued.

Rdstemp Write IOs PS

Disk IORequests per second

os.diskIO.rdstemp.writeIOsPS

The number of write operations per second.

Rdstemp Write KB

Disk IOKilobytes

os.diskIO.rdstemp.writeKb

The total number of kilobytes written.

Rdstemp Write KB PS

Disk IOKilobytes per second

os.diskIO.rdstemp.writeKbPS

The number of kilobytes written per second.

Rdstemp Wrqm PS

Disk IORequests per second

os.diskIO.rdstemp.wrqmPS

The number of merged write requests queued per second.BlockedTasksTasksos.tasks.blockedThe number of tasks that are blocked.RunningTasksTasksos.tasks.runningThe number of tasks that are running.SleepingTasksTasksos.tasks.sleepingThe number of tasks that are sleeping.StoppedTasksTasksos.tasks.stoppedThe number of tasks that are stopped.TotalTasksTasksos.tasks.totalThe total number of tasks.ZombieTasksTasksos.tasks.zombieThe number of child tasks that are inactive with an active parent task.OneLoad Average MinuteProcessesos.loadAverageMinute.oneThe number of processes requesting CPU time over the last minute.FifteenLoad Average MinuteProcessesos.loadAverageMinute.fifteenThe number of processes requesting CPU time over the last 15 minutes.FiveLoad Average MinuteProcessesos.loadAverageMinute.fiveThe number of processes requesting CPU time over the last 5 minutes.CachedSwapKilobytesos.swap.cachedThe amount of swap memory, in kilobytes, used as cache memory.FreeSwapKilobytesos.swap.freeThe amount of swap memory free, in kilobytes. InSwapKilobytesos.swap.inThe amount of memory, in kilobytes, swapped in from disk.OutSwapKilobytesos.swap.outThe amount of memory, in kilobytes, swapped out to disk.TotalSwapKilobytesos.swap.total

The total amount of swap memory available in kilobytes.

Max FilesFile SysFilesos.fileSys.maxFilesThe maximum number of files that can be created for the file system across all storage
volumes.Used FilesFile SysFilesos.fileSys.usedFilesThe number of files in the file system across all storage volumes.Used File PercentFile SysFilesos.fileSys.usedFilePercentThe percentage of available files in use across all storage volumes.Used PercentFile SysPercentageos.fileSys.usedPercentThe percentage of the file-system disk space in use across all storage volumes.UsedFile SysKilobytesos.fileSys.usedThe amount of disk space used by files in the file system across all storage volumes,
in kilobytes.TotalFile SysKilobytesos.fileSys.totalThe total disk space available for the file system across all storage volumes, in
kilobytes.Max FilesFile SysFilesos.fileSys.<volumeName>.maxFilesThe maximum number of files that can be created for the storage volume.Used FilesFile SysFilesos.fileSys.<volumeName>.usedFilesThe number of files in the storage volume.Used File PercentFile SysFilesos.fileSys.<volumeName>.usedFilePercentThe percentage of available files in use in the storage
volume.Used PercentFile SysPercentageos.fileSys.<volumeName>.usedPercentThe percentage of the storage volume disk space in use.UsedFile SysKilobytesos.fileSys.<volumeName>.usedThe amount of disk space used by files in the storage volume, in
kilobytes.TotalFile SysKilobytesos.fileSys.<volumeName>.totalThe total disk space available in the storage volume, in kilobytes.RxNetworkBytes per secondos.network.rxThe number of bytes received per second.TxNetworkBytes per secondos.network.txThe number of bytes uploaded per second.Acu UtilizationGeneralPercentageos.general.acuUtilization

The percentage of current capacity out of the maximum configured
capacity.

Max Configured AcuGeneralACUsos.general.maxConfiguredAcu

The maximum capacity configured by the user, in Aurora capacity units (ACUs).

Min Configured AcuGeneralACUsos.general.minConfiguredAcu

The minimum capacity configured by the user, in ACUs.

Num VCPUsGeneralvCPUsos.general.numVCPUsThe number of virtual CPUs (vCPUs) for the DB instance.Serverless Database CapacityGeneralACUsos.general.serverlessDatabaseCapacity

The current capacity of the instance, in ACUs.

## Performance Insights counters for Aurora MySQL

The following database counters are available with Performance Insights for
Aurora MySQL.

###### Topics

- [Native counters for Aurora MySQL](#USER_PerfInsights_Counters.Aurora_MySQL.Native)

- [Non-native counters for Aurora MySQL](#USER_PerfInsights_Counters.Aurora_MySQL.NonNative)

### Native counters for Aurora MySQL

Native metrics are defined by the database engine and not by Amazon Aurora. You can find definitions for these native metrics in
[Server status variables](https://dev.mysql.com/doc/refman/8.0/en/server-status-variables.html) in the
MySQL documentation.

CounterTypeUnitMetricCom\_analyzeSQLQueries per seconddb.SQL.Com\_analyzeCom\_optimizeSQLQueries per seconddb.SQL.Com\_optimizeCom\_selectSQLQueries per seconddb.SQL.Com\_selectInnodb\_rows\_deletedSQLRows per seconddb.SQL.Innodb\_rows\_deletedInnodb\_rows\_insertedSQLRows per seconddb.SQL.Innodb\_rows\_insertedInnodb\_rows\_readSQLRows per seconddb.SQL.Innodb\_rows\_readInnodb\_rows\_updatedSQLRows per seconddb.SQL.Innodb\_rows\_updatedQueriesSQLQueries per seconddb.SQL.QueriesQuestionsSQLQueries per seconddb.SQL.QuestionsSelect\_full\_joinSQLQueries per seconddb.SQL.Select\_full\_joinSelect\_full\_range\_joinSQLQueries per seconddb.SQL.Select\_full\_range\_joinSelect\_rangeSQLQueries per seconddb.SQL.Select\_rangeSelect\_range\_checkSQLQueries per seconddb.SQL.Select\_range\_checkSelect\_scanSQLQueries per seconddb.SQL.Select\_scanSlow\_queriesSQLQueries per seconddb.SQL.Slow\_queriesSort\_merge\_passesSQLQueries per seconddb.SQL.Sort\_merge\_passesSort\_rangeSQLQueries per seconddb.SQL.Sort\_rangeSort\_rowsSQLQueries per seconddb.SQL.Sort\_rowsSort\_scanSQLQueries per seconddb.SQL.Sort\_scanTotal\_query\_timeSQLMillisecondsdb.SQL.Total\_query\_timeTable\_locks\_immediateLocksRequests per seconddb.Locks.Table\_locks\_immediateTable\_locks\_waitedLocksRequests per seconddb.Locks.Table\_locks\_waitedInnodb\_row\_lock\_timeLocksMilliseconds (average)db.Locks.Innodb\_row\_lock\_timeAborted\_clientsUsersConnectionsdb.Users.Aborted\_clientsAborted\_connectsUsersConnectionsdb.Users.Aborted\_connectsConnectionsUsersConnectionsdb.Users.ConnectionsExternal\_threads\_connectedUsersConnectionsdb.Users.External\_threads\_connectedmax\_connectionsUsersConnectionsdb.Users.max\_connectionsThreads\_connectedUsersConnectionsdb.Users.Threads\_connectedThreads\_createdUsersConnectionsdb.Users.Threads\_createdThreads\_runningUsersConnectionsdb.Users.Threads\_runningCreated\_tmp\_disk\_tablesTempTables per seconddb.Temp.Created\_tmp\_disk\_tablesCreated\_tmp\_tablesTempTables per seconddb.Temp.Created\_tmp\_tablesInnodb\_buffer\_pool\_pages\_dataCachePagesdb.Cache.Innodb\_buffer\_pool\_pages\_dataInnodb\_buffer\_pool\_pages\_totalCachePagesdb.Cache.Innodb\_buffer\_pool\_pages\_totalInnodb\_buffer\_pool\_read\_requestsCachePages per seconddb.Cache.Innodb\_buffer\_pool\_read\_requestsInnodb\_buffer\_pool\_readsCachePages per seconddb.Cache.Innodb\_buffer\_pool\_readsOpened\_tablesCacheTablesdb.Cache.Opened\_tablesOpened\_table\_definitionsCacheTablesdb.Cache.Opened\_table\_definitionsQcache\_hitsCacheQueriesdb.Cache.Qcache\_hits

### Non-native counters for Aurora MySQL

Non-native counter metrics are counters defined by Amazon RDS. A non-native metric
can be a metric that you get with a specific query. A non-native metric also can
be a derived metric, where two or more native counters are used in calculations
for ratios, hit rates, or latencies.

CounterTypeUnitMetricDescriptionDefinitionactive\_transactionsTransactionsdb.Transactions.active\_transactionsThe total active transactions.`SELECT COUNT(1) AS active_transactions FROM
								INFORMATION_SCHEMA.INNODB_TRX`innodb\_buffer\_pool\_hit\_rateCachedb.Cache.innoDB\_buffer\_pool\_hit\_rateThe percentage of reads that InnoDB could satisfy from the buffer
pool.`100 * innodb_buffer_pool_read_requests /
								(innodb_buffer_pool_read_requests +
							innodb_buffer_pool_reads)`innodb\_buffer\_pool\_hitsCachePages per seconddb.Cache.innoDB\_buffer\_pool\_hitsThe number of reads that InnoDB could satisfy from the buffer pool.`innodb_buffer_pool_read_requests - innodb_buffer_pool_reads`innodb\_buffer\_pool\_usageCachePercentagedb.Cache.innoDB\_buffer\_pool\_usage

The percentage of the InnoDB buffer pool that contains data (pages).

###### Note

When using compressed tables, this value can vary. For more information, see the information about
`Innodb_buffer_pool_pages_data` and `Innodb_buffer_pool_pages_total` in [Server status\
sariables](https://dev.mysql.com/doc/refman/8.0/en/server-status-variables.html) in the MySQL documentation.

`Innodb_buffer_pool_pages_data / Innodb_buffer_pool_pages_total *
								100.0`innodb\_deadlocksLocksdb.Locks.innodb\_deadlocksThe total number of deadlocks.`SELECT COUNT AS innodb_deadlocks FROM
								INFORMATION_SCHEMA.INNODB_METRICS WHERE
							NAME='lock_deadlocks'`innodb\_lock\_timeoutsLocksdb.Locks.innodb\_lock\_timeoutsThe total number of deadlocks that timed out.`SELECT COUNT AS innodb_lock_timeouts FROM
								INFORMATION_SCHEMA.INNODB_METRICS WHERE
							NAME='lock_timeouts'`innodb\_row\_lock\_waitsLocksdb.Locks.innodb\_row\_lock\_waitsThe total number of row locks that resulted in a wait.`SELECT COUNT AS innodb_row_lock_waits FROM
								INFORMATION_SCHEMA.INNODB_METRICS WHERE
								NAME='lock_row_lock_waits'`innodb\_rows\_changedSQLdb.SQL.innodb\_rows\_changedThe total InnoDB row operations.`db.SQL.Innodb_rows_inserted + db.SQL.Innodb_rows_deleted +
								db.SQL.Innodb_rows_updated`query\_cache\_hit\_rateCachePercentagedb.Cache.query\_cache\_hit\_rateThe hit ratio for the MySQL result set cache (query cache).`Qcache_hits / (QCache_hits + Com_select) * 100`temp\_disk\_tables\_percentTempdb.Temp.temp\_disk\_tables\_percentThe percentage of temporary tables that are created on disk by the
server when running statements.`(db.Temp.Created_tmp_disk_tables / db.Temp.Created_tmp_tables)
								* 100`trx\_rseg\_history\_lenTransactionsNonedb.Transactions.trx\_rseg\_history\_lenA list of the undo log pages for committed transactions that is maintained by the InnoDB transaction system
to implement multi-version concurrency control. For more information about undo log records details, see
[https://dev.mysql.com/doc/refman/8.0/en/innodb-multi-versioning.html](https://dev.mysql.com/doc/refman/8.0/en/innodb-multi-versioning.html) in the MySQL documentation.`SELECT COUNT AS trx_rseg_history_len FROM INFORMATION_SCHEMA.INNODB_METRICS WHERE NAME='trx_rseg_history_len'`

## Performance Insights counters for Aurora PostgreSQL

The following database counters are available with Performance Insights for
Aurora PostgreSQL.

###### Topics

- [Native counters for Aurora PostgreSQL](#USER_PerfInsights_Counters.Aurora_PostgreSQL.Native)

- [Non-native counters for Aurora PostgreSQL](#USER_PerfInsights_Counters.Aurora_PostgreSQL.NonNative)

### Native counters for Aurora PostgreSQL

Native metrics are defined by the database engine and not by Amazon Aurora. You can find definitions for these native metrics in [Viewing Statistics](https://www.postgresql.org/docs/current/monitoring-stats.html) in the PostgreSQL
documentation.

CounterTypeUnitMetrictup\_deletedSQLTuples per seconddb.SQL.tup\_deletedtup\_fetchedSQLTuples per seconddb.SQL.tup\_fetchedtup\_insertedSQLTuples per seconddb.SQL.tup\_insertedtup\_returnedSQLTuples per seconddb.SQL.tup\_returnedtup\_updatedSQLTuples per seconddb.SQL.tup\_updatedblks\_hitCacheBlocks per seconddb.Cache.blks\_hitbuffers\_allocCacheBlocks per seconddb.Cache.buffers\_allocbuffers\_checkpointCheckpointBlocks per seconddb.Checkpoint.buffers\_checkpointcheckpoints\_reqCheckpointCheckpoints per minutedb.Checkpoint.checkpoints\_reqcheckpoint\_sync\_timeCheckpointMilliseconds per checkpointdb.Checkpoint.checkpoint\_sync\_timecheckpoints\_timedCheckpointCheckpoints per minutedb.Checkpoint.checkpoints\_timedcheckpoint\_write\_timeCheckpointMilliseconds per checkpointdb.Checkpoint.checkpoint\_write\_timemaxwritten\_cleanCheckpointBgwriter clean stops per minutedb.Checkpoint.maxwritten\_cleandeadlocksConcurrencyDeadlocks per minutedb.Concurrency.deadlocksblk\_read\_timeI/OMillisecondsdb.IO.blk\_read\_timeblks\_readI/OBlocks per seconddb.IO.blks\_readbuffers\_backendI/OBlocks per seconddb.IO.buffers\_backendbuffers\_backend\_fsyncI/OBlocks per seconddb.IO.buffers\_backend\_fsyncbuffers\_cleanI/OBlocks per seconddb.IO.buffers\_cleantemp\_bytesTempBytes per seconddb.Temp.temp\_bytestemp\_filesTempFiles per minutedb.Temp.temp\_filesxact\_commitTransactionsCommits per seconddb.Transactions.xact\_commitxact\_rollbackTransactionsRollbacks per seconddb.Transactions.xact\_rollbacknumbackendsUserConnectionsdb.User.numbackendsarchived\_countWALFiles per minutedb.WAL.archived\_count

### Non-native counters for Aurora PostgreSQL

Non-native counter metrics are counters defined by Amazon Aurora. A non-native metric can be a metric that you get with a
specific query. A non-native metric also can be a derived metric, where two or more native counters are
used in calculations for ratios, hit rates, or latencies.

CounterTypeUnitMetricDescriptionDefinitioncheckpoint\_sync\_latencyCheckpointMillisecondsdb.Checkpoint.checkpoint\_sync\_latencyThe total amount of time that has been spent in the portion
of checkpoint processing where files are synchronized to
disk.`checkpoint_sync_time / (checkpoints_timed +
                            checkpoints_req)`checkpoint\_write\_latencyCheckpointMillisecondsdb.Checkpoint.checkpoint\_write\_latencyThe total amount of time that has been spent in the portion
of checkpoint processing where files are written to
disk.`checkpoint_write_time / (checkpoints_timed +
                            checkpoints_req)`local\_blks\_readI/OBlocksdb.IO.local\_blks\_readTotal number of local blocks read.Not applicablelocal\_blk\_read\_timeI/OMillisecondsdb.IO.local\_blk\_read\_timeIf `track_io_timing` is enabled, it tracks the total time spent reading local data file blocks, in milliseconds, otherwise the value is zero. For more information, see
[track\_io\_timing](https://www.postgresql.org/docs/current/runtime-config-statistics.html).Not applicablenum\_blocked\_sessionsLocksdb.Locks.num\_blocked\_sessionsThe number of blocked sessions.–orcache\_blks\_hitI/OQueriesdb.IO.orcache\_blks\_hitTotal number of shared blocks hits from optimized reads cache.Not applicableorcache\_blk\_read\_timeI/OMillisecondsdb.IO.orcache\_blk\_read\_timeIf `track_io_timing` is enabled, it tracks the total time spent reading data
file blocks from Optimized Reads cache, in milliseconds, otherwise the value is zero. For more information, see
[track\_io\_timing](https://www.postgresql.org/docs/current/runtime-config-statistics.html).Not applicableread\_latencyI/OMillisecondsdb.IO.read\_latencyThe time spent reading data file blocks by backends in this
instance.`blk_read_time / blks_read`storage\_blks\_readI/OBlocksdb.IO.storage\_blks\_readTotal number of shared blocks read from aurora storage.Not applicablestorage\_blk\_read\_timeI/OMillisecondsdb.IO.storage\_blk\_read\_timeIf `track_io_timing` is enabled, it tracks the total time spent reading data
file blocks from Aurora storage, in milliseconds, otherwise the value is zero. For more information, see
[track\_io\_timing](https://www.postgresql.org/docs/current/runtime-config-statistics.html).Not applicablenum\_blocked\_sessionsLocksdb.Locks.num\_blocked\_sessionsThe number of blocked sessions.–active\_countStateSessionsdb.state.active\_countThe number of sessions in the `active` state.Not applicableidle\_countStateSessionsdb.state.idle\_countThe number of sessions in the `idle` state.Not applicableidle\_in\_transaction\_aborted\_countStateSessionsdb.state.idle\_in\_transaction\_aborted\_countThe number of sessions in the `idle in transaction (aborted)` state.Not applicableidle\_in\_transaction\_countStateSessionsdb.state.idle\_in\_transaction\_countThe number of sessions in the `idle in transaction` state.Not applicableidle\_in\_transaction\_max\_timeStateSecondsdb.state.idle\_in\_transaction\_max\_timeThe duration of the longest running transaction in the `idle in transaction` state, in seconds.Not applicablelogical\_readsSQLBlocksdb.SQL.logical\_readsThe total number of blocks hit and read.`blks_hit + blks_read`queries\_startedSQLQueriesdb.SQL.queriesThe number of queries started.Not applicablequeries\_finishedSQLQueriesdb.SQL.queriesThe number of queries finished.Not applicabletotal\_query\_timeSQLMillisecondsdb.SQL.total\_query\_timeThe total time spent executing statements, in milliseconds.Not applicableactive\_transactionsTransactionsTransactionsdb.Transactions.active\_transactionsThe number of active transactions.Not applicableblocked\_transactionsTransactionsTransactionsdb.Transactions.blocked\_transactionsThe number of blocked transactions.Not applicablecommit\_latencyTransactionsMicrosecondsdb.Transactions.commit\_latencyThe average duration of commit operations.`db.Transactions.duration_commits / db.Transactions.xact_commit`duration\_commitsTransactionsMillisecondsdb.Transactions.duration\_commitsThe total transaction time spent in the last minute, in milliseconds.Not applicablemax\_used\_xact\_idsTransactionsTransactionsdb.Transactions.max\_used\_xact\_idsThe number of transactions that haven't been vacuumed.Not applicableoldest\_inactive\_logical\_replication\_slot\_xid\_ageTransactionsLengthdb.Transactions.oldest\_inactive\_logical\_replication\_slot\_xid\_ageThe age of the oldest transaction in an inactive logical replication slot.Not applicableoldest\_active\_logical\_replication\_slot\_xid\_ageTransactionsLengthdb.Transactions.oldest\_active\_logical\_replication\_slot\_xid\_ageThe age of the oldest transaction in an active logical replication slot.Not applicableoldest\_reader\_feedback\_xid\_ageTransactionsLengthdb.Transactions.oldest\_reader\_feedback\_xid\_ageThe age of the oldest transaction of a long‐running transaction on an Aurora reader instance or Aurora global DB reader instance.Not applicableoldest\_prepared\_transaction\_xid\_ageTransactionsLengthdb.Transactions.oldest\_prepared\_transaction\_xid\_ageThe age of the oldest prepared transaction.Not applicableoldest\_running\_transaction\_xid\_ageTransactionsLengthdb.Transactions.oldest\_running\_transaction\_xid\_ageThe age of the oldest running transaction.Not applicablemax\_connectionsUsersUsersdb.User.max\_connectionsThe maximum number of connections allowed for a database as configured in `max_connections` parameter.Not applicabletotal\_auth\_attemptsUsersUsersdb.User.total\_auth\_attemptsThe number of connection attempts to this instance.Not applicablearchive\_failed\_countWALFiles per minutedb.WAL.archive\_failed\_countThe number of failed attempts for archiving WAL files, in files per minute.Not applicable

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CloudWatch metrics for Performance Insights

SQL statistics for Performance Insights

All content copied from https://docs.aws.amazon.com/.
