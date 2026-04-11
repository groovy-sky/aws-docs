# TransitGatewayConnect

Describes a transit gateway Connect attachment.

## Contents

**creationTime**

The creation time.

Type: Timestamp

Required: No

**options**

The Connect attachment options.

Type: [TransitGatewayConnectOptions](api-transitgatewayconnectoptions.md) object

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

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/transitgatewayconnect.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/transitgatewayconnect.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/transitgatewayconnect.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

TransitGatewayAttachmentPropagation

TransitGatewayConnectOptions
