# CreateVerifiedAccessEndpointLoadBalancerOptions

Describes the load balancer options when creating an AWS Verified Access endpoint using the
`load-balancer` type.

## Contents

**LoadBalancerArn**

The ARN of the load balancer.

Type: String

Required: No

**Port**

The IP port number.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 65535.

Required: No

**PortRange.N**

The port ranges.

Type: Array of [CreateVerifiedAccessEndpointPortRange](api-createverifiedaccessendpointportrange.md) objects

Required: No

**Protocol**

The IP protocol.

Type: String

Valid Values: `http | https | tcp`

Required: No

**SubnetId.N**

The IDs of the subnets. You can specify only one subnet per Availability Zone.

Type: Array of strings

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/createverifiedaccessendpointloadbalanceroptions.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/createverifiedaccessendpointloadbalanceroptions.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/createverifiedaccessendpointloadbalanceroptions.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateVerifiedAccessEndpointEniOptions

CreateVerifiedAccessEndpointPortRange
