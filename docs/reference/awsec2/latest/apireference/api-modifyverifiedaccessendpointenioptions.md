# ModifyVerifiedAccessEndpointEniOptions

Describes the options when modifying a Verified Access endpoint with the
`network-interface` type.

## Contents

**Port**

The IP port number.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 65535.

Required: No

**PortRange.N**

The port ranges.

Type: Array of [ModifyVerifiedAccessEndpointPortRange](api-modifyverifiedaccessendpointportrange.md) objects

Required: No

**Protocol**

The IP protocol.

Type: String

Valid Values: `http | https | tcp`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/modifyverifiedaccessendpointenioptions.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/modifyverifiedaccessendpointenioptions.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/modifyverifiedaccessendpointenioptions.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ModifyVerifiedAccessEndpointCidrOptions

ModifyVerifiedAccessEndpointLoadBalancerOptions
