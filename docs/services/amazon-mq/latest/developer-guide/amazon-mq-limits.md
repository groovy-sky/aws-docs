# Quotas in Amazon MQ

This topic lists limits within Amazon MQ. Many of the following limits can be changed for
specific AWS accounts. To request an increase for a limit, see [AWS Service Quotas](../../../../general/latest/gr/aws-service-limits.md) in the
_Amazon Web Services General Reference_. Updated limits will not be visible even after the limit increase has been applied.
For more information on viewing current connection limits in Amazon CloudWatch, see [Monitoring Amazon MQ brokers using Amazon CloudWatch](security-logging-monitoring-cloudwatch.md).

###### Topics

- [Brokers](#broker-limits)

- [Configurations](#configuration-limits)

- [Users](#activemq-user-limits)

- [Data Storage](#data-storage-limits)

- [API Throttling](#api-throttling-limits)

## Brokers

The following table lists quotas related to Amazon MQ brokers.

LimitDescriptionBroker name

- Must be unique in your AWS account.

- Must be 1-50 characters long.

- Must contain only characters specified in the [ASCII Printable Character Set](https://en.wikipedia.org/wiki/ASCII).

- Can contain only alphanumeric characters, dashes, periods,
underscores, and tildes ( `- . _ ~`).

Number of brokers, per region200Wire-level connections per protocol for smaller broker

###### Important

Does not apply to RabbitMQ brokers.

300 for `mq.*.micro` instance type brokers. Wire-level connections per protocol for larger broker

###### Important

Does not apply to RabbitMQ brokers.

2,000 for `mq.*.*large` instance type brokers. Security groups per broker5ActiveMQ destinations (queues, and topics) monitored in CloudWatchCloudWatch monitors only the first 1000 destinations. RabbitMQ destinations (queues) monitored in CloudWatchCloudWatch monitors only the first 500 destinations, ordered by number of consumers.Tags per broker50

## Configurations

The following table lists quotas related to Amazon MQ configurations.

LimitDescriptionConfiguration name

- Must be 1-150 characters long.

- Must contain only characters specified in the [ASCII Printable Character Set](https://en.wikipedia.org/wiki/ASCII).

- Can contain only alphanumeric characters, dashes, periods,
underscores, and tildes ( `- . _ ~`).

Revisions per configuration300

## Users

The following table lists quotas related to Amazon MQ ActiveMQ broker users.

LimitDescriptionUsername

- Must be 1-100 characters long.

- Must contain only characters specified in the [ASCII Printable Character Set](https://en.wikipedia.org/wiki/ASCII).

- Can contain only alphanumeric characters, dashes, periods,
underscores, and tildes ( `- . _ ~`).

- Must not contain commas ( `,`).

Password

- Must be 12-250 characters long.

- Must contain only characters specified in the [ASCII Printable Character Set](https://en.wikipedia.org/wiki/ASCII).

- Must contain at least 4 unique characters.

- Must not contain commas ( `,`).

Users per broker (simple auth)250Groups per user (simple auth)20

## Data Storage

The following table lists quotas related to Amazon MQ data storage.

LimitDescriptionStorage capacity per smaller broker20 GB for `mq.*.micro` instance type brokers. For more information regarding
Amazon MQ instance types, see [Broker instance types](broker-instance-types.md).Storage capacity per larger broker200 GB for `mq.m5.*` instance type brokers.
For more information regarding
Amazon MQ instance types, see [Broker instance types](broker-instance-types.md).Job scheduler usage limit per broker [backed by Amazon EBS](broker-storage.md)

###### Important

Does not apply to RabbitMQ brokers.

50 GB. For more information about job scheduler usage, see `JobSchedulerUsage` in the _Apache ActiveMQ API Documentation_.
Temporary storage capacity per smaller broker.

###### Important

Does not apply to RabbitMQ brokers.

5 GB for `mq.*.micro` instance type brokers.
Temporary storage capacity per larger broker.

###### Important

Does not apply to RabbitMQ brokers.

50 GB for `mq.m5.*` instance type brokers.

## API Throttling

The following throttling quotas are aggregated per AWS account, _across all_
_Amazon MQ APIs_ to maintain service bandwidth. For more information about
Amazon MQ APIs, see the _[Amazon MQ REST API Reference](../api-reference.md)_.

###### Important

These quotas don't apply to Amazon MQ for ActiveMQ or Amazon MQ for RabbitMQ broker messaging APIs. For example, Amazon MQ
doesn't throttle the sending or receiving of messages.

API burst limitAPI rate limit10015

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Troubleshooting

Troubleshooting

All content copied from https://docs.aws.amazon.com/.
