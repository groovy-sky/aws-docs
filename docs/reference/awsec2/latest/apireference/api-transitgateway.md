# TransitGateway

Describes a transit gateway.

## Contents

**creationTime**

The creation time.

Type: Timestamp

Required: No

**description**

The description of the transit gateway.

Type: String

Required: No

**options**

The transit gateway options.

Type: [TransitGatewayOptions](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_TransitGatewayOptions.html) object

Required: No

**ownerId**

The ID of the AWS account that owns the transit gateway.

Type: String

Required: No

**state**

The state of the transit gateway.

Type: String

Valid Values: `pending | available | modifying | deleting | deleted`

Required: No

**TagSet.N**

The tags for the transit gateway.

Type: Array of [Tag](api-tag.md) objects

Required: No

**transitGatewayArn**

The Amazon Resource Name (ARN) of the transit gateway.

Type: String

Required: No

**transitGatewayId**

The ID of the transit gateway.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/TransitGateway)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/TransitGateway)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/TransitGateway)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TrafficMirrorTarget

TransitGatewayAssociation
