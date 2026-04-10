# StartResourceScan

Starts a scan of the resources in this account in this Region. You can the status of a
scan using the `ListResourceScans` API action.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**ClientRequestToken**

A unique identifier for this `StartResourceScan` request. Specify this token if
you plan to retry requests so that CloudFormation knows that you're not attempting to start a new
resource scan.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Pattern: `[a-zA-Z0-9][-a-zA-Z0-9]*`

Required: No

**ScanFilters.member.N**

The scan filters to use.

Type: Array of [ScanFilter](api-scanfilter.md) objects

Array Members: Fixed number of 1 item.

Required: No

## Response Elements

The following element is returned by the service.

**ResourceScanId**

The Amazon Resource Name (ARN) of the resource scan. The format is
`arn:${Partition}:cloudformation:${Region}:${Account}:resourceScan/${Id}`. An
example is
`arn:aws:cloudformation:us-east-1:123456789012:resourceScan/f5b490f7-7ed4-428a-aa06-31ff25db0772
                  `.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ResourceScanInProgress**

A resource scan is currently in progress. Only one can be run at a time for an account in a Region.

HTTP Status Code: 400

**ResourceScanLimitExceeded**

The limit on resource scans has been exceeded. Reasons include:

- Exceeded the daily quota for resource scans.

- A resource scan recently failed. You must wait 10 minutes before starting a new resource scan.

- The last resource scan failed after exceeding 100,000 resources. When this happens, you must wait 24 hours
before starting a new resource scan.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/cloudformation-2010-05-15/startresourcescan.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/cloudformation-2010-05-15/startresourcescan.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/startresourcescan.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/cloudformation-2010-05-15/startresourcescan.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/startresourcescan.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/cloudformation-2010-05-15/startresourcescan.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/cloudformation-2010-05-15/startresourcescan.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/cloudformation-2010-05-15/startresourcescan.md)

- [AWS SDK for Python](../../../../services/goto/boto3/cloudformation-2010-05-15/startresourcescan.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/startresourcescan.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SignalResource

StopStackSetOperation

All content copied from https://docs.aws.amazon.com/.
