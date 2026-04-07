# ModifyVerifiedAccessEndpointLoadBalancerOptions

Describes a load balancer when creating an AWS Verified Access endpoint using the
`load-balancer` type.

## Contents

**Port**

The IP port number.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 65535.

Required: No

**PortRange.N**

The port ranges.

Type: Array of [ModifyVerifiedAccessEndpointPortRange](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyVerifiedAccessEndpointPortRange.html) objects

Required: No

**Protocol**

The IP protocol.

Type: String

Valid Values: `http | https | tcp`

Required: No

**SubnetId.N**

The IDs of the subnets.

Type: Array of strings

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/ModifyVerifiedAccessEndpointLoadBalancerOptions)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/ModifyVerifiedAccessEndpointLoadBalancerOptions)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/ModifyVerifiedAccessEndpointLoadBalancerOptions)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ModifyVerifiedAccessEndpointEniOptions

ModifyVerifiedAccessEndpointPortRange
