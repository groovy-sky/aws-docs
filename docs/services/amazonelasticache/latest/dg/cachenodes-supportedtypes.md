# Supported node types

ElastiCache supports the following node types.
Generally speaking, the current generation types provide more memory and computational power
at lower cost when compared to their equivalent previous generation counterparts.

For more information on performance details for each node type, see [Amazon EC2 Instance Types](https://aws.amazon.com/ec2/instance-types).

###### Note

Amazon ElastiCache is transitioning T2 instances to previous generation status.
You will no longer be able to create new ElastiCache clusters using T2 instances or purchase new T2 reserved nodes.
There is no impact to existing T2 clusters or reservations. We recommend upgrading to newer instance types such as T3 or T4g instances for better performance and cost efficiency.

###### Note

The following instance types are supported in the AWS Asia Pacific (Thailand) and Mexico (Central) Regions:

- **m7g/r7g:** large, xl, 2xl, 4xl, 8xl, 12xl, and 16xl.

- **t3/t4g:** micro, small, and medium.

For information on which node size to use, see [Choosing your node size](cachenodes-selectsize.md).

###### Topics

- [Current Generation (Memcached)](#CacheNodes.CurrentGen-Memcached)

- [Current Generation (Valkey or Redis OSS)](#CacheNodes.CurrentGen)

- [Supported node types by AWS Region](#CacheNodes.SupportedTypesByRegion)

- [Burstable Performance Instances](#CacheNodes.Burstable)

- [Related Information](#CacheNodes.RelatedInfo)

## Current Generation (Memcached)

The following tables show the baseline and burst bandwidth for instance types that
use the network I/O credit mechanism to burst beyond their baseline
bandwidth.

###### Note

Instance types with burstable network performance use a network I/O credit
mechanism to burst beyond their baseline bandwidth on a best-effort
basis.

**General**

Instance typeMinimum supported Memcached versionBaseline bandwidth (Gbps)Burst bandwidth (Gbps)cache.m7g.large0.93712.5cache.m7g.xlarge1.87612.5cache.m7g.2xlarge3.7515cache.m7g.4xlarge7.515cache.m7g.8xlarge15N/Acache.m7g.12xlarge22.5N/Acache.m7g.16xlarge30N/Acache.m6g.large1.5.160.7510.0cache.m6g.xlarge1.5.161.2510.0cache.m6g.2xlarge1.5.162.510.0cache.m6g.4xlarge1.5.165.010.0cache.m6g.8xlarge1.5.1612N/Acache.m6g.12xlarge1.5.1620N/Acache.m6g.16xlarge1.5.1625N/Acache.m5.large1.5.160.7510.0cache.m5.xlarge1.5.161.2510.0cache.m5.2xlarge1.5.162.510.0cache.m5.4xlarge1.5.165.010.0cache.m5.12xlarge1.5.16N/AN/Acache.m5.24xlarge1.5.16N/AN/Acache.m4.large1.5.160.451.2cache.m4.xlarge1.5.160.752.8cache.m4.2xlarge1.5.161.010.0cache.m4.4xlarge1.5.162.010.0cache.m4.10xlarge1.5.165.010.0cache.t4g.micro1.5.160.0645.0cache.t4g.small1.5.160.1285.0cache.t4g.medium1.5.160.2565.0cache.t3.micro1.5.160.0645.0cache.t3.small1.5.160.1285.0cache.t3.medium1.5.160.2565.0cache.t2.micro1.5.160.0641.024cache.t2.small1.5.160.1281.024cache.t2.medium1.5.160.2561.024

**Memory optimized for Memcached**

Instance typeMinimum supported versionBaseline bandwidth (Gbps)Burst bandwidth (Gbps)cache.r7g.large0.93712.5cache.r7g.xlarge1.87612.5cache.r7g.2xlarge3.7515cache.r7g.4xlarge7.515cache.r7g.8xlarge15N/Acache.r7g.12xlarge22.5N/Acache.r7g.16xlarge30N/Acache.r6g.large1.5.160.7510.0cache.r6g.xlarge1.5.161.2510.0cache.r6g.2xlarge1.5.162.510.0cache.r6g.4xlarge1.5.165.010.0cache.r6g.8xlarge1.5.1612N/Acache.r6g.12xlarge1.5.1620N/Acache.r6g.16xlarge1.5.1625N/Acache.r5.large1.5.160.7510.0cache.r5.xlarge1.5.161.2510.0cache.r5.2xlarge1.5.162.510.0cache.r5.4xlarge1.5.165.010.0cache.r5.12xlarge1.5.1620N/Acache.r5.24xlarge1.5.1625N/Acache.r4.large1.5.160.7510.0cache.r4.xlarge1.5.161.2510.0cache.r4.2xlarge1.5.162.510.0cache.r4.4xlarge1.5.165.010.0cache.r4.8xlarge1.5.1612N/Acache.r4.16xlarge1.5.1625N/A

**Network optimized for Memcached**

Instance typeMinimum supported versionBaseline bandwidth (Gbps)Burst bandwidth (Gbps)cache.c7gn.large1.6.66.2530cache.c7gn.xlarge1.6.612.540cache.c7gn.2xlarge1.6.62550cache.c7gn.4xlarge1.6.650N/Acache.c7gn.8xlarge1.6.6100N/Acache.c7gn.12xlarge1.6.6150N/Acache.c7gn.16xlarge1.6.6200N/A

## Current Generation (Valkey or Redis OSS)

For more information on Previous Generation, please refer to [Previous\
Generation Nodes](https://aws.amazon.com/elasticache/previous-generation).

###### Note

Instance types with burstable network performance use a network I/O credit
mechanism to burst beyond their baseline bandwidth on a best-effort
basis.

**General**

Instance typeMinimum supported Redis OSS versionEnhanced I/O with Redis OSS 5.0.6+TLS Offloading with Redis OSS 6.2.5+Enhanced I/O Multiplexing with Redis OSS 7.0.4+Baseline bandwidth (Gbps)Burst bandwidth (Gbps)cache.m7g.large6.2NNN0.93712.5cache.m7g.xlarge6.2YYY1.87612.5cache.m7g.2xlarge6.2YYY3.7515cache.m7g.4xlarge6.2YYY7.515cache.m7g.8xlarge6.2YYY15N/Acache.m7g.12xlarge6.2YYY22.5N/Acache.m7g.16xlarge6.2YYY30N/Acache.m6g.large5.0.6NNN0.7510.0cache.m6g.xlarge5.0.6YYY1.2510.0cache.m6g.2xlarge5.0.6YYY2.510.0cache.m6g.4xlarge5.0.6YYY5.010.0cache.m6g.8xlarge5.0.6YYY12N/Acache.m6g.12xlarge5.0.6YYY20N/Acache.m6g.16xlarge5.0.6YYY25N/Acache.m5.large3.2.4NNN0.7510.0cache.m5.xlarge3.2.4YNN1.2510.0cache.m5.2xlarge3.2.4YYY2.510.0cache.m5.4xlarge3.2.4YYY5.010.0cache.m5.12xlarge3.2.4YYY12N/Acache.m5.24xlarge3.2.4YYY25N/Acache.m4.large3.2.4NNN0.451.2cache.m4.xlarge3.2.4YNN0.752.8cache.m4.2xlarge3.2.4YYY1.010.0cache.m4.4xlarge3.2.4YYY2.010.0cache.m4.10xlarge3.2.4YYY5.010.0cache.t4g.micro3.2.4NNN0.0645.0cache.t4g.small5.0.6NNN0.1285.0cache.t4g.medium5.0.6NNN0.2565.0cache.t3.micro3.2.4NNN0.0645.0cache.t3.small3.2.4NNN0.1285.0cache.t3.medium3.2.4NNN0.2565.0cache.t2.micro3.2.4NNN0.0641.024cache.t2.small3.2.4NNN0.1281.024cache.t2.medium3.2.4NNN0.2561.024

**Memory optimized**

Instance typeMinimum supported Redis OSS versionEnhanced I/O with Redis OSS 5.0.6+TLS Offloading with Redis OSS 6.2.5+Enhanced I/O Multiplexing with Redis OSS 7.0.4+Baseline bandwidth (Gbps)Burst bandwidth (Gbps)cache.r7g.large6.2NNN0.93712.5cache.r7g.xlarge6.2YYY1.87612.5cache.r7g.2xlarge6.2YYY3.7515cache.r7g.4xlarge6.2YYY7.515cache.r7g.8xlarge6.2YYY15N/Acache.r7g.12xlarge6.2YYY22.5N/Acache.r7g.16xlarge6.2YYY30N/Acache.r6g.large5.0.6NNN0.7510.0cache.r6g.xlarge5.0.6YYY1.2510.0cache.r6g.2xlarge5.0.6YYY2.510.0cache.r6g.4xlarge5.0.6YYY5.010.0cache.r6g.8xlarge5.0.6YYY12N/Acache.r6g.12xlarge5.0.6YYY20N/Acache.r6g.16xlarge5.0.6YYY25N/Acache.r5.large3.2.4NNN0.7510.0cache.r5.xlarge3.2.4YNN1.2510.0cache.r5.2xlarge3.2.4YYY2.510.0cache.r5.4xlarge3.2.4YYY5.010.0cache.r5.12xlarge3.2.4YYY12N/Acache.r5.24xlarge3.2.4YYY25N/Acache.r4.large3.2.4NNN0.7510.0cache.r4.xlarge3.2.4YNN1.2510.0cache.r4.2xlarge3.2.4YYY2.510.0cache.r4.4xlarge3.2.4YYY5.010.0cache.r4.8xlarge3.2.4YYY12N/Acache.r4.16xlarge3.2.4YYY25N/A

**Memory optimized with data tiering**

Instance typeMinimum supported Redis OSS versionEnhanced I/O with Redis OSS 5.0.6+TLS Offloading with Redis OSS 6.2.5+Enhanced I/O Multiplexing with Redis OSS 7.0.4+Baseline bandwidth (Gbps)Burst bandwidth (Gbps)cache.r6gd.xlarge6.2.0YNN1.2510cache.r6gd.2xlarge6.2.0YYY2.510cache.r6gd.4xlarge6.2.0YYY5.010cache.r6gd.8xlarge6.2.0YYY12N/Acache.r6gd.12xlarge6.2.0YYY20N/Acache.r6gd.16xlarge6.2.0YYY25N/A

**Network optimized**

Instance typeMinimum supported Redis OSS versionEnhanced I/O with Redis OSS 5.0.6+TLS Offloading with Redis OSS 6.2.5+Enhanced I/O Multiplexing with Redis OSS 7.0.4+Baseline bandwidth (Gbps)Burst bandwidth (Gbps)cache.c7gn.large6.2NNN6.2530cache.c7gn.xlarge6.2YYY12.540cache.c7gn.2xlarge6.2YYY2550cache.c7gn.4xlarge6.2YYY50N/Acache.c7gn.8xlarge6.2YYY100N/Acache.c7gn.12xlarge6.2YYY150N/Acache.c7gn.16xlarge6.2YYY200N/A

## Supported node types by AWS Region

Supported node types may vary between AWS Regions. For more details, see [Amazon ElastiCache\
pricing](https://aws.amazon.com/elasticache/pricing).

## Burstable Performance Instances

You can launch general-purpose burstable T4g, T3-Standard and T2-Standard cache
nodes in Amazon ElastiCache. These nodes provide a baseline level of CPU performance with the
ability to burst CPU usage at any time until the accrued credits are exhausted. A
_CPU credit_ provides the performance of a full
CPU core for one minute.

Amazon ElastiCache's T4g, T3 and T2 nodes are configured as standard and suited for
workloads with an average CPU utilization that is consistently below the baseline
performance of the instance. To burst above the baseline, the node spends credits
that it has accrued in its CPU credit balance. If the node is running low on accrued
credits, performance is gradually lowered to the baseline performance level. This
gradual lowering ensures the node doesn't experience a sharp performance drop-off
when its accrued CPU credit balance is depleted. For more information, see [CPU Credits and Baseline Performance for Burstable Performance\
Instances](../../../ec2/latest/userguide/burstable-credits-baseline-concepts.md) in the _Amazon EC2 User_
_Guide._

The following table lists the burstable performance node types, the rate at which
CPU credits are earned per hour. It also shows the maximum number of earned CPU
credits that a node can accrue and the number of vCPUs per node. In addition, it
gives the baseline performance level as a percentage of a full core performance
(using a single vCPU).

Node typeCPU credits earned per hour Maximum earned credits that can be
accrued\* vCPUs  Baseline performance per vCPU  Memory (GiB)  Network performance t4g.micro`12`288210%0.5Up to 5 Gigabitt4g.small`24`576220%1.37Up to 5 Gigabitt4g.medium`24`576220%3.09Up to 5 Gigabitt3.micro`12`288210%0.5Up to 5 Gigabitt3.small`24`576220%1.37Up to 5 Gigabitt3.medium`24`576220%3.09Up to 5 Gigabitt2.micro`6`144110%0.5Low to moderatet2.small`12`288120%1.55Low to moderatet2.medium`24`576220%3.22Low to moderate

\\* The number of credits that can be accrued is equivalent to the number of credits
that can be earned in a 24-hour period.

\\*\\* The baseline performance in the table is per vCPU. Some node sizes that have
more than one vCPU. For these, calculate the baseline CPU utilization for the node
by multiplying the vCPU percentage by the number of vCPUs.

The following CPU credit metrics are available for T3 and T4g burstable
performance instances:

###### Note

These metrics are not available for T2 burstable performance instances.

- `CPUCreditUsage`

- `CPUCreditBalance`

For more information on these metrics, see [CPU Credit Metrics](../../../ec2/latest/userguide/viewing-metrics-with-cloudwatch.md#cpu-credit-metrics).

In addition, be aware of these details:

- All current generation node types are created in a virtual private cloud
(VPC) based on Amazon VPC by default.

- Redis OSS append-only files (AOF) aren't supported for T2 instances.
Redis OSS configuration variables `appendonly` and
`appendfsync` aren't supported on Redis OSS version 2.8.22 and
later.

## Related Information

- [Amazon ElastiCache Product\
Features and Details](https://aws.amazon.com/elasticache/details)

- [Memcached Node-Type Specific Parameters for Memcached](parametergroups-engine.md#ParameterGroups.Memcached)

- [Valkey and Redis OSS parameters](parametergroups-engine.md#ParameterGroups.Redis)

- [In\
Transit Encryption (TLS)](in-transit-encryption.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Connecting to nodes

Rebooting nodes

All content copied from https://docs.aws.amazon.com/.
