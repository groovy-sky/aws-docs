# Preparing for a cross-VPC active-active cluster

You can configure an active-active cluster with Amazon RDS for MySQL DB instances in more
than one VPC. The VPCs can be in the same AWS Region or different
AWS Regions.

###### Note

Sending traffic between multiple AWS Regions might incur additional costs. For more information,
see [Overview of Data Transfer Costs for Common Architectures](https://aws.amazon.com/blogs/architecture/overview-of-data-transfer-costs-for-common-architectures).

If you are configuring an active-active cluster in a single VPC, you can skip these steps and move on to
[Setting up an active-active cluster with new DB instances](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/mysql-active-active-clusters-setting-up.html).

###### To prepare for an active-active cluster with DB instances in more than one VPC

1. Make sure the IPv4 address ranges in the CIDR blocks meet the following requirements:

- The IPv4 address ranges in the CIDR blocks of the VPCs can't overlap.

- All of the IPv4 address ranges in the CIDR blocks either must be lower than
`128.0.0.0/subnet_mask`
or higher than 128.0.0.0/ `subnet_mask`.

The following ranges illustrate these requirements:

- `10.1.0.0/16` in one VPC and `10.2.0.0/16` in the other
VPC is supported.

- `172.1.0.0/16` in one VPC and `172.2.0.0/16` in the other
VPC is supported.

- `10.1.0.0/16` in one VPC and `10.1.0.0/16` in the other
VPC _is not_ supported because the ranges overlap.

- `10.1.0.0/16` in one VPC and `172.1.0.0/16` in the other
VPC _is not_ supported because one is below
`128.0.0.0/subnet_mask` and the other is
above `128.0.0.0/subnet_mask`.

For information about CIDR blocks, see [VPC CIDR blocks](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-cidr-blocks.html) in the _Amazon VPC User Guide_.

2. In each VPC, make sure DNS resolution and DNS hostnames are both enabled.

For instructions, see [View and update DNS attributes for your VPC](https://docs.aws.amazon.com/vpc/latest/userguide/vpc-dns.html#vpc-dns-updating) in the _Amazon VPC User Guide_.

3. Configure the VPCs so that you can route traffic between them in one of the following ways:

- Create a VPC peering connection between the VPCs.

For instructions, see [Create a VPC peering connection](https://docs.aws.amazon.com/vpc/latest/peering/create-vpc-peering-connection.html) in the _Amazon VPC Peering Guide_. In each VPC, make sure there are
inbound rules for your security groups that reference security groups in the peered VPC. Doing so allows traffic
to flow to and from instances that are associated with the referenced security group in the peered VPC.
For instructions, see [Update your security groups to reference peer security groups](https://docs.aws.amazon.com/vpc/latest/peering/vpc-peering-security-groups.html) in the _Amazon VPC Peering Guide_.

- Create a transit gateway between the VPCs.

For instructions, see [Getting started with transit gateways](https://docs.aws.amazon.com/vpc/latest/tgw/tgw-getting-started.html) in _Amazon VPC Transit Gateways_. In each VPC, make sure there are
inbound rules for your security groups that allow traffic from the other VPC, such as inbound rules that specify
the CIDR of the other VPC. Doing so allows traffic to flow to and from instances that are associated with the referenced
security group in the active-active cluster. For more information, see [Control traffic to your AWS resources using security groups](../../../vpc/latest/userguide/vpc-security-groups.md#working-with-security-groups) in the _Amazon VPC User Guide_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Limitations and considerations for active-active clusters

Required parameter settings for active-active clusters
