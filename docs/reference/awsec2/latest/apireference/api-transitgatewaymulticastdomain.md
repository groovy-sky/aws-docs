# TransitGatewayMulticastDomain

Describes the transit gateway multicast domain.

## Contents

**creationTime**

The time the transit gateway multicast domain was created.

Type: Timestamp

Required: No

**options**

The options for the transit gateway multicast domain.

Type: [TransitGatewayMulticastDomainOptions](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_TransitGatewayMulticastDomainOptions.html) object

Required: No

**ownerId**

The ID of the AWS account that owns the transit gateway multicast domain.

Type: String

Required: No

**state**

The state of the transit gateway multicast domain.

Type: String

Valid Values: `pending | available | deleting | deleted`

Required: No

**TagSet.N**

The tags for the transit gateway multicast domain.

Type: Array of [Tag](api-tag.md) objects

Required: No

**transitGatewayId**

The ID of the transit gateway.

Type: String

Required: No

**transitGatewayMulticastDomainArn**

The Amazon Resource Name (ARN) of the transit gateway multicast domain.

Type: String

Required: No

**transitGatewayMulticastDomainId**

The ID of the transit gateway multicast domain.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/TransitGatewayMulticastDomain)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/TransitGatewayMulticastDomain)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/TransitGatewayMulticastDomain)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TransitGatewayMulticastDeregisteredGroupSources

TransitGatewayMulticastDomainAssociation
