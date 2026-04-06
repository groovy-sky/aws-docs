# Amazon EBS volume types

Amazon EBS provides the following volume types, which differ in performance characteristics and
price, so that you can tailor your storage performance and cost to the needs of your
applications.

###### Important

There are several factors that can affect the performance of EBS volumes, such as instance
configuration, I/O characteristics, and workload demand. To fully use the IOPS provisioned on
an EBS volume, use [EBS–optimized instances](../../../ec2/latest/userguide/ebs-optimized.md).
For more information about getting the most out of your EBS volumes, see [Amazon EBS volume performance](https://docs.aws.amazon.com/ebs/latest/userguide/ebs-performance.html).

For more information about pricing, see [Amazon EBS\
Pricing](https://aws.amazon.com/ebs/pricing).

**Volume types**

- [Solid state drive (SSD) volumes](#vol-type-ssd)

- [Hard disk drive (HDD) volumes](#vol-type-hdd)

- [Previous generation volumes](#vol-type-prev)

## Solid state drive (SSD) volumes

SSD-backed volumes are optimized for transactional workloads involving frequent read/write
operations with small I/O size, where the dominant performance attribute is IOPS. SSD-backed
volume types include **General Purpose SSD** and **Provisioned IOPS SSD**. The following is a summary of the use cases and characteristics of SSD-backed
volumes.

[Amazon EBS General Purpose SSD volumes](https://docs.aws.amazon.com/ebs/latest/userguide/general-purpose.html)[Amazon EBS Provisioned IOPS SSD volumes](https://docs.aws.amazon.com/ebs/latest/userguide/provisioned-iops.html)**Volume type**`gp3` 6`gp2``io2` Block Express`io1`**Durability**99.8% - 99.9% durability (0.1% - 0.2% annual failure rate)99.999% durability (0.001% annual failure rate)99.8% - 99.9% durability (0.1% - 0.2% annual failure rate)**Use cases**

- Transactional workloads

- Virtual desktops

- Medium-sized, single-instance databases

- Low-latency interactive applications

- Boot volumes

- Development and test environments

Workloads that require:

- Consistent sub-millisecond latency with average latency under 500 microseconds 5

- Sustained IOPS performance

- More than 80,000 IOPS or 2,000 MiB/s of throughput

- Workloads that require sustained IOPS performance or more than
16,000 IOPS

- I/O-intensive database workloads

**Volume size**1 GiB - 64 TiB 1 GiB - 16 TiB 4 GiB - 64 TiB 4 GiB - 16 TiB **Max IOPS**80,000 3 (64 KiB I/O 4)16,000 (16 KiB I/O 4)256,000 3 (16 KiB I/O 4) 64,000 (16 KiB I/O 4)**Max throughput**2,000 MiB/s250 MiB/s 14,000 MiB/s1,000 MiB/s 2**Amazon EBS Multi-attach**Not supportedSupported**NVMe reservations**Not supportedSupportedNot supported**Boot volume**Supported

1 The throughput limit is between 128 MiB/s and 250
MiB/s, depending on the volume size. For more information, see [gp2 volume performance](https://docs.aws.amazon.com/ebs/latest/userguide/general-purpose.html#gp2-performance). Volumes created before **December 3,**
**2018** that have not been modified since creation might not reach full performance
unless you [modify the volume](https://docs.aws.amazon.com/ebs/latest/userguide/ebs-modify-volume.html).

2 To achieve maximum throughput of 1,000 MiB/s, the volume must
be provisioned with 64,000 IOPS and it must be attached to a [Nitro-based instance](https://docs.aws.amazon.com/ec2/latest/instancetypes/ec2-nitro-instances.html). Volumes created before **December**
**6, 2017** that have not been modified since creation might not reach full performance
unless you [modify the volume](https://docs.aws.amazon.com/ebs/latest/userguide/ebs-modify-volume.html).

3 [Nitro-based instances](https://docs.aws.amazon.com/ec2/latest/instancetypes/ec2-nitro-instances.html) support volumes provisioned with up to 256,000 IOPS. Other instance
types can be attached to volumes provisioned with up to 64,000 IOPS, but can achieve up to 32,000
IOPS.

4 Represents the required I/O size to reach maximum IOPS within
the volume's throughput limit.

5 `io2` Block Express volumes are designed to deliver an average
latency of under 500 microseconds for 16KiB I/O operations.

6 On Outposts, gp3 volumes support sizes up to 16 TiB, IOPS up to
16,000, and throughput up to 1,000 MiB/s.

For more information about the SSD-backed volume types, see the following:

- [Amazon EBS General Purpose SSD volumes](https://docs.aws.amazon.com/ebs/latest/userguide/general-purpose.html)

- [Amazon EBS Provisioned IOPS SSD volumes](https://docs.aws.amazon.com/ebs/latest/userguide/provisioned-iops.html)

## Hard disk drive (HDD) volumes

HDD-backed volumes are optimized for large streaming workloads where the dominant
performance attribute is throughput. HDD volume types include **Throughput Optimized HDD** and **Cold HDD**. The following
is a summary of the use cases and characteristics of HDD-backed volumes.

[Throughput Optimized HDD volumes](https://docs.aws.amazon.com/ebs/latest/userguide/hdd-vols.html#EBSVolumeTypes_st1)[Cold HDD volumes](https://docs.aws.amazon.com/ebs/latest/userguide/hdd-vols.html#EBSVolumeTypes_sc1)**Volume type**`st1``sc1`**Durability**99.8% - 99.9% durability (0.1% - 0.2% annual failure rate)**Use cases**

- Big data

- Data warehouses

- Log processing

- Throughput-oriented storage for data that is infrequently accessed

- Scenarios where the lowest storage cost is important

**Volume size**125 GiB - 16 TiB**Max IOPS per volume** (1 MiB I/O)500250**Max throughput per volume**500 MiB/s250 MiB/s**Amazon EBS Multi-attach**Not supported**Boot volume**Not supported

For more information about the Hard disk drives (HDD) volumes, see [Amazon EBS Throughput Optimized HDD and Cold HDD volumes](https://docs.aws.amazon.com/ebs/latest/userguide/hdd-vols.html).

## Previous generation volumes

Magnetic ( `standard`) volumes are previous generation volumes that are backed by magnetic
drives. They are suited for workloads with small datasets where data is accessed infrequently
and performance is not of primary importance. These volumes deliver approximately 100 IOPS on
average, with burst capability of up to hundreds of IOPS, and they can range in size from 1 GiB
to 1 TiB.

###### Tip

Magnetic is a previous generation volume type. If you need higher performance or
performance consistency than previous-generation volumes can provide, we recommend
using one of the current generation volume types.

The following table describes previous-generation EBS volume types.

Magnetic**Volume type**`standard`**Use cases**Workloads where data is infrequently accessed**Volume size**1 GiB-1 TiB**Max IOPS per volume**40–200**Max throughput per volume**40–90 MiB/s**Boot volume**Supported

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Features and benefits

General Purpose SSD volumes
