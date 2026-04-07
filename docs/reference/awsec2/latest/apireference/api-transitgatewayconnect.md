# TransitGatewayConnect

Describes a transit gateway Connect attachment.

## Contents

**creationTime**

The creation time.

Type: Timestamp

Required: No

**options**

The Connect attachment options.

Type: [TransitGatewayConnectOptions](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_TransitGatewayConnectOptions.html) object

Required: No

**state**

The state of the attachment.

Type: String

Valid Values: `initiating | initiatingRequest | pendingAcceptance | rollingBack | pending | available | modifying | deleting | deleted | failed | rejected | rejecting | failing`

Required: No

**TagSet.N**

The tags for the attachment.

Type: Array of [Tag](api-tag.md) objects

Required: No

**transitGatewayAttachmentId**

The ID of the Connect attachment.

Type: String

Required: No

**transitGatewayId**

The ID of the transit gateway.

Type: String

Required: No

**transportTransitGatewayAttachmentId**

The ID of the attachment from which the Connect attachment was created.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/TransitGatewayConnect)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/TransitGatewayConnect)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/TransitGatewayConnect)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TransitGatewayAttachmentPropagation

TransitGatewayConnectOptions
