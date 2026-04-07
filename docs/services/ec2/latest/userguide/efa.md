# Elastic Fabric Adapter for AI/ML and HPC workloads on Amazon EC2

An Elastic Fabric Adapter (EFA) is a network device that you can attach to your Amazon EC2 instance to
accelerate Artificial Intelligence (AI), Machine Learning (ML), and High Performance Computing
(HPC) applications. EFA enables you to achieve the application performance of an on-premises
AI/ML or HPC cluster, with the scalability, flexibility, and elasticity provided by the AWS
Cloud.

EFA provides lower and more consistent latency and higher throughput than the TCP transport
traditionally used in cloud-based HPC systems. It enhances the performance of inter-instance
communication that is critical for scaling AI/ML and HPC applications. It is optimized to work
on the existing AWS network infrastructure and it can scale depending on application
requirements.

EFA integrates with Libfabric, and it supports Nvidia Collective
Communications Library (NCCL) and NVIDIA Inference Xfer Library (NIXL) for AI and ML applications,
and Open MPI 4.1 and later and Intel MPI 2019 Update 5 and later for HPC
applications. NCCL and MPI integrate with Libfabric 1.7.0 and later. NIXL integrates
with Libfabric 1.21.0 and later.

EFA supports RDMA (Remote Direct Memory Access) write on most supported instance types that have
Nitro version 4 and later. RDMA read is supported on all instances with Nitro version 4 and later.
For more information, see [Supported instance types](#efa-instance-types).

###### Contents

- [EFA basics](#efa-basics)

- [Supported interfaces and libraries](#efa-mpi)

- [Supported instance types](#efa-instance-types)

- [Supported operating systems](#efa-os)

- [EFA limitations](#efa-limits)

- [EFA pricing](#efa-pricing)

- [Get started with EFA and MPI](efa-start.md)

- [Get started with EFA and NCCL](efa-start-nccl.md)

- [Get started with EFA and NIXL](efa-start-nixl.md)

- [Maximize network bandwidth](efa-acc-inst-types.md)

- [Create and attach an EFA](create-efa.md)

- [Detach and delete an EFA](detach-efa.md)

- [Monitor an EFA](efa-working-monitor.md)

- [Verify the EFA installer](efa-verify.md)

- [Release notes](efa-changelog.md)

## EFA basics

An EFA device can be attached to an EC2 instance in two ways:

1. Using a traditional EFA interface, also called EFA with ENA, which creates both an EFA
    device and an ENA device.

2. Using an EFA-only interface, which creates just the EFA device.

The EFA device provides capabilities like built-in OS-bypass and congestion control through
the Scalable Reliable Datagram (SRD) protocol. The EFA device features enable low-latency,
reliable transport functionality that allows EFA interface to provide better application
performance for HPC and ML applications on Amazon EC2. While the ENA device offers traditional
IP networking.

![Contrasting a traditional HPC software stack with one that uses an EFA.](https://docs.aws.amazon.com/images/AWSEC2/latest/UserGuide/images/efa_stack.png)

Traditionally, AI/ML applications use NCCL and NIXL (for disaggregated inference).
HPC applications use the Message Passing Interface (MPI) to interface with the system's network
transport. In the AWS cloud, this has meant that applications interface with NCCL, NIXL, or MPI,
which then uses the operating system's TCP/IP stack and the ENA device driver to enable network
communication between instances.

With a traditional EFA (EFA with ENA) or EFA-only interface, AI/ML applications use NCCL and
NIXL (for disaggregated inference). HPC applications use MPI, to interface directly
with the Libfabric API. The Libfabric API bypasses the operating system kernel and communicates
directly with the EFA device to put packets on the network. This reduces overhead and enables AI/ML
and HPC applications to run more efficiently.

###### Note

Libfabric is a core component of the OpenFabrics Interfaces (OFI) framework,
which defines and exports the user-space API of OFI. For more information, see
the [Libfabric OpenFabrics](https://ofiwg.github.io/libfabric)
website.

### Differences between ENA, EFA, and EFA-only network interfaces

Amazon EC2 provides two types of network interfaces:

- **ENA** interfaces provide all of the traditional IP
networking and routing features that are required to support IP networking for a VPC.
For more information, see [Enable enhanced networking with ENA on your EC2 instances](enhanced-networking-ena.md).

- **EFA** (EFA with ENA) interfaces provide both the
ENA device for IP networking and the EFA device for low-latency, high-throughput
communication.

- **EFA-only** interfaces support only the EFA device
capabilities, without the ENA device for traditional IP networking.

The following table provides a comparison of ENA, EFA, and EFA-only network interfaces.

ENAEFA (EFA with ENA)EFA-onlySupports IP networking functionalityYesYesNoCan be assigned IPv4 or IPv6 addressesYesYesNoCan be used as primary network interface for instanceYesYesNoCounts towards ENI attachment limit for instanceYesYesYesInstance type supportSupported on all Nitro-based instances types[Supported instance types](#efa-instance-types)[Supported instance types](#efa-instance-types)Parameter naming in EC2 APIs`interface``efa``efa-only`Field naming in EC2 consoleNo selectionEFA with ENAEFA-only

## Supported interfaces and libraries

EFAs support the following interfaces and libraries:

- Open MPI 4.1 and later

- Intel MPI 2019 Update 5 and later

- NVIDIA Collective Communications Library (NCCL) 2.4.2 and later

- NVIDIA Inference Xfer Library (NIXL) 1.0.0 and later

- AWS Neuron SDK version 2.3 and later

## Supported instance types

All of the following instance types support EFA. Additionally, the tables
indicate RDMA read and RDMA write support for the instance types.

Nitro v6

Instance typeRDMA read supportRDMA write supportGeneral Purposem8a.48xlargeYesYesm8a.metal-48xlYesYesm8azn.24xlargeYesYesm8azn.metal-24xlYesYesm8gb.16xlargeYesYesm8gb.24xlargeYesYesm8gb.48xlargeYesYesm8gb.metal-24xlYesYesm8gb.metal-48xlYesYesm8gn.16xlargeYesYesm8gn.24xlargeYesYesm8gn.48xlargeYesYesm8gn.metal-24xlYesYesm8gn.metal-48xlYesYesm8i.48xlargeYesYesm8i.96xlargeYesYesm8i.metal-48xlYesYesm8i.metal-96xlYesYesm8id.48xlargeYesYesm8id.96xlargeYesYesm8id.metal-48xlYesYesm8id.metal-96xlYesYesCompute Optimizedc8a.48xlargeYesYesc8a.metal-48xlYesYesc8gb.16xlargeYesYesc8gb.24xlargeYesYesc8gb.48xlargeYesYesc8gb.metal-24xlYesYesc8gb.metal-48xlYesYesc8gn.16xlargeYesYesc8gn.24xlargeYesYesc8gn.48xlargeYesYesc8gn.metal-24xlYesYesc8gn.metal-48xlYesYesc8i.48xlargeYesYesc8i.96xlargeYesYesc8i.metal-48xlYesYesc8i.metal-96xlYesYesc8id.48xlargeYesYesc8id.96xlargeYesYesc8id.metal-48xlYesYesc8id.metal-96xlYesYesMemory Optimizedr8a.48xlargeYesYesr8a.metal-48xlYesYesr8gb.16xlargeYesYesr8gb.24xlargeYesYesr8gb.48xlargeYesYesr8gb.metal-24xlYesYesr8gb.metal-48xlYesYesr8gn.16xlargeYesYesr8gn.24xlargeYesYesr8gn.48xlargeYesYesr8gn.metal-24xlYesYesr8gn.metal-48xlYesYesr8i.48xlargeYesYesr8i.96xlargeYesYesr8i.metal-48xlYesYesr8i.metal-96xlYesYesr8id.48xlargeYesYesr8id.96xlargeYesYesr8id.metal-48xlYesYesr8id.metal-96xlYesYesx8aedz.24xlargeYesYesx8aedz.metal-24xlYesYesx8i.48xlargeYesYesx8i.64xlargeYesYesx8i.96xlargeYesYesx8i.metal-48xlYesYesx8i.metal-96xlYesYesStorage Optimizedi8ge.48xlargeYesNoi8ge.metal-48xlYesNoAccelerated Computingg7e.8xlargeYesYesg7e.12xlargeYesYesg7e.24xlargeYesYesg7e.48xlargeYesYesp6-b200.48xlargeYesYesp6-b300.48xlargeYesYesHigh Performance Computinghpc8a.96xlargeYesYes

Nitro v5

Instance typeRDMA read supportRDMA write supportGeneral Purposem8g.24xlargeYesNom8g.48xlargeYesNom8g.metal-24xlYesNom8g.metal-48xlYesNom8gd.24xlargeNoNom8gd.48xlargeNoNom8gd.metal-24xlNoNom8gd.metal-48xlNoNoCompute Optimizedc7gn.16xlargeYesNoc7gn.metalYesNoc8g.24xlargeYesNoc8g.48xlargeYesNoc8g.metal-24xlYesNoc8g.metal-48xlYesNoc8gd.24xlargeNoNoc8gd.48xlargeNoNoc8gd.metal-24xlNoNoc8gd.metal-48xlNoNoMemory Optimizedr8g.24xlargeNoNor8g.48xlargeNoNor8g.metal-24xlNoNor8g.metal-48xlNoNor8gd.24xlargeNoNor8gd.48xlargeNoNor8gd.metal-24xlNoNor8gd.metal-48xlNoNox8g.24xlargeNoNox8g.48xlargeNoNox8g.metal-24xlNoNox8g.metal-48xlNoNoStorage Optimizedi7ie.48xlargeYesNoi7ie.metal-48xlYesNoi8g.48xlargeNoNoi8g.metal-48xlNoNoAccelerated Computingp5en.48xlargeYesYesp6e-gb200.36xlargeYesYestrn2.3xlargeYesYestrn2.48xlargeYesYestrn2u.48xlargeYesYesHigh Performance Computinghpc7g.4xlargeYesNohpc7g.8xlargeYesNohpc7g.16xlargeYesNo

Nitro v4

Instance typeRDMA read supportRDMA write supportGeneral Purposem6a.48xlargeYesYesm6a.metalYesYesm6i.32xlargeYesYesm6i.metalYesYesm6id.32xlargeYesYesm6id.metalYesYesm6idn.32xlargeYesYesm6idn.metalYesYesm6in.32xlargeYesYesm6in.metalYesYesm7a.48xlargeYesNom7a.metal-48xlYesNom7g.16xlargeYesNom7g.metalYesNom7gd.16xlargeYesNom7gd.metalYesNom7i.48xlargeYesNom7i.metal-48xlYesNoCompute Optimizedc6a.48xlargeYesYesc6a.metalYesYesc6gn.16xlargeYesYesc6i.32xlargeYesYesc6i.metalYesYesc6id.32xlargeYesYesc6id.metalYesYesc6in.32xlargeYesYesc6in.metalYesYesc7a.48xlargeYesNoc7a.metal-48xlYesNoc7g.16xlargeYesYesc7g.metalYesYesc7gd.16xlargeYesNoc7gd.metalYesNoc7i.48xlargeYesNoc7i.metal-48xlYesNoMemory Optimizedr6a.48xlargeYesYesr6a.metalYesYesr6i.32xlargeYesYesr6i.metalYesYesr6id.32xlargeYesYesr6id.metalYesYesr6idn.32xlargeYesYesr6idn.metalYesYesr6in.32xlargeYesYesr6in.metalYesYesr7a.48xlargeNoNor7a.metal-48xlNoNor7g.16xlargeNoNor7g.metalNoNor7gd.16xlargeNoNor7gd.metalNoNor7i.48xlargeNoNor7i.metal-48xlNoNor7iz.32xlargeNoNor7iz.metal-32xlNoNou7i-6tb.112xlargeYesYesu7i-8tb.112xlargeYesYesu7i-12tb.224xlargeYesYesu7in-16tb.224xlargeYesYesu7in-24tb.224xlargeYesYesu7in-32tb.224xlargeYesYesu7inh-32tb.480xlargeYesYesx2idn.32xlargeYesYesx2idn.metalYesYesx2iedn.32xlargeYesYesx2iedn.metalYesYesStorage Optimizedi4g.16xlargeYesYesi4i.32xlargeYesYesi4i.metalYesYesi7i.24xlargeYesNoi7i.48xlargeYesNoi7i.metal-48xlYesNoim4gn.16xlargeYesYesAccelerated Computingf2.48xlargeYesYesg6.8xlargeYesYesg6.12xlargeYesYesg6.16xlargeYesYesg6.24xlargeYesYesg6.48xlargeYesYesg6e.8xlargeYesYesg6e.12xlargeYesYesg6e.16xlargeYesYesg6e.24xlargeYesYesg6e.48xlargeYesYesgr6.8xlargeYesYesp5.4xlargeYesYesp5.48xlargeYesYesp5e.48xlargeYesYestrn1.32xlargeYesYestrn1n.32xlargeYesYesHigh Performance Computinghpc6a.48xlargeYesYeshpc6id.32xlargeYesYeshpc7a.12xlargeYesNohpc7a.24xlargeYesNohpc7a.48xlargeYesNohpc7a.96xlargeYesNo

Nitro v3

Instance typeRDMA read supportRDMA write supportGeneral Purposem5dn.24xlargeNoNom5dn.metalNoNom5n.24xlargeNoNom5n.metalNoNom5zn.12xlargeNoNom5zn.metalNoNoCompute Optimizedc5n.9xlargeNoNoc5n.18xlargeNoNoc5n.metalNoNoMemory Optimizedr5dn.24xlargeNoNor5dn.metalNoNor5n.24xlargeNoNor5n.metalNoNox2iezn.12xlargeNoNox2iezn.metalNoNoStorage Optimizedi3en.12xlargeNoNoi3en.24xlargeNoNoi3en.metalNoNoAccelerated Computingdl1.24xlargeYesNodl2q.24xlargeNoNog4dn.8xlargeNoNog4dn.12xlargeNoNog4dn.16xlargeNoNog4dn.metalNoNog5.8xlargeNoNog5.12xlargeNoNog5.16xlargeNoNog5.24xlargeNoNog5.48xlargeNoNoinf1.24xlargeNoNop3dn.24xlargeNoNop4d.24xlargeYesNop4de.24xlargeYesNovt1.24xlargeNoNoPrevious Generationp3dn.24xlargeNoNo

###### To see the available instance types that support EFAs in a specific Region

The available instance types vary by Region. To see the available instance types that
support EFAs in a Region, use the [describe-instance-types](../../../cli/latest/reference/ec2/describe-instance-types.md)
command with the `--region` parameter. Include the `--filters`
parameter to scope the results to the instance types that support EFA and the
`--query` parameter to scope the output to the value of
`InstanceType`.

```nohighlight

aws ec2 describe-instance-types \
    --region us-east-1  \
    --filters Name=network-info.efa-supported,Values=true \
    --query "InstanceTypes[*].[InstanceType]"  \
    --output text | sort
```

## Supported operating systems

Operating system support differs depending on the processor type. The following table
shows the supported operating systems.

Operating systemIntel/AMD ( `x86_64`) instance typesAWS Graviton ( `arm64`) instance typesAmazon Linux 2023✓✓Amazon Linux 2✓✓RHEL 8 and 9✓✓Debian 11, 12, and 13✓✓Rocky Linux 8 and 9✓✓Ubuntu 22.04 and 24.04✓✓SUSE Linux Enterprise 15 SP2 and later✓✓OpenSUSE Leap 15.5 and later✓

###### Note

Some of the listed operating systems might not be supported with Intel MPI. If you are using
Intel MPI, refer to the [Intel MPI documentation](https://www.intel.com/content/www/us/en/developer/articles/system-requirements/mpi-library-system-requirements.html) to verify support for your operating system.

## EFA limitations

EFAs have the following limitations:

- RDMA write is not supported with all instance types. For more information, see
[Supported instance types](#efa-instance-types).

- EFA traffic1 between P4d/P4de/DL1 instances and other instance types is currently not supported.

- [Instance types that support multiple network cards](using-eni.md#network-cards)
can be configured with one EFA per network card. All other supported instance types
support only one EFA per instance.

- `c7g.16xlarge`, `m7g.16xlarge`, and `r7g.16xlarge`
Dedicated Instances and Dedicated Hosts are not supported when an EFA is attached.

- EFA traffic1 can't cross Availability Zones or VPCs. This does not apply to normal
IP traffic from the ENA device of an EFA interface.

- EFA traffic1 is not routable. Normal IP traffic from the ENA device of an EFA
interface remains routable.

- EFA is not supported on AWS Outposts.

- The EFA device of an EFA (EFA with ENA) interface is supported on Windows instances
only for AWS Cloud Digital Interface Software Development Kit (AWS CDI SDK) based applications. If you
attach an EFA (EFA with ENA) interface to a Windows instance for non-CDI SDK based
applications, it functions as an ENA interface, without the added EFA device capabilities.
The EFA-only interface is not supported by AWS CDI based applications on Windows or Linux.
For more information, see the [AWS Cloud Digital Interface Software Development Kit (AWS CDI SDK) User Guide](../../../../reference/cdi-sdk/latest/ug/what-is.md).

1 _EFA traffic_ refers to the traffic transmitted through the EFA device of
either an EFA (EFA with ENA) or EFA-only interface.

## EFA pricing

EFA is available as an optional Amazon EC2 networking feature that you can enable on any supported
instance at no additional cost.

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

Optimize network performance on Windows

Get started with EFA and MPI
