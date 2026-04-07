# VerifiedAccessEndpointEniOptions

Options for a network-interface type endpoint.

## Contents

**networkInterfaceId**

The ID of the network interface.

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

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/verifiedaccessendpointenioptions.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/verifiedaccessendpointenioptions.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/verifiedaccessendpointenioptions.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

VerifiedAccessEndpointCidrOptions

VerifiedAccessEndpointLoadBalancerOptions
