# Hardware specifications for DB instance classesfor Aurora

In the table in this section, you can find hardware details about the Amazon RDS DB instance
classes for Aurora.

For information about Aurora DB engine support for each DB instance class, see
[Supported DB engines for DB instance classes](concepts-dbinstanceclass-supportaurora.md).

###### Topics

- [Hardware terminology for DB instance classesfor Aurora](#Concepts.DBInstanceClass.hardware-terminology)

- [Hardware specifications for the memory-optimized instance classes](#hw-specs-aur.mem-opt)

- [Hardware specifications for the burstable-performance instance classes](#hardware-specifications.burstable-inst-classes)

## Hardware terminology for DB instance classesfor Aurora

The following terminology is used to describe hardware specifications for DB instance
classes:

**vCPU**

The number of virtual central processing units (CPUs). A _virtual CPU_ is a unit of capacity that you can
use to compare DB instance classes. Instead of purchasing or leasing a particular
processor to use for several months or years, you are renting capacity by
the hour. Our goal is to make a consistent and specific amount of CPU
capacity available, within the limits of the actual underlying
hardware.

**ECU**

The relative measure of the integer processing power of an Amazon EC2 instance.
To make it easy for developers to compare CPU capacity between different
instance classes, we have defined an Amazon EC2 Compute Unit. The amount of CPU
that is allocated to a particular instance is expressed in terms of these
EC2 Compute Units. One ECU currently provides CPU capacity equivalent to a
1.0–1.2 GHz 2007 Opteron or 2007 Xeon processor.

**Memory (GiB)**

The RAM, in gibibytes, allocated to the DB instance. There is often a consistent
ratio between memory and vCPU. As an example, take the db.r4 instance class,
which has a memory to vCPU ratio similar to the db.r5 instance class.
However, for most use cases the db.r5 instance class provides better, more
consistent performance than the db.r4 instance class.

**Max. EBS bandwidth (Mbps)**

The maximum EBS bandwidth in megabits per second. Divide by 8 to get the
expected throughput in megabytes per second.

###### Note

This figure refers to I/O bandwidth for local storage within the DB
instance. It doesn't apply to communication with the Aurora cluster
volume.

**Network bandwidth**

The network speed relative to other DB instance classes.

For information on using Amazon CloudWatch metrics to monitor your Aurora DB instance throughput, see [Evaluating DB instance usage for Aurora MySQL with Amazon CloudWatch metrics](auroramysql-bestpractices-cw.md) and [Evaluating DB instance usage for Aurora PostgreSQL with CloudWatch metrics](aurorapostgresql-anayzeresourceusage.md#AuroraPostgreSQL_AnayzeResourceUsage.EvaluateInstanceUsage).

## Hardware specifications for the memory-optimized instance classes

The following tables show the compute, memory, storage, and bandwidth specifications for the memory-optimized instance classes.

**db.x2g – memory-optimized instance classes with AWS Graviton2 processors**

Instance classvCPUECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.x2g.16xlarge64—1024EBS-optimized only19,00025db.x2g.12xlarge48—768EBS-optimized only14,25020db.x2g.8xlarge32—512EBS-optimized only9,50012db.x2g.4xlarge16—256EBS-optimized only4,750Up to 10db.x2g.2xlarge8—128EBS-optimized onlyUp to 4,750Up to 10db.x2g.xlarge4—64EBS-optimized onlyUp to 4,750Up to 10db.x2g.large2—32EBS-optimized onlyUp to 4,750Up to 10

**db.r8gd – memory-optimized instance classes powered by AWS**
**Graviton4 processors and SSD storage**

Instance classvCPUECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.r8gd.48xlarge192—15366 x 1900 NVMe SSD40,00050db.r8gd.24xlarge96—7683 x 1900 NVMe SSD30,00040db.r8gd.16xlarge64—5122 x 1900 NVMe SSD20,00030db.r8gd.12xlarge48—3843 x 950 NVMe SSD15,00022.5db.r8gd.8xlarge32—2561 x 1900 NVMe SSD10,00015db.r8gd.4xlarge16—1281 x 950 NVMe SSDUp to 10,000Up to 15db.r8gd.2xlarge8—641 x 474 NVMe SSDUp to 10,000Up to 15db.r8gd.xlarge4—321 x 237 NVMe SSDUp to 10,000Up to 12.5db.r8gd.large2—161 x 118 NVMe SSDUp to 10,000Up to 12.5

**db.r8g – memory-optimized instance classes powered by AWS Graviton4 processors**

Instance classvCPUECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.r8g.48xlarge192—1536EBS-optimized only40,00050db.r8g.24xlarge96—768EBS-optimized only30,00040db.r8g.16xlarge64—512EBS-optimized only20,00030db.r8g.12xlarge48—384EBS-optimized only15,00022.5db.r8g.8xlarge32—256EBS-optimized only10,00015db.r8g.4xlarge16—128EBS-optimized onlyUp to 10,000Up to 15db.r8g.2xlarge8—64EBS-optimized onlyUp to 10,000Up to 15db.r8g.xlarge4—32EBS-optimized onlyUp to 10,000Up to 12.5db.r8g.large2—16EBS-optimized onlyUp to 10,000Up to 12.5

**db.r7i – memory-optimized instance classes powered by 4th generation Intel Xeon Scalable processors**

Instance classvCPUECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.r7i.48xlarge192—1536EBS-optimized only40,00050db.r7i.24xlarge96—768EBS-optimized only30,00037.5db.r7i.16xlarge64—512EBS-optimized only20,00025db.r7i.12xlarge48—384EBS-optimized only15,00018.75db.r7i.8xlarge32—256EBS-optimized only10,00012.5db.r7i.4xlarge16—128EBS-optimized onlyUp to 10,000Up to 12.5db.r7i.2xlarge8—64EBS-optimized onlyUp to 10,000Up to 12.5db.r7i.xlarge4—32EBS-optimized onlyUp to 10,000Up to 12.5db.r7i.large2—16EBS-optimized onlyUp to 10,000Up to 12.5

**db.r7g – memory-optimized instance classes with AWS Graviton3 processors**

Instance classvCPUECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.r7g.16xlarge64—512EBS-optimized only20,00030db.r7g.12xlarge48—384EBS-optimized only15,00022.5db.r7g.8xlarge32—256EBS-optimized only10,00015db.r7g.4xlarge16—128EBS-optimized onlyUp to 10,000Up to 15db.r7g.2xlarge8—64EBS-optimized onlyUp to 10,000Up to 15db.r7g.xlarge4—32EBS-optimized onlyUp to 10,000Up to 12.5db.r7g.large2—16EBS-optimized onlyUp to 10,000Up to 12.5

**db.r6id – memory-optimized instance classes with 3rd generation Intel Xeon Scalable processors and SSD**
**storage**

Instance classvCPUECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.r6id.32xlarge128—1,0244x1900 NVMe SSD40,00050db.r6id.24xlarge96—7684x1425 NVMe SSD30,00037.5

**db.r6gd – memory-optimized instance classes with AWS Graviton2 processors and SSD storage**

Instance classvCPUECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.r6gd.16xlarge64—5122 x 1900 NVMe SSD19,00025db.r6gd.12xlarge48—3842 x 1425 NVMe SSD13,50020db.r6gd.8xlarge32—2561 x 1900 NVMe SSD9,00012db.r6gd.4xlarge16—1281 x 950 NVMe SSD4,750Up to 10 db.r6gd.2xlarge8—641 x 474 NVMe SSDUp to 4,750Up to 10 db.r6gd.xlarge4—321 x 237 NVMe SSDUp to 4,750Up to 10

**db.r6g – memory-optimized instance classes with AWS Graviton2 processors**

Instance classvCPUECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.r6g.16xlarge64—512EBS-optimized only19,00025db.r6g.12xlarge48—384EBS-optimized only13,50020db.r6g.8xlarge32—256EBS-optimized only9,00012db.r6g.4xlarge16—128EBS-optimized only4,750Up to 10 db.r6g.2xlarge8—64EBS-optimized onlyUp to 4,750Up to 10 db.r6g.xlarge4—32EBS-optimized onlyUp to 4,750Up to 10 db.r6g.large2—16EBS-optimized onlyUp to 4,750Up to 10

**db.r6i – memory-optimized instance classes with 3rd Generation Intel Xeon Scalable processors**

Instance classvCPUECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.r6i.32xlarge128—1,024EBS-optimized only40,00050db.r6i.24xlarge96—768EBS-optimized only30,00037.5db.r6i.16xlarge64—512EBS-optimized only20,00025db.r6i.12xlarge48—384EBS-optimized only15,00018.75db.r6i.8xlarge32—256EBS-optimized only10,00012.5db.r6i.4xlarge16—128EBS-optimized onlyUp to 10,000Up to 12.5db.r6i.2xlarge8—64EBS-optimized onlyUp to 10,000Up to 12.5db.r6i.xlarge4—32EBS-optimized onlyUp to 10,000Up to 12.5db.r6i.large2—16EBS-optimized onlyUp to 10,000Up to 12.5

**db.r5 – memory-optimized instance classes**

Instance classvCPUECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.r5.24xlarge96347768EBS-optimized only19,00025db.r5.16xlarge64264512EBS-optimized only13,60020db.r5.12xlarge48173384EBS-optimized only9,50012db.r5.8xlarge32132256EBS-optimized only6,80010db.r5.4xlarge1671128EBS-optimized only4,750Up to 10db.r5.2xlarge83864EBS-optimized onlyUp to 4,750Up to 10db.r5.xlarge41932EBS-optimized onlyUp to 4,750Up to 10db.r5.large21016EBS-optimized onlyUp to 4,750Up to 10

**db.r4 – memory-optimized instance classes with Intel Xeon Scalable processors**

Instance classvCPUECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.r4.16xlarge64195488EBS-optimized only14,00025db.r4.8xlarge3299244EBS-optimized only7,00010db.r4.4xlarge1653122EBS-optimized only3,500Up to 10db.r4.2xlarge82761EBS-optimized only1,700Up to 10db.r4.xlarge413.530.5EBS-optimized only850Up to 10db.r4.large2715.25EBS-optimized only425Up to 10

## Hardware specifications for the burstable-performance instance classes

The following tables show the compute, memory, storage, and bandwidth specifications for the burstable-performance instance classes.

**db.t4g – burstable-performance instance classes powered by AWS Graviton2 processors**

Instance classvCPUECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.t4g.large2—8EBS-optimized onlyUp to 2,780Up to 5db.t4g.medium2—4EBS-optimized onlyUp to 2,085Up to 5

**db.t3 – burstable-performance instance classes**

Instance classvCPUECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.t3.large2Variable8EBS-optimized onlyUp to 2,048Up to 5db.t3.medium2Variable4EBS-optimized onlyUp to 1,536Up to 5db.t3.small2Variable2EBS-optimized onlyUp to 1,536Up to 5

**db.t2 – burstable-performance instance classes**

Instance classvCPUECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.t2.medium2Variable4EBS only—Moderatedb.t2.small1Variable2EBS only—Low

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Determining DB
instance class support in AWS Regions

Storage

All content copied from https://docs.aws.amazon.com/.
