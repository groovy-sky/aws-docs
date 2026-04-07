# CreateCapacityManagerDataExport

Creates a new data export configuration for EC2 Capacity Manager. This allows you to automatically export capacity usage data to an S3 bucket on a scheduled basis.
The exported data includes metrics for On-Demand, Spot, and Capacity Reservations usage across your organization.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ClientToken**

Unique, case-sensitive identifier that you provide to ensure the idempotency of the request. For more information, see Ensure Idempotency.

Type: String

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.
If you have the required permissions, the error response is `DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**OutputFormat**

The file format for the exported data. Parquet format is recommended for large datasets and better compression.

Type: String

Valid Values: `csv | parquet`

Required: Yes

**S3BucketName**

The name of the S3 bucket where the capacity data export files will be delivered. The bucket must exist and you must have write permissions to it.

Type: String

Required: Yes

**S3BucketPrefix**

The S3 key prefix for the exported data files. This allows you to organize exports in a specific folder structure within your bucket. If not specified, files are placed at the bucket root.

Type: String

Required: No

**Schedule**

The frequency at which data exports are generated.

Type: String

Valid Values: `hourly`

Required: Yes

**TagSpecification.N**

The tags to apply to the data export configuration. You can tag the export for organization and cost tracking purposes.

Type: Array of [TagSpecification](api-tagspecification.md) objects

Required: No

## Response Elements

The following elements are returned by the service.

**capacityManagerDataExportId**

The unique identifier for the created data export configuration. Use this ID to reference the export in other API calls.

Type: String

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/CreateCapacityManagerDataExport)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/CreateCapacityManagerDataExport)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/createcapacitymanagerdataexport.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/createcapacitymanagerdataexport.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/createcapacitymanagerdataexport.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/createcapacitymanagerdataexport.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/createcapacitymanagerdataexport.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/createcapacitymanagerdataexport.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/CreateCapacityManagerDataExport)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/createcapacitymanagerdataexport.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CopyVolumes

CreateCapacityReservation
