# AdditionalDetail

Describes an additional detail for a path analysis. For more information, see [Reachability Analyzer additional detail codes](../../../../services/vpc/latest/reachability/additional-detail-codes.md).

## Contents

**additionalDetailType**

The additional detail code.

Type: String

Required: No

**component**

The path component.

Type: [AnalysisComponent](api-analysiscomponent.md) object

Required: No

**LoadBalancerSet.N**

The load balancers.

Type: Array of [AnalysisComponent](api-analysiscomponent.md) objects

Required: No

**RuleGroupRuleOptionsPairSet.N**

The rule options.

Type: Array of [RuleGroupRuleOptionsPair](api-rulegroupruleoptionspair.md) objects

Required: No

**RuleGroupTypePairSet.N**

The rule group type.

Type: Array of [RuleGroupTypePair](api-rulegrouptypepair.md) objects

Required: No

**RuleOptionSet.N**

The rule options.

Type: Array of [RuleOption](api-ruleoption.md) objects

Required: No

**serviceName**

The name of the VPC endpoint service.

Type: String

Required: No

**vpcEndpointService**

The VPC endpoint service.

Type: [AnalysisComponent](api-analysiscomponent.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/additionaldetail.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/additionaldetail.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/additionaldetail.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AddIpamOrganizationalUnitExclusion

AddPrefixListEntry
