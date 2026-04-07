# TransitGatewayVpcAttachment

Describes a VPC attachment.

## Contents

**creationTime**

The creation time.

Type: Timestamp

Required: No

**options**

The VPC attachment options.

Type: [TransitGatewayVpcAttachmentOptions](api-transitgatewayvpcattachmentoptions.md) object

Required: No

**state**

The state of the VPC attachment. Note that the `initiating` state has been deprecated.

Type: String

Valid Values: `initiating | initiatingRequest | pendingAcceptance | rollingBack | pending | available | modifying | deleting | deleted | failed | rejected | rejecting | failing`

Required: No

**SubnetIds.N**

The IDs of the subnets.

Type: Array of strings

Required: No

**TagSet.N**

The tags for the VPC attachment.

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

**vpcId**

The ID of the VPC.

Type: String

Required: No

**vpcOwnerId**

The ID of the AWS account that owns the VPC.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/transitgatewayvpcattachment.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/transitgatewayvpcattachment.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/transitgatewayvpcattachment.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

TransitGatewayRouteTableRoute

TransitGatewayVpcAttachmentOptions
