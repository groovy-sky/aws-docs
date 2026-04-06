# Default resource limits

Amazon MQ for RabbitMQ supports configuring the broker resource limits from RabbitMQ 4 onwards. When you create a broker, Amazon MQ automatically applies default values to these resource limits. These defaults act as guardrails to protect your broker availability while accommodating common customer usage patterns. You can customize your broker behavior by changing the limit configuration values to better match your specific workload requirements.

Before making changes, please note:

###### Important

1. Configuration changes can impact broker performance and availability

2. Understand the impact before changing any default configuration options

3. Test configuration changes in non-production environments first

4. Monitor broker health after applying changes

###### Important

Default values and supported ranges for these configurations vary by RabbitMQ version, instance type, and broker deployment mode.

###### Important

_Note: Associating or creating a broker with configuration values outside the supported range will result in an error response._

To learn how to customize these default resource limits for your broker, see [Resource Limit Configuration](configure-resource-limits.md).

The default resource limits applied for RabbitMQ 4.2 brokers are

- [Default resource limits for m7g single-instance deployment](#default-values-single-instance)

- [Default resource limits for m7g cluster deployment](#default-values-cluster-brokers)

## Default resource limits

###### Important

Amazon MQ for RabbitMQ 3 brokers, the default is configured with the maximum resource limit and Amazon MQ does not provide the ability to override resource limit configuration.

### Default values for single instance brokers

Instance typeConnections per NodeChannels per NodeConsumers per channelQueuesvhostsShovelsExchangesMessage size in Bytesmq.m7g.medium100500105001030500524288mq.m7g.large1,5004,500101,00050501,000524288mq.m7g.xlarge3,0009,000102,0001001002,000524288mq.m7g.2xlarge6,00018,000104,0001502004,000524288mq.m7g.4xlarge12,00036,000108,0002004008,000524288mq.m7g.8xlarge24,00072,0001016,00025080016,000524288mq.m7g.12xlarge36,000108,0001024,0003001,20024,000524288mq.m7g.16xlarge48,000144,0001032,0003501,60032,000524288

### Default values for cluster brokers

Instance typeConnections per NodeChannels per NodeConsumers per channelQueuesvhostsShovelsExchangesMessage size in Bytesmq.m7g.medium100300101001010100524288mq.m7g.large5001500101,00050301,000524288mq.m7g.xlarge10003000102,000100602,000524288mq.m7g.2xlarge20006000104,0001501204,000524288mq.m7g.4xlarge400012,000108,0002002408,000524288mq.m7g.8xlarge800024,0001016,00025048016,000524288mq.m7g.12xlarge1200036,0001024,00030072024000524288mq.m7g.16xlarge16,00048,0001032,00035096032,000524288

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Sizing guidelines

Maximum resource limit
