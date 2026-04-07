# Comparing node-based Valkey, Memcached, and Redis OSS clusters

Amazon ElastiCache supports the Valkey, Memcached, and Redis OSS cache engines. Each engine provides some
advantages. Use the information in this topic to help you choose the engine and version that
best meets your requirements.

###### Important

After you create a cache, node-based cluster or replication group, you can upgrade to a
newer engine version, but you cannot downgrade to an older engine version. If you want to use an
older engine version, you must delete the existing cache, node-based cluster or replication group
and create it again with the earlier engine version.

On the surface, the engines look similar. Each of them is an in-memory key-value store.
However, in practice there are significant differences.

###### Choose Memcached if the following apply for you:

- You need the simplest model possible.

- You need to run large nodes with multiple cores or threads.

- You need the ability to scale out and in, adding and removing nodes as demand on your
system increases and decreases.

- You need to cache objects.

###### Choose Valkey or Redis OSS with ElastiCache if the following apply for you:

- **ElastiCache version 7.2 for Valkey or version 7.0 (Enhanced) for Redis OSS**

You want to use [Functions](https://valkey.io/topics/functions-intro),
[Sharded Pub/Sub](https://valkey.io/topics/pubsub), or
[ACL improvements](https://valkey.io/topics/acl). For more information,
see [Redis OSS Version 7.0 (Enhanced)](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/engine-versions.html#redis-version-7.0).

- **ElastiCache version 6.2 (Enhanced) for Redis OSS**

You want the ability to tier data between memory and SSD using the r6gd node type. For more information, see
[Data tiering](data-tiering.md).

- **ElastiCache version 6.0 (Enhanced) for Redis OSS**

You want to authenticate users with role-based access control.

For more information, see [Redis OSS Version 6.0 (Enhanced)](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/engine-versions.html#redis-version-6.0).

- **ElastiCache version 5.0.0 (Enhanced) for Redis OSS**

You want to use [Redis OSS streams](https://redis.io/topics/streams-intro), a log data structure that allows producers to append new items in real time and also allows consumers to consume messages either in a blocking or non-blocking fashion.

For more information, see [Redis OSS Version 5.0.0 (Enhanced)](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/engine-versions.html#redis-version-5-0).

- **ElastiCache version 4.0.10 (Enhanced) for Redis OSS**

Supports both encryption and dynamically adding or removing shards from your Valkey or Redis OSS (cluster mode enabled) cluster.

For more information, see [Redis OSS Version 4.0.10 (Enhanced)](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/engine-versions.html#redis-version-4-0-10).

The following versions are deprecated, have reached or soon to reach end of life.

- **ElastiCache version 3.2.10 (Enhanced) for Redis OSS**

Supports the ability to dynamically add or remove shards from your Valkey or Redis OSS (cluster mode enabled) cluster.

###### Important

Currently ElastiCache 3.2.10 for Redis OSS doesn't support encryption.

For more information, see the following:

- [Redis OSS Version 3.2.10 (Enhanced)](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/engine-versions.html#redis-version-3-2-10)

- Online resharding best practices for Redis OSS, For more information, see the following:

- [Best Practices: Online Resharding](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/best-practices-online-resharding.html)

- [Online Resharding and Shard Rebalancing for Valkey or Redis OSS (Cluster Mode Enabled)](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/scaling-redis-cluster-mode-enabled.html#redis-cluster-resharding-online)

- For more information on scaling Redis OSS clusters, see [Scaling](scaling.md).

- **ElastiCache version 3.2.6 (Enhanced) for Redis OSS**

If you need the functionality of earlier Redis OSS versions plus the following features,
choose 3.2.6:

- In-transit encryption. For more information, see [Amazon ElastiCache for Redis OSS In-Transit Encryption](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/in-transit-encryption.html).

- At-rest encryption. For more information, see [Amazon ElastiCache for Redis OSS At-Rest Encryption](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/at-rest-encryption.html).

- **ElastiCache (Cluster mode enabled) version 3.2.4 for Redis OSS**

If you need the functionality of 2.8.x plus the following features, choose 3.2.4 (clustered mode):

- You need to partition your data across two to 500 node groups (clustered mode
only).

- You need geospatial indexing (clustered mode or non-clustered mode).

- You don't need to support multiple databases.

- **ElastiCache (non-clustered mode) 2.8.x and 3.2.4 (Enhanced) for Redis OSS**

If the following apply for you, 2.8.x or 3.2.4 (non-clustered mode):

- You need complex data types, such as strings, hashes, lists, sets, sorted sets, and bitmaps.

- You need to sort or rank in-memory datasets.

- You need persistence of your key store.

- You need to replicate your data from the primary to one or more read replicas for read intensive
applications.

- You need automatic failover if your primary node fails.

- You need publish and subscribe (pub/sub) capabilities—to inform
clients about events on the server.

- You need backup and restore capabilities for node-based clusters as well as serverless
caches.

- You need to support multiple databases.

Comparison summary of Memcached, Valkey or Redis OSS (cluster mode disabled), and Valkey or Redis OSS (cluster mode enabled) Memcached  Valkey or Redis OSS (cluster mode disabled)  Valkey or Redis OSS (cluster mode enabled) Engine versions+1.4.5 and later4.0.10 and later4.0.10 and laterData typesSimple ‡2.8.x - Complex \*3.2.x and later - Complex †Complex †Data partitioningYesNoYesCluster is modifiableYesYes3.2.10 and later - LimitedOnline reshardingNoNo3.2.10 and laterEncryptionin-transit 1.6.12 and later4.0.10 and later4.0.10 and laterData tieringNo6.2 and later6.2 and laterCompliance certificationsCompliance Certification

    FedRAMP

    HIPAA

    PCI DSS

Yes - 1.6.12 and later

Yes - 1.6.12 and later

Yes

4.0.10 and later

4.0.10 and later

4.0.10 and later

4.0.10 and later

4.0.10 and later

4.0.10 and later

Multi-threadedYesNoNoNode type upgradeNoYesYesEngine upgradingYesYesYesHigh availability (replication)NoYesYesAutomatic failoverNoOptionalRequiredPub/Sub capabilitiesNoYesYesSorted setsNoYesYesBackup and restoreFor serverless caches only, not applicable to node-based clustersYesYesGeospatial indexingNo4.0.10 and laterYes**Notes:**‡ string, objects (like databases)\\* string, sets, sorted sets, lists, hashes, bitmaps, hyperloglog† string, sets, sorted sets, lists, hashes, bitmaps,
hyperloglog, geospatial indexes\+ Excludes versions which are deprecated, have reached or soon to reach end of life.

After you choose the engine for your cluster, we recommend that you use the most recent
version of that engine. For more information, see [Supported node types](cachenodes-supportedtypes.md).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Shards in ElastiCache

Online Migration for Valkey or Redis OSS
