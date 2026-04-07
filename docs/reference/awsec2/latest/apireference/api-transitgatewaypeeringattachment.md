# TransitGatewayPeeringAttachment

Describes the transit gateway peering attachment.

## Contents

**accepterTgwInfo**

Information about the accepter transit gateway.

Type: [PeeringTgwInfo](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_PeeringTgwInfo.html) object

Required: No

**accepterTransitGatewayAttachmentId**

The ID of the accepter transit gateway attachment.

Type: String

Required: No

**creationTime**

The time the transit gateway peering attachment was created.

Type: Timestamp

Required: No

**options**

Details about the transit gateway peering attachment.

Type: [TransitGatewayPeeringAttachmentOptions](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_TransitGatewayPeeringAttachmentOptions.html) object

Required: No

**requesterTgwInfo**

Information about the requester transit gateway.

Type: [PeeringTgwInfo](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_PeeringTgwInfo.html) object

Required: No

**state**

The state of the transit gateway peering attachment. Note that the `initiating` state has been deprecated.

Type: String

Valid Values: `initiating | initiatingRequest | pendingAcceptance | rollingBack | pending | available | modifying | deleting | deleted | failed | rejected | rejecting | failing`

Required: No

**status**

The status of the transit gateway peering attachment.

Type: [PeeringAttachmentStatus](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_PeeringAttachmentStatus.html) object

Required: No

**TagSet.N**

The tags for the transit gateway peering attachment.

Type: Array of [Tag](api-tag.md) objects

Required: No

**transitGatewayAttachmentId**

The ID of the transit gateway peering attachment.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/TransitGatewayPeeringAttachment)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/TransitGatewayPeeringAttachment)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/TransitGatewayPeeringAttachment)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TransitGatewayOptions

TransitGatewayPeeringAttachmentOptions
