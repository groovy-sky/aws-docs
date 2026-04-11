# TransitGatewayConnectPeerConfiguration

Describes the Connect peer details.

## Contents

**BgpConfigurations.N**

The BGP configuration details.

Type: Array of [TransitGatewayAttachmentBgpConfiguration](api-transitgatewayattachmentbgpconfiguration.md) objects

Required: No

**InsideCidrBlocks.N**

The range of interior BGP peer IP addresses.

Type: Array of strings

Required: No

**peerAddress**

The Connect peer IP address on the appliance side of the tunnel.

Type: String

Required: No

**protocol**

The tunnel protocol.

Type: String

Valid Values: `gre`

Required: No

**transitGatewayAddress**

The Connect peer IP address on the transit gateway side of the tunnel.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/transitgatewayconnectpeerconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/transitgatewayconnectpeerconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/transitgatewayconnectpeerconfiguration.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

TransitGatewayConnectPeer

TransitGatewayConnectRequestBgpOptions
