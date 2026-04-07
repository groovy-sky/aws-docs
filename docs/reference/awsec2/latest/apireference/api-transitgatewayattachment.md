# TransitGatewayAttachment

Describes an attachment between a resource and a transit gateway.

## Contents

**association**

The association.

Type: [TransitGatewayAttachmentAssociation](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_TransitGatewayAttachmentAssociation.html) object

Required: No

**creationTime**

The creation time.

Type: Timestamp

Required: No

**resourceId**

The ID of the resource.

Type: String

Required: No

**resourceOwnerId**

The ID of the AWS account that owns the resource.

Type: String

Required: No

**resourceType**

The resource type. Note that the `tgw-peering` resource type has been deprecated.

Type: String

Valid Values: `vpc | vpn | vpn-concentrator | direct-connect-gateway | connect | peering | tgw-peering`

Required: No

**state**

The attachment state. Note that the `initiating` state has been deprecated.

Type: String

Valid Values: `initiating | initiatingRequest | pendingAcceptance | rollingBack | pending | available | modifying | deleting | deleted | failed | rejected | rejecting | failing`

Required: No

**TagSet.N**

The tags for the attachment.

Type: Array of [Tag](api-tag.md) objects

Required: No

**transitGatewayAttachmentId**

The ID of the attachment.

Type: String

Required: No

**transitGatewayId**

The ID of the transit gateway.

Type: String

Required: No

**transitGatewayOwnerId**

The ID of the AWS account that owns the transit gateway.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/TransitGatewayAttachment)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/TransitGatewayAttachment)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/TransitGatewayAttachment)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TransitGatewayAssociation

TransitGatewayAttachmentAssociation
