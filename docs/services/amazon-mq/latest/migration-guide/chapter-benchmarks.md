# Amazon MQ for ActiveMQ Throughput benchmarks

Benchmarking can help you choose the correct instance type and size for your workload
messaging requirements. Scenarios for benchmarking include:

- **Cluster Stability**:
understanding how stable your cluster is during increasing, fluctuating, and stable load types.

- **Defining performance limits**:
approximating the maximum performance and throughput capabilities (i.e. cluster limits) of your cluster
to help better scale your broker nodes when the number of messages published to your broker increases.

- **Optimal architecture and parameters**:
determining the most suitable architecture/parameters for your clusters, such as the number of destinations, persistent mode, message size, etc.

Amazon MQ provides benchmarking figures for the different instance types and sizes available for Amazon MQ for ActiveMQ.

Amazon MQ uses the [ActiveMQ Maven 2\
Performance Test](https://activemq.apache.org/activemq-performance-module-users-manual) to calculate benchmarks. Amazon MQ tests with an active/stanby deployment
broker with an EFS volume as a storage type. The results are from a performance test conducted on
a ECS cluster with a Fargate deployment which consists of a configuration of 4vCPUs and 16 GiBs of
memory (equivalent to an EC2 m5.xlarge instance). Concurrent Store And Dispatch Queues (CSAD) are
set to true for all tests performed. The duration for each test conducted is 5 minutes.

###### Note

When run in your own environment, results may differ by 3-6%.

The following tables provide performance and throughput benchmarks for
Amazon MQ supported instance types to help you choose the correct instance sizes for your messaging workload.

###### Topics

- [mq.m4.large](#amazon-mq-broker-instance-mq.m4.large)

- [mq.m5.large](#amazon-mq-broker-instance-mq.m5.large)

- [mq.m5.xlarge](#amazon-mq-broker-instance-mq.m5.xlarge)

- [mq.m5.2xlarge](#amazon-mq-broker-instance-mq.m5.2xlarge)

- [mq.m5.4xlarge](#amazon-mq-broker-instance-mq.m5.4xlarge)

## `mq.m4.large`

Configuration options:

- **Broker Instance -** `mq.m4.large`

- **Persistent -** `TRUE`

- **Client -** `m5.xlarge`

- **CSAD -** `TRUE`

- **Protocol -** Openwire

Producers/ConsumersMessage sizeMetrics25501001KBTPS184933354665CPU%29%37%47%5KBTPS167225612970CPU%33%47%76%10KBTPS158616702268CPU%44%87%89%

## `mq.m5.large`

Configuration options:

- **Broker Instance -** `mq.m5.large`

- **Persistent -** `TRUE`

- **Client -** `m5.xlarge`

- **CSAD -** `TRUE`

- **Protocol -** Openwire

Producers/ConsumersMessage sizeMetrics25501001KBTPS224740417566CPU%26%32%48%5KBTPS163632054443CPU%37%63%58%10KBTPS166831043227CPU%40%53%86%

## `mq.m5.xlarge`

Configuration options:

- **Broker Instance -** `mq.m5.xlarge`

- **Persistent -** `TRUE`

- **Client -** `m5.xlarge`

- **CSAD -** `TRUE`

- **Protocol -** Openwire

Producers/ConsumersMessage sizeMetrics25501001KBTPS225539327453CPU%28%32%54%5KBTPS176634956215CPU%29%51%82%10KBTPS164132405613CPU%36%61%89%

## `mq.m5.2xlarge`

Configuration options:

- **Broker Instance -** `mq.m5.2xlarge`

- **Persistent -** `TRUE`

- **Client -** `m5.xlarge`

- **CSAD -** `TRUE`

- **Protocol -** Openwire

Producers/ConsumersMessage sizeMetrics25501001KBTPS202540898093CPU%12%18%35%5KBTPS186537366845CPU%15%27%54%10KBTPS174735117057CPU%18%36%67%

## `mq.m5.4xlarge`

Configuration options:

- **Broker Instance -** `mq.m5.4xlarge`

- **Persistent -** `TRUE`

- **Client -** `m5.xlarge`

- **CSAD -** `TRUE`

- **Protocol -** Openwire

Producers/ConsumersMessage sizeMetrics25501001KBTPS209440558153CPU%6%9%17%5KBTPS174235867158CPU%7%13%25%10KBTPS173332886671CPU%9%16%31%

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Migrating to Amazon MQ for RabbitMQ

Supported plugins

All content copied from https://docs.aws.amazon.com/.
