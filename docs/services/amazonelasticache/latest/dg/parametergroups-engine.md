# Engine specific parameters

**Valkey and Redis OSS**

Most Valkey 8 parameters are compatible with Redis OSS 7.1 parameters. Valkey 7.2 parameters are the same as Redis OSS 7 parameters.

If you do not specify a parameter group for your Valkey or Redis OSS cluster, then a default
parameter group appropriate to your engine version will be used. You can't change the
values of any parameters in the default parameter group. However, you can create a
custom parameter group and assign it to your cluster at any time as long as the values
of conditionally modifiable parameters are the same in both parameter groups. For more
information, see [Creating an ElastiCache parameter group](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/ParameterGroups.Creating.html).

###### Topics

- [Valkey and Redis OSS parameters](#ParameterGroups.Redis)

- [Memcached specific parameters](#ParameterGroups.Memcached)

## Valkey and Redis OSS parameters

###### Topics

- [Valkey 8.2 parameter changes](#ParameterGroups.Valkey.8.2)

- [Valkey 8.1 parameter changes](#ParameterGroups.Valkey.8.1)

- [Valkey 8.0 parameter changes](#ParameterGroups.Valkey.8)

- [Valkey 7.2 and Redis OSS 7 parameter changes](#ParameterGroups.Redis.7)

- [Redis OSS 6.x parameter changes](#ParameterGroups.Redis.6-x)

- [Redis OSS 5.0.3 parameter changes](#ParameterGroups.Redis.5-0-3)

- [Redis OSS 5.0.0 parameter changes](#ParameterGroups.Redis.5.0)

- [Redis OSS 4.0.10 parameter changes](#ParameterGroups.Redis.4-0-10)

- [Redis OSS 3.2.10 parameter changes](#ParameterGroups.Redis.3-2-10)

- [Redis OSS 3.2.6 parameter changes](#ParameterGroups.Redis.3-2-6)

- [Redis OSS 3.2.4 parameter changes](#ParameterGroups.Redis.3-2-4)

- [Redis OSS 2.8.24 (enhanced) added parameters](#ParameterGroups.Redis.2-8-24)

- [Redis OSS 2.8.23 (enhanced) added parameters](#ParameterGroups.Redis.2-8-23)

- [Redis OSS 2.8.22 (enhanced) added parameters](#ParameterGroups.Redis.2-8-22)

- [Redis OSS 2.8.21 added parameters](#ParameterGroups.Redis.2-8-21)

- [Redis OSS 2.8.19 added parameters](#ParameterGroups.Redis.2-8-19)

- [Redis OSS 2.8.6 added parameters](#ParameterGroups.Redis.2-8-6)

- [Redis OSS 2.6.13 parameters](#ParameterGroups.Redis.2-6-13)

- [Redis OSS node-type specific parameters](#ParameterGroups.Redis.NodeSpecific)

### Valkey 8.2 parameter changes

**Parameter group family:** valkey8

###### Note

- Valkey 8.2 parameter changes don't apply to Valkey 8.1

- Valkey 8.0 and above parameter groups are incompatible with Redis OSS 7.2.4.

- in Valkey 8.2, the following commands are unavailable for serverless caches: `commandlog`,
`commandlog get`, `commandlog help`, `commandlog len`, and `commandlog reset.`

New parameter groups in Valkey 8.2NameDetailsDescriptionsearch-fanout-target-mode (added in 8.2)

Default: client

Type: string

Modifiable: Yes

Changes Take Effect: Immediately

The search-fanout-target-mode configuration parameter controls how search queries are distributed across nodes in a Valkey cluster environment. This setting accepts two values: "throughput" which optimizes for maximum throughput by randomly distributing search queries across all cluster nodes regardless of client type or READONLY status, and "client" which respects client connection characteristics by routing non-READONLY clients to primary nodes only, READONLY clients on replica connections to replica nodes only, and READONLY clients on primary connections randomly across all nodes.

The default behavior is "client' mode, meaning the system will respect client connection types and READONLY status for query routing decisions. Use throughput mode for high-volume search workloads where maximum cluster resource utilization is desired, and client mode when you want to maintain read/write separation and respect application-level READONLY connection patterns.

search-default-timeout-ms

Default: 50000

Permitted values: 1 to 60000

Type: integer

Modifiable: Yes

Changes Take Effect: Immediately

The default Valkey search query timeout (in milliseconds).search-enable-partial-results

Default: yes

Permitted values: yes, no

Type: boolean

Modifiable: Yes

Changes Take Effect: Immediately

Configures the query failure behavior for Valkey search. When enabled, search queries will return partial results if timeouts occur on one or more shards. When disabled, any shard timeout will cause the entire search query to fail and return an error.

### Valkey 8.1 parameter changes

**Parameter group family:** valkey8

###### Note

- Valkey 8.1 parameter changes don't apply to Valkey 8.0

- Valkey 8.0 and above parameter groups are incompatible with Redis OSS 7.2.4.

- in Valkey 8.1, the following commands are unavailable for serverless caches: `commandlog`,
`commandlog get`, `commandlog help`, `commandlog len`, and `commandlog reset.`

New parameter groups in Valkey 8.1NameDetailsDescription

commandlog-request-larger-than (added in 8.1)

Default: 1048576

Type: integer

Modifiable: Yes

Changes Take Effect: Immediately

The maximum size, in bytes, for requests to be logged by the Valkey Command Log feature.

commandlog-large-request-max-len (added in 8.1)

Default: 128

Permitted values: 0-1024

Type: integer

Modifiable: Yes

Changes Take Effect: Immediately

The maximum length of the Valkey Command Log for requests.

commandlog-reply-larger-than (added in 8.1)

Default: 1048576

Type: integer

Modifiable: Yes

Changes Take Effect: Immediately

The maximum size, in bytes, for responses to be logged by the Valkey Command Log feature.

commandlog-large-reply-max-len (added in 8.1)

Default: 128

Permitted values: 0-1024

Type: integer

Modifiable: Yes

Changes Take Effect: Immediately

The maximum length of the Valkey Command Log for responses.

### Valkey 8.0 parameter changes

**Parameter group family:** valkey8

###### Note

Redis OSS 7.2.4 is incompatible with Valkey 8 and above parameter groups.

Specific parameter changes in Valkey 8.0NameDetailsDescription

repl-backlog-size

Default: 10485760

Type: integer

Modifiable: Yes

Changes Take Effect: Immediately

The size, in bytes, of the primary node backlog buffer. The backlog is used for recording updates to data at the primary node. When a read replica connects to the primary, it attempts to perform a partial sync (psync), where it applies data from the backlog to catch up with the primary node. If the psync fails, then a full sync is required.

The minimum value for this parameter is 16384.

Note: Beginning with Redis OSS 2.8.22, this parameter applies to the primary cluster as well as the read replicas.

maxmemory-samples

Default: 3

Permitted values: 1 to 64

Type: integer

Modifiable: Yes

Changes Take Effect: Immediately

For least-recently-used (LRU) and time-to-live (TTL) calculations, this parameter represents the sample size of keys to check. By default, Redis OSS chooses 3 keys and uses the one that was used least recently.

New parameter groups in Valkey 8.0NameDetailsDescription

extended-redis-compatibility

Permitted values: yes, no

Default: yes

Type: boolean

Modifiable: Yes

Changes take place: immediately

Extended Redis OSS compatibility mode makes Valkey pretend to be Redis OSS 7.2. Enable this only if you have problems with tools or clients.

Customer-facing impacts:

- `LOADING` \- Redis OSS is loading the dataset in memory

- `BUSY` \- Redis OSS is busy

- `MISCONF` \- Redis OSS is configured in either of these ways:

- The `HELLO` command returns "server" => "redis" and "version" => "7.2.4" (our Redis OSS compatibility version).

- The `INFO` field for mode is called "redis\_mode".

Removed parameter groups in Valkey 8.0NameDetailsDescription

lazyfree-lazy-eviction

Permitted values: yes, no

Default: no

Type: boolean

Modifiable: Yes

Changes take place: immediately

Performs an asynchronous delete on evictions.

lazyfree-lazy-expire

Permitted values: yes, no

Default: no

Type: boolean

Modifiable: Yes

Changes take place: immediately

Performs an asynchronous delete on expired keys.

lazyfree-lazy-server-del

Permitted values: yes, no

Default: no

Type: boolean

Modifiable: Yes

Changes take place: immediately

Performs an asynchronous delete for commands which update values.

lazyfree-lazy-user-del

Default: no

Type: string

Modifiable: Yes

Changes take effect: Immediately across all nodes in the cluster

When the value is set to yes, the DEL command acts the same as UNLINK.

replica-lazy-flush

Default: yes

Type: boolean

Modifiable: No

Former name: slave-lazy-flush

Performs an asynchronous flushDB during replica sync.

### Valkey 7.2 and Redis OSS 7 parameter changes

**Parameter group family:** valkey7

Valkey 7.2 default parameter groups are as follows:

- `default.valkey7` – Use this parameter group, or one
derived from it, for Valkey (cluster mode disabled) clusters and replication groups.

- `default.valkey7.cluster.on` – Use this parameter group,
or one derived from it, for Valkey (cluster mode enabled) clusters and replication
groups.

**Parameter group family:** redis7

Redis OSS 7 default parameter groups are as follows:

- `default.redis7` – Use this parameter group, or one
derived from it, for Redis OSS (cluster mode disabled) clusters and replication groups.

- `default.redis7.cluster.on` – Use this parameter group,
or one derived from it, for Redis OSS (cluster mode enabled) clusters and replication
groups.

**Specific parameter changes**

Parameters added in Redis OSS 7 are as follows. Valkey 7.2 also supports these parameters.

Name  Details Description `cluster-allow-pubsubshard-when-down`

Permitted values: `yes`, `no`

Default: `yes`

Type: string

Modifiable: Yes

Changes take effect: Immediately across all nodes in the
cluster.

When set to the default of yes, allows nodes to serve
pubsub shard traffic while the cluster is in a down state, as
long as it believes it owns the slots.

`cluster-preferred-endpoint-type`

Permitted values: `ip`,
`tls-dynamic`

Default: `tls-dynamic`

Type: string

Modifiable: Yes

Changes take effect: Immediately across all nodes in the
cluster.

This value controls what endpoint is returned for
MOVED/ASKING requests as well as the endpoint field for
`CLUSTER SLOTS` and `CLUSTER SHARDS`.
When the value is set to ip, the node will advertise its ip
address. When the value is set to tls-dynamic, the node will
advertise a hostname when encryption-in-transit is enabled and
an ip address otherwise.

`latency-tracking`

Permitted values: `yes`, `no`

Default: `no`

Type: string

Modifiable: Yes

Changes take effect: Immediately across all nodes in the
cluster.

When set to yes tracks the per command latencies and
enables exporting the percentile distribution via the
`INFO` latency statistics command, and cumulative
latency distributions (histograms) via the `LATENCY`
command.

`hash-max-listpack-entries`

Permitted values: `0+`

Default: `512`

Type: integer

Modifiable: Yes

Changes take effect: Immediately across all nodes in the
cluster.

The maximum number of hash entries in order for the dataset
to be compressed.

`hash-max-listpack-value`

Permitted values: `0+`

Default: `64`

Type: integer

Modifiable: Yes

Changes take effect: Immediately across all nodes in the
cluster.

The threshold of biggest hash entries in order for the
dataset to be compressed.

`zset-max-listpack-entries`

Permitted values: `0+`

Default: `128`

Type: integer

Modifiable: Yes

Changes take effect: Immediately across all nodes in the
cluster.

The maximum number of sorted set entries in order for the
dataset to be compressed.

`zset-max-listpack-value`

Permitted values: `0+`

Default: `64`

Type: integer

Modifiable: Yes

Changes take effect: Immediately across all nodes in the
cluster.

The threshold of biggest sorted set entries in order for
the dataset to be compressed.

Parameters changed in Redis OSS 7 are as follows.

Name  Details Description `activerehashing`

Modifiable: `no`. In Redis OSS 7, this parameter is
hidden and enabled by default. In order to disable it, you need
to create a [support case](https://console.aws.amazon.com/support/home).

Modifiable was yes.

Parameters removed in Redis OSS 7 are as follows.

Name  Details Description `hash-max-ziplist-entries`

Permitted values: `0+`

Default: `512`

Type: integer

Modifiable: Yes

Changes take effect: Immediately across all nodes in the
cluster.

Use `listpack` instead of `ziplist`
for representing small hash encoding

`hash-max-ziplist-value`

Permitted values: `0+`

Default: `64`

Type: integer

Modifiable: Yes

Changes take effect: Immediately across all nodes in the
cluster.

Use `listpack` instead of `ziplist`
for representing small hash encoding

`zset-max-ziplist-entries`

Permitted values: `0+`

Default: `128`

Type: integer

Modifiable: Yes

Changes take effect: Immediately across all nodes in the
cluster.

Use `listpack` instead of `ziplist`
for representing small hash encoding.

`zset-max-ziplist-value`

Permitted values: `0+`

Default: `64`

Type: integer

Modifiable: Yes

Changes take effect: Immediately across all nodes in the
cluster.

Use `listpack` instead of `ziplist`
for representing small hash encoding.

`list-max-ziplist-size`

Permitted values:

Default: `-2`

Type: integer

Modifiable: Yes

Changes take effect: Immediately across all nodes in the
cluster.

The number of entries allowed per internal list
node.

### Redis OSS 6.x parameter changes

**Parameter group family:** redis6.x

Redis OSS 6.x default parameter groups are as follows:

- `default.redis6.x` – Use this parameter group, or one
derived from it, for Valkey or Redis OSS (cluster mode disabled) clusters and replication groups.

- `default.redis6.x.cluster.on` – Use this parameter
group, or one derived from it, for Valkey or Redis OSS (cluster mode enabled) clusters and replication
groups.

###### Note

In Redis OSS engine version 6.2, when the r6gd node family was introduced for use
with [Data tiering in ElastiCache](data-tiering.md), only
_noeviction_, _volatile-lru_ and
_allkeys-lru_ max-memory policies are supported with r6gd
node types.

For more information, see [ElastiCache version 6.2 for Redis OSS (enhanced)](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/engine-versions.html#redis-version-6.2) and [ElastiCache version 6.0 for Redis OSS (enhanced)](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/engine-versions.html#redis-version-6.0).

Parameters added in Redis OSS 6.x are as follows.

Details Description `acl-pubsub-default (added in 6.2)`

Permitted values: `resetchannels`,
`allchannels`

Default: `allchannels`

Type: string

Modifiable: Yes

Changes take effect: The existing Redis OSS users associated to
the cluster will continue to have existing permissions. Either
update the users or reboot the cluster to update the existing
Redis OSS users.

Default pubsub channel permissions for ACL users deployed
to this cluster.

`cluster-allow-reads-when-down (added in 6.0)`

Default: no

Type: string

Modifiable: Yes

Changes take effect: Immediately across all nodes in the
cluster

When set to yes, a Redis OSS (cluster mode enabled) replication
group continues to process read commands even when a node is not
able to reach a quorum of primaries.

When set to the default of no, the replication group rejects
all commands. We recommend setting this value to yes if you are
using a cluster with fewer than three node groups or your
application can safely handle stale reads.

`tracking-table-max-keys (added in 6.0)`

Default: 1,000,000

Type: number

Modifiable: Yes

Changes take effect: Immediately across all nodes in the
cluster

To assist client-side caching, Redis OSS supports tracking
which clients have accessed which keys.

When the tracked key is modified, invalidation messages are
sent to all clients to notify them their cached values are no
longer valid. This value enables you to specify the upper bound
of this table. After this parameter value is exceeded, clients
are sent invalidation randomly. This value should be tuned to
limit memory usage while still keeping track of enough keys.
Keys are also invalidated under low memory conditions.

`acllog-max-len (added in 6.0)`

Default: 128

Type: number

Modifiable: Yes

Changes take effect: Immediately across all nodes in the
cluster

This value corresponds to the max number of entries in the
ACL log.

`active-expire-effort (added in 6.0)`

Default: 1

Type: number

Modifiable: Yes

Changes take effect: Immediately across all nodes in the
cluster

Redis OSS deletes keys that have exceeded their time to live by
two mechanisms. In one, a key is accessed and is found to be
expired. In the other, a periodic job samples keys and causes
those that have exceeded their time to live to expire. This
parameter defines the amount of effort that Redis OSS uses to expire
items in the periodic job.

The default value of 1 tries to avoid having more than 10
percent of expired keys still in memory. It also tries to avoid
consuming more than 25 percent of total memory and to add
latency to the system. You can increase this value up to 10 to
increase the amount of effort spent on expiring keys. The
tradeoff is higher CPU and potentially higher latency. We
recommend a value of 1 unless you are seeing high memory usage
and can tolerate an increase in CPU utilization.

`lazyfree-lazy-user-del (added in 6.0)`

Default: no

Type: string

Modifiable: Yes

Changes take effect: Immediately across all nodes in the
cluster

When the value is set to yes, the `DEL` command
acts the same as `UNLINK`.

Parameters removed in Redis OSS 6.x are as follows.

Name  Details Description `lua-replicate-commands`

Permitted values: yes/no

Default: yes

Type: boolean

Modifiable: Yes

Changes take effect: Immediately

Always enable Lua effect replication or not in Lua scripts

### Redis OSS 5.0.3 parameter changes

**Parameter group family:** redis5.0

Redis OSS 5.0 default parameter groups

- `default.redis5.0` – Use this parameter group, or one
derived from it, for Valkey or Redis OSS (cluster mode disabled) clusters and replication groups.

- `default.redis5.0.cluster.on` – Use this parameter
group, or one derived from it, for Valkey or Redis OSS (cluster mode enabled) clusters and replication
groups.

Parameters added in Redis OSS 5.0.3 Name  Details Description `rename-commands`

Default: none

Type: string

Modifiable: Yes

Changes take effect: Immediately across all nodes in the
cluster

A space-separated list of renamed Redis OSS commands. The following
is a restricted list of commands available for renaming:

`APPEND AUTH BITCOUNT BITFIELD BITOP BITPOS BLPOP BRPOP
										BRPOPLPUSH BZPOPMIN BZPOPMAX CLIENT CLUSTER COMMAND DBSIZE
										DECR DECRBY DEL DISCARD DUMP ECHO EVAL EVALSHA EXEC EXISTS
										EXPIRE EXPIREAT FLUSHALL FLUSHDB GEOADD GEOHASH GEOPOS
										GEODIST GEORADIUS GEORADIUSBYMEMBER GET GETBIT GETRANGE
										GETSET HDEL HEXISTS HGET HGETALL HINCRBY HINCRBYFLOAT HKEYS
										HLEN HMGET HMSET HSET HSETNX HSTRLEN HVALS INCR INCRBY
										INCRBYFLOAT INFO KEYS LASTSAVE LINDEX LINSERT LLEN LPOP
										LPUSH LPUSHX LRANGE LREM LSET LTRIM MEMORY MGET MONITOR MOVE
										MSET MSETNX MULTI OBJECT PERSIST PEXPIRE PEXPIREAT PFADD
										PFCOUNT PFMERGE PING PSETEX PSUBSCRIBE PUBSUB PTTL PUBLISH
										PUNSUBSCRIBE RANDOMKEY READONLY READWRITE RENAME RENAMENX
										RESTORE ROLE RPOP RPOPLPUSH RPUSH RPUSHX SADD SCARD SCRIPT
										SDIFF SDIFFSTORE SELECT SET SETBIT SETEX SETNX SETRANGE
										SINTER SINTERSTORE SISMEMBER SLOWLOG SMEMBERS SMOVE SORT
										SPOP SRANDMEMBER SREM STRLEN SUBSCRIBE SUNION SUNIONSTORE
										SWAPDB TIME TOUCH TTL TYPE UNSUBSCRIBE UNLINK UNWATCH WAIT
										WATCH ZADD ZCARD ZCOUNT ZINCRBY ZINTERSTORE ZLEXCOUNT
										ZPOPMAX ZPOPMIN ZRANGE ZRANGEBYLEX ZREVRANGEBYLEX
										ZRANGEBYSCORE ZRANK ZREM ZREMRANGEBYLEX ZREMRANGEBYRANK
										ZREMRANGEBYSCORE ZREVRANGE ZREVRANGEBYSCORE ZREVRANK ZSCORE
										ZUNIONSTORE SCAN SSCAN HSCAN ZSCAN XINFO XADD XTRIM XDEL
										XRANGE XREVRANGE XLEN XREAD XGROUP XREADGROUP XACK XCLAIM
										XPENDING GEORADIUS_RO GEORADIUSBYMEMBER_RO LOLWUT XSETID
										SUBSTR`

For more information, see [ElastiCache version 5.0.6 for Redis OSS (enhanced)](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/engine-versions.html#redis-version-5-0.6).

### Redis OSS 5.0.0 parameter changes

**Parameter group family:** redis5.0

Redis OSS 5.0 default parameter groups

- `default.redis5.0` – Use this parameter group, or one
derived from it, for Valkey or Redis OSS (cluster mode disabled) clusters and replication groups.

- `default.redis5.0.cluster.on` – Use this parameter
group, or one derived from it, for Valkey or Redis OSS (cluster mode enabled) clusters and replication
groups.

Parameters added in Redis OSS 5.0 Name  Details Description `stream-node-max-bytes`

Permitted values: 0+

Default: 4096

Type: integer

Modifiable: Yes

Changes take effect: Immediately

The stream data structure is a radix tree of nodes that encode
multiple items inside. Use this configuration to specify the maximum
size of a single node in radix tree in Bytes. If set to 0, the size
of the tree node is unlimited. `stream-node-max-entries`

Permitted values: 0+

Default: 100

Type: integer

Modifiable: Yes

Changes take effect: Immediately

The stream data structure is a radix tree of nodes that encode
multiple items inside. Use this configuration to specify the maximum
number of items a single node can contain before switching to a new
node when appending new stream entries. If set to 0, the number of
items in the tree node is unlimited `active-defrag-max-scan-fields`

Permitted values: 1 to 1000000

Default: 1000

Type: integer

Modifiable: Yes

Changes take effect: Immediately

Maximum number of set/hash/zset/list fields that will be
processed from the main dictionary scan `lua-replicate-commands`

Permitted values: yes/no

Default: yes

Type: boolean

Modifiable: Yes

Changes take effect: Immediately

Always enable Lua effect replication or not in Lua scripts
`replica-ignore-maxmemory`

Default: yes

Type: boolean

Modifiable: No

Determines if replica ignores `maxmemory` setting by
not evicting items independent from the primary

Redis OSS has renamed several parameters in engine version 5.0 in response to
community feedback. For more information, see [What's New in Redis OSS 5?](https://aws.amazon.com/redis/Whats_New_Redis5). The following table lists the new names and how they map to previous
versions.

Parameters renamed in Redis OSS 5.0 Name  Details Description `replica-lazy-flush`

Default: yes

Type: boolean

Modifiable: No

Former name: slave-lazy-flush

Performs an asynchronous flushDB during replica sync.`client-output-buffer-limit-replica-hard-limit`

Default: For values see [Redis OSS node-type specific parameters](#ParameterGroups.Redis.NodeSpecific)

Type: integer

Modifiable: No

Former name:
client-output-buffer-limit-slave-hard-limit

For Redis OSS read replicas: If a client's output buffer reaches the
specified number of bytes, the client will be disconnected.`client-output-buffer-limit-replica-soft-limit`

Default: For values see [Redis OSS node-type specific parameters](#ParameterGroups.Redis.NodeSpecific)

Type: integer

Modifiable: No

Former name:
client-output-buffer-limit-slave-soft-limit

For Redis OSS read replicas: If a client's output buffer reaches the
specified number of bytes, the client will be disconnected, but only
if this condition persists for
`client-output-buffer-limit-replica-soft-seconds`.`client-output-buffer-limit-replica-soft-seconds`

Default: 60

Type: integer

Modifiable: No

Former name:
client-output-buffer-limit-slave-soft-seconds

For Redis OSS read replicas: If a client's output buffer remains at
`client-output-buffer-limit-replica-soft-limit` bytes
for longer than this number of seconds, the client will be
disconnected.`replica-allow-chaining`

Default: no

Type: string

Modifiable: No

Former name: slave-allow-chaining

Determines whether a read replica in Redis OSS can have read replicas
of its own.`min-replicas-to-write`

Default: 0

Type: integer

Modifiable: Yes

Former name: min-slaves-to-write

Changes Take Effect: Immediately

The minimum number of read replicas which must be available
in order for the primary node to accept writes from clients. If
the number of available replicas falls below this number, then
the primary node will no longer accept write requests.

If either this parameter or min-replicas-max-lag is 0, then
the primary node will always accept writes requests, even if no
replicas are available.

`min-replicas-max-lag `

Default: 10

Type: integer

Modifiable: Yes

Former name: min-slaves-max-lag

Changes Take Effect: Immediately

The number of seconds within which the primary node must
receive a ping request from a read replica. If this amount of
time passes and the primary does not receive a ping, then the
replica is no longer considered available. If the number of
available replicas drops below min-replicas-to-write, then the
primary will stop accepting writes at that point.

If either this parameter or min-replicas-to-write is 0, then
the primary node will always accept write requests, even if no
replicas are available.

`close-on-replica-write `

Default: yes

Type: boolean

Modifiable: Yes

Former name: close-on-slave-write

Changes Take Effect: Immediately

If enabled, clients who attempt to write to a read-only
replica will be disconnected.

Parameters removed in Redis OSS 5.0 Name  Details Description `repl-timeout`

Default: 60

Modifiable: No

Parameter is not available in this version.

### Redis OSS 4.0.10 parameter changes

**Parameter group family:** redis4.0

Redis OSS 4.0.x default parameter groups

- `default.redis4.0` – Use this parameter group, or one
derived from it, for Valkey or Redis OSS (cluster mode disabled) clusters and replication groups.

- `default.redis4.0.cluster.on` – Use this parameter
group, or one derived from it, for Valkey or Redis OSS (cluster mode enabled) clusters and replication
groups.

Parameters changed in Redis OSS 4.0.10 Name  Details Description `maxmemory-policy`

Permitted values: `allkeys-lru`,
`volatile-lru`,
`allkeys-lfu`,
`volatile-lfu`,
`allkeys-random`, `volatile-random`,
`volatile-ttl`, `noeviction`

Default: volatile-lru

Type: string

Modifiable: Yes

Changes take place: immediately

`maxmemory-policy` was added in version 2.6.13. In
version 4.0.10 two new permitted values are added:
`allkeys-lfu`, which will evict any key using
approximated LFU, and `volatile-lfu`, which will evict
using approximated LFU among the keys with an expire set. In version
6.2, when the r6gd node family was introduced for use with
data-tiering, only `noeviction`,
`volatile-lru` and `allkeys-lru`
max-memory policies are supported with r6gd node types.

Parameters added in Redis OSS 4.0.10 Name  Details Description **Async**
**deletion parameters**`lazyfree-lazy-eviction`

Permitted values: yes/no

Default: no

Type: boolean

Modifiable: Yes

Changes take place: immediately

Performs an asynchronous delete on evictions.`lazyfree-lazy-expire`

Permitted values: yes/no

Default: no

Type: boolean

Modifiable: Yes

Changes take place: immediately

Performs an asynchronous delete on expired keys.`lazyfree-lazy-server-del`

Permitted values: yes/no

Default: no

Type: boolean

Modifiable: Yes

Changes take place: immediately

Performs an asynchronous delete for commands which update
values.`slave-lazy-flush`

Permitted values: N/A

Default: no

Type: boolean

Modifiable: No

Changes take place: N/APerforms an asynchronous flushDB during slave sync.**LFU**
**parameters**`lfu-log-factor`

Permitted values: any integer > 0

Default: 10

Type: integer

Modifiable: Yes

Changes take place: immediately

Set the log factor, which determines the number of key hits to
saturate the key counter.`lfu-decay-time`

Permitted values: any integer

Default: 1

Type: integer

Modifiable: Yes

Changes take place: immediately

The amount of time in minutes to decrement the key
counter.**Active**
**defragmentation parameters**`activedefrag`

Permitted values: yes/no

Default: no

Type: boolean

Modifiable: Yes

Changes take place: immediately

Enables active defragmentation.

###### Note

In Valkey and Redis OSS versions 7.0 and above, AWS may automatically perform defragmentation when operationally necessary, regardless of this setting.

`active-defrag-ignore-bytes`

Permitted values: 10485760-104857600

Default: 104857600

Type: integer

Modifiable: Yes

Changes take place: immediately

Minimum amount of fragmentation waste to start active
defrag.`active-defrag-threshold-lower`

Permitted values: 1-100

Default: 10

Type: integer

Modifiable: Yes

Changes take place: immediately

Minimum percentage of fragmentation to start active
defrag.`active-defrag-threshold-upper`

Permitted values: 1-100

Default: 100

Type: integer

Modifiable: Yes

Changes take place: immediately

Maximum percentage of fragmentation at which we use maximum
effort.`active-defrag-cycle-min`

Permitted values: 1-75

Default: 25

Type: integer

Modifiable: Yes

Changes take place: immediately

Minimal effort for defrag in CPU percentage.`active-defrag-cycle-max`

Permitted values: 1-75

Default: 75

Type: integer

Modifiable: Yes

Changes take place: immediately

Maximal effort for defrag in CPU percentage.**Client**
**output buffer parameters**`client-query-buffer-limit`

Permitted values: 1048576-1073741824

Default: 1073741824

Type: integer

Modifiable: Yes

Changes take place: immediately

Max size of a single client query buffer.`proto-max-bulk-len`

Permitted values: 1048576-536870912

Default: 536870912

Type: integer

Modifiable: Yes

Changes take place: immediately

Max size of a single element request.

### Redis OSS 3.2.10 parameter changes

**Parameter group family:** redis3.2

ElastiCache for Redis OSS 3.2.10 there are no additional parameters supported.

### Redis OSS 3.2.6 parameter changes

**Parameter group family:** redis3.2

For Redis OSS 3.2.6 there are no additional parameters supported.

### Redis OSS 3.2.4 parameter changes

**Parameter group family:** redis3.2

Beginning with Redis OSS 3.2.4 there are two default parameter groups.

- `default.redis3.2` – When running Redis OSS 3.2.4, specify
this parameter group or one derived from it, if you want to create a
Valkey or Redis OSS (cluster mode disabled) replication group and still use the additional features of
Redis OSS 3.2.4.

- `default.redis3.2.cluster.on` – Specify this parameter
group or one derived from it, when you want to create a Valkey or Redis OSS (cluster mode enabled)
replication group.

###### Topics

- [New parameters for Redis OSS 3.2.4](#ParameterGroups.Redis.3-2-4.New)

- [Parameters changed in Redis OSS 3.2.4 (enhanced)](#ParameterGroups.Redis.3-2-4.Changed)

#### New parameters for Redis OSS 3.2.4

**Parameter group family:** redis3.2

For Redis OSS 3.2.4 the following additional parameters are supported.

Name  Details Description `list-max-ziplist-size`

Default: -2

Type: integer

Modifiable: No

Lists are encoded in a special way to save space. The number
of entries allowed per internal list node can be specified as a
fixed maximum size or a maximum number of elements. For a fixed
maximum size, use -5 through -1, meaning:

- -5: max size: 64 Kb - not recommended for normal
workloads

- -4: max size: 32 Kb - not recommended

- -3: max size: 16 Kb - not recommended

- -2: max size: 8 Kb - recommended

- -1: max size: 4 Kb - recommended

- Positive numbers mean store up to exactly that
number of elements per list node.

`list-compress-depth`

Default: 0

Type: integer

Modifiable: Yes

Changes Take Effect: Immediately

Lists may also be compressed. Compress depth is the number of
quicklist ziplist nodes from each side of the list to exclude
from compression. The head and tail of the list are always
uncompressed for fast push and pop operations. Settings are:

- 0: Disable all compression.

- 1: Start compressing with the 1st node in from the
head and tail.

\[head\]->node->node->...->node->\[tail\]

All nodes except \[head\] and \[tail\]
compress.

- 2: Start compressing with the 2nd node in from the
head and tail.

\[head\]->\[next\]->node->node->...->node->\[prev\]->\[tail\]

\[head\], \[next\], \[prev\], \[tail\] do not compress.
All other nodes compress.

- Etc.

`cluster-enabled`

Default: no/yes \*

Type: string

Modifiable: No

Indicates whether this is a Valkey or Redis OSS (cluster mode enabled) replication
group in cluster mode (yes) or a Valkey or Redis OSS (cluster mode enabled) replication
group in non-cluster mode (no). Valkey or Redis OSS (cluster mode enabled) replication
groups in cluster mode can partition their data across up to
500 node groups.

\\* Redis OSS 3.2. _x_ has two default
parameter groups.

- `default.redis3.2` – default
value `no`.

- `default.redis3.2.cluster.on` –
default value `yes`.

.

`cluster-require-full-coverage`

Default: no

Type: boolean

Modifiable: yes

Changes Take Effect: Immediately

When set to `yes`, Valkey or Redis OSS (cluster mode enabled) nodes in
cluster mode stop accepting queries if they detect there is
at least one hash slot uncovered (no available node is
serving it). This way if the cluster is partially down, the
cluster becomes unavailable. It automatically becomes
available again as soon as all the slots are covered
again.

However, sometimes you want the subset of the cluster
which is working to continue to accept queries for the part
of the key space that is still covered. To do so, just set
the `cluster-require-full-coverage` option to
`no`.

`hll-sparse-max-bytes`

Default: 3000

Type: integer

Modifiable: Yes

Changes Take Effect: Immediately

HyperLogLog sparse representation bytes limit. The
limit includes the 16 byte header. When a HyperLogLog using
the sparse representation crosses this limit, it is
converted into the dense representation.

A value greater than 16000 is not recommended, because at
that point the dense representation is more memory
efficient.

We recommend a value of about 3000 to have the benefits of
the space-efficient encoding without slowing down PFADD too
much, which is O(N) with the sparse encoding. The value can
be raised to ~10000 when CPU is not a concern, but space is,
and the data set is composed of many HyperLogLogs with
cardinality in the 0 - 15000 range.

`reserved-memory-percent`

Default: 25

Type: integer

Modifiable: Yes

Changes Take Effect: Immediately

The percent of a node's memory reserved for nondata use.
By default, the Redis OSS data footprint grows until it consumes
all of the node's memory. If this occurs, then node
performance will likely suffer due to excessive memory
paging. By reserving memory, you can set aside some of the
available memory for non-Redis OSS purposes to help reduce the
amount of paging.

This parameter is specific to ElastiCache, and is not part of
the standard Redis OSS distribution.

For more information, see `reserved-memory` and
[Managing reserved memory for Valkey and Redis OSS](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/redis-memory-management.html).

#### Parameters changed in Redis OSS 3.2.4 (enhanced)

**Parameter group family:** redis3.2

For Redis OSS 3.2.4 the following parameters were changed.

Name  Details Change `activerehashing`

Modifiable: Yes if the parameter group is not
associated with any clusters. Otherwise,
no.

Modifiable was No.

`databases`

Modifiable: Yes if the parameter group is not
associated with any clusters. Otherwise,
no.

Modifiable was No.

`appendonly`

Default: off

Modifiable: No

If you want to upgrade from an earlier Redis OSS version,
you must first turn `appendonly`
off.

`appendfsync`

Default: off

Modifiable: No

If you want to upgrade from an earlier Redis OSS version,
you must first turn `appendfsync`
off.

`repl-timeout`

Default: 60

Modifiable: No

Is now unmodifiable with a default of 60.`tcp-keepalive`

Default: 300

Default was 0.

`list-max-ziplist-entries`

Parameter is no longer available.

`list-max-ziplist-value`

Parameter is no longer available.

### Redis OSS 2.8.24 (enhanced) added parameters

**Parameter group family:** redis2.8

For Redis OSS 2.8.24 there are no additional parameters supported.

### Redis OSS 2.8.23 (enhanced) added parameters

**Parameter group family:** redis2.8

For Redis OSS 2.8.23 the following additional parameter is supported.

Name  Details Description `close-on-slave-write `

Default: yes

Type: string (yes/no)

Modifiable: Yes

Changes Take Effect: Immediately

If enabled, clients who attempt to write to a read-only
replica will be disconnected.

#### How close-on-slave-write works

The `close-on-slave-write` parameter is introduced by Amazon ElastiCache to
give you more control over how your cluster responds when a primary node and a
read replica node swap roles due to promoting a read replica to primary.

![Image: close-on-replica-write, everything working fine](https://docs.aws.amazon.com/images/AmazonElastiCache/latest/dg/images/ElastiCache-close-on-slave-write-01.png)

If the read-replica cluster is promoted to primary for any reason other than a
Multi-AZ enabled replication group failing over, the client will continue trying
to write to endpoint A. Because endpoint A is now the endpoint for a
read-replica, these writes will fail. This is the behavior for Redis OSS before
ElastiCache introducing `close-on-replica-write` and the behavior if you
disable `close-on-replica-write`.

![Image: close-on-slave-write, writes failing](https://docs.aws.amazon.com/images/AmazonElastiCache/latest/dg/images/ElastiCache-close-on-slave-write-02.png)

With `close-on-replica-write` enabled, any time a client attempts
to write to a read-replica, the client connection to the cluster is closed. Your
application logic should detect the disconnection, check the DNS table, and
reconnect to the primary endpoint, which now would be endpoint B.

![Image: close-on-slave-write, writing to new primary cluster](https://docs.aws.amazon.com/images/AmazonElastiCache/latest/dg/images/ElastiCache-close-on-slave-write-03.png)

#### When you might disable close-on-replica-write

If disabling `close-on-replica-write` results in writes to the
failing cluster, why disable `close-on-replica-write`?

As previously mentioned, with `close-on-replica-write` enabled, any
time a client attempts to write to a read-replica the client connection to the
cluster is closed. Establishing a new connection to the node takes time. Thus,
disconnecting and reconnecting as a result of a write request to the replica
also affects the latency of read requests that are served through the same
connection. This effect remains in place until a new connection is established.
If your application is especially read-heavy or very latency-sensitive, you
might keep your clients connected to avoid degrading read performance.

### Redis OSS 2.8.22 (enhanced) added parameters

**Parameter group family:** redis2.8

For Redis OSS 2.8.22 there are no additional parameters supported.

###### Important

- Beginning with Redis OSS version 2.8.22, `repl-backlog-size`
applies to the primary cluster as well as to replica clusters.

- Beginning with Redis OSS version 2.8.22, the `repl-timeout`
parameter is not supported. If it is changed, ElastiCache will overwrite with
the default (60s), as we do with `appendonly`.

The following parameters are no longer supported.

- _appendonly_

- _appendfsync_

- _repl-timeout_

### Redis OSS 2.8.21 added parameters

**Parameter group family:** redis2.8

For Redis OSS 2.8.21, there are no additional parameters supported.

### Redis OSS 2.8.19 added parameters

**Parameter group family:** redis2.8

For Redis OSS 2.8.19 there are no additional parameters supported.

### Redis OSS 2.8.6 added parameters

**Parameter group family:** redis2.8

For Redis OSS 2.8.6 the following additional parameters are supported.

Name  Details  Description `min-slaves-max-lag `

Default: 10

Type: integer

Modifiable: Yes

Changes Take Effect: Immediately

The number of seconds within which the primary node must
receive a ping request from a read replica. If this amount of
time passes and the primary does not receive a ping, then the
replica is no longer considered available. If the number of
available replicas drops below min-slaves-to-write, then the
primary will stop accepting writes at that point.

If either this parameter or min-slaves-to-write is 0, then the
primary node will always accept writes requests, even if no
replicas are available.

`min-slaves-to-write`

Default: 0

Type: integer

Modifiable: Yes

Changes Take Effect: Immediately

The minimum number of read replicas which must be available
in order for the primary node to accept writes from clients. If
the number of available replicas falls below this number, then
the primary node will no longer accept write requests.

If either this parameter or min-slaves-max-lag is 0, then the
primary node will always accept writes requests, even if no
replicas are available.

`notify-keyspace-events`

Default: (an empty string)

Type: string

Modifiable: Yes

Changes Take Effect: Immediately

The types of keyspace events that Redis OSS can notify clients
of. Each event type is represented by a single letter:

- **K** — Keyspace
events, published with a prefix of
_\_\_keyspace@<db>\_\__

- **E** — Key-event
events, published with a prefix of
_\_\_keyevent@<db>\_\__

- **g** — Generic,
non-specific commands such as _DEL_,
_EXPIRE_,
_RENAME_, etc.

- **$** — String
commands

- **l** — List
commands

- **s** — Set
commands

- **h** — Hash
commands

- **z** — Sorted set
commands

- **x** — Expired
events (events generated every time a key
expires)

- **e** — Evicted
events (events generated when a key is evicted for
maxmemory)

- **A** — An alias
for _g$lshzxe_

You can have any combination of these event types. For
example, _AKE_ means that Redis OSS can publish
notifications of all event types.

Do not use any characters other than those listed above;
attempts to do so will result in error messages.

By default, this parameter is set to an empty string, meaning
that keyspace event notification is disabled.

`repl-backlog-size`

Default: 1048576

Type: integer

Modifiable: Yes

Changes Take Effect: Immediately

The size, in bytes, of the primary node backlog buffer. The
backlog is used for recording updates to data at the primary
node. When a read replica connects to the primary, it attempts
to perform a partial sync ( `psync`), where it applies
data from the backlog to catch up with the primary node. If the
`psync` fails, then a full sync is
required.

The minimum value for this parameter is 16384.

###### Note

Beginning with Redis OSS 2.8.22, this parameter applies to the
primary cluster as well as the read replicas.

`repl-backlog-ttl`

Default: 3600

Type: integer

Modifiable: Yes

Changes Take Effect: Immediately

The number of seconds that the primary node will retain the
backlog buffer. Starting from the time the last replica node
disconnected, the data in the backlog will remain intact until
`repl-backlog-ttl` expires. If the replica has
not connected to the primary within this time, then the primary
will release the backlog buffer. When the replica eventually
reconnects, it will have to perform a full sync with the
primary.

If this parameter is set to 0, then the backlog buffer will
never be released.

`repl-timeout`

Default: 60

Type: integer

Modifiable: Yes

Changes Take Effect: Immediately

Represents the timeout period, in seconds, for:

- Bulk data transfer during synchronization, from the
read replica's perspective

- Primary node timeout from the replica's
perspective

- Replica timeout from the primary node's
perspective

### Redis OSS 2.6.13 parameters

**Parameter group family:** redis2.6

Redis OSS 2.6.13 was the first version of Redis OSS supported by ElastiCache. The following
table shows the Redis OSS 2.6.13 parameters that ElastiCache supports.

Name  Details  Description `activerehashing`

Default: yes

Type: string (yes/no)

Modifiable: Yes

Changes take place: At Creation

Determines whether to enable Redis' active rehashing
feature. The main hash table is rehashed ten times per second;
each rehash operation consumes 1 millisecond of CPU time.

This value is set when you create the parameter group. When
assigning a new parameter group to a cluster, this value must be
the same in both the old and new parameter
groups.

`appendonly`

Default: no

Type: string

Modifiable: Yes

Changes Take Effect: Immediately

Enables or disables Redis' append only file feature (AOF).
AOF captures any Redis OSS commands that change data in the cache,
and is used to recover from certain node failures.

The default value is _no_, meaning AOF is
turned off. Set this parameter to _yes_ to
enable AOF.

For more information, see [Mitigating Failures](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/disaster-recovery-resiliency.html#FaultTolerance).

###### Note

Append Only Files (AOF) is not supported for
cache.t1.micro and cache.t2.\* nodes. For nodes of this type,
the `appendonly` parameter value is ignored.

###### Note

For Multi-AZ replication groups, AOF is not
allowed.

`appendfsync`

Default: everysec

Type: string

Modifiable: Yes

Changes Take Effect: Immediately

When `appendonly` is set to yes, controls how often
the AOF output buffer is written to disk:

- _no_ — the buffer is flushed
to disk on an as-needed basis.

- _everysec_ — the buffer is
flushed once per second. This is the default.

- _always_ — the buffer is
flushed every time that data in the cluster is
modified.

- Appendfsync is not supported for versions 2.8.22 and
later.

`client-output-buffer-limit-normal-hard-limit`

Default: 0

Type: integer

Modifiable: Yes

Changes Take Effect: Immediately

If a client's output buffer reaches the specified number of
bytes, the client will be disconnected. The default is zero (no
hard limit).

`client-output-buffer-limit-normal-soft-limit`

Default: 0

Type: integer

Modifiable: Yes

Changes Take Effect: Immediately

If a client's output buffer reaches the specified number of
bytes, the client will be disconnected, but only if this condition
persists for
`client-output-buffer-limit-normal-soft-seconds`. The
default is zero (no soft limit).`client-output-buffer-limit-normal-soft-seconds`

Default: 0

Type: integer

Modifiable: Yes

Changes Take Effect: Immediately

If a client's output buffer remains at
`client-output-buffer-limit-normal-soft-limit` bytes
for longer than this number of seconds, the client will be
disconnected. The default is zero (no time limit).`client-output-buffer-limit-pubsub-hard-limit`

Default: 33554432

Type: integer

Modifiable: Yes

Changes Take Effect: Immediately

For Redis OSS publish/subscribe clients: If a client's output
buffer reaches the specified number of bytes, the client will be
disconnected.

`client-output-buffer-limit-pubsub-soft-limit`

Default: 8388608

Type: integer

Modifiable: Yes

Changes Take Effect: Immediately

For Redis OSS publish/subscribe clients: If a client's output buffer
reaches the specified number of bytes, the client will be
disconnected, but only if this condition persists for
`client-output-buffer-limit-pubsub-soft-seconds`.`client-output-buffer-limit-pubsub-soft-seconds`

Default: 60

Type: integer

Modifiable: Yes

Changes Take Effect: Immediately

For Redis OSS publish/subscribe clients: If a client's output buffer
remains at `client-output-buffer-limit-pubsub-soft-limit`
bytes for longer than this number of seconds, the client will be
disconnected.`client-output-buffer-limit-slave-hard-limit`

Default: For values see [Redis OSS node-type specific parameters](#ParameterGroups.Redis.NodeSpecific)

Type: integer

Modifiable: No

For Redis OSS read replicas: If a client's output buffer reaches the
specified number of bytes, the client will be disconnected.`client-output-buffer-limit-slave-soft-limit`

Default: For values see [Redis OSS node-type specific parameters](#ParameterGroups.Redis.NodeSpecific)

Type: integer

Modifiable: No

For Redis OSS read replicas: If a client's output buffer reaches the
specified number of bytes, the client will be disconnected, but only
if this condition persists for
`client-output-buffer-limit-slave-soft-seconds`.`client-output-buffer-limit-slave-soft-seconds`

Default: 60

Type: integer

Modifiable: No

For Redis OSS read replicas: If a client's output buffer remains at
`client-output-buffer-limit-slave-soft-limit` bytes
for longer than this number of seconds, the client will be
disconnected.`databases`

Default: 16

Type: integer

Modifiable: No

Changes take place: At Creation

The number of logical partitions the databases is split
into. We recommend keeping this value low.

This value is set when you create the parameter group. When
assigning a new parameter group to a cluster, this value must be
the same in both the old and new parameter
groups.

`hash-max-ziplist-entries`

Default: 512

Type: integer

Modifiable: Yes

Changes Take Effect: Immediately

Determines the amount of memory used for hashes. Hashes with
fewer than the specified number of entries are stored using a
special encoding that saves space.`hash-max-ziplist-value`

Default: 64

Type: integer

Modifiable: Yes

Changes Take Effect: Immediately

Determines the amount of memory used for hashes. Hashes with
entries that are smaller than the specified number of bytes are
stored using a special encoding that saves space.`list-max-ziplist-entries`

Default: 512

Type: integer

Modifiable: Yes

Changes Take Effect: Immediately

Determines the amount of memory used for lists. Lists with fewer
than the specified number of entries are stored using a special
encoding that saves space.`list-max-ziplist-value`

Default: 64

Type: integer

Modifiable: Yes

Changes Take Effect: Immediately

Determines the amount of memory used for lists. Lists with
entries that are smaller than the specified number of bytes are
stored using a special encoding that saves space.`lua-time-limit`

Default: 5000

Type: integer

Modifiable: No

The maximum execution time for a Lua script, in milliseconds,
before ElastiCache takes action to stop the script.

If
`lua-time-limit` is exceeded, all Redis OSS commands
will return an error of the form _\_\_\_\_-BUSY_.
Since this state can cause interference with many essential
Redis OSS operations, ElastiCache will first issue a _SCRIPT_
_KILL_ command. If this is unsuccessful, ElastiCache will
forcibly restart Redis OSS.

`maxclients` This value applies to all
instance types except those explicity specified

Default: 65000

Type: integer

Modifiable: No

The maximum number of clients that can be connected
at one time.

t2.medium Default: 20000

Type: integer

Modifiable: No

t2.small Default: 20000

Type: integer

Modifiable: No

t2.micro Default: 20000

Type: integer

Modifiable: No

t4g.micro Default: 20000

Type: integer

Modifiable: No

t3.medium Default: 46000

Type: integer

Modifiable: No

t3.small Default: 46000

Type: integer

Modifiable: No

t3.micro Default: 20000

Type: integer

Modifiable: No

`maxmemory-policy`

Default: volatile-lru

Type: string

Modifiable: Yes

Changes Take Effect: Immediately

The eviction policy for keys when maximum memory usage is
reached.

Valid values are: `volatile-lru | allkeys-lru |
										volatile-random | allkeys-random | volatile-ttl |
										noeviction`

For more information, see [Using Valkey or Redis OSS as an\
LRU cache](https://valkey.io/topics/lru-cache).

`maxmemory-samples`

Default: 3

Type: integer

Modifiable: Yes

Changes Take Effect: Immediately

For least-recently-used (LRU) and time-to-live (TTL)
calculations, this parameter represents the sample size of keys to
check. By default, Redis OSS chooses 3 keys and uses the one that was
used least recently.`reserved-memory`

Default: 0

Type: integer

Modifiable: Yes

Changes Take Effect: Immediately

The total memory, in bytes, reserved for non-data usage. By
default, the Redis OSS node will grow until it consumes the node's
`maxmemory` (see [Redis OSS node-type specific parameters](#ParameterGroups.Redis.NodeSpecific)). If
this occurs, then node performance will likely suffer due to
excessive memory paging. By reserving memory you can set aside
some of the available memory for non-Redis OSS purposes to help
reduce the amount of paging.

This parameter is specific to ElastiCache, and is not part of the
standard Redis OSS distribution.

For more information, see `reserved-memory-percent`
and [Managing reserved memory for Valkey and Redis OSS](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/redis-memory-management.html).

`set-max-intset-entries`

Default: 512

Type: integer

Modifiable: Yes

Changes Take Effect: Immediately

Determines the amount of memory used for certain kinds of sets
(strings that are integers in radix 10 in the range of 64 bit signed
integers). Such sets with fewer than the specified number of entries
are stored using a special encoding that saves space.`slave-allow-chaining`

Default: no

Type: string

Modifiable: No

Determines whether a read replica in Redis OSS can have read replicas
of its own.`slowlog-log-slower-than`

Default: 10000

Type: integer

Modifiable: Yes

Changes Take Effect: Immediately

The maximum execution time, in microseconds, for commands to be
logged by the Redis OSS Slow Log feature.`slowlog-max-len`

Default: 128

Type: integer

Modifiable: Yes

Changes Take Effect: Immediately

The maximum length of the Redis OSS Slow Log.`tcp-keepalive`

Default: 0

Type: integer

Modifiable: Yes

Changes Take Effect: Immediately

If this is set to a nonzero value (N), node clients are polled
every N seconds to ensure that they are still connected. With the
default setting of 0, no such polling occurs.

###### Important

Some aspects of this parameter changed in Redis OSS version
3.2.4. See [Parameters changed in Redis OSS 3.2.4 (enhanced)](#ParameterGroups.Redis.3-2-4.Changed).

`timeout`

Default: 0

Type: integer

Modifiable: Yes

Changes Take Effect: Immediately

The number of seconds a node waits before timing out. Values are:

- `0` – never disconnect an idle
client.

- `1-19` – invalid values.

- `>=20` – the number of seconds a node
waits before disconnecting an idle client.

`zset-max-ziplist-entries`

Default: 128

Type: integer

Modifiable: Yes

Changes Take Effect: Immediately

Determines the amount of memory used for sorted sets. Sorted sets
with fewer than the specified number of elements are stored using a
special encoding that saves space.`zset-max-ziplist-value`

Default: 64

Type: integer

Modifiable: Yes

Changes Take Effect: Immediately

Determines the amount of memory used for sorted sets. Sorted sets
with entries that are smaller than the specified number of bytes are
stored using a special encoding that saves space.

###### Note

If you do not specify a parameter group for your Redis OSS 2.6.13 cluster, then a
default parameter group ( `default.redis2.6`) will be used. You cannot
change the values of any parameters in the default parameter group; however, you
can always create a custom parameter group and assign it to your cluster at any
time.

### Redis OSS node-type specific parameters

Although most parameters have a single value, some parameters have different
values depending on the node type used. The following table shows the default values
for the `maxmemory`,
`client-output-buffer-limit-slave-hard-limit`, and
`client-output-buffer-limit-slave-soft-limit` parameters for each
node type. The value of `maxmemory` is the maximum number of bytes
available to you for use, data and other uses, on the node. For more information,
see [Available memory](https://aws.amazon.com/premiumsupport/knowledge-center/available-memory-elasticache-redis-node).

###### Note

The `maxmemory` parameter cannot be modified.

Node type Maxmemory Client-output-buffer-limit-slave-hard-limitClient-output-buffer-limit-slave-soft-limitcache.t1.micro1426063361426063314260633cache.t2.micro5819596805819596858195968cache.t2.small1665138688166513868166513868cache.t2.medium3461349376346134937346134937cache.t3.micro5368709125368709153687091cache.t3.small1471026299147102629147102629cache.t3.medium3317862236331786223331786223cache.t4g.micro5368709125368709153687091cache.t4g.small1471026299147102629147102629cache.t4g.medium3317862236331786223331786223cache.m1.small9437184009437184094371840cache.m1.medium3093299200309329920309329920cache.m1.large7025459200702545920702545920cache.m1.xlarge1488977920014889779201488977920cache.m2.xlarge1709178880017091788801709178880cache.m2.2xlarge3502243840035022438403502243840cache.m2.4xlarge7088373760070883737607088373760cache.m3.medium2988441600309329920309329920cache.m3.large6501171200650117120650117120cache.m3.xlarge1426063360014260633601426063360cache.m3.2xlarge2998927360029989273602998927360cache.m4.large6892593152689259315689259315cache.m4.xlarge1532850176015328501761532850176cache.m4.2xlarge3188912635931889126363188912636cache.m4.4xlarge6525729062965257290636525729063cache.m4.10xlarge1660476142391660476142416604761424cache.m5.large6854542746685454275 685454275cache.m5.xlarge1389192171513891921721389192172cache.m5.2xlarge2796666921027966669212796666921cache.m5.4xlarge5611617812556116178125611617812cache.m5.12xlarge1687159719941687159719916871597199cache.m5.24xlarge3375005628423375005628433750056284cache.m6g.large6854542746685454275685454275cache.m6g.xlarge1389192171513891921721389192172cache.m6g.2xlarge2796666921027966669212796666921cache.m6g.4xlarge5611617812556116178125611617812cache.m6g.8xlarge1113255523121113255523111132555231cache.m6g.12xlarge1687159719941687159719916871597199cache.m6g.16xlarge2250003752282250003752322500037523cache.c1.xlarge6501171200650117120650117120cache.r3.large1447034880014680064001468006400cache.r3.xlarge3051356160030408704003040870400cache.r3.2xlarge6249512960060817408006081740800cache.r3.4xlarge1264582656001226833920012268339200cache.r3.8xlarge2543845376002453667840024536678400cache.r4.large1320178155613201781551320178155cache.r4.xlarge2689822883926898228832689822883cache.r4.2xlarge5419753799754197537995419753799cache.r4.4xlarge1088585465861088585465810885854658cache.r4.8xlarge2182554320902182554320921825543209cache.r4.16xlarge4370215731204370215731243702157312cache.r5.large1403718103014037181031403718103cache.r5.xlarge2826184970228261849702826184970cache.r5.2xlarge5671118356556711183565671118356cache.r5.4xlarge1136098652161136098652211360986522cache.r5.12xlarge3412063465473412063465534120634655cache.r5.24xlarge6824859738116824859738168248597381cache.r6g.large1403718103014037181031403718103cache.r6g.xlarge2826184970228261849702826184970cache.r6g.2xlarge5671118356556711183565671118356cache.r6g.4xlarge1136098652161136098652211360986522cache.r6g.8xlarge2250003752282250003752322500037523cache.r6g.12xlarge3412063465473412063465534120634655cache.r6g.16xlarge4500007504564500007504645000075046cache.r6gd.xlarge2826184970228261849702826184970cache.r6gd.2xlarge5671118356556711183565671118356cache.r6gd.4xlarge1136098652161136098652211360986522cache.r6gd.8xlarge2250003752282250003752322500037523cache.r6gd.12xlarge3412063465473412063465534120634655cache.r6gd.16xlarge4500007504564500007504645000075046cache.r7g.large1403718103014037181031403718103cache.r7g.xlarge2826184970228261849702826184970cache.r7g.2xlarge5671118356556711183565671118356cache.r7g.4xlarge1136098652161136098652211360986522cache.r7g.8xlarge2250003752282250003752322500037523cache.r7g.12xlarge3412063465473412063465534120634655cache.r7g.16xlarge4500007504564500007504645000075046cache.m7g.large6854542746685454275685454275cache.m7g.xlarge1389192171513891921721389192172cache.m7g.2xlarge2796666921027966669212796666921cache.m7g.4xlarge5611617812556116178125611617812cache.m7g.8xlarge1113255523121113255523111132555231cache.m7g.12xlarge1687159719941687159719916871597199cache.m7g.16xlarge2250003752282250003752322500037523cache.c7gn.large331786223614037181031403718103cache.c7gn.xlarge685454274628261849702826184970cache.c7gn.2xlarge1389192171556711183565671118356cache.c7gn.4xlarge279666692101136098652211360986522cache.c7gn.8xlarge561161781252250003752322500037523cache.c7gn.12xlarge843579859973412063465534120634655cache.c7gn.16xlarge1136098652164500007504645000075046

###### Note

All current generation instance types are created in an Amazon Virtual Private Cloud VPC by
default.

T1 instances do not support Multi-AZ.

T1 and T2 instances do not support Redis OSS AOF.

Redis OSS configuration variables `appendonly` and
`appendfsync` are not supported on Redis OSS version 2.8.22 and
later.

## Memcached specific parameters

**Memcached**

If you do not specify a parameter group for your Memcached cluster, then a default
parameter group appropriate to your engine version will be used. You can't change the
values of any parameters in a default parameter group. However, you can create a custom
parameter group and assign it to your cluster at any time. For more information, see
[Creating an ElastiCache parameter group](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/ParameterGroups.Creating.html).

###### Topics

- [Memcached 1.6.17 changes](#ParameterGroups.Memcached.1.6.17)

- [Memcached 1.6.6 added parameters](#ParameterGroups.Memcached.1-6-6)

- [Memcached 1.5.10 parameter changes](#ParameterGroups.Memcached.1-5-10)

- [Memcached 1.4.34 added parameters](#ParameterGroups.Memcached.1-4-34)

- [Memcached 1.4.33 added parameters](#ParameterGroups.Memcached.1-4-33)

- [Memcached 1.4.24 added parameters](#ParameterGroups.Memcached.1-4-24)

- [Memcached 1.4.14 added parameters](#ParameterGroups.Memcached.1-4-14)

- [Memcached 1.4.5 supported parameters](#ParameterGroups.Memcached.1-4-5)

- [Memcached connection overhead](#ParameterGroups.Memcached.Overhead)

- [Memcached node-type specific parameters](#ParameterGroups.Memcached.NodeSpecific)

### Memcached 1.6.17 changes

From Memcached 1.6.17, we no longer support these administrative commands:
`lru_crawler`, `lru`, and `slabs`. With these
changes, you will not be able to enable/disable `lru_crawler` at runtime
via commands. Please enable/disable `lru_crawler` by modifying your
custom parameter group.

### Memcached 1.6.6 added parameters

For Memcached 1.6.6, no additional parameters are supported.

**Parameter group family:** memcached1.6

### Memcached 1.5.10 parameter changes

For Memcached 1.5.10, the following additional parameters are supported.

**Parameter group family:** memcached1.5

NameDetailsDescription`no_modern `

Default: 1

Type: boolean

Modifiable: Yes

Allowed\_Values: 0,1

Changes Take Effect: At launch

An alias for disabling `slab_reassign`,
`lru_maintainer_thread`, `lru_segmented`, and `maxconns_fast`
commands.

When using Memcached 1.5 and higher, `no_modern` also sets the hash\_algorithm to
`jenkins`.

In addition, when using Memcached 1.5.10, `inline_ascii_reponse` is controlled by the parameter `parallelly`. This means that if `no_modern` is disabled then `inline_ascii_reponse` is disabled. From Memcached engine 1.5.16 onward the `inline_ascii_response` parameter no longer applies, so `no_modern` being abled or disabled has no effect on `inline_ascii_reponse`.

If `no_modern` is disabled, then `slab_reassign`, `lru_maintainer_thread`, `lru_segmented`, and `maxconns_fast` WILL be enabled. Since `slab_automove` and `hash_algorithm` parameters are not SWITCH parameters, their setting is based on the configurations in the parameter group.

If you want to disable `no_modern` and revert to `modern`, you must configure a custom parameter group to disable this parameter and then reboot for these changes to take effect.

###### Note

The default configuration value for this parameter has
been changed from 0 to 1 as of August 20, 2021. The updated
default value will get automatically picked up by new
ElastiCache users for each regions after August 20th, 2021.
Existing ElastiCache users in the regions before August 20th, 2021
need to manually modify their custom parameter groups in
order to pick up this new change.

`inline_ascii_resp `

Default: 0

Type: boolean

Modifiable: Yes

Allowed\_Values: 0,1

Changes Take Effect: At launch

Stores numbers from `VALUE` response, inside an
item, using up to 24 bytes. Small slowdown for ASCII
`get`, `faster` sets.

For Memcached 1.5.10, the following parameters are removed.

NameDetailsDescription`expirezero_does_not_evict `

Default: 0

Type: boolean

Modifiable: Yes

Allowed\_Values: 0,1

Changes Take Effect: At launch

No longer supported in this version.

`modern `

Default: 1

Type: boolean

Modifiable: Yes (requires re-launch if set to
`no_modern`)

Allowed\_Values: 0,1

Changes Take Effect: At launch

No longer supported in this version. Starting with this
version, `no-modern` is enabled by default with every
launch or re-launch.

### Memcached 1.4.34 added parameters

For Memcached 1.4.34, no additional parameters are supported.

**Parameter group family:** memcached1.4

### Memcached 1.4.33 added parameters

For Memcached 1.4.33, the following additional parameters are supported.

**Parameter group family:** memcached1.4

NameDetailsDescription` modern `

Default: enabled

Type: boolean

Modifiable: Yes

Changes Take Effect: At launch

An alias to multiple features. Enabling `modern` is
equivalent to turning following commands on and using a murmur3
hash algorithm: `slab_reassign`,
`slab_automove`, `lru_crawler`,
`lru_maintainer`, `maxconns_fast`, and
`hash_algorithm=murmur3`.

` watch `

Default: enabled

Type: boolean

Modifiable: Yes

Changes Take Effect: Immediately

Logs can get dropped if user hits their
`watcher_logbuf_size` and
`worker_logbuf_size` limits.

Logs fetches, evictions or mutations. When, for example, user
turns `watch` on, they can see logs when
`get`, `set`, `delete`, or
`update` occur.

` idle_timeout `

Default: 0 (disabled)

Type: integer

Modifiable: Yes

Changes Take Effect: At Launch

The minimum number of seconds a client will be allowed to idle
before being asked to close. Range of values: 0 to
86400.

` track_sizes `

Default: disabled

Type: boolean

Modifiable: Yes

Changes Take Effect: At Launch

Shows the sizes each slab group has consumed.

Enabling `track_sizes` lets you run `stats
										sizes` without the need to run `stats
										sizes_enable`.

` watcher_logbuf_size `

Default: 256 (KB)

Type: integer

Modifiable: Yes

Changes Take Effect: At Launch

The `watch` command turns on stream logging for
Memcached. However `watch` can drop logs if the rate
of evictions, mutations or fetches are high enough to cause the
logging buffer to become full. In such situations, users can
increase the buffer size to reduce the chance of log
losses.

` worker_logbuf_size `

Default: 64 (KB)

Type: integer

Modifiable: Yes

Changes Take Effect: At Launch

The `watch` command turns on stream logging for
Memcached. However `watch` can drop logs if the rate
of evictions, mutations or fetches are high enough to cause
logging buffer get full. In such situations, users can increase
the buffer size to reduce the chance of log
losses.

` slab_chunk_max `

Default: 524288 (bytes)

Type: integer

Modifiable: Yes

Changes Take Effect: At Launch

Specifies the maximum size of a slab. Setting smaller slab
size uses memory more efficiently. Items larger than
`slab_chunk_max` are split over multiple
slabs.

` lru_crawler metadump [all|1|2|3]`

Default: disabled

Type: boolean

Modifiable: Yes

Changes Take Effect: Immediately

if lru\_crawler is enabled this command dumps all keys.

`all|1|2|3` \- all slabs, or specify a particular
slab number

### Memcached 1.4.24 added parameters

For Memcached 1.4.24, the following additional parameters are supported.

**Parameter group family:** memcached1.4

NameDetailsDescription` disable_flush_all `

Default: 0 (disabled)

Type: boolean

Modifiable: Yes

Changes Take Effect: At launch

Add parameter ( `-F`) to disable flush\_all. Useful
if you never want to be able to run a full flush on production
instances.

Values: 0, 1 (user can do a `flush_all` when the
value is 0).

` hash_algorithm `

Default: jenkins

Type: string

Modifiable: Yes

Changes Take Effect: At launch

The hash algorithm to be used. Permitted values: murmur3 and
jenkins.` lru_crawler `

Default: 0 (disabled)

Type: boolean

Modifiable: Yes

Changes Take Effect: After restart

###### Note

You can temporarily enable `lru_crawler` at
runtime from the command line. For more information, see the
Description column.

Cleans slab classes of items that have expired. This is a low
impact process that runs in the background. Currently requires
initiating a crawl using a manual command.

To temporarily enable, run `lru_crawler enable` at
the command line.

`lru_crawler 1,3,5` crawls slab classes 1, 3, and 5
looking for expired items to add to the freelist.

Values: 0,1

###### Note

Enabling `lru_crawler` at the command line
enables the crawler until either disabled at the command
line or the next reboot. To enable permanently, you must
modify the parameter value. For more information, see [Modifying an ElastiCache parameter group](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/ParameterGroups.Modifying.html).

` lru_maintainer `

Default: 0 (disabled)

Type: boolean

Modifiable: Yes

Changes Take Effect: At launch

A background thread that shuffles items between the LRUs as
capacities are reached. Values: 0, 1.

` expirezero_does_not_evict `

Default: 0 (disabled)

Type: boolean

Modifiable: Yes

Changes Take Effect: At launch

When used with `lru_maintainer`, makes items with
an expiration time of 0 unevictable.

###### Warning

This can crowd out memory available for other evictable
items.

Can be set to disregard
`lru_maintainer`.

### Memcached 1.4.14 added parameters

For Memcached 1.4.14, the following additional parameters are supported.

**Parameter group family:** memcached1.4

Parameters added in Memcached 1.4.14 Name  Details  Description `config_max`

Default: 16

Type: integer

Modifiable: No

The maximum number of ElastiCache configuration entries.`config_size_max`

Default: 65536

Type: integer

Modifiable: No

The maximum size of the configuration entries, in bytes.`hashpower_init`

Default: 16

Type: integer

Modifiable: No

The initial size of the ElastiCache hash table, expressed as a power of
two. The default is 16 (2^16), or 65536 keys.`maxconns_fast`

Default: 0 (false)

Type: Boolean

Modifiable: Yes

Changes Take Effect: After restart

Changes the way in which new connections requests are handled
when the maximum connection limit is reached. If this parameter is
set to 0 (zero), new connections are added to the backlog queue and
will wait until other connections are closed. If the parameter is
set to 1, ElastiCache sends an error to the client and immediately closes
the connection.`slab_automove`

Default: 0

Type: integer

Modifiable: Yes

Changes Take Effect: After restart

Adjusts the slab automove algorithm: If this parameter is set to
0 (zero), the automove algorithm is disabled. If it is set to 1,
ElastiCache takes a slow, conservative approach to automatically moving
slabs. If it is set to 2, ElastiCache aggressively moves slabs whenever
there is an eviction. (This mode is not recommended except for
testing purposes.)`slab_reassign`

Default: 0 (false)

Type: Boolean

Modifiable: Yes

Changes Take Effect: After restart

Enable or disable slab reassignment. If this parameter is set to
1, you can use the "slabs reassign" command to manually reassign
memory.

### Memcached 1.4.5 supported parameters

**Parameter group family:** memcached1.4

For Memcached 1.4.5, the following parameters are supported.

Parameters added in Memcached 1.4.5 Name  Details  Description `backlog_queue_limit`

Default: 1024

Type: integer

Modifiable: No

The backlog queue limit.`binding_protocol`

Default: auto

Type: string

Modifiable: Yes

Changes Take Effect: After restart

The binding protocol.

Permissible values are:
`ascii` and `auto`.

For guidance on modifying the value of
`binding_protocol`, see [Modifying an ElastiCache parameter group](https://docs.aws.amazon.com/AmazonElastiCache/latest/dg/ParameterGroups.Modifying.html).

`cas_disabled`

Default: 0 (false)

Type: Boolean

Modifiable: Yes

Changes Take Effect: After restart

If `1` (true), check and set (CAS) operations will be
disabled, and items stored will consume 8 fewer bytes than with CAS
enabled.`chunk_size`

Default: 48

Type: integer

Modifiable: Yes

Changes Take Effect: After restart

The minimum amount, in bytes, of space to allocate for the
smallest item's key, value, and flags.`chunk_size_growth_factor`

Default: 1.25

Type: float

Modifiable: Yes

Changes Take Effect: After restart

The growth factor that controls the size of each successive
Memcached chunk; each chunk will be
`chunk_size_growth_factor` times larger than the
previous chunk.`error_on_memory_exhausted`

Default: 0 (false)

Type: Boolean

Modifiable: Yes

Changes Take Effect: After restart

If `1` (true), when there is no more memory to store
items, Memcached will return an error rather than evicting
items.`large_memory_pages`

Default: 0 (false)

Type: Boolean

Modifiable: No

If `1` (true), ElastiCache will try to use large memory
pages.`lock_down_paged_memory`

Default: 0 (false)

Type: Boolean

Modifiable: No

If `1` (true), ElastiCache will lock down all paged
memory.`max_item_size`

Default: 1048576

Type: integer

Modifiable: Yes

Changes Take Effect: After restart

The size, in bytes, of the largest item that can be stored in the
cluster.`max_simultaneous_connections`

Default: 65000

Type: integer

Modifiable: No

The maximum number of simultaneous connections.`maximize_core_file_limit`

Default: 0 (false)

Type: Boolean

Modifiable:

Changes Take Effect: After restart

If `1` (true), ElastiCache will maximize the core file
limit.`memcached_connections_overhead`

Default: 100

Type: integer

Modifiable: Yes

Changes Take Effect: After restart

The amount of memory to be reserved for Memcached connections and
other miscellaneous overhead. For information about this parameter,
see [Memcached connection overhead](#ParameterGroups.Memcached.Overhead).`requests_per_event`

Default: 20

Type: integer

Modifiable: No

The maximum number of requests per event for a given connection.
This limit is required to prevent resource starvation.

### Memcached connection overhead

On each node, the memory made available for storing items is the total available
memory on that node (which is stored in the `max_cache_memory` parameter)
minus the memory used for connections and other overhead (which is stored in the
`memcached_connections_overhead` parameter). For example, a node of
type `cache.m1.small` has a `max_cache_memory` of 1300MB. With
the default `memcached_connections_overhead` value of 100MB, the
Memcached process will have 1200MB available to store items.

The default values for the `memcached_connections_overhead` parameter
satisfy most use cases; however, the required amount of allocation for connection
overhead can vary depending on multiple factors, including request rate, payload
size, and the number of connections.

You can change the value of the `memcached_connections_overhead` to
better suit the needs of your application. For example, increasing the value of the
`memcached_connections_overhead` parameter will reduce the amount of
memory available for storing items and provide a larger buffer for connection
overhead. Decreasing the value of the `memcached_connections_overhead`
parameter will give you more memory to store items, but can increase your risk of
swap usage and degraded performance. If you observe swap usage and degraded
performance, try increasing the value of the
`memcached_connections_overhead` parameter.

###### Important

For the `cache.t1.micro` node type, the value for
`memcached_connections_overhead` is determined as follows:

- If you cluster is using the default parameter group, ElastiCache will set
the value for `memcached_connections_overhead` to
13MB.

- If your cluster is using a parameter group that you have created
yourself, you can set the value of
`memcached_connections_overhead` to a value of your
choice.

### Memcached node-type specific parameters

Although most parameters have a single value, some parameters have different
values depending on the node type used. The following table shows the default values
for the `max_cache_memory` and `num_threads` parameters for
each node type. The values on these parameters cannot be modified.

Node type max\_cache\_memory (in megabytes) num\_threads cache.t1.micro213 1cache.t2.micro5551cache.t2.small15881cache.t2.medium33012cache.t3.micro5122cache.t3.small14022cache.t3.medium33642cache.t4g.micro5122cache.t4g.small14022cache.t4g.medium31642cache.m1.small13011cache.m1.medium33501cache.m1.large71002cache.m1.xlarge14600 4cache.m2.xlarge338002cache.m2.2xlarge304124cache.m2.4xlarge68000 16cache.m3.medium28501cache.m3.large62002cache.m3.xlarge136004cache.m3.2xlarge286008cache.m4.large65732cache.m4.xlarge11496 4cache.m4.2xlarge304128cache.m4.4xlarge6223416cache.m4.10xlarge15835540cache.m5.large65372cache.m5.xlarge132484cache.m5.2xlarge266718cache.m5.4xlarge5351616cache.m5.12xlarge16090048cache.m5.24xlarge321865 96cache.m6g.large65372cache.m6g.xlarge132484cache.m6g.2xlarge266718cache.m6g.4xlarge5351616cache.m6g.8xlarge10700032cache.m6g.12xlarge16090048cache.m6g.16xlarge21457764cache.c1.xlarge66008cache.r3.large138002cache.r3.xlarge291004cache.r3.2xlarge596008cache.r3.4xlarge12060016cache.r3.8xlarge12060032cache.r4.large125902cache.r4.xlarge256524cache.r4.2xlarge516868cache.r4.4xlarge10381516cache.r4.8xlarge20814432cache.r4.16xlarge41677664cache.r5.large133872cache.r5.xlarge269534cache.r5.2xlarge540848cache.r5.4xlarge10834716cache.r5.12xlarge32540048cache.r5.24xlarge65086996cache.r6g.large133872cache.r6g.xlarge269534cache.r6g.2xlarge540848cache.r6g.4xlarge10834716cache.r6g.8xlarge21457732cache.r6g.12xlarge32540048cache.r6g.16xlarge42915464cache.c7gn.large31642cache.c7gn.xlarge65374cache.c7gn.2xlarge132488cache.c7gn.4xlarge2667116cache.c7gn.8xlarge5351632cache.c7gn.12xlarge32540048cache.c7gn.16xlarge10834764

###### Note

All T2 instances are created in an Amazon Virtual Private Cloud (Amazon VPC).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Deleting an ElastiCache parameter group

Connecting an EC2 instance and an ElastiCache cache automatically
