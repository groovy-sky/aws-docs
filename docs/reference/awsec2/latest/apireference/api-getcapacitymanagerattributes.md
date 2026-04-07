# GetCapacityManagerAttributes

Retrieves the current configuration and status of EC2 Capacity Manager for your account, including enablement status, Organizations access settings, and data ingestion status.

## Request Parameters

For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.
If you have the required permissions, the error response is `DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

## Response Elements

The following elements are returned by the service.

**capacityManagerStatus**

The current status of Capacity Manager.

Type: String

Valid Values: `enabled | disabled`

**dataExportCount**

The number of active data export configurations for this account. This count includes all data exports regardless of their current delivery status.

Type: Integer

**earliestDatapointTimestamp**

The timestamp of the earliest data point available in Capacity Manager, in milliseconds since epoch. This indicates how far back historical data is available for queries.

Type: Timestamp

**ingestionStatus**

The current data ingestion status. Initial ingestion may take several hours after enabling Capacity Manager.

Type: String

Valid Values: `initial-ingestion-in-progress | ingestion-complete | ingestion-failed`

**ingestionStatusMessage**

A descriptive message providing additional details about the current ingestion status. This may include error information if ingestion has
failed or progress details during initial setup.

Type: String

**latestDatapointTimestamp**

The timestamp of the most recent data point ingested by Capacity Manager, in milliseconds since epoch. This indicates how current your capacity data is.

Type: Timestamp

**organizationsAccess**

Indicates whether Organizations access is enabled for cross-account data aggregation.

Type: Boolean

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/GetCapacityManagerAttributes)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/GetCapacityManagerAttributes)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/getcapacitymanagerattributes.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/getcapacitymanagerattributes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/getcapacitymanagerattributes.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/getcapacitymanagerattributes.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/getcapacitymanagerattributes.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/getcapacitymanagerattributes.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/GetCapacityManagerAttributes)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/getcapacitymanagerattributes.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetAwsNetworkPerformanceData

GetCapacityManagerMetricData
