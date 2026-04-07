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

Type: [TransitGatewayOptions](api-transitgatewayoptions.md) object

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

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/transitgateway.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/transitgateway.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/transitgateway.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

TrafficMirrorTarget

TransitGatewayAssociation
