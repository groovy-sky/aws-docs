# IpamPolicyAllocationRule

Information about an IPAM policy allocation rule.

Allocation rules are optional configurations within an IPAM policy that map AWS resource types to specific IPAM pools. If no rules are defined, the resource types default to using Amazon-provided IP addresses.

## Contents

**sourceIpamPoolId**

The ID of the source IPAM pool for the allocation rule.

An IPAM pool is a collection of IP addresses in IPAM that can be allocated to AWS resources.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/ipampolicyallocationrule.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/ipampolicyallocationrule.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/ipampolicyallocationrule.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

IpamPolicy

IpamPolicyAllocationRuleRequest
