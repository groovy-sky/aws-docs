# ResourceScanSummary

A summary of the resource scan. This is returned by the `ListResourceScan` API
action.

## Contents

**EndTime**

The time that the resource scan was finished.

Type: Timestamp

Required: No

**PercentageCompleted**

The percentage of the resource scan that has been completed.

Type: Double

Required: No

**ResourceScanId**

The Amazon Resource Name (ARN) of the resource scan.

Type: String

Required: No

**ScanType**

The scan type that has been completed.

Type: String

Valid Values: `FULL | PARTIAL`

Required: No

**StartTime**

The time that the resource scan was started.

Type: Timestamp

Required: No

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

Required: No

**StatusReason**

The reason for the resource scan status, providing more information if a failure
happened.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/resourcescansummary.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/resourcescansummary.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/resourcescansummary.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResourceMapping

ResourceTargetDefinition

All content copied from https://docs.aws.amazon.com/.
