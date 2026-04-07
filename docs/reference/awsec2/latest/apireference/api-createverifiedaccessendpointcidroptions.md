# CreateVerifiedAccessEndpointCidrOptions

Describes the CIDR options for a Verified Access endpoint.

## Contents

**Cidr**

The CIDR.

Type: String

Required: No

**PortRange.N**

The port ranges.

Type: Array of [CreateVerifiedAccessEndpointPortRange](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateVerifiedAccessEndpointPortRange.html) objects

Required: No

**Protocol**

The protocol.

Type: String

Valid Values: `http | https | tcp`

Required: No

**SubnetId.N**

The IDs of the subnets.

Type: Array of strings

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/CreateVerifiedAccessEndpointCidrOptions)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/CreateVerifiedAccessEndpointCidrOptions)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/CreateVerifiedAccessEndpointCidrOptions)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CreateTransitGatewayVpcAttachmentRequestOptions

CreateVerifiedAccessEndpointEniOptions
