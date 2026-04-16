---
title: "Quotas for your IPAM"
---

# Quotas for your IPAM

This section lists the quotas related to IPAM. The Service Quotas console also provides
information about IPAM quotas. You can use the Service Quotas console to view default quotas and
[request quota increases](https://console.aws.amazon.com/servicequotas/home/services/ec2-ipam/quotas) for
adjustable quotas. For more information, see [Requesting a quota\
increase](../../../servicequotas/latest/userguide/request-quota-increase.md) in the _Service Quotas User Guide_.

NameDefaultAdjustableAmazon-provided contiguous public IPv4 CIDR blocks2Yes. Contact the AWS Support Center as described in [AWS service quotas](../../../../general/latest/gr/aws-service-limits.md) in the _AWS General Reference_. Amazon-provided contiguous public IPv4 CIDR block netmask
length/29Acceptable size is between /29 and /30. To request an increase,
contact the AWS Support Center as described in [AWS service\
quotas](../../../../general/latest/gr/aws-service-limits.md) in the _AWS General Reference_. Amazon-provided IPv6 CIDR block netmask length/52Yes. Contact the AWS Support Center as described in [AWS service quotas](../../../../general/latest/gr/aws-service-limits.md) in the _AWS General Reference_. Amazon-provided IPv6 CIDR blocks per Regional pool1Yes. Contact the AWS Support Center as described in [AWS service quotas](../../../../general/latest/gr/aws-service-limits.md) in the _AWS General Reference_.
Autonomous System Numbers (ASNs) that you can bring to IPAM5Yes. Contact the AWS Support Center as described in [AWS service quotas](../../../../general/latest/gr/aws-service-limits.md) in the _AWS General Reference_.
CIDRs per pool50[Yes](https://console.aws.amazon.com/servicequotas/home/services/ec2-ipam/quotas/L-0BC051D6)Enabled targets per IPAM policy100Yes. To request an adjustment to the quota, contact the AWS Support Center as described in [AWS service quotas](../../../../general/latest/gr/aws-service-limits.md) in the _AWS General Reference_.IPAM administrators per organization1NoIPAMs per Region1[No](https://console.aws.amazon.com/servicequotas/home/services/ec2-ipam/quotas/L-F8B4A9E6)IPAM policies per IPAM10Yes. To request an adjustment to the quota, contact the AWS Support Center as described in [AWS service quotas](../../../../general/latest/gr/aws-service-limits.md) in the _AWS General Reference_.IPAM policy allocation rules per resource-locale pair\*10Yes. To request an adjustment to the quota, contact the AWS Support Center as described in [AWS service quotas](../../../../general/latest/gr/aws-service-limits.md) in the _AWS General Reference_.[Organizational unit exclusions per resource discovery](exclude-ous.md)10Yes. Contact the AWS Support Center as described in [AWS service quotas](../../../../general/latest/gr/aws-service-limits.md) in the _AWS General Reference_. Pool depth (the number of pools within pools)10[Yes](https://console.aws.amazon.com/servicequotas/home/services/ec2-ipam/quotas/L-047C0565)Pools per scope50[Yes](https://console.aws.amazon.com/servicequotas/home/services/ec2-ipam/quotas/L-7319AFC3)Prefix list resolvers per IPAM10[Yes](https://console.aws.amazon.com/servicequotas/home/services/ec2-ipam/quotas)Prefix list resolver targets per prefix list resolver50Yes. Contact the AWS Support Center as described in [AWS service quotas](../../../../general/latest/gr/aws-service-limits.md) in the _AWS General Reference_.Rules per prefix list resolver100Yes. Contact the AWS Support Center as described in [AWS service quotas](../../../../general/latest/gr/aws-service-limits.md) in the _AWS General Reference_.CIDR entries per prefix list resolver version1000Yes. Contact the AWS Support Center as described in [AWS service quotas](../../../../general/latest/gr/aws-service-limits.md) in the _AWS General Reference_.Resource discovery associations per IPAM5[Yes](https://console.aws.amazon.com/servicequotas/home/services/ec2-ipam/quotas/L-037D1B6C)Resource discoveries per Region1[No](https://console.aws.amazon.com/servicequotas/home/services/ec2-ipam/quotas/L-F0D8E837)[Resource utilization metrics](cloudwatch-ipam-res-util.md)50Yes. Contact the AWS Support Center as described in [AWS service quotas](../../../../general/latest/gr/aws-service-limits.md) in the _AWS General Reference_. Scopes per IPAM5[Yes](https://console.aws.amazon.com/servicequotas/home/services/ec2-ipam/quotas/L-F493CFD2). When you create an IPAM, a private and public default scope
are created for you. If you want to create additional scopes, they will be
private scopes. You cannot create additional public scopes.

\\* _Resource-locale pair:_ When setting allocation rules, you must specify both a resource type (the AWS resource like EIPs, ALBs, or RDS clusters) and a locale (the AWS Region or Local Zone where the rule applies). Allocation rules are scoped to this resource type and locale combination. For example, if you're setting a policy for EIPs in us-east-1, you can set up to 10 rules for that specific resource-locale pair\*.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Example policy

Pricing

All content copied from https://docs.aws.amazon.com/.
