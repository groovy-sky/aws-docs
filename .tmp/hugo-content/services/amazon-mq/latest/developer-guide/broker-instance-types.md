# Amazon MQ for ActiveMQ broker instance types

The combined description of the
broker instance _class_ ( `m5`) and
_size_ ( `large`, `medium`) is called the
_broker instance type_ (for example, `mq.m5.large`). The following table lists the available
Amazon MQ broker instance types for ActiveMQ brokers.

Amazon MQ provides at least a 90 day notice before an instance type reaches end of support.
We recommend upgrading your broker to a new instance type
before the end-of-support date to prevent any disruptions.

###### Important

You cannot create brokers on `t2.micro` or `mq.m4.large`
after March 17, 2025.

Instance TypevCPUMemory (GiB)Recommended UseStorageEnd of support on Amazon MQ`mq.t3.micro`21

Evaluation

EFS`mq.m5.large`28

Production

EFS or EBS`mq.m5.xlarge`416ProductionEFS or EBS`mq.m5.2xlarge`832ProductionEFS or EBS`mq.m5.4xlarge`1664ProductionEFS or EBS

For more information about throughput considerations, see [Choose the Correct Broker Instance Type for the Best Throughput](best-practices-activemq.md#broker-instance-types-choosing).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Network of brokers

Broker configurations

All content copied from https://docs.aws.amazon.com/.
