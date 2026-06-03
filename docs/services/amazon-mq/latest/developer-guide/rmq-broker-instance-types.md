---
title: "Amazon MQ for RabbitMQ broker instance types"
---

# Amazon MQ for RabbitMQ broker instance types

The combined description of the broker instance class (m7g) and size (large, medium) is called the broker instance type (for example, mq.m7g.large).

We recommend using mq.m7g instance types for both cluster and single-instance deployments.

Amazon MQ provides at least a 90 day notice before an instance type reaches end of support. We recommend upgrading your broker to a new instance type before the end-of-support date to prevent any disruptions.

###### Important

You cannot downgrade a broker from an `mq.m7g` or `mq.m5` instance type to a `mq.t3.micro` instance type.

The `mq.t3.micro` instance type does not support cluster deployment.

## Instance types for m7g cluster deployment

We recommending using `mq.m7g.x` instance types with cluster deployment.
The following table shows the available `mq.m7g.x` instance types for cluster deployment.

Instance TypevCPUMemory (GiB)Network Baseline / Burst bandwidth (Gbps) Recommended useStorageDisk volume size per node(GB)mq.m7g.medium140.52 / 12.5

Evaluation

EBS5mq.m7g.large280.937 / 12.5

Production

EBS15mq.m7g.xlarge4161.876 / 12.5

Production

EBS25mq.m7g.2xlarge8323.75 / 15.0

Production

EBS45mq.m7g.4xlarge16647.5 / 15.0

Production

EBS90mq.m7g.8xlarge3212815 Gigabit

Production

EBS175mq.m7g.12xlarge4819222.5 Gigabit

Production

EBS260mq.m7g.16xlarge6425630 Gigabit

Production

EBS345

## Instance types for m7g single instance deployment

The following table shows the available `mq.m7g.x` instance types
for single instance deployment.

Instance TypevCPUMemory (GiB)Network Baseline / Burst bandwidth (Gbps) Recommended useStorageDisk volume size per node(GB)mq.m7g.medium140.52 / 12.5

Evaluation

EBS200mq.m7g.large280.937 / 12.5

Production

EBS200mq.m7g.xlarge4161.876 / 12.5

Production

EBS200mq.m7g.2xlarge8323.75 / 15.0

Production

EBS200mq.m7g.4xlarge16647.5 / 15.0

Production

EBS200mq.m7g.8xlarge3212815 Gigabit

Production

EBS200mq.m7g.12xlarge4819222.5 Gigabit

Production

EBS200mq.m7g.16xlarge6425639 Gigabit

Production

EBS200

## Instance types for `mq.m5` single instance deployment

The following tables show the available `mq.m5.x` instance types for single instance deployment

Instance TypevCPUMemory (GiB)Network Baseline / Burst bandwidth (Gbps) Recommended useStorageDisk volume size per node(GB)mq.t3.micro210.064 / 5.0EvaluationEBS20mq.m5.large280.75 / 10.0ProductionEBS200mq.m5.xlarge4161.25 / 10.0ProductionEBS200mq.m5.2xlarge8322.5 / 10.0ProductionEBS200mq.m5.4xlarge16645.0 / 10.0ProductionEBS200

## Instance types for `mq.m5` cluster deployment

The following tables show the available `mq.m5.x` instance types for cluster deployment

Instance TypevCPUMemory (GiB)Network Baseline / Burst bandwidth (Gbps) Recommended useStorageDisk volume size per node(GB)mq.m5.large280.75 / 10.0ProductionEBS200mq.m5.xlarge4161.25 / 10.0ProductionEBS200mq.m5.2xlarge8322.5 / 10.0ProductionEBS200mq.m5.4xlarge16645.0 / 10.0ProductionEBS200

## Memory and disk alarms

Amazon MQ configures memory and disk thresholds on each RabbitMQ broker to protect
against resource exhaustion. When a threshold is exceeded, RabbitMQ triggers an
[alarm](https://www.rabbitmq.com/docs/alarms)
and blocks publishers from sending messages. Consumers on separate connections
continue to operate normally. However, if a publisher and consumer share the same
connection, the consumer is also blocked.

###### Important

Amazon MQ manages these thresholds and you cannot modify them. When the alarm
condition clears, publishers are unblocked automatically. For troubleshooting
information, see
[Amazon MQ for RabbitMQ: High memory alarm](troubleshooting-action-required-codes-rabbitmq-memory-alarm.md) and
[RabbitMQ on Amazon MQ: Disk limit alarm](troubleshooting-action-required-codes-disk-limit-alarm.md).

### Memory alarm

The `vm_memory_high_watermark` parameter defines the maximum amount
of memory that a RabbitMQ broker can use before it blocks publishers from sending
messages. When memory usage exceeds this threshold, RabbitMQ triggers a memory
alarm. For more information, see
[Memory Alarms](https://www.rabbitmq.com/docs/memory) on the
RabbitMQ website.

For `mq.m7g` instance types, Amazon MQ sets the following absolute memory
high watermark values:

Instance TypeMemory High Watermark (GiB)mq.m7g.medium1.8mq.m7g.large4.3mq.m7g.xlarge9.3mq.m7g.2xlarge19.3mq.m7g.4xlarge39.4mq.m7g.8xlarge79.7mq.m7g.12xlarge119.8mq.m7g.16xlarge160.1

For `mq.m5` instance types, Amazon MQ sets a relative memory high
watermark of 0.4 (40% of the available memory).

The higher memory thresholds on `mq.m7g` instances allow RabbitMQ to use
more available memory before triggering an alarm. For more information about
performance improvements with `mq.m7g` instances, see
[Improve\
RabbitMQ performance on Amazon MQ with AWS Graviton3-based M7g instances](https://aws.amazon.com/blogs/big-data/improve-rabbitmq-performance-on-amazon-mq-with-aws-graviton3-based-m7g-instances)
on the AWS Blog.

### Disk alarm

The `disk_free_limit` parameter defines the minimum amount of free
disk space that a RabbitMQ node requires. When free disk space on any node drops
below this limit, RabbitMQ triggers a disk alarm and blocks publishers from
sending messages. For more information, see
[Disk\
Alarms](https://www.rabbitmq.com/docs/disk-alarms) on the RabbitMQ website.

For `mq.m7g` instance types, Amazon MQ sets the following disk free
limits. Single-instance brokers have a higher disk free limit to provide
additional protection because they do not have other nodes to serve traffic
if disk space is exhausted.

Deployment ModeDisk Free Limit (GiB)Single-instance10Cluster2

For `mq.m5` instance types, Amazon MQ sets the following disk free
limits. These values apply to both single-instance and cluster deployments.

Instance TypeDisk Free Limit (GiB)mq.m5.large12mq.m5.xlarge20mq.m5.2xlarge36mq.m5.4xlarge69

Because `mq.m7g` instances have a lower disk free limit,
more of the provisioned disk volume is available for message storage compared
to equivalent `mq.m5` instances.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Deploying a RabbitMQ broker

Sizing guidelines

All content copied from https://docs.aws.amazon.com/.
