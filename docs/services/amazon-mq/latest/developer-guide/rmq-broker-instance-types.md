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

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Deploying a RabbitMQ broker

Sizing guidelines
