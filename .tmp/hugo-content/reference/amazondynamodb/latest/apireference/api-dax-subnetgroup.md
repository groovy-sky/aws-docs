# SubnetGroup

Represents the output of one of the following actions:

- _CreateSubnetGroup_

- _ModifySubnetGroup_

## Contents

###### Note

In the following list, the required parameters are described first.

**Description**

The description of the subnet group.

Type: String

Required: No

**SubnetGroupName**

The name of the subnet group.

Type: String

Required: No

**Subnets**

A list of subnets associated with the subnet group.

Type: Array of [Subnet](api-dax-subnet.md) objects

Required: No

**SupportedNetworkTypes**

The network types supported by this subnet. Returns an array of strings that can
include `ipv4`, `ipv6`, or both, indicating whether the subnet
group supports IPv4 only, IPv6 only, or dual-stack deployments.

Type: Array of strings

Valid Values: `ipv4 | ipv6 | dual_stack`

Required: No

**VpcId**

The Amazon Virtual Private Cloud identifier (VPC ID) of the subnet group.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dax-2017-04-19/subnetgroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dax-2017-04-19/subnetgroup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dax-2017-04-19/subnetgroup.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Subnet

Tag

All content copied from https://docs.aws.amazon.com/.
