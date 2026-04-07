# InternetGatewayAttachment

Describes the attachment of a VPC to an internet gateway or an egress-only internet gateway.

## Contents

**state**

The current state of the attachment. For an internet gateway, the state is
`available` when attached to a VPC; otherwise, this value is not
returned.

Type: String

Valid Values: `attaching | attached | detaching | detached`

Required: No

**vpcId**

The ID of the VPC.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/internetgatewayattachment.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/internetgatewayattachment.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/internetgatewayattachment.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

InternetGateway

InterruptibleCapacityAllocation
