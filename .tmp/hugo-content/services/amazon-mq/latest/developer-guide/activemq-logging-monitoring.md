# Available CloudWatch metrics Amazon MQ for ActiveMQ brokers

## Amazon MQ for ActiveMQ metrics

MetricUnitDescription`AmqpMaximumConnections`CountThe maximum number of clients you can connect to your broker using AMQP. For more information on connection quotas, see [Quotas in Amazon MQ](amazon-mq-limits.md).`BurstBalance`PercentThe percentage of burst credits remaining on the Amazon EBS volume
used to persist message data for throughput-optimized brokers. If
this balance reaches zero, the IOPS provided by the Amazon EBS volume
will decrease until the Burst Balance refills. For more information
on how Burst Balances work in Amazon EBS, see: [I/O\
Credits and Burst Performance](../../../ec2/latest/userguide/ebs-volume-types.md#IOcredit).`CpuCreditBalance`Credits (vCPU-minutes)

###### Important

This metric is available only for the
`mq.t2.micro` broker instance type.

CPU credit metrics are available only at five-minute
intervals.

The number of earned CPU credits that an instance has accrued
since it was launched or started (including the number of launch
credits). The credit balance is available for the broker
instance to spend on bursts beyond the baseline CPU
utilization.

Credits are accrued in the credit balance after they're
earned and removed from the credit balance after they're
spent. The credit balance has a maximum limit. Once the limit is
reached, any newly earned credits are discarded.

`CpuUtilization`PercentThe percentage of allocated Amazon EC2 compute units that the broker
currently uses.`CurrentConnectionsCount`CountThe current number of active connections on the current
broker.`EstablishedConnectionsCount`CountThe total number of connections, active and inactive, that have
been established on the broker.`HeapUsage`PercentThe percentage of the ActiveMQ JVM memory limit that the broker
currently uses.`InactiveDurableTopicSubscribersCount`CountThe number of inactive durable topic subscribers, up to a maximum
of 2000. `JobSchedulerStorePercentUsage`PercentThe percentage of disk space used by the job scheduler
store.`JournalFilesForFastRecovery`CountThe number of journal files that will be replayed after a clean
shutdown.`JournalFilesForFullRecovery`CountThe number of journal files that will be replayed after an
unclean shutdown.`MqttMaximumConnections`CountThe maximum number of clients you can connect to your broker using MQTT. For more information on connection quotas, see [Quotas in Amazon MQ](amazon-mq-limits.md).`NetworkConnectorConnectionCount`CountThe number of nodes connected to the broker in a [network of brokers](network-of-brokers.md) using NetworkConnector.`NetworkIn`BytesThe volume of incoming traffic for the broker.`NetworkOut`BytesThe volume of outgoing traffic for the broker.`OpenTransactionCount`CountThe total number of transactions in progress.`OpenwireMaximumConnections`CountThe maximum number of clients you can connect to your broker using OpenWire. For more information on connection quotas, see [Quotas in Amazon MQ](amazon-mq-limits.md).`StompMaximumConnections`CountThe maximum number of clients you can connect to your broker using STOMP. For more information on connection quotas, see [Quotas in Amazon MQ](amazon-mq-limits.md).`StorePercentUsage`PercentThe percent used by the storage limit. If this reaches 100, the
broker will refuse messages.`TempPercentUsage`PercentThe percentage of available temporary storage used by
non-persistent messages. `TotalConsumerCount`CountThe number of message consumers subscribed to destinations on the
current broker.`TotalMessageCount`CountThe number of messages stored on the broker.`TotalProducerCount`CountThe number of message producers active on destinations on the
current broker.`VolumeReadOps`CountThe number of read operations performed on the Amazon EBS
volume.`VolumeWriteOps`CountThe number of write operations performed on the Amazon EBS
volume.`WsMaximumConnections`CountThe maximum number of clients you can connect to your broker using WebSocket. For more information on connection quotas, see [Quotas in Amazon MQ](amazon-mq-limits.md).

### Dimensions for ActiveMQ broker metrics

DimensionDescription`Broker`

The name of the broker

###### Note

A single-instance broker has the suffix -1. An active/standby broker for high availability has the suffixes -1 and -2 for its redundant pair.

## ActiveMQ destination (queue and topic) metrics

###### Important

The following metrics include per-minute counts for the CloudWatch polling
period.

- `EnqueueCount`

- `ExpiredCount`

- `DequeueCount`

- `DispatchCount`

- `InFlightCount`

For example, in a five-minute [CloudWatch period](../../../amazoncloudwatch/latest/monitoring/cloudwatch-concepts.md#CloudWatchPeriods), `EnqueueCount` has five count values, each
for a one-minute portion of the period. The `Minimum` and
`Maximum` statistics provide the lowest and highest per-minute
value during the specified period.

MetricUnitDescription`ConsumerCount`CountThe number of consumers subscribed to the destination.`EnqueueCount`CountThe number of messages sent to the destination, per
minute.`EnqueueTime`Time (milliseconds)The end-to-end latency from when a message arrives at a broker
until it is delivered to a consumer.

###### Note

`EnqueueTime` does not measure the end-to-end latency from when a message is sent by a producer until
it reaches the broker, nor the latency from when a message is received by a broker until it is acknowledged
by the broker. Rather, `EnqueueTime` is the number of milliseconds from the moment a message is received by the broker
until it is successfully delivered to a consumer.

`ExpiredCount`CountThe number of messages that couldn't be delivered because
they expired, per minute.`DispatchCount`CountThe number of messages sent to consumers, per minute.`DequeueCount`CountThe number of messages acknowledged by consumers, per
minute.`InFlightCount`CountThe number of messages sent to consumers that have not been
acknowledged.`ReceiveCount`CountThe number of messages that have been received from the remote
broker for a duplex network connector.`MemoryUsage`PercentThe percentage of the memory limit that the destination currently
uses.`ProducerCount`CountThe number of producers for the destination.`QueueSize`CountThe number of messages in the queue.

###### Important

This metric applies only to queues.

`TotalEnqueueCount`CountThe total number of messages that have been sent to the
broker.`TotalDequeueCount`CountThe total number of messages that have been consumed by
clients.

###### Note

`TotalEnqueueCount` and `TotalDequeueCount` metrics
include messages for advisory topics. For more information about advisory topic
messages, see the [ActiveMQ documentation](https://activemq.apache.org/advisory-message.html).

### Dimensions for ActiveMQ destination (queue and topic) metrics

DimensionDescription`Broker`

The name of the broker.

###### Note

A single-instance broker has the suffix
`-1`. An active/standby broker for high
availability has the suffixes `-1` and
`-2` for its redundant pair.

`Topic` or `Queue`The name of the topic or queue.`NetworkConnector`The name of the network connector.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Accessing CloudWatch metrics

Metrics for RabbitMQ

All content copied from https://docs.aws.amazon.com/.
