# TransitGatewayMeteringPolicyEntry

Describes an entry in a transit gateway metering policy.

## Contents

**meteredAccount**

The AWS account ID to which the metered traffic is attributed.

Type: String

Valid Values: `source-attachment-owner | destination-attachment-owner | transit-gateway-owner`

Required: No

**meteringPolicyRule**

The metering policy rule that defines traffic matching criteria.

Type: [TransitGatewayMeteringPolicyRule](api-transitgatewaymeteringpolicyrule.md) object

Required: No

**policyRuleNumber**

The rule number of the metering policy entry.

Type: String

Required: No

**state**

The state of the metering policy entry.

Type: String

Valid Values: `available | deleted`

Required: No

**updatedAt**

The date and time when the metering policy entry was last updated.

Type: Timestamp

Required: No

**updateEffectiveAt**

The date and time when the metering policy entry update becomes effective.

Type: Timestamp

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/transitgatewaymeteringpolicyentry.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/transitgatewaymeteringpolicyentry.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/transitgatewaymeteringpolicyentry.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

TransitGatewayMeteringPolicy

TransitGatewayMeteringPolicyRule
