# VerifiedAccessEndpointLoadBalancerOptions

Describes a load balancer when creating an AWS Verified Access endpoint using the
`load-balancer` type.

## Contents

**loadBalancerArn**

The ARN of the load balancer.

Type: String

Required: No

**port**

The IP port number.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 65535.

Required: No

**PortRangeSet.N**

The port ranges.

Type: Array of [VerifiedAccessEndpointPortRange](api-verifiedaccessendpointportrange.md) objects

Required: No

**protocol**

The IP protocol.

Type: String

Valid Values: `http | https | tcp`

Required: No

**SubnetIdSet.N**

The IDs of the subnets.

Type: Array of strings

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/verifiedaccessendpointloadbalanceroptions.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/verifiedaccessendpointloadbalanceroptions.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/verifiedaccessendpointloadbalanceroptions.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

VerifiedAccessEndpointEniOptions

VerifiedAccessEndpointPortRange
