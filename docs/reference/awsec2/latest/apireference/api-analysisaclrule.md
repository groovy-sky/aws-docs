# AnalysisAclRule

Describes a network access control (ACL) rule.

## Contents

**cidr**

The IPv4 address range, in CIDR notation.

Type: String

Required: No

**egress**

Indicates whether the rule is an outbound rule.

Type: Boolean

Required: No

**portRange**

The range of ports.

Type: [PortRange](api-portrange.md) object

Required: No

**protocol**

The protocol.

Type: String

Required: No

**ruleAction**

Indicates whether to allow or deny traffic that matches the rule.

Type: String

Required: No

**ruleNumber**

The rule number.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/analysisaclrule.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/analysisaclrule.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/analysisaclrule.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AlternatePathHint

AnalysisComponent
