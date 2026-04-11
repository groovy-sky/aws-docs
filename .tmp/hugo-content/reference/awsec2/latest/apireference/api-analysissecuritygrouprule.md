# AnalysisSecurityGroupRule

Describes a security group rule.

## Contents

**cidr**

The IPv4 address range, in CIDR notation.

Type: String

Required: No

**direction**

The direction. The following are the possible values:

- egress

- ingress

Type: String

Required: No

**portRange**

The port range.

Type: [PortRange](api-portrange.md) object

Required: No

**prefixListId**

The prefix list ID.

Type: String

Required: No

**protocol**

The protocol name.

Type: String

Required: No

**securityGroupId**

The security group ID.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/analysissecuritygrouprule.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/analysissecuritygrouprule.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/analysissecuritygrouprule.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AnalysisRouteTableRoute

AsnAssociation
