# DescribeResourceScan

Describes details of a resource scan.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**ResourceScanId**

The Amazon Resource Name (ARN) of the resource scan.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**EndTime**

The time that the resource scan was finished.

Type: Timestamp

**PercentageCompleted**

The percentage of the resource scan that has been completed.

Type: Double

**ResourceScanId**

The Amazon Resource Name (ARN) of the resource scan. The format is
`arn:${Partition}:cloudformation:${Region}:${Account}:resourceScan/${Id}`. An
example is
`arn:aws:cloudformation:us-east-1:123456789012:resourceScan/f5b490f7-7ed4-428a-aa06-31ff25db0772
                  `.

Type: String

**ResourcesRead**

The number of resources that were read. This is only available for scans with a
`Status` set to `COMPLETE`, `EXPIRED`, or
`FAILED`.

###### Note

This field may be 0 if the resource scan failed with a
`ResourceScanLimitExceededException`.

Type: Integer

**ResourcesScanned**

The number of resources that were listed. This is only available for scans with a
`Status` set to `COMPLETE`, `EXPIRED`, or `FAILED
      `.

Type: Integer

**ResourceTypes.member.N**

The list of resource types for the specified scan. Resource types are only available for
scans with a `Status` set to `COMPLETE` or `FAILED `.

Type: Array of strings

Length Constraints: Minimum length of 1. Maximum length of 256.

**ScanFilters.member.N**

The scan filters that were used.

Type: Array of [ScanFilter](api-scanfilter.md) objects

Array Members: Fixed number of 1 item.

**StartTime**

The time that the resource scan was started.

Type: Timestamp

**Status**

Status of the resource scan.

**IN\_PROGRESS**

The resource scan is still in progress.

**COMPLETE**

The resource scan is complete.

**EXPIRED**

The resource scan has expired.

**FAILED**

The resource scan has failed.

Type: String

Valid Values: `IN_PROGRESS | FAILED | COMPLETE | EXPIRED`

**StatusReason**

The reason for the resource scan status, providing more information if a failure
happened.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ResourceScanNotFound**

The resource scan was not found.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/cloudformation-2010-05-15/describeresourcescan.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/cloudformation-2010-05-15/describeresourcescan.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/describeresourcescan.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/cloudformation-2010-05-15/describeresourcescan.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/describeresourcescan.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/cloudformation-2010-05-15/describeresourcescan.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/cloudformation-2010-05-15/describeresourcescan.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/cloudformation-2010-05-15/describeresourcescan.md)

- [AWS SDK for Python](../../../../services/goto/boto3/cloudformation-2010-05-15/describeresourcescan.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/describeresourcescan.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribePublisher

DescribeStackDriftDetectionStatus
