# SecondaryNetworkIpv4CidrBlockAssociation

Describes an IPv4 CIDR block associated with a secondary network.

## Contents

**associationId**

The association ID for the IPv4 CIDR block.

Type: String

Required: No

**cidrBlock**

The IPv4 CIDR block.

Type: String

Required: No

**state**

The state of the CIDR block association.

Type: String

Valid Values: `associating | associated | association-failed | disassociating | disassociated | disassociation-failed`

Required: No

**stateReason**

The reason for the current state of the CIDR block association.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/secondarynetworkipv4cidrblockassociation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/secondarynetworkipv4cidrblockassociation.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/secondarynetworkipv4cidrblockassociation.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

SecondaryNetwork

SecondarySubnet
