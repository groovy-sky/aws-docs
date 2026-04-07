# TransitGatewayPolicyRule

Describes a rule associated with a transit gateway policy.

## Contents

**destinationCidrBlock**

The destination CIDR block for the transit gateway policy rule.

Type: String

Required: No

**destinationPortRange**

The port range for the transit gateway policy rule. Currently this is set to \* (all).

Type: String

Required: No

**metaData**

The meta data tags used for the transit gateway policy rule.

Type: [TransitGatewayPolicyRuleMetaData](api-transitgatewaypolicyrulemetadata.md) object

Required: No

**protocol**

The protocol used by the transit gateway policy rule.

Type: String

Required: No

**sourceCidrBlock**

The source CIDR block for the transit gateway policy rule.

Type: String

Required: No

**sourcePortRange**

The port range for the transit gateway policy rule. Currently this is set to \* (all).

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/transitgatewaypolicyrule.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/transitgatewaypolicyrule.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/transitgatewaypolicyrule.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

TransitGatewayPeeringAttachmentOptions

TransitGatewayPolicyRuleMetaData
