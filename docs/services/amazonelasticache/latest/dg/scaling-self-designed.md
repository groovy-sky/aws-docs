---
title: "Scaling node-based clusters"
---

# Scaling node-based clusters

The amount of data your application needs to process is seldom static.
It increases and decreases as your business grows or experiences normal fluctuations in demand.
If you self-manage your cache, you need to provision sufficient hardware for your demand peaks,
which can be expensive. By using Amazon ElastiCache you can scale to meet current demand, paying only for
what you use. ElastiCache enables you to scale your cache to match demand.

###### Note

If a Valkey or Redis OSS cluster is replicated across one or more Regions, then those Regions are scaled in order. When scaling up, secondary Regions are scaled first and then the primary Region. When scaling down, the primary Region is first and then any secondary Regions follow.

When updating the engine version, the order is secondary Region and then primary Region.

###### Topics

- [On-demand scaling for Memcached clusters](scaling-self-designed-mem-heading.md)

- [Manual scaling for Memcached clusters](scaling-memcached-manually.md)

- [Scaling for Valkey or Redis OSS (Cluster Mode Disabled) clusters](scaling-redis-classic.md)

- [Scaling replica nodes for Valkey or Redis OSS (Cluster Mode Disabled)](scaling-redisreplgrps.md)

- [Scaling Valkey or Redis OSS (Cluster Mode Enabled) clusters](scaling-redis-cluster-mode-enabled.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Scaling ElastiCache Serverless clusters

On-demand scaling for Memcached clusters

All content copied from https://docs.aws.amazon.com/.
