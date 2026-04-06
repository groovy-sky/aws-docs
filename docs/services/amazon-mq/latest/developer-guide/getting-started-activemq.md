# Getting started: Creating and connecting to an ActiveMQ broker

A _broker_ is a message broker environment running on Amazon MQ. It is the basic building block of Amazon MQ. The combined description of the
broker instance _class_ ( `m5`) and
_size_ ( `large`, `medium`) is called the
_broker instance type_ (for example, `mq.m5.large`). For more information, see [What is an Amazon MQ for ActiveMQ broker?](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/amazon-mq-basic-elements.html#broker).

## Create an ActiveMQ broker

The first and most common Amazon MQ task is creating a broker. The following example shows how you can use the AWS Management Console to create a basic broker.

1. Sign in to the [Amazon MQ console](https://console.aws.amazon.com/amazon-mq).

2. On the **Select broker engine** page, choose **Apache ActiveMQ**.

3. On the **Select deployment and storage** page, in the **Deployment mode and storage type** section, do the following:
1. Choose the **Deployment mode** (for example, **Active/standby broker**). For more information, see
       [Deployment options for Amazon MQ for ActiveMQ brokers](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/amazon-mq-broker-architecture.html).

- A **Single-instance broker**
is comprised of one broker in one Availability Zone. The broker communicates with your application and with an Amazon EBS or Amazon EFS storage volume. For more information, see
[Option 1: Amazon MQ single-instance brokers](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/amazon-mq-broker-architecture.html#single-broker-deployment).

- An **Active/standby broker for high availability**
is comprised of two brokers in two different Availability Zones,
configured in a _redundant pair_. These brokers communicate synchronously with your application, and with Amazon EFS. For more information, see
[Option 2: Amazon MQ active/standby brokers for high availability](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/amazon-mq-broker-architecture.html#active-standby-broker-deployment).

2. Choose the **Storage type** (for example, **EBS**). For more information, see
       [Storage](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/broker-storage.html).

      ###### Note

      Amazon EBS replicates data within a single Availability Zone and doesn't support the
      [ActiveMQ active/standby](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/amazon-mq-broker-architecture.html#active-standby-broker-deployment) deployment mode.

3. Choose **Next**.
4. On the **Configure settings** page, in the **Details** section, do the following:
1. Enter the **Broker name**.

      ###### Important

      Do not add personally identifiable information (PII) or other confidential or sensitive information in broker names.
      Broker names are accessible to other AWS services, including CloudWatch Logs. Broker names are not intended to be used for
      private or sensitive data.

      ###### Note

      In the **Additional settings** section, you can also configure the following:

- [Configurations](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/amazon-mq-broker-configuration-parameters.html)

- [CloudWatch logs](security-logging-monitoring.md)

- Private access

- [Broker maintenance window](maintaining-brokers.md)

2. Choose the **Broker instance type** (for example, **mq.m5.large**). For more information, see
       [Broker instance types](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/broker-instance-types.html).
5. In the **ActiveMQ Web Console access** section, provide a **Username**
    and **Password**. The following restrictions apply to broker usernames and passwords:

- Your username can contain only alphanumeric characters, dashes, periods, underscores, and tildas (- . \_ ~).

- Your password must be at least 12 characters long, contain at least 4 unique characters and must not contain commas, colons, or equal signs (,:=).

###### Important

Do not add personally identifiable information (PII) or other confidential or sensitive information in broker usernames.
Broker usernames are accessible to other AWS services, including CloudWatch Logs. Broker usernames are not intended to be used for
private or sensitive data.

6. Choose **Deploy**.

While Amazon MQ creates your broker, it displays the **Creation in progress** status.

Creating the broker takes about 15 minutes.

When your broker is created successfully, Amazon MQ displays the **Running** status.

7. Choose **`MyBroker`**.

On the **`MyBroker`** page,
    in the **Connect** section, note your
    broker's **[ActiveMQ web console](http://activemq.apache.org/web-console.html)** URL, for example:

```nohighlight

https://b-1234a5b6-78cd-901e-2fgh-3i45j6k178l9-1.mq.us-east-2.amazonaws.com:8162
```

Also, note your broker's [wire-level\
    protocol **Endpoints**](http://activemq.apache.org/configuring-transports.html). The following is an example
    of an OpenWire endpoint:

```nohighlight

ssl://b-1234a5b6-78cd-901e-2fgh-3i45j6k178l9-1.mq.us-east-2.amazonaws.com:61617
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Setting up

Getting started: Creating and connecting to a RabbitMQ broker
