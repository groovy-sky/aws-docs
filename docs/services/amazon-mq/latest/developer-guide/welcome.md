# What is Amazon MQ?

Amazon MQ is a managed message broker service for [Apache ActiveMQ](http://activemq.apache.org/) Classic
and [RabbitMQ](https://www.rabbitmq.com/) that manages the setup, operation, and maintenance of message brokers.
You can create a new Amazon MQ broker using industry standard messaging protocols,
or migrate existing message brokers to Amazon MQ without rewriting messaging code.

A _broker_ is a message broker environment running on Amazon MQ.
It is the basic building block of Amazon MQ.
A _message_ broker allows software applications and components
to communicate using various programming languages, operating systems, and formal messaging protocols.
You can use Amazon MQ brokers for communication between large scale,
cloud native applications and components.

###### Topics

- [Amazon MQ features](#amazonmq-features)

- [How can I get started with Amazon MQ?](#get-started)

- [How can I provide feedback to Amazon MQ?](#amazon-mq-we-want-to-hear-from-you)

## Amazon MQ features

**Managed maintenance and version upgrades**

Amazon MQ performs [maintenance](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/maintaining-brokers.html) and
[version upgrades](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/upgrading-brokers.html)
for a message broker during your scheduled [maintenance window](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/maintaining-brokers.html).

**Monitor brokers with CloudWatch**

Amazon MQ is integrated with [Amazon CloudWatch](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/security-logging-monitoring.html)
so you can view and analyze metrics for your brokers and queues.
You can view and analyze metrics from the Amazon MQ console, the CloudWatch console, command line, and API.
Metrics are automatically collected and pushed to CloudWatch every minute.

**Security**

Amazon MQ provides [encryption](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/data-protection.html) of your messages at rest and in transit.
Connections to the broker use SSL, and access can be restricted to
a private endpoint within your Amazon VPC.
Additonality, you can use [AWS Identity and Access Management](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/security-iam.html) (IAM) to control the actions your
IAM users and groups can take on specific Amazon MQ brokers.

**Quorum queues for RabbitMQ on Amazon MQ**

[Quorum queues](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/quorum-queues.html) are a replicated queue type made up of a leader node (primary replica) and follower nodes (other replicas).
Each node is in a different availability zone, so if one node is temporarily unavailable,
message delivery continues with a newly elected leader replica in another availability zone.
Quorum queues are useful for handling poison messages,
which occur when a message fails and is requeued multiple times.

**Cross-Region data replication for ActiveMQ on Amazon MQ**

[Cross-Region data replication](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/crdr-for-active-mq.html) (CRDR) allows for asynchronous message replication from
the primary broker in a primary AWS Region to the replica broker in a replica Region.
By issuing a failover request to the Amazon MQ API,
the current replica broker is promoted to the primary broker role,
and the current primary broker is demoted to the replica role.

## How can I get started with Amazon MQ?

To get started with _ActiveMQ on Amazon MQ_, review the following documentation:

- [Getting started: Creating and connecting to an ActiveMQ broker](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/getting-started-activemq.html)

- [Deployment options for Amazon MQ for ActiveMQ brokers](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/amazon-mq-broker-architecture.html)

- [ActiveMQ tutorials](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/activemq-on-amazon-mq.html)

- [Amazon MQ for ActiveMQ best practices](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/best-practices-activemq.html)

To get started with _RabbitMQ on Amazon MQ_, review the following documentation:

- [Getting started: Creating and connecting to a RabbitMQ broker](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/getting-started-rabbitmq.html)

- [Deployment options for Amazon MQ for RabbitMQ brokers](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/rabbitmq-broker-architecture.html)

- [RabbitMQ tutorials](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/rabbitmq-on-amazon-mq.html)

- [Amazon MQ for RabbitMQ best practices](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/best-practices-rabbitmq.html)

To learn about Amazon MQ REST APIs, see the _[Amazon MQ REST API Reference](https://docs.aws.amazon.com/amazon-mq/latest/api-reference)_.

To learn about Amazon MQ AWS CLI commands, see [Amazon MQ in the\
_AWS CLI Command Reference_](https://docs.aws.amazon.com/cli/latest/reference/mq/index.html).

## How can I provide feedback to Amazon MQ?

We welcome and encourgae your feedback on the documentation.
You can use the thumbs up and thumbs down icons on the right hand side to submit feedback, or
you can use the
"Provide feedback" form linked below.

To contact the Amazon MQ team, use the [Amazon MQ Discussion\
Forum](https://forums.aws.amazon.com/forum.jspa?forumID=279).

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Setting up
