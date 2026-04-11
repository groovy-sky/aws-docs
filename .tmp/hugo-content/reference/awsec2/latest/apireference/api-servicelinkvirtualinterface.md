# ServiceLinkVirtualInterface

Describes the service link virtual interfaces that establish connectivity between AWS Outpost and on-premises networks.

## Contents

**configurationState**

The current state of the service link virtual interface.

Type: String

Valid Values: `pending | available | deleting | deleted`

Required: No

**localAddress**

The IPv4 address assigned to the local gateway virtual interface on the Outpost side.

Type: String

Required: No

**outpostArn**

The Outpost Amazon Resource Number (ARN) for the service link virtual interface.

Type: String

Required: No

**outpostId**

The Outpost ID for the service link virtual interface.

Type: String

Required: No

**outpostLagId**

The link aggregation group (LAG) ID for the service link virtual interface.

Type: String

Required: No

**ownerId**

The ID of the AWS account that owns the service link virtual interface..

Type: String

Required: No

**peerAddress**

The IPv4 peer address for the service link virtual interface.

Type: String

Required: No

**peerBgpAsn**

The ASN for the Border Gateway Protocol (BGP) associated with the service link virtual interface.

Type: Long

Required: No

**serviceLinkVirtualInterfaceArn**

The Amazon Resource Number (ARN) for the service link virtual interface.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1283.

Required: No

**serviceLinkVirtualInterfaceId**

The ID of the service link virtual interface.

Type: String

Required: No

**TagSet.N**

The tags associated with the service link virtual interface.

Type: Array of [Tag](api-tag.md) objects

Required: No

**vlan**

The virtual local area network for the service link virtual interface.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/servicelinkvirtualinterface.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/servicelinkvirtualinterface.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/servicelinkvirtualinterface.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ServiceDetail

ServiceTypeDetail
