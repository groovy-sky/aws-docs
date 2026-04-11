# Amazon MQ for RabbitMQ broker defaults

When you create an Amazon MQ for RabbitMQ broker, Amazon MQ applies a default set of
broker policies and vhost limits to optimize your broker's performance. Amazon MQ applies vhost
limits only to the default ( `/`) vhost. Amazon MQ will not apply default policies to newly created vhosts.
We recommend keeping these defaults for all new and existing brokers. However,
you can modify, override, or delete these defaults at any time.

Amazon MQ creates different broker policies and vhost limits for Amazon MQ for RabbitMQ 3 and RabbitMQ 4. The differences will be discussed in detail in the following subsections.

Amazon MQ creates policies and limits based on the instance type and broker deployment mode
that you choose when you create your broker. The default policies are named according to the
deployment mode, as follows:

**Amazon MQ for RabbitMQ 3:**

- Single-instance – `AWS-DEFAULT-POLICY-SINGLE-INSTANCE`

- Cluster deployment – `AWS-DEFAULT-POLICY-CLUSTER-MULTI-AZ` && `AWS-DEFAULT-QUORUM-QUEUES-POLICY-CLUSTER-MULTI-AZ`

**Amazon MQ for RabbitMQ 4:**

- Single-instance – `AWS-DEFAULT-POLICY-SINGLE-INSTANCE`

- Cluster deployment – `AWS-DEFAULT-POLICY-CLUSTER` && `AWS-DEFAULT-QUORUM-QUEUES-POLICY-CLUSTER-MULTI-AZ`

For [single-instance\
brokers](rabbitmq-broker-architecture.md#rabbitmq-broker-architecture-single-instance), Amazon MQ sets the policy priority value to `0`. To override the
default priority value, you can create your own custom policies with higher priority values.
For [cluster deployments](rabbitmq-broker-architecture.md#rabbitmq-broker-architecture-cluster), Amazon MQ
sets the priority value to `1` for broker defaults. To create your own custom
policy for clusters, assign a priority value greater than `1`.

###### Note

In cluster deployments, `ha-mode` and `ha-sync-mode` broker
policies are required for classic mirroring and high availability (HA). These settings are only applicable for Amazon MQ for RabbitMQ 3 and are not configured for RabbitMQ 4.

If you delete the default `AWS-DEFAULT-POLICY-CLUSTER-MULTI-AZ` policy,
Amazon MQ uses the `ha-all-AWS-OWNED-DO-NOT-DELETE` policy with a priority value
of `0`. This ensures that the required `ha-mode` and
`ha-sync-mode` policies are still in effect. If you create your own
custom policy, Amazon MQ automatically appends `ha-mode` and
`ha-sync-mode` to your policy definitions.

###### Topics

- [Policy and limit descriptions](#rabbitmq-defaults-descriptions)

- [Recommended default values](#rabbitmq-defaults-values)

## Policy and limit descriptions

The following list describes the default policies and limits that Amazon MQ applies to a
newly created broker. The values for `max-length`, `max-queues`,
and `max-connections` vary based on your broker's instance type and
deployment mode. These values are listed in the [Recommended default values](#rabbitmq-defaults-values) section.

**Settings on both RabbitMQ 3 and RabbitMQ 4 brokers**

- `queue-mode: lazy` (policy) –
Enables lazy queues. By default, queues keep an in-memory cache of messages,
enabling the broker to deliver messages to consumers as fast as possible. This
can lead to the broker running out of memory and raising a high-memory alarm.
Lazy queues attempt to move messages to disk as early as is practical. This
means that fewer messages are kept in memory under normal operating conditions.
Using lazy queues, Amazon MQ for RabbitMQ can support much larger messaging loads
and longer queues. Note that for certain use cases, brokers with lazy queues
might perform marginally slower. This is because messages are moved from disk to
broker, as opposed to delivering messages from an in-memory cache.

###### Deployment modes

Single-instance, cluster

- `max-length:
                                  number-of-messages`
(policy) – Sets a limit for the number of messages in a queue. In cluster deployments, the limit
prevents paused queue synchronization in cases such as broker reboots, or
following a maintenance window.

###### Deployment modes

Cluster

- `overflow: reject-publish` (policy)
– Enforces queues with a `max-length` policy to reject new
messages after the number of messages in the queue reaches the
`max-length` value. To ensure that messages aren't lost if a
queue is in an overflow state, client applications that publish messages to the
broker must implement [publisher\
confirms](best-practices-message-reliability.md#configure-confirmation-acknowledgement). For information about implementing publisher confirms, see
[Publisher Confirms](https://www.rabbitmq.com/confirms.html) on the RabbitMQ website.

###### Deployment modes

Cluster

**Settings specific to RabbitMQ 3**

- `max-queues:
                                  number-of-queues-per-vhost`
(vhost limit) – Sets the limit for the number of queues in a
broker. Similar to the `max-length` policy definition, limiting the number of queues
in cluster deployments prevents paused queue synchronization following broker reboots or maintenance windows.
Limiting queues also prevents excessive amounts of CPU usage for maintaining queues.

###### Deployment modes

Single-instance, cluster

- `max-connections:
                                  number-of-connections-per-vhost`
(vhost limit) – Sets the limit for the number of client connections to
the broker. Limiting the number of connections according to the recommended values
prevents excessive broker memory usage, which could result in the broker raising a high memory
alarm and pausing operations.

###### Deployment modes

Single-instance, cluster

## Recommended default values

###### Important

`max-queues` and `max-connections` are only applied to Amazon MQ for RabbitMQ 3.

###### Note

The `max-length` and `max-queue` default limits are tested and evaluated based on an average message size of 5 kB.
If your messages are significantly larger than 5 kB, you will need to adjust and reduce the `max-length` and `max-queue` limits.

The following table lists the default limit values for a newly created broker. Amazon MQ
applies these values according to the broker's instance type and deployment mode.

Instance typeDeployment mode`max-length``max-queues``max-connections`mq.m7g.mediumSingle-instanceN/A2,500100Cluster500,000100100mq.m7g.largeSingle-instanceN/A20,0005,000Cluster8,000,00010,0005,000mq.m7g.xlargeSingle-instanceN/A30,00010,000Cluster9,000,00015,00010,000mq.m7g.2xlargeSingle-instanceN/A40,00020,000Cluster10,000,00040,00020,000mq.m7g.4xlargeSingle-instanceN/A60,00040,000Cluster12,000,00030,00040,000mq.m7g.8xlargeSingle-instanceN/A80,00080,000Cluster20,000,00040,00080,000mq.m7g.12xlargeSingle-instanceN/A100,000120,000Cluster30,000,00020,000120,000mq.m7g.16xlargeSingle-instanceN/A120,000160,000Cluster40,000,00050,000160,000

Instance typeDeployment mode`max-length``max-queues``max-connections`t3.microSingle-instanceN/A500500m5.largeSingle-instanceN/A20,0004,000m5.largeCluster8,000,00010,00015,000m5.xlargeSingle-instanceN/A30,0008,000m5.xlargeCluster9,000,00010,00020,000m5.2xlargeSingle-instanceN/A60,00015,000m5.2xlargeCluster10,000,00010,00040,000m5.4xlargeSingle-instanceN/A150,00030,000m5.4xlargeCluster12,000,00010,000100,000

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Maximum resource limit

Broker configurations

All content copied from https://docs.aws.amazon.com/.
