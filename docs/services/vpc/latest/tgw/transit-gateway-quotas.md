---
title: "AWS Transit Gateway Quotas"
---

# AWS Transit Gateway Quotas

Your AWS account has the following quotas (previously referred to as
_limits_) related to transit gateways. Unless otherwise noted, each quota
is Region-specific.

The Service Quotas console provides information about the quotas for your account. You can use
the Service Quotas console to view default quotas and [request quota increases](https://console.aws.amazon.com/servicequotas/home?) for
adjustable quotas. For more information, see [Requesting a quota\
increase](../../../servicequotas/latest/userguide/request-quota-increase.md) in the _Service Quotas User Guide_.

If an adjustable quota is not yet available in Service Quotas, you can open a support case.

## General

NameDefaultAdjustableTransit gateways per account5[Yes](https://console.aws.amazon.com/servicequotas/home/services/ec2/quotas/L-A2478D36)CIDR blocks per transit gateway5No

The CIDR blocks are used in the [Connect attachments and Connect peers in AWS Transit Gateway](tgw-connect.md) feature.

## Routing

NameDefaultAdjustableTransit gateway route tables per transit gateway20[Yes](https://console.aws.amazon.com/servicequotas/home/services/ec2/quotas/L-43872EB7)Total combined routes (dynamic and static) across all route tables
for a single transit gateway10,000Contact your Solutions Architect (SA) or Technical Account Manager
(TAM) for further assistance.Dynamic routes advertised from a virtual router appliance to a Connect peer1,000Contact your Solutions Architect (SA) or Technical Account Manager
(TAM) for further assistance.Routes advertised from a Connect peer on a transit gateway to a virtual router
appliance5,000NoStatic routes for a prefix to a single attachment1No

Advertised routes come from the route table that's associated with the Connect attachment.

## Transit gateway attachments

A transit gateway cannot have more than one VPC attachment to the same VPC.

NameDefaultAdjustableAttachments per transit gateway5,000[Yes](https://console.aws.amazon.com/servicequotas/home/services/ec2/quotas/L-E0233F82)Transit gateways per VPC5NoPeering attachments per transit gateway50[Yes](https://console.aws.amazon.com/servicequotas/home/services/ec2/quotas/L-A1B5A36F)Pending peering attachments per transit gateway10[Yes](https://console.aws.amazon.com/servicequotas/home/services/ec2/quotas/L-62499967)Peering attachments between two transit gateways or between one transit gateway and a
Cloud WAN core network edge (CNE)1NoConnect peers (GRE tunnels) per Connect attachment4NoVPN Concentrators per transit gateway5NoVPN connections per VPN Concentrator100No

## Bandwidth

There are many factors that can affect realized bandwidth through a Site-to-Site VPN connection,
including but not limited to: packet size, traffic mix (TCP/UDP), shaping or throttling
policies on intermediate networks, internet weather, and specific application
requirements. For VPC attachments, Direct Connect gateways, or peered transit gateway attachments, we
will attempt to provide additional bandwidth beyond the default value.

NameDefaultAdjustableBandwidth per VPC attachment per Availability ZoneUp to 100 Gbps each direction (i.e., 100 Gbps ingress and 100 Gbps egress)Contact your Solutions Architect (SA) or Technical Account Manager
(TAM) for further assistance.Packets per second per transit gateway VPC attachment per Availability
ZoneUp to 7,500,000Contact your Solutions Architect (SA) or Technical Account Manager
(TAM) for further assistance.Bandwidth for Direct Connect gateway or peered transit gateway connection
per available Availability Zone in the RegionUp to 100 Gbps each direction (i.e., 100 Gbps ingress and 100 Gbps egress)Contact your Solutions Architect (SA) or Technical Account Manager
(TAM) for further assistance.Packets per second per transit gateway attachment (Direct Connect and
peering attachments) per available Availability Zone in the RegionUp to 7,500,000Contact your Solutions Architect (SA) or Technical Account Manager
(TAM) for further assistance.Maximum bandwidth per Connect peer (GRE tunnel) per Connect attachmentUp to 5 GbpsNoMaximum packets per second per Connect peerUp to 300,000No

You can use equal-cost multipath routing (ECMP) to get
higher VPN bandwidth by aggregating multiple VPN tunnels. To use ECMP, the VPN
connection must be configured for dynamic routing. ECMP is not supported on VPN
connections that use static routing.

You can create up to 4 Connect peers per Connect
attachment (up to 20 Gbps in total bandwidth per Connect attachment), as long
as the underlying transport (VPC or Direct Connect) attachment supports the required
bandwidth. You can use ECMP to get higher bandwidth by scaling horizontally
across multiple Connect peers of the same Connect attachment or across
multiple Connect attachments on the same transit gateway. The transit gateway cannot use ECMP
between the BGP peerings of the same Connect peer.

For bandwidth and packet limits with VPN tunnel, please refer to
[VPN bandwidth and throughput](../../../vpn/latest/s2svpn/vpn-limits.md#vpn-quotas-bandwidth) .

## Direct Connect gateways

NameDefaultAdjustableDirect Connect gateways per transit gateway20NoTransit gateways per Direct Connect gateway6No

## Maximum transmission unit (MTU)

- The MTU of a network connection is the size, in bytes, of the largest
permissible packet that can be passed over the connection. The larger the MTU of
a connection, the more data that can be passed in a single packet. A transit gateway
supports an MTU of 8500 bytes for traffic between VPCs, Direct Connect, Transit
Gateway Connect, and peering attachments (intra-Region, inter-Region, and Cloud
WAN peering attachments). Traffic over VPN connections can have an MTU of 1500
bytes.

- When migrating from VPC peering to use a transit gateway, an MTU size mismatch between
VPC peering and the transit gateway might result in some asymmetric traffic packets
dropping. Update both VPCs at the same time to avoid jumbo packets dropping due
to a size mismatch.

- The transit gateway enforces Maximum Segment Size (MSS) clamping for all packets. For
more information, see [RFC879](https://tools.ietf.org/html/rfc879).

- For details about Site-to-Site VPN quotas for MTU, see [Maximum transmission unit (MTU)](../../../vpn/latest/s2svpn/vpn-limits.md#vpn-quotas-mtu) in the _AWS Site-to-Site VPN User Guide_.

- Transit gateways support Path MTU Discovery (PMTUD) for traffic ingressing on
VPC and Connect attachments. Transit gateway generates the
`FRAG_NEEDED` for ICMPv4 packets and `Packet Too Big
                          (PTB)` for ICMPv6 packets. Transit gateways does not support PMTUD on
Site-to-site VPN, Direct Connect, and Peering attachments. For more information
about Path MTU Discovery, see [Path MTU\
Discovery](../userguide/path-mtu-discovery.md) in the _Amazon VPC User Guide_

## Multicast

###### Note

Transit gateway multicast may not be suitable for high-frequency trading or
performance-sensitive applications. We strongly recommend that you review the
following multicast limits. Contact your account or Solution Architect team for a
detailed review of your performance requirements.

NameDefaultAdjustableMulticast domains per transit gateway20Contact your Solutions Architect (SA) or Technical Account Manager
(TAM) for further assistance.Multicast network interfaces per transit gateway10,000Contact your Solutions Architect (SA) or Technical Account Manager
(TAM) for further assistance.Multicast domain associations per VPC20Contact your Solutions Architect (SA) or Technical Account Manager
(TAM) for further assistance.Static and IGMPv2 multicast group members and sources per transit gateway 10,000NoStatic and IGMPv2 multicast group members per transit gateway multicast
group100NoMaximum multicast throughput per flow1 GbpsNoMaximum aggregate multicast throughput per Availability Zone20 GbpsNoMaximum packets per second per flow (less than 10 receivers)75,000NoMaximum packets per second per flow (greater than 10
receivers)15,000NoMaximum aggregate packets per second (less than 10 receivers)2,500,000NoMaximum aggregate packets per second (greater than 10
receivers)500,000No

## AWS Network Manager

NameDefaultAdjustableGlobal networks per AWS account5YesDevices per global network200YesLinks per global network200YesSites per global network200YesConnections per global network500No

## Additional quota resources

For more information, see the following:

- [Site-to-Site VPN quotas](../../../vpn/latest/s2svpn/vpn-limits.md) in the
_AWS Site-to-Site VPN User Guide_

- [Amazon VPC quotas](../userguide/amazon-vpc-limits.md) in the
_Amazon VPC User Guide_

- [Direct Connect \
quotas](../../../directconnect/latest/userguide/limits.md) in the _AWS Direct Connect User Guide_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Network ACLs

Document history

All content copied from https://docs.aws.amazon.com/.
