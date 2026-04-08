# EnableInsightRules

Enables the specified Contributor Insights rules. When rules are enabled, they
immediately begin analyzing log data.

## Request Parameters

**RuleNames**

An array of the rule names to enable. If you need to find out the names of your rules,
use [DescribeInsightRules](api-describeinsightrules.md).

Type: Array of strings

Length Constraints: Minimum length of 1. Maximum length of 128.

Pattern: `[\x20-\x7E]+`

Required: Yes

## Response Elements

The following element is returned by the service.

**Failures**

An array listing the rules that could not be enabled. You cannot disable or enable
built-in rules.

Type: Array of [PartialFailure](api-partialfailure.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterValue**

The value of an input parameter is bad or out-of-range.

**message**

HTTP Status Code: 400

**LimitExceededException**

The operation exceeded one or more limits.

HTTP Status Code: 400

**MissingParameter**

An input parameter that is required is missing.

**message**

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/monitoring-2010-08-01/enableinsightrules.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/monitoring-2010-08-01/enableinsightrules.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/monitoring-2010-08-01/enableinsightrules.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/monitoring-2010-08-01/enableinsightrules.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/monitoring-2010-08-01/enableinsightrules.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/monitoring-2010-08-01/enableinsightrules.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/monitoring-2010-08-01/enableinsightrules.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/monitoring-2010-08-01/enableinsightrules.md)

- [AWS SDK for Python](../../../../services/goto/boto3/monitoring-2010-08-01/enableinsightrules.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/monitoring-2010-08-01/enableinsightrules.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

EnableAlarmActions

GetAlarmMuteRule
