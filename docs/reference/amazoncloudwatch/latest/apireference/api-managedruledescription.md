# ManagedRuleDescription

Contains information about managed Contributor Insights rules, as returned by
`ListManagedInsightRules`.

## Contents

**ResourceARN**

If a managed rule is enabled, this is the ARN for the related AWS
resource.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

**RuleState**

Describes the state of a managed rule. If present, it contains information about the
Contributor Insights rule that contains information about the related AWS
resource.

Type: [ManagedRuleState](api-managedrulestate.md) object

Required: No

**TemplateName**

The template name for the managed rule. Used to enable managed rules using
`PutManagedInsightRules`.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Pattern: `[0-9A-Za-z][\-\.\_0-9A-Za-z]{0,126}[0-9A-Za-z]`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/monitoring-2010-08-01/managedruledescription.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/monitoring-2010-08-01/managedruledescription.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/monitoring-2010-08-01/managedruledescription.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ManagedRule

ManagedRuleState

All content copied from https://docs.aws.amazon.com/.
