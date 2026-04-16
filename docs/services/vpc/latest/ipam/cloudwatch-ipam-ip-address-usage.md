---
title: "IPAM metrics"
---

# IPAM metrics

IPAM publishes data about your IPAM, pools, and scopes to Amazon CloudWatch. You can use these
metrics to create alarms for IPAM pools to notify you if the address pools are nearing
exhaustion or if resources fail to comply with allocation rules set on a pool. Creating
alarms and setting up notifications with Amazon CloudWatch is outside the scope of this section.
For more information, see [Using\
Amazon CloudWatch alarms](../../../amazoncloudwatch/latest/monitoring/alarmthatsendsemail.md) in the _Amazon CloudWatch User Guide_.

The metrics and dimensions that IPAM sends to Amazon CloudWatch are listed below.

## IPAM metrics

The `AWS/IPAM` namespace includes the following IPAM metrics.

Metric nameDescriptionTotalActiveIpCount

The total active IP count is the number of active IP
addresses in your IPAM that you would be charged if you switched from the Free Tier
to the Advanced Tier. An active IP address is defined as an IP address or a prefix associated with an Elastic Network Interface (ENI) that is attached to a resource such as an EC2 Instance.

- This metric is only available to customers in the Free Tier.

- If your IPAM is [integrated with AWS Organizations](enable-integ-ipam.md), the active IP count covers all the Organization accounts.

- You cannot view a breakdown of the active IP count by IP type (public/private) or class (IPv4/IPv6).

- IPAM only counts IPs from ENIs owned by monitored accounts. The count may be inaccurate for shared subnets. IP addresses are excluded if the subnet owner or ENI owner is not covered by IPAM.

## IPAM pool metrics

The `AWS/IPAM` namespace includes the following pool metrics for IPAM.

Metric nameDescriptionCompliantResourceCidrsThe number of managed resource CIDRs that comply with the allocation rules of the IPAM pool. For more information about allocation rules, see [Create a top-level IPv4 pool](create-top-ipam.md).NoncompliantResourceCidrsThe number of managed resource CIDRs that do not comply with the allocation rules of the IPAM pool. For more information about allocation rules, see [Create a top-level IPv4 pool](create-top-ipam.md).PercentAllocatedThe percentage of a pool's IP space that has been allocated to other pools.PercentAssignedThe percentage of a pools IP space that has been allocated to resources, including manual allocations.PercentAvailableThe percentage of a pool's IP space that has not been allocated to other pools or resources.

## IPAM scope metrics

The `AWS/IPAM` namespace includes the following scope metrics for IPAM.

Metric nameDescriptionCompliantResourceCidrsThe number of resource CIDRs that comply with the allocation rules for IPAM pools in the scope.ManagedResourceCidrsThe number of resource CIDRs for manageable resources (VPCs or public IPv4 pools) that are allocated from an IPAM pool in the scope.NoncompliantResourceCidrsThe number of resource CIDRs that do not comply with the allocation rules for the IPAM pools in the scope.OverlappingResourceCidrsThe number of resource CIDRs that overlap in the scope.UnmanagedResourceCidrsThe number of resource CIDRs in the scope that are currently associated with manageable resources but are not managed by IPAM.

## IPAM public IP metrics

The `AWS/IPAM` namespace includes the following public IP metrics for IPAM.

Metric nameDescriptionAmazonOwnedContigIPsThe number of IP addresses within CIDRs that are provisioned to
Amazon-provided contiguous public IPv4 pools owned by the IPAM.AllocatedAmazonOwnedContigIPsThe number of IP addresses that have been allocated from an
Amazon-provided contiguous public IPv4 pool CIDR block.UnallocatedAmazonOwnedContigIPsThe number of IP addresses within the Amazon-provided contiguous
public IPv4 pool CIDR block owned by the IPAM.AssociatedAmazonOwnedContigIPsThe number of Elastic IP addresses that have been allocated from an
Amazon-provided contiguous public IPv4 pool CIDR block that are
associated with an elastic network interface.UnassociatedAmazonOwnedContigIPsThe number of Elastic IP addresses that have been allocated from an
Amazon-provided contiguous public IPv4 pool CIDR block that are not
associated with an elastic network interface.

## IPAM prefix list resolver metrics

We encourage you to set CloudWatch alarms on failure metrics as you may need to reassess and adjust [IPAM prefix list resolver rules](automate-prefix-list-updates.md) to stay within the limits for version and prefix list size.

Metric nameDescriptionIpamPrefixListResolverSyncFailurePrefix list resolver failed to sync with target. This may happen if a quota such as 'CIDR entries per prefix list resolver version' is exceeded, the target prefix list is not found, or sync is disabled on the target managed prefix list.IpamPrefixListResolverSyncSuccessPrefix list resolver successfully synced with target.IpamPrefixListResolverVersionCreationSuccessVersion creation succeeded.IpamPrefixListResolverVersionCreationFailureVersion creation failed. This may happen if you've reached your 'CIDR entries per prefix list resolver version' quota.

## Metric dimensions

To filter the IPAM metrics, use the following dimensions.

DimensionDescriptionAddressFamilyThe IP address family for resource CIDRs (IPv4 or IPv6).LocaleThe AWS Region where an IPAM pool is available for allocations.PoolIDThe ID of a pool.ScopeIDThe ID of a scope.

For information about monitoring VPCs with Amazon CloudWatch, see [CloudWatch metrics for your VPCs](../userguide/vpc-cloudwatch.md) in the
_Amazon Virtual Private Cloud User Guide_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Manage alarms

Resource utilization metrics

All content copied from https://docs.aws.amazon.com/.
