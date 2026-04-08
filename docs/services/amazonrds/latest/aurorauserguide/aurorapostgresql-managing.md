# Performance and scaling for Amazon Aurora PostgreSQL

The following section discusses managing performance and scaling for an Amazon Aurora PostgreSQL
DB cluster. It also includes information about other maintenance tasks.

###### Topics

- [Scaling Aurora PostgreSQL DB instances](#AuroraPostgreSQL.Managing.Performance.InstanceScaling)

- [Maximum connections to an Aurora PostgreSQL DB instance](#AuroraPostgreSQL.Managing.MaxConnections)

- [Temporary storage limits for Aurora PostgreSQL](#AuroraPostgreSQL.Managing.TempStorage)

- [Huge pages for Aurora PostgreSQL](#AuroraPostgreSQL.Managing.HugePages)

- [Testing Amazon Aurora PostgreSQL by using fault injection queries](aurorapostgresql-managing-faultinjectionqueries.md)

- [Displaying volume status for an Aurora PostgreSQL DB cluster](aurorapostgresql-managing-volumestatus.md)

- [Specifying the RAM disk for the stats\_temp\_directory](aurorapostgresql-managing-ramdisk.md)

- [Managing temporary files with PostgreSQL](postgresql-managingtempfiles.md)

## Scaling Aurora PostgreSQL DB instances

You can scale Aurora PostgreSQL DB instances in two ways, instance scaling and read
scaling. For more information about read scaling, see [Read scaling](aurora-managing-performance.md#Aurora.Managing.Performance.ReadScaling).

You can scale your Aurora PostgreSQL DB cluster by modifying the DB instance class for
each DB instance in the DB cluster. Aurora PostgreSQL supports several DB instance classes
optimized for Aurora. Don't use db.t2 or db.t3 instance classes for larger Aurora
clusters of size greater than 40 terabytes (TB).

###### Note

We recommend using the T DB instance classes only for development and test
servers, or other non-production servers. For more details on the T instance
classes, see [DB instance class types](concepts-dbinstanceclass-types.md).

Scaling isn't instantaneous. It can take 15 minutes or more to complete the
change to a different DB instance class. If you use this approach to modify the DB
instance class, you apply the change during the next scheduled maintenance window
(rather than immediately) to avoid affecting users.

As an alternative to modifying the DB instance class directly, you can minimize
downtime by using the high availability features of Amazon Aurora. First, add an Aurora
Replica to your cluster. When creating the replica, choose the DB instance class size
that you want to use for your cluster. When the Aurora Replica is synchronized with the
cluster, you then failover to the newly added Replica. To learn more, see [Aurora Replicas](aurora-replication.md#Aurora.Replication.Replicas)
and [Fast failover with Amazon Aurora PostgreSQL](aurorapostgresql-bestpractices-fastfailover.md).

For detailed specifications of the DB instance classes supported by Aurora PostgreSQL, see
[Supported DB engines for DB instance classes](concepts-dbinstanceclass-supportaurora.md).

## Maximum connections to an Aurora PostgreSQL DB instance

An Aurora PostgreSQL DB cluster allocates resources based on the DB instance class and its
available memory. Each connection to the DB cluster consumes incremental amounts of
these resources, such as memory and CPU. Memory consumed per connection varies based on
query type, count, and whether temporary tables are used. Even an idle connection
consumes memory and CPU. That's because when queries run on a connection, more
memory is allocated for each query and it's not released completely, even when
processing stops. Thus, we recommend that you make sure your applications aren't
holding on to idle connections: each one of these wastes resources and affects
performance negatively. For more information, see [Resources consumed\
by idle PostgreSQL connections](https://aws.amazon.com/blogs/database/resources-consumed-by-idle-postgresql-connections).

The maximum number of connections allowed by an Aurora PostgreSQL DB instance is
determined by the `max_connections` parameter value specified in the
parameter group for that DB instance. The ideal setting for the
`max_connections` parameter is one that supports all the client
connections your application needs, without an excess of unused connections, plus at
least 3 more connections to support AWS automation. Before modifying the
`max_connections` parameter setting, we recommend that you consider the
following:

- If the `max_connections` value is too low, the Aurora PostgreSQL DB
instance might not have sufficient connections available when clients attempt to
connect. If this happens, attempts to connect using `psql` raise
error messages such as the following:

```nohighlight

psql: FATAL: remaining connection slots are reserved for non-replication superuser connections
```

- If the `max_connections` value exceeds the number of connections
that are actually needed, the unused connections can cause performance to
degrade.

The default value of `max_connections` is derived from the following
Aurora PostgreSQL `LEAST` function:

`LEAST({DBInstanceClassMemory/9531392},5000)`.

If you want to change the value for `max_connections`, you need to create a
custom DB cluster parameter group and change its value there. After applying your custom
DB parameter group to your cluster, be sure to reboot the primary instance so the new
value takes effect. For more information, see [Amazon Aurora PostgreSQL parameters](aurorapostgresql-reference-parametergroups.md) and [Creating a DB cluster parameter groupin Amazon Aurora](user-workingwithparamgroups-creatingcluster.md).

###### Tip

If your applications frequently open and close connections, or keep a large number
of long-lived connections open, we recommend that you use Amazon RDS Proxy. RDS Proxy is a
fully managed, highly available database proxy that uses connection pooling to share
database connections securely and efficiently. To learn more about RDS Proxy, see
[Amazon RDS Proxyfor Aurora](rds-proxy.md).

For details about how Aurora Serverless v2 instances handle this parameter, see [Maximum connections for Aurora Serverless v2](aurora-serverless-v2-setting-capacity.md#aurora-serverless-v2.max-connections).

## Temporary storage limits for Aurora PostgreSQL

Aurora PostgreSQL stores tables and indexes in the Aurora storage subsystem. Aurora PostgreSQL
uses separate temporary storage for non-persistent temporary files. This includes files
that are used for such purposes as sorting large data sets during query processing or
for index build operations. For more information, see the article [How can I\
troubleshoot local storage issues in Aurora PostgreSQL-Compatible\
instances?](https://repost.aws/knowledge-center/postgresql-aurora-storage-issue).

These local storage volumes are backed by Amazon Elastic Block Store and can be extended by using a
larger DB instance class. For more information about storage, see [Amazon Aurora storage](aurora-overview-storagereliability.md). You can also increase your
local storage for temporary objects by using an NVMe enabled instance type and Aurora
Optimized Reads-enabled temporary objects. For more information, see [Improving query performance for Aurora PostgreSQL with Aurora Optimized Reads](aurorapostgresql-optimized-reads.md).

###### Note

You might see `storage-optimization` events when scaling DB instances,
for example, from db.r5.2xlarge to db.r5.4xlarge.

The following table shows the maximum amount of temporary storage available for each
Aurora PostgreSQL DB instance class. For more information on DB instance class support for
Aurora, see [Amazon AuroraDB instance classes](concepts-dbinstanceclass.md).

DB instance classMaximum temporary storage available (GiB)db.x2g.16xlarge1829db.x2g.12xlarge1606db.x2g.8xlarge1071db.x2g.4xlarge535db.x2g.2xlarge268db.x2g.xlarge134db.x2g.large67db.r8g.48xlarge3072db.r8g.24xlarge1536db.r8g.16xlarge998db.r8g.12xlarge749db.r8g.8xlarge499db.r8g.4xlarge250db.r8g.2xlarge125db.r8g.xlarge63db.r8g.large31db.r7g.16xlarge1008db.r7g.12xlarge756db.r7g.8xlarge504db.r7g.4xlarge252db.r7g.2xlarge126db.r7g.xlarge63db.r7g.large32db.r7i.48xlarge3072db.r7i.24xlarge1500db.r7i.16xlarge1008db.r7i.12xlarge748db.r7i.8xlarge504db.r7i.4xlarge249db.r7i.2xlarge124db.r7i.xlarge62db.r7i.large31db.r6g.16xlarge1008db.r6g.12xlarge756db.r6g.8xlarge504db.r6g.4xlarge252db.r6g.2xlarge126db.r6g.xlarge63db.r6g.large32db.r6i.32xlarge1829db.r6i.24xlarge1500db.r6i.16xlarge1008db.r6i.12xlarge748db.r6i.8xlarge504db.r6i.4xlarge249db.r6i.2xlarge124db.r6i.xlarge62db.r6i.large31db.r5.24xlarge1500db.r5.16xlarge1008db.r5.12xlarge748db.r5.8xlarge504db.r5.4xlarge249db.r5.2xlarge124db.r5.xlarge62db.r5.large31db.r4.16xlarge960db.r4.8xlarge480db.r4.4xlarge240db.r4.2xlarge120db.r4.xlarge60db.r4.large30db.t4g.large16.5db.t4g.medium8.13db.t3.large16db.t3.medium7.5

###### Note

NVMe enabled instance types can increase the temporary space available by up to
the total NVMe size. For more information, see [Improving query performance for Aurora PostgreSQL with Aurora Optimized Reads](aurorapostgresql-optimized-reads.md).

You can monitor the temporary storage available for a DB instance with the
`FreeLocalStorage` CloudWatch metric,
 -->
described in [Amazon CloudWatch metrics for Amazon Aurora](aurora-auroramonitoring-metrics.md). (This doesn't apply to
Aurora Serverless v2.)

For some workloads, you can reduce the amount of temporary storage by allocating more
memory to the processes that are performing the operation. To increase the memory
available to an operation, increasing the values of the [work\_mem](https://www.postgresql.org/docs/current/runtime-config-resource.html) or [maintenance\_work\_mem](https://www.postgresql.org/docs/current/runtime-config-resource.html) PostgreSQL parameters.

## Huge pages for Aurora PostgreSQL

_Huge pages_ are a memory management feature that
reduces overhead when a DB instance is working with large contiguous chunks of memory,
such as that used by shared buffers. This PostgreSQL feature is supported by all
currently available Aurora PostgreSQL versions.

`Huge_pages` parameter is turned on by default for all DB instance classes
other than t3.medium,db.t3.large,db.t4g.medium,db.t4g.large instance classes. You can't
change the `huge_pages` parameter value or turn off this feature in the
supported instance classes of Aurora PostgreSQL.

On Aurora PostgreSQL DB instances that don't support the huge pages memory feature, specific
process memory usage might increase without corresponding workload changes.

The system allocates shared memory segments like the buffer cache during server
startup. When huge memory pages aren't available, the system doesn't charge these
allocations to the postmaster process. Instead, it includes the memory in the process
that first accessed each 4KB page in the shared memory segment.

###### Note

Active connections share allocated memory as needed, regardless of how shared
memory usage is tracked across processes.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Babelfish supports XML datatype methods

Testing Amazon Aurora PostgreSQL by using fault injection queries

All content copied from https://docs.aws.amazon.com/.
