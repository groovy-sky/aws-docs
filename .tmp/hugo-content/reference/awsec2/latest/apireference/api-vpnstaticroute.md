# VpnStaticRoute

Describes a static route for a VPN connection.

## Contents

**destinationCidrBlock**

The CIDR block associated with the local subnet of the customer data center.

Type: String

Required: No

**source**

Indicates how the routes were provided.

Type: String

Valid Values: `Static`

Required: No

**state**

The current state of the static route.

Type: String

Valid Values: `pending | available | deleting | deleted`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/vpnstaticroute.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/vpnstaticroute.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/vpnstaticroute.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

VpnGateway

VpnTunnelLogOptions
