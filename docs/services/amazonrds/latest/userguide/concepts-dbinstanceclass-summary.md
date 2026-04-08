# Hardware specifications for DB instance classes

In the tables in this section, you can find hardware details about the Amazon RDS DB instance
classes.

For information about Amazon RDS DB engine support for each DB instance class,
see [Supported DB engines for DB instance classes](concepts-dbinstanceclass-support.md).

###### Topics

- [Hardware terminology for DB instance classes](#Concepts.DBInstanceClass.hardware-terminology)

- [Hardware specifications for the general-purpose instance classes](#hardware-specifications.gen-purpose-inst-classes)

- [Hardware specifications for the memory-optimized instance classes](#hardware-specifications.mem-opt-inst-classes)

- [Hardware specifications for the compute-optimized instance classes](#hardware-specifications.compute-opt-inst-classes)

- [Hardware specifications for the burstable-performance instance classes](#hardware-specifications.burstable-inst-classes)

###### Note

RDS for SQL Server supports Optimize CPU starting with 7th generation instance classes (such as db.m7i and db.r7i). The vCPU counts documented below may differ for these instance classes. For accurate vCPU counts, refer to [DB instance classes that support Optimize CPU](sqlserver-concepts-general-optimizecpu-support.md).

## Hardware terminology for DB instance classes

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

**EBS-optimized**

The DB instance uses an optimized configuration stack and provides additional,
dedicated capacity for I/O. This optimization provides the best performance
by minimizing contention between I/O and other traffic from your instance.
For more information about Amazon EBS–optimized instances, see [Amazon EBS–Optimized\
instances](../../../ec2/latest/userguide/ebsoptimized.md) in the _Amazon EC2 User Guide._

EBS-optimized instances have a baseline and maximum IOPS rate. The maximum IOPS rate is
enforced at the DB instance level. A set of EBS volumes that combine to have an IOPS rate that is
higher than the maximum can't exceed the instance-level threshold. For example, if the maximum
IOPS for a particular DB instance class is 40,000, and you attach four 64,000 IOPS EBS volumes, the
maximum IOPS is 40,000 rather than 256,000. For the IOPS maximum specific to each EC2 instance
type, see [Supported instance\
types](../../../ec2/latest/userguide/ebs-optimized.md#ebs-optimization-support) in the _Amazon EC2 User Guide for Linux Instances_.

**Max. EBS bandwidth (Mbps)**

The maximum EBS bandwidth in megabits per second. Divide by 8 to get the
expected throughput in megabytes per second.

###### Important

General Purpose SSD (gp2) volumes for Amazon RDS DB instances have a throughput
limit of 250 MiB/s in most cases. However, the throughput limit can vary
depending on volume size. For more information, see [Amazon EBS volume\
types](../../../ec2/latest/userguide/ebsvolumetypes.md) in the _Amazon EC2 User Guide._

**Network bandwidth**

The network speed relative to other DB instance classes.

## Hardware specifications for the general-purpose instance classes

The following tables show the compute, memory, storage, and bandwidth specifications for
the general-purpose instance classes.

**db.m8g – general-purpose instance classes powered by AWS**
**Graviton4 processors**

Instance classvCPUECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.m8g.48xlarge192—768EBS-optimized only40,00050db.m8g.24xlarge96—384EBS-optimized only30,00040db.m8g.16xlarge64—256EBS-optimized only20,00030db.m8g.12xlarge48—192EBS-optimized only15,00022.5db.m8g.8xlarge32—128EBS-optimized only10,00015db.m8g.4xlarge\*16—64EBS-optimized onlyUp to 10,000Up to 15db.m8g.2xlarge\*8—32EBS-optimized onlyUp to 10,000Up to 15db.m8g.xlarge\*4—16EBS-optimized onlyUp to 10,000Up to 12.5db.m8g.large\*2—8EBS-optimized onlyUp to 10,000Up to 12.5

**db.m7i – general-purpose instance classes powered by 4th**
**generation Intel Xeon Scalable processors**

Instance classvCPUProcessor coresSocketsECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.m7i.metal-48xl192962—768EBS-optimized only40,00050db.m7i.metal-24xl96481—384EBS-optimized only30,00037.5db.m7i.48xlarge192———768EBS-optimized only40,00050db.m7i.24xlarge96———384EBS-optimized only30,00037.5db.m7i.16xlarge64———256EBS-optimized only20,00025db.m7i.12xlarge48———192EBS-optimized only15,00018.75db.m7i.8xlarge32———128EBS-optimized only10,00012.5db.m7i.4xlarge16———64EBS-optimized onlyUp to 10,000Up to 12.5db.m7i.2xlarge8———32EBS-optimized onlyUp to 10,000Up to 12.5db.m7i.xlarge4———16EBS-optimized onlyUp to 10,000Up to 12.5db.m7i.large2———8EBS-optimized onlyUp to 10,000Up to 12.5

**db.m7g – general-purpose instance classes powered by AWS**
**Graviton3 processors**

Instance classvCPUECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.m7g.16xlarge64—256EBS-optimized only20,00030db.m7g.12xlarge48—192EBS-optimized only15,00022.5db.m7g.8xlarge32—128EBS-optimized only10,00015db.m7g.4xlarge16—64EBS-optimized onlyUp to 10,000Up to 15db.m7g.2xlarge\*8—32EBS-optimized onlyUp to 10,000Up to 15db.m7g.xlarge\*4—16EBS-optimized onlyUp to 10,000Up to 12.5db.m7g.large\*2—8EBS-optimized onlyUp to 10,000Up to 12.5

**db.m6g – general-purpose instance classes powered by AWS**
**Graviton2 processors**

Instance classvCPUECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.m6g.16xlarge64—256EBS-optimized only19,00025db.m6g.12xlarge48—192EBS-optimized only14,25020db.m6g.8xlarge32—128EBS-optimized only9,50012db.m6g.4xlarge16—64EBS-optimized only4,750Up to 10db.m6g.2xlarge\*8—32EBS-optimized onlyUp to 4,750Up to 10db.m6g.xlarge\*4—16EBS-optimized onlyUp to 4,750Up to 10db.m6g.large\*2—8EBS-optimized onlyUp to 4,750Up to 10

**db.m6gd – general-purpose instance classes powered by AWS**
**Graviton2 processors and SSD storage**

Instance classvCPUECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.m6gd.16xlarge64—2562 x 1900 NVMe SSD19,00025db.m6gd.12xlarge48—1922 x 1425 NVMe SSD14,25020db.m6gd.8xlarge32—1281 x 1900 NVMe SSD9,50012db.m6gd.4xlarge\*16—641 x 950 NVMe SSD4,750Up to 10db.m6gd.2xlarge\*8—321 x 474 NVMe SSDUp to 4,750Up to 10db.m6gd.xlarge\*4—161 x 237 NVMe SSDUp to 4,750Up to 10db.m6gd.large\*2—81 x 118 NVMe SSDUp to 4,750Up to 10

**db.m6id – general-purpose instance classes powered by 3rd**
**generation Intel Xeon Scalable processors and SSD storage**

Instance classvCPUPhysical coresSocketsECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.m6id.metal128642—5124 x 1900 NVMe SSD40,00050db.m6id.32xlarge128———5124 x 1900 NVMe SSD40,00050db.m6id.24xlarge96———3844 x 1425 NVMe SSD30,00037.5db.m6id.16xlarge64———2562 x 1900 NVMe SSD20,00025db.m6id.12xlarge48———1922 x 1425 NVMe SSD15,00018.75db.m6id.8xlarge32———1281 x 1900 NVMe SSD10,00012.5db.m6id.4xlarge\*16———641 x 950 NVMe SSDUp to 10,000Up to 12.5db.m6id.2xlarge\*8———321 x 474 NVMe SSDUp to 10,000Up to 12.5db.m6id.xlarge\*4———161 x 237 NVMe SSDUp to 10,000Up to 12.5db.m6id.large\*2———81 x 118 NVMe SSDUp to 10,000Up to 12.5

**db.m6idn – general-purpose instance classes with 3rd**
**Generation Intel Xeon Scalable processors, SSD storage, and network**
**optimization**

Instance classvCPUECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.m6idn.32xlarge128—5124 x 1900 NVMe SSD100,000200db.m6idn.24xlarge96—3844 x 1425 NVMe SSD75,000150db.m6idn.16xlarge64—2562 x 1900 NVMe SSD50,000100db.m6idn.12xlarge48—1922 x 1425 NVMe SSD37,50075db.m6idn.8xlarge32—1281 x 1900 NVMe SSD25,00050db.m6idn.4xlarge\*16—641 x 950 NVMe SSDUp to 25,000Up to 50db.m6idn.2xlarge\*8—321 x 474 NVMe SSDUp to 25,000Up to 40db.m6idn.xlarge\*4—161 x 237 NVMe SSDUp to 25,000Up to 30db.m6idn.large\*2—81 x 118 NVMe SSDUp to 25,000Up to 25

**db.m6in – general-purpose instance classes powered by 3rd**
**generation Intel Xeon Scalable processors and network optimization**

Instance classvCPUProcessor coresSocketsECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.m6in.metal128642—512EBS-optimized only100,000200db.m6in.32xlarge128—512EBS-optimized only100,000200db.m6in.24xlarge96—384EBS-optimized only75,000150db.m6in.16xlarge64—256EBS-optimized only50,000100db.m6in.12xlarge48—192EBS-optimized only37,50075db.m6in.8xlarge32—128EBS-optimized only25,00050db.m6in.4xlarge\*16—64EBS-optimized onlyUp to 25,000Up to 50db.m6in.2xlarge\*8—32EBS-optimized onlyUp to 25,000Up to 40db.m6in.xlarge\*4—16EBS-optimized onlyUp to 25,000Up to 30db.m6in.large\*2—8EBS-optimized onlyUp to 25,000Up to 25

**db.m6i – general-purpose instance classes powered by 3rd**
**generation Intel Xeon Scalable processors**

Instance classvCPUProcessor coresSocketsECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.m6i.metal128642—512EBS-optimized only40,00050db.m6i.32xlarge128—512EBS-optimized only40,00050db.m6i.24xlarge96—384EBS-optimized only30,00037.5db.m6i.16xlarge64—256EBS-optimized only20,00025db.m6i.12xlarge48—192EBS-optimized only15,00018.75db.m6i.8xlarge32—128EBS-optimized only10,00012.5db.m6i.4xlarge\*16—64EBS-optimized onlyUp to 10,000Up to 12.5db.m6i.2xlarge\*8—32EBS-optimized onlyUp to 10,000Up to 12.5db.m6i.xlarge\*4—16EBS-optimized onlyUp to 10,000Up to 12.5db.m6i.large\*2—8EBS-optimized onlyUp to 10,000Up to 12.5

**db.m5d – general-purpose instance classes powered by Intel**
**Xeon Platinum processors and SSD storage**

Instance classvCPUECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.m5d.24xlarge963453844 x 900 NVMe SSD19,00025db.m5d.16xlarge642622564 x 600 NVMe SSD13,60020db.m5d.12xlarge481731922 x 900 NVMe SSD9,50012db.m5d.8xlarge321311282 x 600 NVMe SSD6,80010db.m5d.4xlarge1661642 x 300 NVMe SSD4,750Up to 10db.m5d.2xlarge\*831321 x 300 NVMe SSDUp to 4,750Up to 10db.m5d.xlarge\*415161 x 150 NVMe SSDUp to 4,750Up to 10db.m5d.large\*21081 x 75 NVMe SSDUp to 4,750Up to 10

**db.m5 – general-purpose instance classes with Intel Xeon**
**Platinum processors**

Instance classvCPUECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.m5.24xlarge96345384EBS-optimized only19,00025db.m5.16xlarge64262256EBS-optimized only13,60020db.m5.12xlarge48173192EBS-optimized only9,50012db.m5.8xlarge32131128EBS-optimized only6,80010db.m5.4xlarge166164EBS-optimized only4,750Up to 10db.m5.2xlarge\*83132EBS-optimized onlyUp to 4,750Up to 10db.m5.xlarge\*41516EBS-optimized onlyUp to 4,750Up to 10db.m5.large\*2108EBS-optimized onlyUp to 4,750Up to 10

**db.m4 – general-purpose instance classes with Intel Xeon**
**Scalable processors**

Instance classvCPUECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.m4.16xlarge64188256EBS-optimized only10,00025db.m4.10xlarge40124.5160EBS-optimized only4,00010db.m4.4xlarge1653.564EBS-optimized only2,000Highdb.m4.2xlarge825.532EBS-optimized only1,000Highdb.m4.xlarge41316EBS-optimized only750Highdb.m4.large26.58EBS-optimized only450Moderate

**db.m3 – general-purpose instance classes**

Instance classvCPUECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.m3.2xlarge82630EBS-optimized only1,000Highdb.m3.xlarge41315EBS-optimized only500Highdb.m3.large26.57.5EBS only—Moderatedb.m3.medium133.75EBS only—Moderate

\\* These DB instance classes can support maximum performance for 30 minutes
at least once every 24 hours. For more information on baseline performance of the underlying
EC2 instance types, see [Amazon EBS-optimized\
instances](../../../ec2/latest/userguide/ebsoptimized.md) in the _Amazon EC2 User Guide._

## Hardware specifications for the memory-optimized instance classes

The following tables show the compute, memory, storage, and bandwidth specifications
for the memory-optimized instance classes.

**db.z1d – memory-optimized instance**
**classes**

Instance classvCPUECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.z1d.12xlarge482713842 x 900 NVMe SSD19,00025db.z1d.6xlarge241341921 x 900 NVMe SSD9,50012db.z1d.3xlarge1275961 x 450 NVMe SSD4,750Up to 10db.z1d.2xlarge85364

1 x 300 NVMe SSD

3,170Up to 10db.z1d.xlarge\*428321 x 150 NVMe SSDUp to 3,170Up to 10db.z1d.large\*215161 x 75 NVMe SSDUp to 3,170Up to 10

**db.x2g – memory-optimized instance classes with AWS**
**Graviton2 processors**

Instance classvCPUECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.x2g.16xlarge64—1024EBS-optimized only19,00025db.x2g.12xlarge48—768EBS-optimized only14,25020db.x2g.8xlarge32—512EBS-optimized only9,50012db.x2g.4xlarge16—256EBS-optimized only4,750Up to 10db.x2g.2xlarge8—128EBS-optimized onlyUp to 4,750Up to 10db.x2g.xlarge4—64EBS-optimized onlyUp to 4,750Up to 10db.x2g.large2—32EBS-optimized onlyUp to 4,750Up to 10

**db.x2idn – memory-optimized instance classes with 3rd**
**generation Intel Xeon Scalable processors**

Instance classvCPUProcessor coresSocketsECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.x2idn.metal128642—2,0482 x 1900 NVMe SSD80,000100db.x2idn.32xlarge128———2,0482 x 1900 NVMe SSD80,000100db.x2idn.24xlarge96———1,5362 x 1425 NVMe SSD60,00075db.x2idn.16xlarge64———1,0241 x 1900 NVMe SSD40,00050

**db.x2iedn – memory-optimized instance classes with local**
**NVMe-based SSDs, with 3rd generation Intel Xeon Scalable**
**processors**

Instance classvCPUProcessor coresSocketsECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.x2iedn.metal128642—4,0962 x 1900 NVMe SSD80,000100db.x2iedn.32xlarge128———4,0962 x 1900 NVMe SSD80,000100db.x2iedn.24xlarge96———3,0722 x 1425 NVMe SSD60,00075db.x2iedn.16xlarge64———2,0481 x 1900 NVMe SSD40,00050db.x2iedn.8xlarge32———1,0241 x 950 NVMe SSD20,00025db.x2iedn.4xlarge16———5121 x 475 NVMe SSDUp to 20,000Up to 25db.x2iedn.2xlarge8———2561 x 237 NVMe SSDUp to 20,000Up to 25db.x2iedn.xlarge4———1281 x 118 NVMe SSDUp to 20,000Up to 25

**db.x2iezn – memory-optimized instance classes with 2nd**
**generation Intel Xeon Scalable processors**

Instance classvCPUProcessor coresSocketsECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.x2iezn.metal48242—1,536EBS-optimized only19,000100db.x2iezn.12xlarge>48———1,536EBS-optimized only19,000100db.x2iezn.8xlarge32———1,024EBS-optimized only12,00075db.x2iezn.6xlarge24———768EBS-optimized onlyUp to 9,50050db.x2iezn.4xlarge16———512EBS-optimized onlyUp to 4,750Up to 25db.x2iezn.2xlarge8———256EBS-optimized onlyUp to 3,170Up to 25

**db.x1e – memory-optimized instance**
**classes**

Instance classvCPUECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.x1e.32xlarge1283403,904EBS-optimized only14,00025db.x1e.16xlarge641791,952EBS-optimized only7,00010db.x1e.8xlarge3291976EBS-optimized only3,500Up to 10db.x1e.4xlarge1647488EBS-optimized only1,750Up to 10db.x1e.2xlarge823244EBS-optimized only1,000Up to 10db.x1e.xlarge412122EBS-optimized only500Up to 10

**db.x1 – memory-optimized instance**
**classes**

Instance classvCPUECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.x1.32xlarge1283491,952EBS-optimized only14,00025db.x1.16xlarge64174.5976EBS-optimized only7,00010

**db.m8gd – memory-optimized instance classes powered by AWS**
**Graviton4 processors and SSD storage**

Instance classvCPUECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.m8gd.48xlarge192—7686 x 1900 NVMe SSD40,00050db.m8gd.24xlarge96—3843 x 1900 NVMe SSD30,00040db.m8gd.16xlarge64—2562 x 1900 NVMe SSD20,00030db.m8gd.12xlarge48—1923 x 950 NVMe SSD15,00022.5db.m8gd.8xlarge32—1281 x 1900 NVMe SSD10,00015db.m8gd.4xlarge16—641 x 950 NVMe SSDUp to 10,000Up to 15db.m8gd.2xlarge8—321 x 474 NVMe SSDUp to 10,000Up to 15db.m8gd.xlarge4—161 x 237 NVMe SSDUp to 10,000Up to 12.5db.m8gd.large2—81 x 118 NVMe SSDUp to 10,000Up to 12.5

**db.r8gd – memory-optimized instance classes with AWS**
**Graviton4 processors**

Instance classvCPUECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.r8gd.48xlarge192—15366 x 1900 NVMe SSD40,00050db.r8gd.24xlarge96—7683 x 1900 NVMe SSD30,00040db.r8gd.16xlarge64—5122 x 1900 NVMe SSD20,00030db.r8gd.12xlarge48—3843 x 950 NVMe SSD15,00022.5db.r8gd.8xlarge32—2561 x 1900 NVMe SSD10,00015db.r8gd.4xlarge16—1281 x 950 NVMe SSDUp to 10,000Up to 15db.r8gd.2xlarge8—641 x 474 NVMe SSDUp to 10,000Up to 15db.r8gd.xlarge4—321 x 237 NVMe SSDUp to 10,000Up to 12.5db.r8gd.large2—161 x 118 NVMe SSDUp to 10,000Up to 12.5

**db.r8g – memory-optimized instance classes with AWS**
**Graviton4 processors**

Instance classvCPUECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.r8g.48xlarge192—1536EBS-optimized only40,00050db.r8g.24xlarge96—768EBS-optimized only30,00040db.r8g.16xlarge64—512EBS-optimized only20,00030db.r8g.12xlarge48—384EBS-optimized only15,00022.5db.r8g.8xlarge32—256EBS-optimized only10,00015db.r8g.4xlarge\*16—128EBS-optimized onlyUp to 10,000Up to 15db.r8g.2xlarge\*8—64EBS-optimized onlyUp to 10,000Up to 15db.r8g.xlarge\*4—32EBS-optimized onlyUp to 10,000Up to 12.5db.r8g.large\*2—16EBS-optimized onlyUp to 10,000Up to 12.5

**db.r7i – memory-optimized instance classes powered by 4th**
**generation Intel Xeon Scalable processors**

Instance classvCPUProcessor coresSocketsECUMemory (GiB)Normalized unitsInstance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.r7i.metal-48xl192962—1536192EBS-optimized only40,00050db.r7i.metal-24xl96481—76896EBS-optimized only30,00037.5db.r7i.48xlarge192———1536192EBS-optimized only40,00050db.r7i.24xlarge96———76896EBS-optimized only30,00037.5db.r7i.16xlarge64———51264EBS-optimized only20,00025db.r7i.12xlarge48———38448EBS-optimized only15,00018.75db.r7i.8xlarge.tpc2.mem3x32———76896EBS-optimized only30,00012.5db.r7i.8xlarge.tpc2.mem2x32———51264EBS-optimized only20,00012.5db.r7i.8xlarge32———25632EBS-optimized only10,00012.5db.r7i.6xlarge.tpc2.mem4x24———76896EBS-optimized only30,000Up to 12.5db.r7i.6xlarge.tpc2.mem2x24———38448EBS-optimized only15,000Up to 12.5db.r7i.4xlarge.tpc2.mem4x16———51264EBS-optimized only20,000Up to 12.5db.r7i.4xlarge.tpc2.mem3x16———38448EBS-optimized only15,000Up to 12.5db.r7i.4xlarge.tpc2.mem2x16———25632EBS-optimized only10,000Up to 12.5db.r7i.4xlarge16———12816EBS-optimized onlyUp to 10,000Up to 12.5db.r7i.3xlarge.tpc2.mem4x12———38448EBS-optimized only15,000Up to 12.5db.r7i.2xlarge.tpc2.mem8x8———51264EBS-optimized only20,000Up to 12.5db.r7i.2xlarge.tpc2.mem4x8———25632EBS-optimized only10,000Up to 12.5db.r7i.2xlarge8———648EBS-optimized onlyUp to 10,000Up to 12.5db.r7i.xlarge.tpc2.mem4x4———12816EBS-optimized onlyUp to 10,000Up to 12.5db.r7i.xlarge.tpc2.mem2x4———648EBS-optimized onlyUp to 10,000Up to 12.5db.r7i.xlarge4———324EBS-optimized onlyUp to 10,000Up to 12.5db.r7i.large2———162EBS-optimized onlyUp to 10,000Up to 12.5

**db.r7g – memory-optimized instance classes with AWS**
**Graviton3 processors**

Instance classvCPUECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.r7g.16xlarge64—512EBS-optimized only20,00030db.r7g.12xlarge48—384EBS-optimized only15,00022.5db.r7g.8xlarge32—256EBS-optimized only10,00015db.r7g.4xlarge16—128EBS-optimized onlyUp to 10,000Up to 15db.r7g.2xlarge\*8—64EBS-optimized onlyUp to 10,000Up to 15db.r7g.xlarge\*4—32EBS-optimized onlyUp to 10,000Up to 12.5db.r7g.large\*2—16EBS-optimized onlyUp to 10,000Up to 12.5

**db.r6g – memory-optimized instance classes with AWS**
**Graviton2 processors**

Instance classvCPUECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.r6g.16xlarge64—512EBS-optimized only19,00025db.r6g.12xlarge48—384EBS-optimized only14,25020db.r6g.8xlarge32—256EBS-optimized only9,50012db.r6g.4xlarge16—128EBS-optimized only4,750Up to 10 db.r6g.2xlarge\*8—64EBS-optimized onlyUp to 4,750Up to 10 db.r6g.xlarge\*4—32EBS-optimized onlyUp to 4,750Up to 10 db.r6g.large\*2—16EBS-optimized onlyUp to 4,750Up to 10

**db.r6gd – memory-optimized instance classes with AWS**
**Graviton2 processors and SSD storage**

Instance classvCPUECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.r6gd.16xlarge64—5122 x 1900 NVMe SSD19,00025db.r6gd.12xlarge48—3842 x 1425 NVMe SSD14,25020db.r6gd.8xlarge32—2561 x 1900 NVMe SSD9,50012db.r6gd.4xlarge16—1281 x 950 NVMe SSD4,750Up to 10 db.r6gd.2xlarge8—641 x 474 NVMe SSDUp to 4,750Up to 10 db.r6gd.xlarge4—321 x 237 NVMe SSDUp to 4,750Up to 10 db.r6gd.large2—161 x 118 NVMe SSDUp to 4,750Up to 10

**db.r6id – memory-optimized instance classes with 3rd**
**generation Intel Xeon Scalable processors and SSD storage**

Instance classvCPUProcessor coresSocketsECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.r6id.metal128642—1,0244 x 1900 NVMe SSD40,00050db.r6id.32xlarge128———1,0244x1900 NVMe SSD40,00050db.r6id.24xlarge96———7684x1425 NVMe SSD30,00037.5db.r6id.16xlarge64———5122x1900 NVMe SSD20,00025db.r6id.12xlarge48———3842x1425 NVMe SSD15,00018.75db.r6id.8xlarge32———2561x1900 NVMe SSD10,00012.5db.r6id.4xlarge\*16———1281x950 NVMe SSDUp to 10,000Up to 12.5db.r6id.2xlarge\*8———641x474 NVMe SSDUp to 10,000Up to 12.5db.r6id.xlarge\*4———321x237 NVMe SSDUp to 10,000Up to 12.5db.r6id.large\*2———161x118 NVMe SSDUp to 10,000Up to 12.5

**db.r6idn – memory-optimized instance classes with 3rd**
**generation Intel Xeon Scalable processors, SSD storage, and network**
**optimization**

Instance classvCPUECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.r6idn.32xlarge128—1,0244x1900 NVMe SSD100,000200db.r6idn.24xlarge96—7684x1425 NVMe SSD75,000150db.r6idn.16xlarge64—5122x1900 NVMe SSD50,000100db.r6idn.12xlarge48—3842x1425 NVMe SSD37,50075db.r6idn.8xlarge32—2561x1900 NVMe SSD25,00050db.r6idn.4xlarge\*16—1281x950 NVMe SSDUp to 25,000Up to 50db.r6idn.2xlarge\*8—641x474 NVMe SSDUp to 25,000Up to 40db.r6idn.xlarge\*4—321x237 NVMe SSDUp to 25,000Up to 30db.r6idn.large\*2—161x118 NVMe SSDUp to 25,000Up to 25

**db.r6in – memory-optimized instance classes with 3rd**
**generation Intel Xeon Scalable processors and network**
**optimization**

Instance classvCPUProcessor coresSocketsECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.r6in.metal128642—1,024EBS-optimized only100,000200db.r6in.32xlarge128———1,024EBS-optimized only100,000200db.r6in.24xlarge96———768EBS-optimized only75,000150db.r6in.16xlarge64———512EBS-optimized only50,000100db.r6in.12xlarge48———384EBS-optimized only37,50075db.r6in.8xlarge32———256EBS-optimized only25,00050db.r6in.4xlarge\*16———128EBS-optimized onlyUp to 25,000Up to 50db.r6in.2xlarge\*8———64EBS-optimized onlyUp to 25,000Up to 40db.r6in.xlarge\*4———32EBS-optimized onlyUp to 25,000Up to 30db.r6in.large\*2———16EBS-optimized onlyUp to 25,000Up to 25

**db.r6i – Oracle memory-optimized instance classes**
**preconfigured for high memory, storage, and I/O**

Instance classvCPUECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.r6i.8xlarge.tpc2.mem4x32—1024EBS-optimized only40,00050db.r6i.8xlarge.tpc2.mem3x32—768EBS-optimized only30,00037.5db.r6i.6xlarge.tpc2.mem4x24—768EBS-optimized only30,00037.5db.r6i.4xlarge.tpc2.mem4x16—512EBS-optimized only20,00025db.r6i.4xlarge.tpc2.mem3x16—384EBS-optimized only15,00018.75db.r6i.4xlarge.tpc2.mem2x16—256EBS-optimized only10,00012.5db.r6i.2xlarge.tpc2.mem8x8—512EBS-optimized only20,00025db.r6i.2xlarge.tpc2.mem4x8—256EBS-optimized only10,00012.5db.r6i.2xlarge.tpc1.mem2x8—128EBS-optimized onlyUp to 10,00012.5db.r6i.xlarge.tpc2.mem4x4—128EBS-optimized onlyUp to 10,00012.5db.r6i.xlarge.tpc2.mem2x4—64EBS-optimized onlyUp to 10,00012.5db.r6i.large.tpc1.mem2x2—32EBS-optimized onlyUp to 10,00012.5

**db.r6i – memory-optimized instance classes with 3rd**
**Generation Intel Xeon Scalable processors**

Instance classvCPUProcessor coresSocketsECUMemory (GiB)Normalized unitsInstance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.r6i.metal128642—1,024256EBS-optimized only40,00050db.r6i.32xlarge128———1,024256EBS-optimized only40,00050db.r6i.24xlarge96———768192EBS-optimized only30,00037.5db.r6i.16xlarge64———512128EBS-optimized only20,00025db.r6i.12xlarge48———38496EBS-optimized only15,00018.75db.r6i.8xlarge32———25664EBS-optimized only10,00012.5db.r6i.4xlarge\*16———12832EBS-optimized onlyUp to 10,000Up to 12.5db.r6i.2xlarge\*8———6416EBS-optimized onlyUp to 10,000Up to 12.5db.r6i.xlarge\*4———328EBS-optimized onlyUp to 10,000Up to 12.5db.r6i.large\*2———164EBS-optimized onlyUp to 10,000Up to 12.5

**db.r5d – memory-optimized instance classes with Intel**
**Xeon Platinum processors and SSD storage**

Instance classvCPUECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.r5d.24xlarge963477684 x 900 NVMe SSD19,00025db.r5d.16xlarge642645124 x 600 NVMe SSD13,60020db.r5d.12xlarge481733842 x 900 NVMe SSD9,50012db.r5d.8xlarge321322562 x 600 NVMe SSD6,80010db.r5d.4xlarge16711282 x 300 NVMe SSD4,750Up to 10db.r5d.2xlarge\*838641 x 300 NVMe SSDUp to 4,750Up to 10db.r5d.xlarge\*419321 x 150 NVMe SSDUp to 4,750Up to 10db.r5d.large\*210161 x 75 NVMe SSDUp to 4,750Up to 10

**db.r5b – Oracle memory-optimized instance classes**
**preconfigured for high memory, storage, and I/O**

Instance classvCPUECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.r5b.8xlarge.tpc2.mem3x32—768EBS-optimized only60,00025db.r5b.6xlarge.tpc2.mem4x24—768EBS-optimized only60,00025db.r5b.4xlarge.tpc2.mem4x16—512EBS-optimized only40,00020db.r5b.4xlarge.tpc2.mem3x16—384EBS-optimized only30,00012db.r5b.4xlarge.tpc2.mem2x16—256EBS-optimized only20,00010db.r5b.2xlarge.tpc2.mem8x8—512EBS-optimized only40,00020db.r5b.2xlarge.tpc2.mem4x8—256EBS-optimized only20,00010db.r5b.2xlarge.tpc1.mem2x8—128EBS-optimized only10,000Up to 10db.r5b.xlarge.tpc2.mem4x4—128EBS-optimized only10,000Up to 10db.r5b.xlarge.tpc2.mem2x4—64EBS-optimized onlyUp to 10,000Up to 10db.r5b.large.tpc1.mem2x2—32EBS-optimized onlyUp to 10,000Up to 10

**db.r5b – memory-optimized instance classes with Intel**
**Xeon Platinum processors and EBS optimization**

Instance classvCPUECUMemory (GiB)Normalized unitsInstance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.r5b.24xlarge96347768192EBS-optimized only60,00025db.r5b.16xlarge64264512128EBS-optimized only40,00020db.r5b.12xlarge4817338496EBS-optimized only30,00012db.r5b.8xlarge3213225664EBS-optimized only20,00010db.r5b.4xlarge167112832EBS-optimized only10,000Up to 10db.r5b.2xlarge\*8386416EBS-optimized onlyUp to 10,000Up to 10db.r5b.xlarge\*419328EBS-optimized onlyUp to 10,000Up to 10db.r5b.large\*210164EBS-optimized onlyUp to 10,000Up to 10

**db.r5 – Oracle memory-optimized instance classes**
**preconfigured for high memory, storage, and I/O**

Instance classvCPUECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.r5.12xlarge.tpc2.mem2x48—768EBS-optimized only19,00025db.r5.8xlarge.tpc2.mem3x32—768EBS-optimized only19,00025db.r5.6xlarge.tpc2.mem4x24—768EBS-optimized only19,00025db.r5.4xlarge.tpc2.mem4x16—512EBS-optimized only13,60020db.r5.4xlarge.tpc2.mem3x16—384EBS-optimized only9,50012db.r5.4xlarge.tpc2.mem2x16—256EBS-optimized only6,80010db.r5.2xlarge.tpc2.mem8x8—512EBS-optimized only13,60020db.r5.2xlarge.tpc2.mem4x8—256EBS-optimized only6,80010db.r5.2xlarge.tpc1.mem2x8—128EBS-optimized only4,750Up to 10db.r5.xlarge.tpc2.mem4x4—128EBS-optimized only4,750Up to 10db.r5.xlarge.tpc2.mem2x4—64EBS-optimized onlyUp to 4,750Up to 10db.r5.large.tpc1.mem2x2—32EBS-optimized onlyUp to 4,750Up to 10

**db.r5 – memory-optimized instance**
**classes**

Instance classvCPUECUMemory (GiB)Normalized unitsInstance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.r5.24xlarge96347768192EBS-optimized only19,00025db.r5.16xlarge64264512128EBS-optimized only13,60020db.r5.12xlarge4817338496EBS-optimized only9,50012db.r5.8xlarge3213225664EBS-optimized only6,80010db.r5.4xlarge167112832EBS-optimized only4,750Up to 10db.r5.2xlarge\*8386416EBS-optimized onlyUp to 4,750Up to 10db.r5.xlarge\*419328EBS-optimized onlyUp to 4,750Up to 10db.r5.large\*210164EBS-optimized onlyUp to 4,750Up to 10

**db.r4 – memory-optimized instance classes with Intel Xeon**
**Scalable processors**

Instance classvCPUECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.r4.16xlarge64195488EBS-optimized only14,00025db.r4.8xlarge3299244EBS-optimized only7,00010db.r4.4xlarge1653122EBS-optimized only3,500Up to 10db.r4.2xlarge82761EBS-optimized only1,700Up to 10db.r4.xlarge413.530.5EBS-optimized only850Up to 10db.r4.large2715.25EBS-optimized only425Up to 10

**db.r3 – memory-optimized instance**
**classes**

Instance classvCPUECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.r3.8xlarge\*\*32104244EBS only—10db.r3.4xlarge1652122EBS-optimized only2,000Highdb.r3.2xlarge82661EBS-optimized only1,000Highdb.r3.xlarge41330.5EBS-optimized only500Moderatedb.r3.large26.515.25EBS-optimized only—Moderate

\\* These DB instance classes can support maximum performance for 30
minutes at least once every 24 hours. For more information on baseline performance
of the underlying EC2 instance types, see [Amazon EBS-optimized instances](../../../ec2/latest/userguide/ebsoptimized.md) in
the _Amazon EC2 User Guide._

\\*\\* The r3.8xlarge DB instance class doesn't have dedicated EBS
bandwidth and therefore doesn't offer EBS optimization. For this instance class,
network traffic and Amazon EBS traffic share the same 10-gigabit network
interface.

## Hardware specifications for the compute-optimized instance classes

The following tables show the compute, memory, storage, and bandwidth specifications for
the compute-optimized instance classes.

**db.c6gd – compute-optimized instance classes (for Multi-AZ DB cluster**
**deployments only)**

Instance classvCPUECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.c6gd.16xlarge64—1282 x 1900 NVMe SSD19,00025db.c6gd.12xlarge48—962 x 1425 NVMe SSD14,25020db.c6gd.8xlarge32—641 x 1900 NVMe SSD9,50012db.c6gd.4xlarge16—321 x 950 NVMe SSD4,750Up to 10db.c6gd.2xlarge8—161 x 474 NVMe SSDUp to 4,750Up to 10db.c6gd.xlarge4—81 x 237 NVMe SSDUp to 4,750Up to 10db.c6gd.large2—41 x 118 NVMe SSDUp to 4,750Up to 10db.c6gd.medium1—21 x 59 NVMe SSDUp to 4,750Up to 10

## Hardware specifications for the burstable-performance instance classes

The following tables show the compute, memory, storage, and bandwidth specifications for
the burstable-performance instance classes.

**db.t4g – burstable-performance instance classes powered by**
**AWS Graviton2 processors**

Instance classvCPUECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.t4g.2xlarge\*8—32EBS-optimized onlyUp to 2,780Up to 5db.t4g.xlarge\*4—16EBS-optimized onlyUp to 2,780Up to 5db.t4g.large\*2—8EBS-optimized onlyUp to 2,780Up to 5db.t4g.medium\*2—4EBS-optimized onlyUp to 2,085Up to 5db.t4g.small\*2—2EBS-optimized onlyUp to 2,085Up to 5db.t4g.micro\*2—1EBS-optimized onlyUp to 2,085Up to 5

**db.t3 – burstable-performance instance**
**classes**

Instance classvCPUECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.t3.2xlarge\*8Variable32EBS-optimized onlyUp to 2,780Up to 5db.t3.xlarge\*4Variable16EBS-optimized onlyUp to 2,780Up to 5db.t3.large\*2Variable8EBS-optimized onlyUp to 2,780Up to 5db.t3.medium\*2Variable4EBS-optimized onlyUp to 2,085Up to 5db.t3.small\*2Variable2EBS-optimized onlyUp to 2,085Up to 5db.t3.micro\*2Variable1EBS-optimized onlyUp to 2,085Up to 5

**db.t2 – burstable-performance instance**
**classes**

Instance classvCPUECUMemory (GiB)Instance storage (GiB)Max. EBS bandwidth (Mbps)Network bandwidth (Gbps)db.t2.2xlarge8Variable32EBS only—Moderatedb.t2.xlarge4Variable16EBS only—Moderatedb.t2.large2Variable8EBS only—Moderatedb.t2.medium2Variable4EBS only—Moderatedb.t2.small1Variable2EBS only—Lowdb.t2.micro1Variable1EBS only—Low

\\* These DB instance classes can support maximum performance for 30 minutes
at least once every 24 hours. For more information on baseline performance of the underlying
EC2 instance types, see [Amazon EBS-optimized\
instances](../../../ec2/latest/userguide/ebsoptimized.md) in the _Amazon EC2 User Guide._

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Configuring the processor for RDS for Oracle

DB instance storage

All content copied from https://docs.aws.amazon.com/.
