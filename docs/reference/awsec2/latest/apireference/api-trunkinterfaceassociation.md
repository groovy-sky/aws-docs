# TrunkInterfaceAssociation

Information about an association between a branch network interface with a trunk network interface.

## Contents

**associationId**

The ID of the association.

Type: String

Required: No

**branchInterfaceId**

The ID of the branch network interface.

Type: String

Required: No

**greKey**

The application key when you use the GRE protocol.

Type: Integer

Required: No

**interfaceProtocol**

The interface protocol. Valid values are `VLAN` and `GRE`.

Type: String

Valid Values: `VLAN | GRE`

Required: No

**TagSet.N**

The tags for the trunk interface association.

Type: Array of [Tag](api-tag.md) objects

Required: No

**trunkInterfaceId**

The ID of the trunk network interface.

Type: String

Required: No

**vlanId**

The ID of the VLAN when you use the VLAN protocol.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/trunkinterfaceassociation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/trunkinterfaceassociation.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/trunkinterfaceassociation.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

TransitGatewayVpcAttachmentOptions

TunnelOption
