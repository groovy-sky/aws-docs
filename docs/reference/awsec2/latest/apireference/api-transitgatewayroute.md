# TransitGatewayRoute

Describes a route for a transit gateway route table.

## Contents

**destinationCidrBlock**

The CIDR block used for destination matches.

Type: String

Required: No

**prefixListId**

The ID of the prefix list used for destination matches.

Type: String

Required: No

**state**

The state of the route.

Type: String

Valid Values: `pending | active | blackhole | deleting | deleted`

Required: No

**TransitGatewayAttachments.N**

The attachments.

Type: Array of [TransitGatewayRouteAttachment](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_TransitGatewayRouteAttachment.html) objects

Required: No

**transitGatewayRouteTableAnnouncementId**

The ID of the transit gateway route table announcement.

Type: String

Required: No

**type**

The route type.

Type: String

Valid Values: `static | propagated`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/TransitGatewayRoute)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/TransitGatewayRoute)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/TransitGatewayRoute)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TransitGatewayRequestOptions

TransitGatewayRouteAttachment
