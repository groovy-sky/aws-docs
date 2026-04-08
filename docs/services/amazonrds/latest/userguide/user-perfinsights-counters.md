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

- [Performance Insights counters for Amazon RDS for MariaDB and MySQL](#USER_PerfInsights_Counters.MySQL)

- [Performance Insights counters for Amazon RDS for Microsoft SQL Server](#USER_PerfInsights_Counters.SQLServer)

- [Performance Insights counters for Amazon RDS for Oracle](#USER_PerfInsights_Counters.Oracle)

- [Performance Insights counters for Amazon RDS for PostgreSQL](#USER_PerfInsights_Counters.PostgreSQL)

## Performance Insights operating system counters

The following operating system counters, which are prefixed with `os`, are
available with Performance Insights for all RDS engines
except RDS for SQL Server
.

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

Read IOs PS

Disk IORequests per second

os.diskIO.<devicename>.readIOsPS

The number of read operations per second.

Write IOs PS

Disk IORequests per second

os.diskIO.<devicename>.writeIOsPS

The number of write operations per second.

Avg Queue Len

Disk IORequests

os.diskIO.<devicename>.avgQueueLen

The number of requests waiting in the I/O device's queue.

Avg Req Sz

Disk IORequests

os.diskIO.<devicename>.avgReqSz

The number of requests waiting in the I/O device's queue.

Await

Disk IOMilliseconds

os.diskIO.<devicename>.await

The number of milliseconds required to respond to requests, including queue time and service time.

Read IOs PS

Disk IORequests

os.diskIO.<devicename>.readIOsPS

The number of read operations per second.

Read KB

Disk IOKilobytes

os.diskIO.<devicename>.readKb

The total number of kilobytes read.

Read KB PS

Disk IOKilobytes per second

os.diskIO.<devicename>.readKbPS

The number of kilobytes read per second.

Rrqm PS

Disk IORequests per second

os.diskIO.<devicename>.rrqmPS

The number of merged read requests queued per second.

TPS

Disk IOTransactions per second

os.diskIO.<devicename>.tps

The number of I/O transactions per second.

Util

Disk IOPercentage

os.diskIO.<devicename>.util

The percentage of CPU time during which requests were issued.

Write KB

Disk IOKilobytes

os.diskIO.<devicename>.writeKb

The total number of kilobytes written.

Write KB PS

Disk IOKilobytes per second

os.diskIO.<devicename>.writeKbPS

The number of kilobytes written per second.

Wrqm PS

Disk IORequests per second

os.diskIO.<devicename>.wrqmPS

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

## Performance Insights counters for Amazon RDS for MariaDB and MySQL

The following database counters are available with Performance Insights for Amazon RDS
for MariaDB and MySQL.

###### Topics

- [Native counters for RDS for MariaDB and RDS for MySQL](#USER_PerfInsights_Counters.MySQL.Native)

- [Non-native counters for Amazon RDS for MariaDB and MySQL](#USER_PerfInsights_Counters.MySQL.NonNative)

### Native counters for RDS for MariaDB and RDS for MySQL

Native metrics are defined by the database engine and not by Amazon RDS. For definitions of these native metrics, see [Server Status Variables](https://dev.mysql.com/doc/refman/8.0/en/server-status-variables.html) (for 8.0) and [Server Status Variables](https://dev.mysql.com/doc/refman/8.4/en/server-status-variables.html) (for 8.4) in the MySQL documentation.

CounterTypeUnitMetricCom\_analyzeSQLQueries per seconddb.SQL.Com\_analyzeCom\_optimizeSQLQueries per seconddb.SQL.Com\_optimizeCom\_selectSQLQueries per seconddb.SQL.Com\_selectConnectionsSQLThe number of connection attempts per minute (successful or not) to the MySQL
serverdb.Users.ConnectionsInnodb\_rows\_deletedSQLRows per seconddb.SQL.Innodb\_rows\_deletedInnodb\_rows\_insertedSQLRows per seconddb.SQL.Innodb\_rows\_insertedInnodb\_rows\_readSQLRows per seconddb.SQL.Innodb\_rows\_readInnodb\_rows\_updatedSQLRows per seconddb.SQL.Innodb\_rows\_updatedSelect\_full\_joinSQLQueries per seconddb.SQL.Select\_full\_joinSelect\_full\_range\_joinSQLQueries per seconddb.SQL.Select\_full\_range\_joinSelect\_rangeSQLQueries per seconddb.SQL.Select\_rangeSelect\_range\_checkSQLQueries per seconddb.SQL.Select\_range\_checkSelect\_scanSQLQueries per seconddb.SQL.Select\_scanSlow\_queriesSQLQueries per seconddb.SQL.Slow\_queriesSort\_merge\_passesSQLQueries per seconddb.SQL.Sort\_merge\_passesSort\_rangeSQLQueries per seconddb.SQL.Sort\_rangeSort\_rowsSQLQueries per seconddb.SQL.Sort\_rowsSort\_scanSQLQueries per seconddb.SQL.Sort\_scanQuestionsSQLQueries per seconddb.SQL.QuestionsInnodb\_row\_lock\_timeLocksMilliseconds (average)db.Locks.Innodb\_row\_lock\_timeTable\_locks\_immediateLocksRequests per seconddb.Locks.Table\_locks\_immediateTable\_locks\_waitedLocksRequests per seconddb.Locks.Table\_locks\_waitedAborted\_clientsUsersConnectionsdb.Users.Aborted\_clientsAborted\_connectsUsersConnectionsdb.Users.Aborted\_connectsmax\_connectionsUsersConnectionsdb.User.max\_connectionsThreads\_createdUsersConnectionsdb.Users.Threads\_createdThreads\_runningUsersConnectionsdb.Users.Threads\_runningInnodb\_data\_writesI/OOperations per seconddb.IO.Innodb\_data\_writesInnodb\_dblwr\_writesI/OOperations per seconddb.IO.Innodb\_dblwr\_writesInnodb\_log\_write\_requestsI/OOperations per seconddb.IO.Innodb\_log\_write\_requestsInnodb\_log\_writesI/OOperations per seconddb.IO.Innodb\_log\_writesInnodb\_pages\_writtenI/OPages per seconddb.IO.Innodb\_pages\_writtenCreated\_tmp\_disk\_tablesTempTables per seconddb.Temp.Created\_tmp\_disk\_tablesCreated\_tmp\_tablesTempTables per seconddb.Temp.Created\_tmp\_tablesInnodb\_buffer\_pool\_pages\_dataCachePagesdb.Cache.Innodb\_buffer\_pool\_pages\_dataInnodb\_buffer\_pool\_pages\_totalCachePagesdb.Cache.Innodb\_buffer\_pool\_pages\_totalInnodb\_buffer\_pool\_read\_requestsCachePages per seconddb.Cache.Innodb\_buffer\_pool\_read\_requestsInnodb\_buffer\_pool\_readsCachePages per seconddb.Cache.Innodb\_buffer\_pool\_readsOpened\_tablesCacheTablesdb.Cache.Opened\_tablesOpened\_table\_definitionsCacheTablesdb.Cache.Opened\_table\_definitionsQcache\_hitsCacheQueriesdb.Cache.Qcache\_hits

### Non-native counters for Amazon RDS for MariaDB and MySQL

Non-native counter metrics are counters defined by Amazon RDS. A non-native metric
can be a metric that you get with a specific query. A non-native metric also can
be a derived metric, where two or more native counters are used in calculations
for ratios, hit rates, or latencies.

CounterTypeUnitMetricDescriptionDefinitioninnodb\_buffer\_pool\_hitsCacheReadsdb.Cache.innoDB\_buffer\_pool\_hitsThe number of reads that InnoDB could satisfy from the buffer pool.`innodb_buffer_pool_read_requests - innodb_buffer_pool_reads`innodb\_buffer\_pool\_hit\_rateCachePercentagedb.Cache.innoDB\_buffer\_pool\_hit\_rateThe percentage of reads that InnoDB could satisfy from the buffer pool.`100 * innodb_buffer_pool_read_requests / (innodb_buffer_pool_read_requests +
									innodb_buffer_pool_reads)`innodb\_buffer\_pool\_usageCachePercentagedb.Cache.innoDB\_buffer\_pool\_usage

The percentage of the InnoDB buffer pool that contains data (pages).

###### Note

When using compressed tables, this value can vary. For more information, see the
information about `Innodb_buffer_pool_pages_data`
and `Innodb_buffer_pool_pages_total` in [Server Status Variables](https://dev.mysql.com/doc/refman/8.0/en/server-status-variables.html) (for 8.0) and [Server Status Variables](https://dev.mysql.com/doc/refman/8.4/en/server-status-variables.html) (for 8.4) in the MySQL
documentation.

`Innodb_buffer_pool_pages_data / Innodb_buffer_pool_pages_total *
								100.0`query\_cache\_hit\_rateCachePercentagedb.Cache.query\_cache\_hit\_rateMySQL result set cache (query cache) hit ratio.`Qcache_hits / (QCache_hits + Com_select) * 100`innodb\_datafile\_writes\_to\_diskI/OWritesdb.IO.innoDB\_datafile\_writes\_to\_diskThe number of InnoDB data file writes to disk, excluding double write and redo logging
write operations.`Innodb_data_writes - Innodb_log_writes - Innodb_dblwr_writes`innodb\_rows\_changedSQLRowsdb.SQL.innodb\_rows\_changedThe total InnoDB row operations.`db.SQL.Innodb_rows_inserted + db.SQL.Innodb_rows_deleted +
									db.SQL.Innodb_rows_updated`active\_transactionsTransactionsTransactionsdb.Transactions.active\_transactionsThe total active transactions.`SELECT COUNT(1) AS active_transactions FROM
								INFORMATION_SCHEMA.INNODB_TRX`trx\_rseg\_history\_lenTransactionsNonedb.Transactions.trx\_rseg\_history\_lenA list of the undo log pages for committed transactions that is maintained by the
InnoDB transaction system to implement multi-version concurrency
control. For more information about undo log records details, see
[InnoDB Multi-Versioning](https://dev.mysql.com/doc/refman/8.0/en/innodb-multi-versioning.html) (for 8.0) and [InnoDB Multi-Versioning](https://dev.mysql.com/doc/refman/8.4/en/innodb-multi-versioning.html) (for 8.4) in the MySQL
documentation.`SELECT COUNT AS trx_rseg_history_len FROM INFORMATION_SCHEMA.INNODB_METRICS WHERE NAME='trx_rseg_history_len'`innodb\_deadlocksLocksLocksdb.Locks.innodb\_deadlocksThe total number of deadlocks.`SELECT COUNT AS innodb_deadlocks FROM INFORMATION_SCHEMA.INNODB_METRICS WHERE
									NAME='lock_deadlocks'`innodb\_lock\_timeoutsLocksLocksdb.Locks.innodb\_lock\_timeoutsThe total number of locks that timed out.`SELECT COUNT AS innodb_lock_timeouts FROM INFORMATION_SCHEMA.INNODB_METRICS WHERE
									NAME='lock_timeouts'`innodb\_row\_lock\_waitsLocksLocksdb.Locks.innodb\_row\_lock\_waitsThe total number of row locks that resulted in a wait.`SELECT COUNT AS innodb_row_lock_waits FROM INFORMATION_SCHEMA.INNODB_METRICS WHERE
									NAME='lock_row_lock_waits'`

## Performance Insights counters for Amazon RDS for Microsoft SQL Server

The following database counters are available with Performance Insights for RDS
for Microsoft SQL Server.

### Native counters for RDS for Microsoft SQL Server

Native metrics are defined by the database engine and not by Amazon RDS. You can find definitions for
these native metrics in [Use SQL Server Objects](https://docs.microsoft.com/en-us/sql/relational-databases/performance-monitor/use-sql-server-objects?view=sql-server-2017) in the Microsoft SQL Server documentation.

CounterTypeUnitMetricForwarded Records[Access Methods](https://docs.microsoft.com/en-us/sql/relational-databases/performance-monitor/sql-server-access-methods-object?view=sql-server-2017)Records per seconddb.Access Methods.Forwarded RecordsPage Splits[Access Methods](https://docs.microsoft.com/en-us/sql/relational-databases/performance-monitor/sql-server-access-methods-object?view=sql-server-2017)Splits per seconddb.Access Methods.Page SplitsBuffer cache hit ratio[Buffer Manager](https://docs.microsoft.com/en-us/sql/relational-databases/performance-monitor/sql-server-buffer-manager-object?view=sql-server-2017)Ratiodb.Buffer Manager.Buffer cache hit ratioPage life expectancy[Buffer Manager](https://docs.microsoft.com/en-us/sql/relational-databases/performance-monitor/sql-server-buffer-manager-object?view=sql-server-2017)Expectancy in secondsdb.Buffer Manager.Page life expectancyPage lookups[Buffer Manager](https://docs.microsoft.com/en-us/sql/relational-databases/performance-monitor/sql-server-buffer-manager-object?view=sql-server-2017)Lookups per seconddb.Buffer Manager.Page lookupsPage reads[Buffer Manager](https://docs.microsoft.com/en-us/sql/relational-databases/performance-monitor/sql-server-buffer-manager-object?view=sql-server-2017)Reads per seconddb.Buffer Manager.Page readsPage writes[Buffer Manager](https://docs.microsoft.com/en-us/sql/relational-databases/performance-monitor/sql-server-buffer-manager-object?view=sql-server-2017)Writes per seconddb.Buffer Manager.Page writesActive Transactions[Databases](https://docs.microsoft.com/en-us/sql/relational-databases/performance-monitor/sql-server-databases-object?view=sql-server-2017)Transactionsdb.Databases.Active Transactions (\_Total)Log Bytes Flushed[Databases](https://docs.microsoft.com/en-us/sql/relational-databases/performance-monitor/sql-server-databases-object?view=sql-server-2017)Bytes flushed per seconddb.Databases.Log Bytes Flushed (\_Total)Log Flush Waits[Databases](https://docs.microsoft.com/en-us/sql/relational-databases/performance-monitor/sql-server-databases-object?view=sql-server-2017)Waits per seconddb.Databases.Log Flush Waits (\_Total)Log Flushes[Databases](https://docs.microsoft.com/en-us/sql/relational-databases/performance-monitor/sql-server-databases-object?view=sql-server-2017)Flushes per seconddb.Databases.Log Flushes (\_Total)Write Transactions[Databases](https://docs.microsoft.com/en-us/sql/relational-databases/performance-monitor/sql-server-databases-object?view=sql-server-2017)Transactions per seconddb.Databases.Write Transactions (\_Total)Processes blocked[General Statistics](https://docs.microsoft.com/en-us/sql/relational-databases/performance-monitor/sql-server-general-statistics-object?view=sql-server-2017)Processes blockeddb.General Statistics.Processes blockedUser Connections[General Statistics](https://docs.microsoft.com/en-us/sql/relational-databases/performance-monitor/sql-server-general-statistics-object?view=sql-server-2017)Connectionsdb.General Statistics.User ConnectionsLatch Waits[Latches](https://docs.microsoft.com/en-us/sql/relational-databases/performance-monitor/sql-server-latches-object?view=sql-server-2017)Waits per seconddb.Latches.Latch WaitsNumber of Deadlocks[Locks](https://docs.microsoft.com/en-us/sql/relational-databases/performance-monitor/sql-server-locks-object?view=sql-server-2017)Deadlocks per seconddb.Locks.Number of Deadlocks (\_Total)Memory Grants Pending[Memory Manager](https://docs.microsoft.com/en-us/sql/relational-databases/performance-monitor/sql-server-memory-manager-object?view=sql-server-2017)Memory grantsdb.Memory Manager.Memory Grants PendingBatch Requests[SQL Statistics](https://docs.microsoft.com/en-us/sql/relational-databases/performance-monitor/sql-server-sql-statistics-object?view=sql-server-2017)Requests per seconddb.SQL Statistics.Batch RequestsSQL Compilations[SQL Statistics](https://docs.microsoft.com/en-us/sql/relational-databases/performance-monitor/sql-server-sql-statistics-object?view=sql-server-2017)Compilations per seconddb.SQL Statistics.SQL CompilationsSQL Re-Compilations[SQL Statistics](https://docs.microsoft.com/en-us/sql/relational-databases/performance-monitor/sql-server-sql-statistics-object?view=sql-server-2017)Re-compilations per seconddb.SQL Statistics.SQL Re-Compilations

## Performance Insights counters for Amazon RDS for Oracle

The following database counters are available with Performance Insights for RDS
for Oracle.

### Native counters for RDS for Oracle

Native metrics are defined by the database engine and not by Amazon RDS. You can find definitions for
these native metrics in [Statistics Descriptions](https://docs.oracle.com/en/database/oracle/oracle-database/12.2/refrn/statistics-descriptions-2.html) in the Oracle documentation.

###### Note

For the `CPU used by this session` counter metric, the unit has
been transformed from the native centiseconds to active sessions to make the
value easier to use. For example, CPU send in the DB Load chart represents
the demand for CPU. The counter metric `CPU used by this session`
represents the amount of CPU used by Oracle sessions. You can compare CPU
send to the `CPU used by this session` counter metric. When
demand for CPU is higher than CPU used, sessions are waiting for CPU
time.

CounterTypeUnitMetricCPU used by this sessionUserActive sessionsdb.User.CPU used by this sessionSQL\*Net roundtrips to/from clientUserRoundtrips per seconddb.User.SQL\*Net roundtrips to/from clientBytes received via SQL\*Net from clientUserBytes per seconddb.User.bytes received via SQL\*Net from clientUser commitsUserCommits per seconddb.User.user commitsLogons cumulativeUserLogons per seconddb.User.logons cumulativeUser callsUserCalls per seconddb.User.user callsBytes sent via SQL\*Net to clientUserBytes per seconddb.User.bytes sent via SQL\*Net to clientUser rollbacksUserRollbacks per seconddb.User.user rollbacksRedo sizeRedoBytes per seconddb.Redo.redo sizeParse count (total)SQLParses per seconddb.SQL.parse count (total)Parse count (hard)SQLParses per seconddb.SQL.parse count (hard)Table scan rows gottenSQLRows per seconddb.SQL.table scan rows gottenSorts (memory)SQLSorts per seconddb.SQL.sorts (memory)Sorts (disk)SQLSorts per seconddb.SQL.sorts (disk)Sorts (rows)SQLSorts per seconddb.SQL.sorts (rows)Physical read bytesCacheBytes per seconddb.Cache.physical read bytesDB block getsCacheBlocks per seconddb.Cache.db block getsDBWR checkpointsCacheCheckpoints per minutedb.Cache.DBWR checkpointsPhysical readsCacheReads per seconddb.Cache.physical readsConsistent gets from cacheCacheGets per seconddb.Cache.consistent gets from cacheDB block gets from cacheCacheGets per seconddb.Cache.db block gets from cacheConsistent getsCacheGets per seconddb.Cache.consistent gets

## Performance Insights counters for Amazon RDS for PostgreSQL

The following database counters are available with Performance Insights for Amazon RDS
for PostgreSQL.

###### Topics

- [Native counters for Amazon RDS for PostgreSQL](#USER_PerfInsights_Counters.PostgreSQL.Native)

- [Non-native counters for Amazon RDS for PostgreSQL](#USER_PerfInsights_Counters.PostgreSQL.NonNative)

### Native counters for Amazon RDS for PostgreSQL

Native metrics are defined by the database engine and not by Amazon RDS. You can find definitions for
these native metrics in [Viewing\
Statistics](https://www.postgresql.org/docs/current/monitoring-stats.html) in the PostgreSQL documentation.

CounterTypeUnitMetricblks\_hitCacheBlocks per seconddb.Cache.blks\_hitbuffers\_allocCacheBlocks per seconddb.Cache.buffers\_allocbuffers\_checkpointCheckpointBlocks per seconddb.Checkpoint.buffers\_checkpointcheckpoint\_sync\_timeCheckpointMilliseconds per checkpointdb.Checkpoint.checkpoint\_sync\_timecheckpoint\_write\_timeCheckpointMilliseconds per checkpointdb.Checkpoint.checkpoint\_write\_timecheckpoints\_reqCheckpointCheckpoints per minutedb.Checkpoint.checkpoints\_reqcheckpoints\_timedCheckpointCheckpoints per minutedb.Checkpoint.checkpoints\_timedmaxwritten\_cleanCheckpointBgwriter clean stops per minute db.Checkpoint.maxwritten\_cleandeadlocksConcurrencyDeadlocks per minutedb.Concurrency.deadlocksblk\_read\_timeI/OMillisecondsdb.IO.blk\_read\_timeblks\_readI/OBlocks per seconddb.IO.blks\_readbuffers\_backendI/OBlocks per seconddb.IO.buffers\_backendbuffers\_backend\_fsyncI/OBlocks per seconddb.IO.buffers\_backend\_fsyncbuffers\_cleanI/OBlocks per seconddb.IO.buffers\_cleantup\_deletedSQLTuples per seconddb.SQL.tup\_deletedtup\_fetchedSQLTuples per seconddb.SQL.tup\_fetchedtup\_insertedSQLTuples per seconddb.SQL.tup\_insertedtup\_returnedSQLTuples per seconddb.SQL.tup\_returnedtup\_updatedSQLTuples per seconddb.SQL.tup\_updatedtemp\_bytesTempBytes per seconddb.Temp.temp\_bytestemp\_filesTempFiles per minutedb.Temp.temp\_filesxact\_commitTransactionsCommits per seconddb.Transactions.xact\_commitxact\_rollbackTransactionsRollbacks per seconddb.Transactions.xact\_rollbacknumbackendsUserConnectionsdb.User.numbackendsarchived\_countWrite-ahead log (WAL)Files per minutedb.WAL.archived\_count

### Non-native counters for Amazon RDS for PostgreSQL

Non-native counter metrics are counters defined by Amazon RDS. A non-native metric
can be a metric that you get with a specific query. A non-native metric also can
be a derived metric, where two or more native counters are used in calculations
for ratios, hit rates, or latencies.

CounterTypeUnitMetricDescriptionDefinitioncheckpoint\_sync\_latencyCheckpointdb.Checkpoint.checkpoint\_sync\_latencyThe total amount of time that has been spent in the portion of checkpoint processing where
files are synchronized to disk.`checkpoint_sync_time / (checkpoints_timed + checkpoints_req)`checkpoint\_write\_latencyCheckpointdb.Checkpoint.checkpoint\_write\_latencyThe total amount of time that has been spent in the portion of checkpoint processing where
files are written to disk.`checkpoint_write_time / (checkpoints_timed + checkpoints_req)`read\_latencyI/Odb.IO.read\_latencyThe time spent reading data file blocks by backends in this instance.`blk_read_time / blks_read`idle\_in\_transaction\_aborted\_countStateSessionsdb.state.idle\_in\_transaction\_aborted\_countThe number of sessions in the `idle in transaction (aborted)` state.Not applicableidle\_in\_transaction\_countStateSessionsdb.state.idle\_in\_transaction\_countThe number of sessions in the `idle in transaction` state.Not applicableidle\_in\_transaction\_max\_timeStateSecondsdb.state.idle\_in\_transaction\_max\_timeThe duration of the longest running transaction in the `idle in transaction` state, in seconds.Not applicableactive\_transactionsTransactionsTransactionsdb.Transactions.active\_transactionsThe number of active transactions.Not applicableblocked\_transactionsTransactionsTransactionsdb.Transactions.blocked\_transactionsThe number of blocked transactions.Not applicableoldest\_active\_logical\_replication\_slot\_xid\_ageTransactionsdb.Transactions.oldest\_active\_logical\_replication\_slot\_xid\_age

The age of the oldest transaction in an active logical replication slot.

For more information, see [Logical replication slot](appendix-postgresql-commondbatasks-autovacuum-monitoring-resolving-identifiableblockers.md#Appendix.PostgreSQL.CommonDBATasks.Autovacuum_Monitoring.Logical_replication_slot).

`–`oldest\_inactive\_logical\_replication\_slot\_xid\_ageTransactionsdb.Transactions.oldest\_inactive\_logical\_replication\_slot\_xid\_age

The age of the oldest transaction in an inactive logical replication slot.

For more information, see [Logical replication slot](appendix-postgresql-commondbatasks-autovacuum-monitoring-resolving-identifiableblockers.md#Appendix.PostgreSQL.CommonDBATasks.Autovacuum_Monitoring.Logical_replication_slot).

`–`oldest\_prepared\_transaction\_xid\_ageTransactionsdb.Transactions.oldest\_prepared\_transaction\_xid\_age

The age of the oldest prepared transaction.

For more information, see [Prepared transaction](appendix-postgresql-commondbatasks-autovacuum-monitoring-resolving-identifiableblockers.md#Appendix.PostgreSQL.CommonDBATasks.Autovacuum_Monitoring.Prepared_transaction).

`–`oldest\_running\_transaction\_xid\_ageTransactionsdb.Transactions.oldest\_running\_transaction\_xid\_age

The age of the oldest running transaction.

For more information, see [Active statement](appendix-postgresql-commondbatasks-autovacuum-monitoring-resolving-identifiableblockers.md#Appendix.PostgreSQL.CommonDBATasks.Autovacuum_Monitoring.Active_statement) for the oldest running active transaction and [Idle in transaction](appendix-postgresql-commondbatasks-autovacuum-monitoring-resolving-identifiableblockers.md#Appendix.PostgreSQL.CommonDBATasks.Autovacuum_Monitoring.Idle_in_transaction) for the oldest running idle-in-transaction.

`–`oldest\_hot\_standby\_feedback\_xid\_ageTransactionsdb.Transactions.oldest\_hot\_standby\_feedback\_xid\_age

The age of the oldest running transaction on a read replica with `hot_standby_feedback` enabled.

For more information, see [Read replicas](appendix-postgresql-commondbatasks-autovacuum-monitoring-resolving-identifiableblockers.md#Appendix.PostgreSQL.CommonDBATasks.Autovacuum_Monitoring.Read_replicas).

`–`max\_used\_xact\_idsTransactionsTransactionsdb.Transactions.max\_used\_xact\_idsThe number of transactions that haven't been vacuumed.Not applicablemax\_connectionsUsersConnectionsdb.User.max\_connectionsThe maximum number of connections allowed for a DB instance as configured in `max_connections` parameter.Not applicablearchive\_failed\_countWALFiles per minutedb.WAL.archive\_failed\_countThe number of failed attempts for archiving WAL files, in files per minute.Not applicable

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CloudWatch metrics for Performance Insights

SQL statistics for Performance Insights

All content copied from https://docs.aws.amazon.com/.
