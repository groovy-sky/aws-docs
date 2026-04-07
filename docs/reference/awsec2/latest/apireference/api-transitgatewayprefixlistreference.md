# TransitGatewayPrefixListReference

Describes a prefix list reference.

## Contents

**blackhole**

Indicates whether traffic that matches this route is dropped.

Type: Boolean

Required: No

**prefixListId**

The ID of the prefix list.

Type: String

Required: No

**prefixListOwnerId**

The ID of the prefix list owner.

Type: String

Required: No

**state**

The state of the prefix list reference.

Type: String

Valid Values: `pending | available | modifying | deleting`

Required: No

**transitGatewayAttachment**

Information about the transit gateway attachment.

Type: [TransitGatewayPrefixListAttachment](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_TransitGatewayPrefixListAttachment.html) object

Required: No

**transitGatewayRouteTableId**

The ID of the transit gateway route table.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/TransitGatewayPrefixListReference)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/TransitGatewayPrefixListReference)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/TransitGatewayPrefixListReference)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TransitGatewayPrefixListAttachment

TransitGatewayPropagation
