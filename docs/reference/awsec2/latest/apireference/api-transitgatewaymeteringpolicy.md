# TransitGatewayMeteringPolicy

Describes a transit gateway metering policy.

## Contents

**MiddleboxAttachmentIdSet.N**

The IDs of the middlebox attachments associated with the metering policy.

Type: Array of strings

Required: No

**state**

The state of the transit gateway metering policy.

Type: String

Valid Values: `available | deleted | pending | modifying | deleting`

Required: No

**TagSet.N**

The tags assigned to the transit gateway metering policy.

Type: Array of [Tag](api-tag.md) objects

Required: No

**transitGatewayId**

The ID of the transit gateway associated with the metering policy.

Type: String

Required: No

**transitGatewayMeteringPolicyId**

The ID of the transit gateway metering policy.

Type: String

Required: No

**updateEffectiveAt**

The date and time when the metering policy update becomes effective.

Type: Timestamp

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/TransitGatewayMeteringPolicy)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/TransitGatewayMeteringPolicy)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/TransitGatewayMeteringPolicy)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TransitGatewayConnectRequestBgpOptions

TransitGatewayMeteringPolicyEntry
