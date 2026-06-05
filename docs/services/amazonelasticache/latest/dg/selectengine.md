---
title: "Comparing node-based Valkey, Memcached, and Redis OSS clusters"
---

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

- **ElastiCache version 9.0 for Valkey**

You want built-in [full-text search, aggregations, and hybrid search](https://github.com/valkey-io/valkey-search) capabilities that extend the vector search introduced in Valkey 8.2, [hash field expiration](https://github.com/valkey-io/valkey/pull/2089) for per-field TTLs, [multi-database support in cluster mode](https://github.com/valkey-io/valkey/pull/1671), [polygon-based geospatial queries](https://github.com/valkey-io/valkey/pull/1809), up to [40% higher throughput with pipelining](https://github.com/valkey-io/valkey/pull/2092), or [atomic slot migration](https://github.com/valkey-io/valkey/pull/1949) for safer cluster scaling, or _durability_ via a Multi-AZ transactional log for zero data loss during failures. For more information, see [Valkey 9.0](versionmanagement.md#valkey-version-9.0).

- **ElastiCache version 8.2 for Valkey**

You want native [vector search](search.md) with microsecond latency and 95%+ recall rate. For more information, see [Valkey 8.2](versionmanagement.md#valkey-version-8.2).

- **ElastiCache version 8.1 for Valkey**

You want a [new memory-efficient hash table](https://valkey.io/blog/new-hash-table) that reduces memory overhead by up to 20%, native [Bloom filter](https://valkey.io/topics/bloomfilters) support, the [COMMANDLOG](https://valkey.io/commands/commandlog-get) command, and up to 514% higher throughput for BITCOUNT. For more information, see [Valkey 8.1](versionmanagement.md#valkey-version-8.1).

- **ElastiCache version 8.0 for Valkey**

You want up to 20% better memory efficiency through [embedded keys](https://github.com/valkey-io/valkey/pull/541), [per-slot metrics](https://github.com/valkey-io/valkey/pull/20), [dual-channel replication](https://github.com/valkey-io/valkey/pull/60), and [async I/O threading](https://github.com/valkey-io/valkey/pull/763). For more information, see [Valkey 8.0](versionmanagement.md#valkey-version-8).

- **ElastiCache version 7.2 for Valkey or version 7.0 (Enhanced) for Redis OSS**

You want to use [Functions](https://valkey.io/topics/functions-intro),
[Sharded Pub/Sub](https://valkey.io/topics/pubsub), or
[ACL improvements](https://valkey.io/topics/acl). For more information,
see [Redis OSS Version 7.0 (Enhanced)](engine-versions.md#redis-version-7.0).

- **ElastiCache version 6.2 (Enhanced) for Redis OSS**

You want the ability to tier data between memory and SSD using the r6gd node type. For more information, see
[Data tiering](data-tiering.md).

- **ElastiCache version 6.0 (Enhanced) for Redis OSS**

You want to authenticate users with role-based access control.

For more information, see [Redis OSS Version 6.0 (Enhanced)](engine-versions.md#redis-version-6.0).

Comparison summary of Memcached, Valkey or Redis OSS (cluster mode disabled), and Valkey or Redis OSS (cluster mode enabled) Memcached  Valkey or Redis OSS (cluster mode disabled)  Valkey or Redis OSS (cluster mode enabled) Engine versions+1.4.5 and later4.0.10 and later4.0.10 and laterData typesSimple ‡2.8.x - Complex \*3.2.x and later - Complex †Complex †Data partitioningYesNoYesCluster is modifiableYesYes3.2.10 and later - LimitedOnline reshardingNoNo3.2.10 and laterEncryptionin-transit 1.6.12 and later4.0.10 and later4.0.10 and laterData tieringNo6.2 and later6.2 and laterMemory efficienciesNoValkey 8.0 and laterValkey 8.0 and laterBloom filtersNoValkey 8.1 and laterValkey 8.1 and laterVector searchNoValkey 8.2 and laterValkey 8.2 and laterFull-text searchNoValkey 9.0 and laterValkey 9.0 and laterHybrid search (text + vector)NoValkey 9.0 and laterValkey 9.0 and laterAggregation pipelinesNoValkey 9.0 and laterValkey 9.0 and laterHash field expirationNoValkey 9.0 and laterValkey 9.0 and laterNumbered databases in cluster modeNoN/AValkey 9.0 and laterDurabilityNoNoValkey 9.0 and later (sync/async)Compliance certificationsCompliance Certification

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

Multi-threadedYesNoNoNode type upgradeYes (engine version 1.5+)YesYesEngine upgradingYesYesYesHigh availability (replication)NoYesYesAutomatic failoverNoOptionalRequiredPub/Sub capabilitiesNoYesYesSorted setsNoYesYesBackup and restoreFor serverless caches only, not applicable to node-based clustersYesYesGeospatial indexingNo4.0.10 and laterYes**Notes:**‡ string, objects (like databases)\\* string, sets, sorted sets, lists, hashes, bitmaps, hyperloglog† string, sets, sorted sets, lists, hashes, bitmaps,
hyperloglog, geospatial indexes\+ Excludes versions which are deprecated, have reached or soon to reach end of life.

After you choose the engine for your cluster, we recommend that you use the most recent
version of that engine. For more information, see [Supported node types](cachenodes-supportedtypes.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Shards in ElastiCache

Online Migration for Valkey or Redis OSS

All content copied from https://docs.aws.amazon.com/.
