# IpamPolicyDocument

Information about an IPAM policy.

## Contents

**AllocationRuleSet.N**

The allocation rules in the IPAM policy document.

Allocation rules are optional configurations within an IPAM policy that map AWS resource types to specific IPAM pools. If no rules are defined, the resource types default to using Amazon-provided IP addresses.

Type: Array of [IpamPolicyAllocationRule](api-ipampolicyallocationrule.md) objects

Required: No

**ipamPolicyId**

The ID of the IPAM policy.

Type: String

Required: No

**locale**

The locale of the IPAM policy document.

Type: String

Required: No

**resourceType**

The resource type of the IPAM policy document.

The AWS service or resource type that can use IP addresses through IPAM policies. Supported services and resource types include:

- Elastic IP addresses

Type: String

Valid Values: `alb | eip | rds | rnat`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/ipampolicydocument.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/ipampolicydocument.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/ipampolicydocument.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

IpamPolicyAllocationRuleRequest

IpamPolicyOrganizationTarget
