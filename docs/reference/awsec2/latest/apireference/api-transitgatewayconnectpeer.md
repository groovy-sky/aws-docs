# TransitGatewayConnectPeer

Describes a transit gateway Connect peer.

## Contents

**connectPeerConfiguration**

The Connect peer details.

Type: [TransitGatewayConnectPeerConfiguration](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_TransitGatewayConnectPeerConfiguration.html) object

Required: No

**creationTime**

The creation time.

Type: Timestamp

Required: No

**state**

The state of the Connect peer.

Type: String

Valid Values: `pending | available | deleting | deleted`

Required: No

**TagSet.N**

The tags for the Connect peer.

Type: Array of [Tag](api-tag.md) objects

Required: No

**transitGatewayAttachmentId**

The ID of the Connect attachment.

Type: String

Required: No

**transitGatewayConnectPeerId**

The ID of the Connect peer.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/TransitGatewayConnectPeer)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/TransitGatewayConnectPeer)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/TransitGatewayConnectPeer)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TransitGatewayConnectOptions

TransitGatewayConnectPeerConfiguration
