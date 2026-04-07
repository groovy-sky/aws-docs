# RevokedSecurityGroupRule

A security group rule removed with [RevokeSecurityGroupEgress](api-revokesecuritygroupegress.md) or [RevokeSecurityGroupIngress](api-revokesecuritygroupingress.md).

## Contents

**cidrIpv4**

The IPv4 CIDR of the traffic source.

Type: String

Required: No

**cidrIpv6**

The IPv6 CIDR of the traffic source.

Type: String

Required: No

**description**

A description of the revoked security group rule.

Type: String

Required: No

**fromPort**

The 'from' port number of the security group rule.

Type: Integer

Required: No

**groupId**

A security group ID.

Type: String

Required: No

**ipProtocol**

The security group rule's protocol.

Type: String

Required: No

**isEgress**

Defines if a security group rule is an outbound rule.

Type: Boolean

Required: No

**prefixListId**

The ID of a prefix list that's the traffic source.

Type: String

Required: No

**referencedGroupId**

The ID of a referenced security group.

Type: String

Required: No

**securityGroupRuleId**

A security group rule ID.

Type: String

Required: No

**toPort**

The 'to' port number of the security group rule.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/revokedsecuritygrouprule.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/revokedsecuritygrouprule.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/revokedsecuritygrouprule.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ResponseLaunchTemplateData

Route
