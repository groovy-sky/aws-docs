# FirewallStatelessRule

Describes a stateless rule.

## Contents

**DestinationPortSet.N**

The destination ports.

Type: Array of [PortRange](api-portrange.md) objects

Required: No

**DestinationSet.N**

The destination IP addresses, in CIDR notation.

Type: Array of strings

Required: No

**priority**

The rule priority.

Type: Integer

Valid Range: Minimum value of -1. Maximum value of 65535.

Required: No

**ProtocolSet.N**

The protocols.

Type: Array of integers

Valid Range: Minimum value of 0. Maximum value of 255.

Required: No

**ruleAction**

The rule action. The possible values are `pass`, `drop`, and
`forward_to_site`.

Type: String

Required: No

**ruleGroupArn**

The ARN of the stateless rule group.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1283.

Required: No

**SourcePortSet.N**

The source ports.

Type: Array of [PortRange](api-portrange.md) objects

Required: No

**SourceSet.N**

The source IP addresses, in CIDR notation.

Type: Array of strings

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/firewallstatelessrule.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/firewallstatelessrule.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/firewallstatelessrule.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

FirewallStatefulRule

FleetBlockDeviceMappingRequest
