# PutManagedInsightRules

Creates a managed Contributor Insights rule for a specified AWS
resource. When you enable a managed rule, you create a Contributor Insights rule that
collects data from AWS services. You cannot edit these rules with
`PutInsightRule`. The rules can be enabled, disabled, and deleted using
`EnableInsightRules`, `DisableInsightRules`, and
`DeleteInsightRules`. If a previously created managed rule is currently
disabled, a subsequent call to this API will re-enable it. Use
`ListManagedInsightRules` to describe all available rules.

## Request Parameters

**ManagedRules**

A list of `ManagedRules` to enable.

Type: Array of [ManagedRule](api-managedrule.md) objects

Required: Yes

## Response Elements

The following element is returned by the service.

**Failures**

An array that lists the rules that could not be enabled.

Type: Array of [PartialFailure](api-partialfailure.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/monitoring-2010-08-01/putmanagedinsightrules.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/monitoring-2010-08-01/putmanagedinsightrules.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/monitoring-2010-08-01/putmanagedinsightrules.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/monitoring-2010-08-01/putmanagedinsightrules.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/monitoring-2010-08-01/putmanagedinsightrules.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/monitoring-2010-08-01/putmanagedinsightrules.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/monitoring-2010-08-01/putmanagedinsightrules.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/monitoring-2010-08-01/putmanagedinsightrules.md)

- [AWS SDK for Python](../../../../services/goto/boto3/monitoring-2010-08-01/putmanagedinsightrules.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/monitoring-2010-08-01/putmanagedinsightrules.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PutInsightRule

PutMetricAlarm

All content copied from https://docs.aws.amazon.com/.
