# Subnet

Represents the subnet associated with a cluster. This parameter refers to subnets
defined in Amazon Virtual Private Cloud (Amazon VPC) and used with ElastiCache.

## Contents

###### Note

In the following list, the required parameters are described first.

**SubnetAvailabilityZone**

The Availability Zone associated with the subnet.

Type: [AvailabilityZone](api-availabilityzone.md) object

Required: No

**SubnetIdentifier**

The unique identifier for the subnet.

Type: String

Required: No

**SubnetOutpost**

The outpost ARN of the subnet.

Type: [SubnetOutpost](api-subnetoutpost.md) object

Required: No

**SupportedNetworkTypes.member.N**

Either `ipv4` \| `ipv6` \| `dual_stack`. IPv6 is
supported for workloads using Valkey 7.2 and above, Redis OSS engine version 6.2
to 7.1 or Memcached engine version 1.6.6 and above on all instances built on the [Nitro system](http://aws.amazon.com/ec2/nitro).

Type: Array of strings

Valid Values: `ipv4 | ipv6 | dual_stack`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/subnet.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/subnet.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/subnet.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Snapshot

SubnetOutpost

All content copied from https://docs.aws.amazon.com/.
