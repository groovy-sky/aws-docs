# ListResourceScans

List the resource scans from newest to oldest. By default it will return up to 10 resource
scans.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**MaxResults**

If the number of available results exceeds this maximum, the response includes a
`NextToken` value that you can use for the `NextToken` parameter to
get the next set of results. The default value is 10. The maximum value is 100.

Type: Integer

Required: No

**NextToken**

The token for the next set of items to return. (You received this token from a previous
call.)

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

**ScanTypeFilter**

The scan type that you want to get summary information about. The default is
`FULL`.

Type: String

Valid Values: `FULL | PARTIAL`

Required: No

## Response Elements

The following elements are returned by the service.

**NextToken**

If the request doesn't return all the remaining results, `NextToken` is set to
a token. To retrieve the next set of results, call `ListResourceScans` again and
use that value for the `NextToken` parameter. If the request returns all results,
`NextToken` is set to an empty string.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

**ResourceScanSummaries.member.N**

The list of scans returned.

Type: Array of [ResourceScanSummary](api-resourcescansummary.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/cloudformation-2010-05-15/listresourcescans.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/cloudformation-2010-05-15/listresourcescans.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/listresourcescans.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/cloudformation-2010-05-15/listresourcescans.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/listresourcescans.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/cloudformation-2010-05-15/listresourcescans.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/cloudformation-2010-05-15/listresourcescans.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/cloudformation-2010-05-15/listresourcescans.md)

- [AWS SDK for Python](../../../../services/goto/boto3/cloudformation-2010-05-15/listresourcescans.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/listresourcescans.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListResourceScanResources

ListStackInstanceResourceDrifts

All content copied from https://docs.aws.amazon.com/.
