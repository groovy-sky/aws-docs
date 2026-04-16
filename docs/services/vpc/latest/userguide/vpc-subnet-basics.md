---
title: "VPC basics"
---

# VPC basics

A VPC spans all of the Availability Zones in a Region. After you create a VPC, you
can add one or more subnets in each Availability Zone. For more information, see [Subnets for your VPC](configure-subnets.md).

###### Contents

- [VPC IP address range](#vpc-ip-address-range)

- [VPC diagram](#vpc-diagram)

- [VPC resources](#vpc-resources)

## VPC IP address range

When you create a VPC, you specify its IP addresses as follows:

- IPv4 only – The VPC has an
IPv4 CIDR block but does not have an IPv6 CIDR block.

- Dual stack – The VPC has both an
IPv4 CIDR block and an IPv6 CIDR block.

For more information, see [IP addressing for your VPCs and subnets](vpc-ip-addressing.md).

## VPC diagram

The following diagram shows a VPC with no additional VPC resources.
For example VPC configurations, see [VPC examples](vpc-examples-intro.md).

![A VPC that spans the Availability Zones for its Region.](https://docs.aws.amazon.com/images/vpc/latest/userguide/images/vpc-diagram.png)

## VPC resources

Each VPC automatically comes with the following resources:

- [Default DHCP option set](dhcpoptionsetconcepts.md#ArchitectureDiagram)

- [Default network ACL](default-network-acl.md)

- [Default security group](default-security-group.md)

- [Main route table](subnet-route-tables.md#main-route-table)

You can create the following resources for your VPC:

- [Network ACLs](vpc-network-acls.md)

- [Custom route tables](vpc-route-tables.md)

- [Security groups](vpc-security-groups.md)

- [Internet gateway](vpc-internet-gateway.md)

- [NAT gateways](vpc-nat-gateway.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Virtual private clouds

VPC configuration options

All content copied from https://docs.aws.amazon.com/.
