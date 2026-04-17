---
title: "Subnet"
---

# Subnet

Represents the subnet associated with a DAX cluster. This parameter
refers to subnets defined in Amazon Virtual Private Cloud (Amazon VPC) and used with
DAX.

## Contents

###### Note

In the following list, the required parameters are described first.

**SubnetAvailabilityZone**

The Availability Zone (AZ) for the subnet.

Type: String

Required: No

**SubnetIdentifier**

The system-assigned identifier for the subnet.

Type: String

Required: No

**SupportedNetworkTypes**

The network types supported by this subnet. Returns an array of strings that can
include `ipv4`, `ipv6`, or both, indicating whether the subnet
supports IPv4 only, IPv6 only, or dual-stack deployments.

Type: Array of strings

Valid Values: `ipv4 | ipv6 | dual_stack`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/dax-2017-04-19/Subnet)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/dax-2017-04-19/Subnet)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/dax-2017-04-19/Subnet)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SSESpecification

SubnetGroup

All content copied from https://docs.aws.amazon.com/.
