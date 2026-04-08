# InsightRule

This structure contains the definition for a Contributor Insights rule. For more
information about this rule, see [Using Constributor Insights to analyze high-cardinality data](../../../../services/amazoncloudwatch/latest/monitoring/contributorinsights.md) in the
_Amazon CloudWatch User Guide_.

## Contents

**Definition**

The definition of the rule, as a JSON object. The definition contains the keywords
used to define contributors, the value to aggregate on if this rule returns a sum
instead of a count, and the filters. For details on the valid syntax, see [Contributor Insights Rule Syntax](../../../../services/amazoncloudwatch/latest/monitoring/contributorinsights-rulesyntax.md).

Type: String

Length Constraints: Minimum length of 1. Maximum length of 8192.

Pattern: `[\x00-\x7F]+`

Required: Yes

**Name**

The name of the rule.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Pattern: `[\x20-\x7E]+`

Required: Yes

**Schema**

For rules that you create, this is always `{"Name": "CloudWatchLogRule",
            "Version": 1}`. For managed rules, this is `{"Name": "ServiceLogRule",
                "Version": 1}`

Type: String

Required: Yes

**State**

Indicates whether the rule is enabled or disabled.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 32.

Pattern: `[\x20-\x7E]+`

Required: Yes

**ApplyOnTransformedLogs**

Displays whether the rule is evaluated on the transformed versions of logs, for log groups
that have [Log transformation](../../../../services/amazoncloudwatch/latest/logs/cloudwatch-logs-transformation.md) enabled. If this is `false`, log events are evaluated before they are transformed.

Type: Boolean

Required: No

**ManagedRule**

An optional built-in rule that AWS manages.

Type: Boolean

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/monitoring-2010-08-01/insightrule.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/monitoring-2010-08-01/insightrule.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/monitoring-2010-08-01/insightrule.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EvaluationCriteria

InsightRuleContributor

All content copied from https://docs.aws.amazon.com/.
