---
title: "Understand traffic mirror source and target connectivity options"
---

# Understand traffic mirror source and target connectivity options

Connecting traffic mirror sources and targets in different VPCs can be useful for monitoring and analyzing network traffic across multiple environments. This allows you to centralize your network monitoring and security analysis, even if your application infrastructure is spread across different VPCs or AWS accounts.

The traffic mirror source and the traffic mirror target (monitoring appliance) can be in the
same VPC or different VPCs and can be connected using the following resources:

- An intra-Region VPC peering connection.

- A transit gateway.

- A Gateway Load Balancer endpoint.

The traffic mirror target can be owned by an AWS account that is different from the
traffic mirror source.

The mirrored traffic is sent to the traffic mirror target using the source VPC route table.
Before you configure Traffic Mirroring, make sure that the traffic mirror source can route to the
traffic mirror target.

The following table describes the available resource configurations.

Source ownerSource VPCTarget ownerTarget VPCConnectivity optionAccount AVPC 1Account AVPC1No additional configurationAccount AVPC 1Account AVPC 2Intra-Region peering or a transit gateway or Gateway Load Balancer endpointAccount AVPC 1Account BVPC 2Cross-account intra-Region peering connection, a transit gateway, or a Gateway Load Balancer endpointAccount AVPC 1Account BVPC 1VPC sharing

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Sessions

Packet format

All content copied from https://docs.aws.amazon.com/.
