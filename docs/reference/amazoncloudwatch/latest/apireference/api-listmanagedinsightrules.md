# ListManagedInsightRules

Returns a list that contains the number of managed Contributor Insights rules in your
account.

## Request Parameters

**MaxResults**

The maximum number of results to return in one operation. If you omit this parameter,
the default number is used. The default number is `100`.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 500.

Required: No

**NextToken**

Include this value to get the next set of rules if the value was returned by the
previous operation.

Type: String

Required: No

**ResourceARN**

The ARN of an AWS resource that has managed Contributor Insights
rules.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: Yes

## Response Elements

The following elements are returned by the service.

**ManagedRules**

The managed rules that are available for the specified AWS resource.

Type: Array of [ManagedRuleDescription](api-managedruledescription.md) objects

**NextToken**

Include this value to get the next set of rules if the value was returned by the
previous operation.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidNextToken**

The next token specified is invalid.

**message**

HTTP Status Code: 400

**InvalidParameterValue**

The value of an input parameter is bad or out-of-range.

**message**

HTTP Status Code: 400

**MissingParameter**

An input parameter that is required is missing.

**message**

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/monitoring-2010-08-01/listmanagedinsightrules.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/monitoring-2010-08-01/listmanagedinsightrules.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/monitoring-2010-08-01/listmanagedinsightrules.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/monitoring-2010-08-01/listmanagedinsightrules.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/monitoring-2010-08-01/listmanagedinsightrules.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/monitoring-2010-08-01/listmanagedinsightrules.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/monitoring-2010-08-01/listmanagedinsightrules.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/monitoring-2010-08-01/listmanagedinsightrules.md)

- [AWS SDK for Python](../../../../services/goto/boto3/monitoring-2010-08-01/listmanagedinsightrules.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/monitoring-2010-08-01/listmanagedinsightrules.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListDashboards

ListMetrics

All content copied from https://docs.aws.amazon.com/.
