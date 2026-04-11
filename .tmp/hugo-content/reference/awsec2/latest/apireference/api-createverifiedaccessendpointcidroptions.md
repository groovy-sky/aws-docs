# CreateVerifiedAccessEndpointCidrOptions

Describes the CIDR options for a Verified Access endpoint.

## Contents

**Cidr**

The CIDR.

Type: String

Required: No

**PortRange.N**

The port ranges.

Type: Array of [CreateVerifiedAccessEndpointPortRange](api-createverifiedaccessendpointportrange.md) objects

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

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/createverifiedaccessendpointcidroptions.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/createverifiedaccessendpointcidroptions.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/createverifiedaccessendpointcidroptions.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateTransitGatewayVpcAttachmentRequestOptions

CreateVerifiedAccessEndpointEniOptions
